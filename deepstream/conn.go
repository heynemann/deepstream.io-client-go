// deepstream.io-client-go
// https://github.com/heynemann/deepstream.io-client-go
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Bernardo Heynemann <heynemann@gmail.com>

package deepstream

import (
	"fmt"
	"time"

	"github.com/heynemann/deepstream.io-client-go/interfaces"
)

//ClientOptions used to connect to deepstream
type ClientOptions struct {
	AutoReconnect       bool
	AutoLogin           bool
	ConnectionTimeoutMs int
	WriteTimeoutMs      int
	ReadTimeoutMs       int
	Username            string
	Password            string
	HeartbeatIntervalMs int
	ErrorHandler        func(error)
}

//DefaultOptions to connect to deepstream
func DefaultOptions() *ClientOptions {
	return &ClientOptions{
		AutoReconnect:       true,
		AutoLogin:           true,
		Username:            "",
		Password:            "",
		ErrorHandler:        nil,
		ConnectionTimeoutMs: 100,
		WriteTimeoutMs:      100,
		ReadTimeoutMs:       100,
		HeartbeatIntervalMs: 2000,
	}
}

//Client represents a connection to a deepstream.io server
type Client struct {
	AuthParams     map[string]interface{}
	Connector      *Connector
	Options        *ClientOptions
	Event          *EventManager
	loginRequested bool
	lastHeartbeat  time.Time
}

//New creates a new client
func New(url string, optionsOrNil ...*ClientOptions) (*Client, error) {
	options := DefaultOptions()
	if len(optionsOrNil) == 1 {
		options = optionsOrNil[0]
	}
	username := options.Username
	password := options.Password
	authParams := map[string]interface{}{}
	if username != "" {
		authParams = map[string]interface{}{
			"username": username,
			"password": password,
		}
	}
	conn := NewConnector(url, options.ConnectionTimeoutMs, options.WriteTimeoutMs, options.ReadTimeoutMs)
	cli := &Client{
		Connector:      conn,
		Options:        options,
		AuthParams:     authParams,
		loginRequested: false,
	}
	cli.Event = NewEventManager(cli)

	cli.Connector.AddMessageHandler(cli.onMessage)

	err := cli.startMonitoringConnection()
	if err != nil {
		return cli, err
	}

	return cli, nil
}

func (c *Client) startMonitoringConnection() error {
	err := c.Connector.Connect()
	if err != nil {
		return err
	}
	c.lastHeartbeat = time.Now()
	c.startMonitoringHeartbeat()
	return nil
}

func (c *Client) startMonitoringHeartbeat() {
	tolerance := time.Duration(c.Options.HeartbeatIntervalMs*2) * time.Millisecond

	go func() {
		for {
			if time.Now().Sub(c.lastHeartbeat) > tolerance {
				//TODO: Change this to a typed error
				c.Error(fmt.Errorf("Two connections heartbeats missed successively"))
				c.Close()
				return
			}
			time.Sleep(time.Duration(c.Options.HeartbeatIntervalMs) * time.Millisecond)
		}
	}()
}

//var heartBeatTolerance = this._options.heartbeatInterval * 2;

//if( Date.now() - this._lastHeartBeat > heartBeatTolerance ) {
//clearInterval( this._heartbeatInterval );
//this._endpoint.close();
//this._onError( 'Two connections heartbeats missed successively' );
//}

//GetConnectionState returns the connection state for the connector
func (c *Client) GetConnectionState() interfaces.ConnectionState {
	return c.Connector.ConnectionState
}

func (c *Client) onMessage(msg *Message) {
	//fmt.Println(msg)
	var err error
	switch {
	case msg.Topic == "C":
		err = c.handleConnectionMessages(msg)
	case msg.Topic == "A":
		err = c.handleAuthenticationMessages(msg)
	case msg.Topic == "E":
		err = c.handleEventMessages(msg)
	}

	if err != nil {
		c.Error(err)
	}
}

func (c *Client) handleConnectionMessages(msg *Message) error {
	switch {
	case msg.Action == "CH":
		return c.handleChallengeRequest(msg)
	case msg.Action == "A":
		if c.Connector.ConnectionState == interfaces.ConnectionStateChallenging {
			return c.handleChallengeAck(msg)
		}
	case msg.Action == "PI":
		return c.handlePing(msg)
	default:
		fmt.Println("Message not understood!")
	}

	return nil
}

func (c *Client) handleChallengeRequest(msg *Message) error {
	c.Connector.ConnectionState = interfaces.ConnectionStateChallenging
	challenge := NewChallengeResponseAction(c.Connector.URL)
	return c.Connector.WriteMessage([]byte(challenge.ToAction()))
}

func (c *Client) handleChallengeAck(msg *Message) error {
	c.Connector.ConnectionState = interfaces.ConnectionStateAwaitingConnection
	if c.Options.AutoLogin || c.loginRequested {
		return c.Login()
	}
	return nil
}

func (c *Client) handlePing(msg *Message) error {
	pong := &PongAction{}
	return c.Connector.WriteMessage([]byte(pong.ToAction()))
}

//Login to deepstream - if connection is still being started, it will login as soon as possible
func (c *Client) Login() error {
	state := c.GetConnectionState()
	if !c.Options.AutoLogin && (state == interfaces.ConnectionStateChallenging ||
		state == interfaces.ConnectionStateAwaitingConnection) {
		c.loginRequested = true
		return nil
	}

	if state != interfaces.ConnectionStateAwaitingConnection {
		return fmt.Errorf("The connection should be restored before logging in (%s).", state)
	}

	c.loginRequested = false

	authRequestAction, err := NewAuthRequestAction(c.AuthParams)
	if err != nil {
		return err
	}

	c.Connector.ConnectionState = interfaces.ConnectionStateAuthenticating

	//Send Authentication Request
	return c.Connector.WriteMessage([]byte(authRequestAction.ToAction()))
}

func (c *Client) handleAuthenticationMessages(msg *Message) error {
	switch {
	case msg.Action == "A":
		if c.Connector.ConnectionState == interfaces.ConnectionStateAuthenticating {
			return c.handleAuthenticationAck(msg)
		}
	case msg.Action == "E":
		return fmt.Errorf(
			"Could not connect to deepstream.io server with the provided credentials (user: %s).",
			c.AuthParams["user"],
		)
	default:
		fmt.Println("Message not understood!")
	}

	return nil
}

func (c *Client) handleAuthenticationAck(msg *Message) error {
	c.Connector.ConnectionState = interfaces.ConnectionStateOpen
	return nil
}

func (c *Client) handleEventMessages(msg *Message) error {
	switch {
	case msg.Action == "A":
		return c.Event.handleEventSubscriptionAck(msg)
	case msg.Action == "EVT":
		return c.Event.handleEventMessageReceived(msg)
	default:
		fmt.Println("Message not understood!")
	}

	return nil
}

//Error handlers errors in client
func (c *Client) Error(err error) error {
	c.Connector.ConnectionState = interfaces.ConnectionStateError
	c.onError(err)
	return err
}

//Close the connection
func (c *Client) Close() error {
	err := c.Connector.Close()
	if err != nil {
		return c.Error(err)
	}
	return nil
}

func (c *Client) onError(err error) {
	if c.Options.ErrorHandler != nil {
		c.Options.ErrorHandler(err)
	}
}
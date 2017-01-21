// deepstream.io-client-go
// https://github.com/heynemann/deepstream.io-client-go
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Bernardo Heynemann <heynemann@gmail.com>

package interfaces

//MessageSeparator invisible message separator
const MessageSeparator = "\u001e"

//MessagePartSeparator invisible message separator
const MessagePartSeparator = "\u001f"

//SourceMessageConnector identifies source message connector key
const SourceMessageConnector = "SOURCE_MESSAGE_CONNECTOR"

//Log Level

//LogLevelDebug identifies log level
const LogLevelDebug = 0

//LogLevelInfo identifies log level
const LogLevelInfo = 1

//LogLevelWarn identifies log level
const LogLevelWarn = 2

//LogLevelError identifies log level
const LogLevelError = 2

//LogLevelOff identifies log level
const LogLevelOff = 100

//Server State

//ServerStateStarting indicates the server is starting
const ServerStateStarting = "starting"

//ServerStateInitialized indicates the server is initialized
const ServerStateInitialized = "initialized"

//ServerStateRunning indicates the server is initialized
const ServerStateRunning = "is-running"

//ServerStateClosing indicates the server is closing
const ServerStateClosing = "closing"

//ServerStateClosed indicates the server is closed
const ServerStateClosed = "closed"

//Connection State

//ConnectionStateClosed indicates the connection has been closed
const ConnectionStateClosed = "CLOSED"

//ConnectionStateAwaitingConnection indicates the connection is waiting for a connection
const ConnectionStateAwaitingConnection = "AWAITING_CONNECTION"

//ConnectionStateChallenging indicates the connection is being challenged for credentials
const ConnectionStateChallenging = "CHALLENGING"

//ConnectionStateAwaitingAuthentication indicates the connection is waiting for authentication
const ConnectionStateAwaitingAuthentication = "AWAITING_AUTHENTICATION"

//ConnectionStateAuthenticating indicates the connection is authenticating with the server
const ConnectionStateAuthenticating = "AUTHENTICATING"

//ConnectionStateOpen indicates the connection is open
const ConnectionStateOpen = "OPEN"

//ConnectionStateError indicates the connection has errored
const ConnectionStateError = "ERROR"

//ConnectionStateReconnecting indicates the connection is reconnecting
const ConnectionStateReconnecting = "RECONNECTING"

//Event

//name	value	server	client
//EVENT.TRIGGER_EVENT	TRIGGER_EVENT	✔
//EVENT.INCOMING_CONNECTION	INCOMING_CONNECTION	✔
//EVENT.INFO	INFO	✔
//EVENT.SUBSCRIBE	SUBSCRIBE	✔
//EVENT.UNSUBSCRIBE	UNSUBSCRIBE	✔
//EVENT.RECORD_DELETION	RECORD_DELETION	✔
//EVENT.INVALID_AUTH_MSG	INVALID_AUTH_MSG	✔
//EVENT.INVALID_AUTH_DATA	INVALID_AUTH_DATA	✔
//EVENT.AUTH_ATTEMPT	AUTH_ATTEMPT	✔
//EVENT.AUTH_ERROR	AUTH_ERROR	✔
//EVENT.TOO_MANY_AUTH_ATTEMPTS	TOO_MANY_AUTH_ATTEMPTS	✔	✔
//EVENT.AUTH_SUCCESSFUL	AUTH_SUCCESSFUL	✔
//EVENT.NOT_AUTHENTICATED	NOT_AUTHENTICATED		✔
//EVENT.CONNECTION_ERROR	CONNECTION_ERROR	✔	✔
//EVENT.MESSAGE_PERMISSION_ERROR	MESSAGE_PERMISSION_ERROR	✔	✔
//EVENT.MESSAGE_PARSE_ERROR	MESSAGE_PARSE_ERROR	✔	✔
//EVENT.MAXIMUM_MESSAGE_SIZE_EXCEEDED	MAXIMUM_MESSAGE_SIZE_EXCEEDED	✔
//EVENT.MESSAGE_DENIED	MESSAGE_DENIED	✔	✔
//EVENT.INVALID_MESSAGE_DATA	INVALID_MESSAGE_DATA	✔
//EVENT.UNKNOWN_TOPIC	UNKNOWN_TOPIC	✔
//EVENT.UNKNOWN_ACTION	UNKNOWN_ACTION	✔
//EVENT.MULTIPLE_SUBSCRIPTIONS	MULTIPLE_SUBSCRIPTIONS	✔
//EVENT.NOT_SUBSCRIBED	NOT_SUBSCRIBED	✔
//EVENT.LISTENER_EXISTS	LISTENER_EXISTS		✔
//EVENT.NOT_LISTENING	NOT_LISTENING		✔
//EVENT.IS_CLOSED	IS_CLOSED		✔
//EVENT.ACK_TIMEOUT	ACK_TIMEOUT	✔	✔
//EVENT.RESPONSE_TIMEOUT	RESPONSE_TIMEOUT	✔	✔
//EVENT.DELETE_TIMEOUT	DELETE_TIMEOUT		✔
//EVENT.UNSOLICITED_MESSAGE	UNSOLICITED_MESSAGE		✔
//EVENT.MULTIPLE_ACK	MULTIPLE_ACK	✔
//EVENT.MULTIPLE_RESPONSE	MULTIPLE_RESPONSE	✔
//EVENT.NO_RPC_PROVIDER	NO_RPC_PROVIDER	✔
//EVENT.RECORD_LOAD_ERROR	RECORD_LOAD_ERROR	✔
//EVENT.RECORD_CREATE_ERROR	RECORD_CREATE_ERROR	✔
//EVENT.RECORD_UPDATE_ERROR	RECORD_UPDATE_ERROR	✔
//EVENT.RECORD_DELETE_ERROR	RECORD_DELETE_ERROR	✔
//EVENT.RECORD_SNAPSHOT_ERROR	RECORD_SNAPSHOT_ERROR	✔
//EVENT.RECORD_NOT_FOUND	RECORD_NOT_FOUND	✔	✔
//EVENT.CACHE_RETRIEVAL_TIMEOUT	CACHE_RETRIEVAL_TIMEOUT	✔
//EVENT.STORAGE_RETRIEVAL_TIMEOUT	STORAGE_RETRIEVAL_TIMEOUT	✔
//EVENT.CLOSED_SOCKET_INTERACTION	CLOSED_SOCKET_INTERACTION	✔
//EVENT.CLIENT_DISCONNECTED	CLIENT_DISCONNECTED	✔
//EVENT.INVALID_MESSAGE	INVALID_MESSAGE	✔
//EVENT.VERSION_EXISTS	VERSION_EXISTS	✔	✔
//EVENT.INVALID_VERSION	INVALID_VERSION	✔
//EVENT.PLUGIN_ERROR	PLUGIN_ERROR	✔
//EVENT.UNKNOWN_CALLEE	UNKNOWN_CALLEE	✔	✔
//Topic

//name	value	server	client
//TOPIC.CONNECTION	C	✔	✔
//TOPIC.AUTH	A	✔	✔
//TOPIC.ERROR	X	✔	✔
//TOPIC.EVENT	E	✔	✔
//TOPIC.RECORD	R	✔	✔
//TOPIC.RPC	P	✔	✔
//TOPIC.PRIVATE	PRIVATE/	✔	✔
//Actions

//name	value	server	client
//ACTIONS.ACK	A	✔	✔
//ACTIONS.READ	R	✔	✔
//ACTIONS.REDIRECT	RED		✔
//ACTIONS.CHALLENGE	CH		✔
//ACTIONS.CHALLENGE_RESPONSE	CHR		✔
//ACTIONS.CREATE	C	✔	✔
//ACTIONS.UPDATE	U	✔	✔
//ACTIONS.PATCH	P	✔	✔
//ACTIONS.DELETE	D	✔	✔
//ACTIONS.SUBSCRIBE	S	✔	✔
//ACTIONS.UNSUBSCRIBE	US	✔	✔
//ACTIONS.HAS	H	✔	✔
//ACTIONS.SNAPSHOT	SN	✔	✔
//ACTIONS.LISTEN_SNAPSHOT	LSN	✔
//ACTIONS.LISTEN	L	✔	✔
//ACTIONS.UNLISTEN	UL	✔	✔
//ACTIONS.LISTEN_ACCEPT	LA	✔	✔
//ACTIONS.LISTEN_REJECT	LR	✔	✔
//ACTIONS.SUBSCRIPTION_HAS_PROVIDER	SH	✔	✔
//ACTIONS.SUBSCRIPTIONS_FOR_PATTERN_FOUND	SF	✔
//ACTIONS.SUBSCRIPTION_FOR_PATTERN_FOUND	SP	✔
//ACTIONS.SUBSCRIPTION_FOR_PATTERN_REMOVED	SR	✔
//ACTIONS.PROVIDER_UPDATE	PU	✔	✔
//ACTIONS.QUERY	Q	✔	✔
//ACTIONS.CREATEORREAD	CR	✔	✔
//ACTIONS.EVENT	EVT	✔	✔
//ACTIONS.ERROR	E	✔	✔
//ACTIONS.REQUEST	REQ	✔	✔
//ACTIONS.RESPONSE	RES	✔	✔
//ACTIONS.REJECTION	REJ	✔	✔
//Data Types

//DataType represents one of the available data types in an action
type DataType string

//TypesString indicates that the data in an action is of type string
const TypesString DataType = "S"

//TypesObject indicates that the data in an action is of type object (interface{})
const TypesObject DataType = "O"

//TypesNumber indicates that the data in an action is of type number
const TypesNumber DataType = "N"

//TypesNull indicates that the data in an action is nil
const TypesNull DataType = "L"

//TypesTrue indicates that the data in an action is true
const TypesTrue DataType = "T"

//TypesFalse indicates that the data in an action is false
const TypesFalse DataType = "F"

//TypesUndefined indicates that the data in an action is undefined
const TypesUndefined DataType = "U"
package event

const (
	// AgentCalled event.
	AgentCalled = "AgentCalled"

	// AgentComplete event.
	AgentComplete = "AgentComplete"

	// AgentConnect event.
	AgentConnect = "AgentConnect"

	// AgentDump event.
	AgentDump = "AgentDump"

	// AgentLogin event.
	AgentLogin = "AgentLogin"

	// AgentLogoff event.
	AgentLogoff = "AgentLogoff"

	// AgentRingNoAnswer event.
	AgentRingNoAnswer = "AgentRingNoAnswer"

	// Agents event.
	Agents = "Agents"

	// AgentsComplete event.
	AgentsComplete = "AgentsComplete"

	// AGIExecEnd event.
	AGIExecEnd = "AGIExecEnd"

	// AGIExecStart event.
	AGIExecStart = "AGIExecStart"

	// Alarm event.
	Alarm = "Alarm"

	// AlarmClear event.
	AlarmClear = "AlarmClear"

	// AOCD event.
	AOCD = "AOC-D"

	// AOCE event.
	AOCE = "AOC-E"

	// AOCS event.
	AOCS = "AOC-S"

	// AorDetail event.
	AorDetail = "AorDetail"

	// AorList event.
	AorList = "AorList"

	// AorListComplete event.
	AorListComplete = "AorListComplete"

	// AsyncAGIEnd event.
	AsyncAGIEnd = "AsyncAGIEnd"

	// AsyncAGIExec event.
	AsyncAGIExec = "AsyncAGIExec"

	// AsyncAGIStart event.
	AsyncAGIStart = "AsyncAGIStart"

	// AttendedTransfer event.
	AttendedTransfer = "AttendedTransfer"

	// AuthDetail event.
	AuthDetail = "AuthDetail"

	// AuthList event.
	AuthList = "AuthList"

	// AuthListComplete event.
	AuthListComplete = "AuthListComplete"

	// AuthMethodNotAllowed event.
	AuthMethodNotAllowed = "AuthMethodNotAllowed"

	// BlindTransfer event.
	BlindTransfer = "BlindTransfer"

	// BridgeCreate event.
	BridgeCreate = "BridgeCreate"

	// BridgeDestroy event.
	BridgeDestroy = "BridgeDestroy"

	// BridgeEnter event.
	BridgeEnter = "BridgeEnter"

	// BridgeInfoChannel event.
	BridgeInfoChannel = "BridgeInfoChannel"

	// BridgeInfoComplete event.
	BridgeInfoComplete = "BridgeInfoComplete"

	// BridgeLeave event.
	BridgeLeave = "BridgeLeave"

	// BridgeMerge event.
	BridgeMerge = "BridgeMerge"

	// BridgeVideoSourceUpdate event.
	BridgeVideoSourceUpdate = "BridgeVideoSourceUpdate"

	// Cdr event.
	Cdr = "Cdr"

	// CEL event.
	CEL = "CEL"

	// ChallengeResponseFailed event.
	ChallengeResponseFailed = "ChallengeResponseFailed"

	// ChallengeSent event.
	ChallengeSent = "ChallengeSent"

	// ChannelTalkingStart event.
	ChannelTalkingStart = "ChannelTalkingStart"

	// ChannelTalkingStop event.
	ChannelTalkingStop = "ChannelTalkingStop"

	// ChanSpyStart event.
	ChanSpyStart = "ChanSpyStart"

	// ChanSpyStop event.
	ChanSpyStop = "ChanSpyStop"

	// ConfbridgeEnd event.
	ConfbridgeEnd = "ConfbridgeEnd"

	// ConfbridgeJoin event.
	ConfbridgeJoin = "ConfbridgeJoin"

	// ConfbridgeLeave event.
	ConfbridgeLeave = "ConfbridgeLeave"

	// ConfbridgeList event.
	ConfbridgeList = "ConfbridgeList"

	// ConfbridgeMute event.
	ConfbridgeMute = "ConfbridgeMute"

	// ConfbridgeRecord event.
	ConfbridgeRecord = "ConfbridgeRecord"

	// ConfbridgeStart event.
	ConfbridgeStart = "ConfbridgeStart"

	// ConfbridgeStopRecord event.
	ConfbridgeStopRecord = "ConfbridgeStopRecord"

	// ConfbridgeTalking event.
	ConfbridgeTalking = "ConfbridgeTalking"

	// ConfbridgeUnmute event.
	ConfbridgeUnmute = "ConfbridgeUnmute"

	// ContactList event.
	ContactList = "ContactList"

	// ContactListComplete event.
	ContactListComplete = "ContactListComplete"

	// ContactStatus event.
	ContactStatus = "ContactStatus"

	// ContactStatusDetail event.
	ContactStatusDetail = "ContactStatusDetail"

	// CoreShowChannel event.
	CoreShowChannel = "CoreShowChannel"

	// CoreShowChannelsComplete event.
	CoreShowChannelsComplete = "CoreShowChannelsComplete"

	// DAHDIChannel event.
	DAHDIChannel = "DAHDIChannel"

	// DeviceStateChange event.
	DeviceStateChange = "DeviceStateChange"

	// DeviceStateListComplete event.
	DeviceStateListComplete = "DeviceStateListComplete"

	// DialBegin event.
	DialBegin = "DialBegin"

	// DialEnd event.
	DialEnd = "DialEnd"

	// DialState event.
	DialState = "DialState"

	// DNDState event.
	DNDState = "DNDState"

	// DTMFBegin event.
	DTMFBegin = "DTMFBegin"

	// DTMFEnd event.
	DTMFEnd = "DTMFEnd"

	// EndpointDetail event.
	EndpointDetail = "EndpointDetail"

	// EndpointDetailComplete event.
	EndpointDetailComplete = "EndpointDetailComplete"

	// EndpointList event.
	EndpointList = "EndpointList"

	// EndpointListComplete event.
	EndpointListComplete = "EndpointListComplete"

	// ExtensionStateListComplete event.
	ExtensionStateListComplete = "ExtensionStateListComplete"

	// ExtensionStatus event.
	ExtensionStatus = "ExtensionStatus"

	// FailedACL event.
	FailedACL = "FailedACL"

	// FAXSession event.
	FAXSession = "FAXSession"

	// FAXSessionsComplete event.
	FAXSessionsComplete = "FAXSessionsComplete"

	// FAXSessionsEntry event.
	FAXSessionsEntry = "FAXSessionsEntry"

	// FAXStats event.
	FAXStats = "FAXStats"

	// FAXStatus event.
	FAXStatus = "FAXStatus"

	// FullyBooted event.
	FullyBooted = "FullyBooted"

	// Hangup event.
	Hangup = "Hangup"

	// HangupHandlerPop event.
	HangupHandlerPop = "HangupHandlerPop"

	// HangupHandlerPush event.
	HangupHandlerPush = "HangupHandlerPush"

	// HangupHandlerRun event.
	HangupHandlerRun = "HangupHandlerRun"

	// HangupRequest event.
	HangupRequest = "HangupRequest"

	// Hold event.
	Hold = "Hold"

	// IdentifyDetail event.
	IdentifyDetail = "IdentifyDetail"

	// InvalidAccountID event.
	InvalidAccountID = "InvalidAccountID"

	// InvalidPassword event.
	InvalidPassword = "InvalidPassword"

	// InvalidTransport event.
	InvalidTransport = "InvalidTransport"

	// Load event.
	Load = "Load"

	// LoadAverageLimit event.
	LoadAverageLimit = "LoadAverageLimit"

	// LocalBridge event.
	LocalBridge = "LocalBridge"

	// LocalOptimizationBegin event.
	LocalOptimizationBegin = "LocalOptimizationBegin"

	// LocalOptimizationEnd event.
	LocalOptimizationEnd = "LocalOptimizationEnd"

	// LogChannel event.
	LogChannel = "LogChannel"

	// MCID event.
	MCID = "MCID"

	// MeetmeEnd event.
	MeetmeEnd = "MeetmeEnd"

	// MeetmeJoin event.
	MeetmeJoin = "MeetmeJoin"

	// MeetmeLeave event.
	MeetmeLeave = "MeetmeLeave"

	// MeetmeMute event.
	MeetmeMute = "MeetmeMute"

	// MeetmeTalking event.
	MeetmeTalking = "MeetmeTalking"

	// MeetmeTalkRequest event.
	MeetmeTalkRequest = "MeetmeTalkRequest"

	// MemoryLimit event.
	MemoryLimit = "MemoryLimit"

	// MessageWaiting event.
	MessageWaiting = "MessageWaiting"

	// MiniVoiceMail event.
	MiniVoiceMail = "MiniVoiceMail"

	// MonitorStart event.
	MonitorStart = "MonitorStart"

	// MonitorStop event.
	MonitorStop = "MonitorStop"

	// MusicOnHoldStart event.
	MusicOnHoldStart = "MusicOnHoldStart"

	// MusicOnHoldStop event.
	MusicOnHoldStop = "MusicOnHoldStop"

	// MWIGet event.
	MWIGet = "MWIGet"

	// MWIGetComplete event.
	MWIGetComplete = "MWIGetComplete"

	// NewAccountCode event.
	NewAccountCode = "NewAccountCode"

	// NewCallerid event.
	NewCallerid = "NewCallerid"

	// Newchannel event.
	Newchannel = "Newchannel"

	// NewConnectedLine event.
	NewConnectedLine = "NewConnectedLine"

	// NewExten event.
	NewExten = "NewExten"

	// Newstate event.
	Newstate = "Newstate"

	// OriginateResponse event.
	OriginateResponse = "OriginateResponse"

	// ParkedCall event.
	ParkedCall = "ParkedCall"

	// ParkedCallGiveUp event.
	ParkedCallGiveUp = "ParkedCallGiveUp"

	// ParkedCallSwap event.
	ParkedCallSwap = "ParkedCallSwap"

	// ParkedCallTimeOut event.
	ParkedCallTimeOut = "ParkedCallTimeOut"

	// PeerStatus event.
	PeerStatus = "PeerStatus"

	// Pickup event.
	Pickup = "Pickup"

	// PresenceStateChange event.
	PresenceStateChange = "PresenceStateChange"

	// PresenceStateListComplete event.
	PresenceStateListComplete = "PresenceStateListComplete"

	// PresenceStatus event.
	PresenceStatus = "PresenceStatus"

	// QueueCallerAbandon event.
	QueueCallerAbandon = "QueueCallerAbandon"

	// QueueCallerJoin event.
	QueueCallerJoin = "QueueCallerJoin"

	// QueueCallerLeave event.
	QueueCallerLeave = "QueueCallerLeave"

	// QueueMemberAdded event.
	QueueMemberAdded = "QueueMemberAdded"

	// QueueMemberPause event.
	QueueMemberPause = "QueueMemberPause"

	// QueueMemberPenalty event.
	QueueMemberPenalty = "QueueMemberPenalty"

	// QueueMemberRemoved event.
	QueueMemberRemoved = "QueueMemberRemoved"

	// QueueMemberRinginuse event.
	QueueMemberRinginuse = "QueueMemberRinginuse"

	// QueueMemberStatus event.
	QueueMemberStatus = "QueueMemberStatus"

	// ReceiveFAX event.
	ReceiveFAX = "ReceiveFAX"

	// Registry event.
	Registry = "Registry"

	// Reload event.
	Reload = "Reload"

	// Rename event.
	Rename = "Rename"

	// RequestBadFormat event.
	RequestBadFormat = "RequestBadFormat"

	// RequestNotAllowed event.
	RequestNotAllowed = "RequestNotAllowed"

	// RequestNotSupported event.
	RequestNotSupported = "RequestNotSupported"

	// RTCPReceived event.
	RTCPReceived = "RTCPReceived"

	// RTCPSent event.
	RTCPSent = "RTCPSent"

	// SendFAX event.
	SendFAX = "SendFAX"

	// SessionLimit event.
	SessionLimit = "SessionLimit"

	// SessionTimeout event.
	SessionTimeout = "SessionTimeout"

	// Shutdown event.
	Shutdown = "Shutdown"

	// SIPQualifyPeerDone event.
	SIPQualifyPeerDone = "SIPQualifyPeerDone"

	// SoftHangupRequest event.
	SoftHangupRequest = "SoftHangupRequest"

	// SpanAlarm event.
	SpanAlarm = "SpanAlarm"

	// SpanAlarmClear event.
	SpanAlarmClear = "SpanAlarmClear"

	// Status event.
	Status = "Status"

	// StatusComplete event.
	StatusComplete = "StatusComplete"

	// SuccessfulAuth event.
	SuccessfulAuth = "SuccessfulAuth"

	// TransportDetail event.
	TransportDetail = "TransportDetail"

	// UnexpectedAddress event.
	UnexpectedAddress = "UnexpectedAddress"

	// Unhold event.
	Unhold = "Unhold"

	// Unload event.
	Unload = "Unload"

	// UnParkedCall event.
	UnParkedCall = "UnParkedCall"

	// UserEvent event.
	UserEvent = "UserEvent"

	// VarSet event.
	VarSet = "VarSet"
)

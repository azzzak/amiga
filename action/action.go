package action

const (
	// AbsoluteTimeout action.
	AbsoluteTimeout = "AbsoluteTimeout"

	// AgentLogoff action.
	AgentLogoff = "AgentLogoff"

	// Agents action.
	Agents = "Agents"

	// AGI action.
	AGI = "AGI"

	// AOCMessage action.
	AOCMessage = "AOCMessage"

	// Atxfer action.
	Atxfer = "Atxfer"

	// BlindTransfer action.
	BlindTransfer = "BlindTransfer"

	// Bridge action.
	Bridge = "Bridge"

	// BridgeDestroy action.
	BridgeDestroy = "BridgeDestroy"

	// BridgeInfo action.
	BridgeInfo = "BridgeInfo"

	// BridgeKick action.
	BridgeKick = "BridgeKick"

	// BridgeList action.
	BridgeList = "BridgeList"

	// BridgeTechnologyList action.
	BridgeTechnologyList = "BridgeTechnologyList"

	// BridgeTechnologySuspend action.
	BridgeTechnologySuspend = "BridgeTechnologySuspend"

	// BridgeTechnologyUnsuspend action.
	BridgeTechnologyUnsuspend = "BridgeTechnologyUnsuspend"

	// CancelAtxfer action.
	CancelAtxfer = "CancelAtxfer"

	// Challenge action.
	Challenge = "Challenge"

	// ChangeMonitor action.
	ChangeMonitor = "ChangeMonitor"

	// Command action.
	Command = "Command"

	// ConfbridgeKick action.
	ConfbridgeKick = "ConfbridgeKick"

	// ConfbridgeList action.
	ConfbridgeList = "ConfbridgeList"

	// ConfbridgeListRooms action.
	ConfbridgeListRooms = "ConfbridgeListRooms"

	// ConfbridgeLock action.
	ConfbridgeLock = "ConfbridgeLock"

	// ConfbridgeMute action.
	ConfbridgeMute = "ConfbridgeMute"

	// ConfbridgeSetSingleVideoSrc action.
	ConfbridgeSetSingleVideoSrc = "ConfbridgeSetSingleVideoSrc"

	// ConfbridgeStartRecord action.
	ConfbridgeStartRecord = "ConfbridgeStartRecord"

	// ConfbridgeStopRecord action.
	ConfbridgeStopRecord = "ConfbridgeStopRecord"

	// ConfbridgeUnlock action.
	ConfbridgeUnlock = "ConfbridgeUnlock"

	// ConfbridgeUnmute action.
	ConfbridgeUnmute = "ConfbridgeUnmute"

	// ControlPlayback action.
	ControlPlayback = "ControlPlayback"

	// CoreSettings action.
	CoreSettings = "CoreSettings"

	// CoreShowChannels action.
	CoreShowChannels = "CoreShowChannels"

	// CoreStatus action.
	CoreStatus = "CoreStatus"

	// CreateConfig action.
	CreateConfig = "CreateConfig"

	// DAHDIDialOffhook action.
	DAHDIDialOffhook = "DAHDIDialOffhook"

	// DAHDIDNDoff action.
	DAHDIDNDoff = "DAHDIDNDoff"

	// DAHDIDNDon action.
	DAHDIDNDon = "DAHDIDNDon"

	// DAHDIHangup action.
	DAHDIHangup = "DAHDIHangup"

	// DAHDIRestart action.
	DAHDIRestart = "DAHDIRestart"

	// DAHDIShowChannels action.
	DAHDIShowChannels = "DAHDIShowChannels"

	// DAHDITransfer action.
	DAHDITransfer = "DAHDITransfer"

	// DataGet action.
	DataGet = "DataGet"

	// DBDel action.
	DBDel = "DBDel"

	// DBDelTree action.
	DBDelTree = "DBDelTree"

	// DBGet action.
	DBGet = "DBGet"

	// DBPut action.
	DBPut = "DBPut"

	// DeviceStateList action.
	DeviceStateList = "DeviceStateList"

	// DialplanExtensionAdd action.
	DialplanExtensionAdd = "DialplanExtensionAdd"

	// DialplanExtensionRemove action.
	DialplanExtensionRemove = "DialplanExtensionRemove"

	// Events action.
	Events = "Events"

	// ExtensionState action.
	ExtensionState = "ExtensionState"

	// ExtensionStateList action.
	ExtensionStateList = "ExtensionStateList"

	// FAXSession action.
	FAXSession = "FAXSession"

	// FAXSessions action.
	FAXSessions = "FAXSessions"

	// FAXStats action.
	FAXStats = "FAXStats"

	// Filter action.
	Filter = "Filter"

	// FilterList action.
	FilterList = "FilterList"

	// GetConfig action.
	GetConfig = "GetConfig"

	// GetConfigJSON action.
	GetConfigJSON = "GetConfigJSON"

	// Getvar action.
	Getvar = "Getvar"

	// Hangup action.
	Hangup = "Hangup"

	// IAXnetstats action.
	IAXnetstats = "IAXnetstats"

	// IAXpeerlist action.
	IAXpeerlist = "IAXpeerlist"

	// IAXpeers action.
	IAXpeers = "IAXpeers"

	// IAXregistry action.
	IAXregistry = "IAXregistry"

	// JabberSendResXMPP action.
	JabberSendResXMPP = "JabberSend_res_xmpp"

	// ListCategories action.
	ListCategories = "ListCategories"

	// ListCommands action.
	ListCommands = "ListCommands"

	// LocalOptimizeAway action.
	LocalOptimizeAway = "LocalOptimizeAway"

	// LoggerRotate action.
	LoggerRotate = "LoggerRotate"

	// Login action.
	Login = "Login"

	// Logoff action.
	Logoff = "Logoff"

	// MailboxCount action.
	MailboxCount = "MailboxCount"

	// MailboxStatus action.
	MailboxStatus = "MailboxStatus"

	// MeetmeList action.
	MeetmeList = "MeetmeList"

	// MeetmeListRooms action.
	MeetmeListRooms = "MeetmeListRooms"

	// MeetmeMute action.
	MeetmeMute = "MeetmeMute"

	// MeetmeUnmute action.
	MeetmeUnmute = "MeetmeUnmute"

	// MessageSend action.
	MessageSend = "MessageSend"

	// MixMonitor action.
	MixMonitor = "MixMonitor"

	// MixMonitorMute action.
	MixMonitorMute = "MixMonitorMute"

	// ModuleCheck action.
	ModuleCheck = "ModuleCheck"

	// ModuleLoad action.
	ModuleLoad = "ModuleLoad"

	// Monitor action.
	Monitor = "Monitor"

	// MuteAudio action.
	MuteAudio = "MuteAudio"

	// MWIDelete action.
	MWIDelete = "MWIDelete"

	// MWIGet action.
	MWIGet = "MWIGet"

	// MWIUpdate action.
	MWIUpdate = "MWIUpdate"

	// Originate action.
	Originate = "Originate"

	// Park action.
	Park = "Park"

	// ParkedCalls action.
	ParkedCalls = "ParkedCalls"

	// Parkinglots action.
	Parkinglots = "Parkinglots"

	// PauseMonitor action.
	PauseMonitor = "PauseMonitor"

	// Ping action.
	Ping = "Ping"

	// PJSIPNotify action.
	PJSIPNotify = "PJSIPNotify"

	// PJSIPQualify action.
	PJSIPQualify = "PJSIPQualify"

	// PJSIPRegister action.
	PJSIPRegister = "PJSIPRegister"

	// PJSIPShowAors action.
	PJSIPShowAors = "PJSIPShowAors"

	// PJSIPShowAuths action.
	PJSIPShowAuths = "PJSIPShowAuths"

	// PJSIPShowContacts action.
	PJSIPShowContacts = "PJSIPShowContacts"

	// PJSIPShowEndpoint action.
	PJSIPShowEndpoint = "PJSIPShowEndpoint"

	// PJSIPShowEndpoints action.
	PJSIPShowEndpoints = "PJSIPShowEndpoints"

	// PJSIPShowRegistrationInboundContactStatuses action.
	PJSIPShowRegistrationInboundContactStatuses = "PJSIPShowRegistrationInboundContactStatuses"

	// PJSIPShowRegistrationsInbound action.
	PJSIPShowRegistrationsInbound = "PJSIPShowRegistrationsInbound"

	// PJSIPShowRegistrationsOutbound action.
	PJSIPShowRegistrationsOutbound = "PJSIPShowRegistrationsOutbound"

	// PJSIPShowResourceLists action.
	PJSIPShowResourceLists = "PJSIPShowResourceLists"

	// PJSIPShowSubscriptionsInbound action.
	PJSIPShowSubscriptionsInbound = "PJSIPShowSubscriptionsInbound"

	// PJSIPShowSubscriptionsOutbound action.
	PJSIPShowSubscriptionsOutbound = "PJSIPShowSubscriptionsOutbound"

	// PJSIPUnregister action.
	PJSIPUnregister = "PJSIPUnregister"

	// PlayDTMF action.
	PlayDTMF = "PlayDTMF"

	// PresenceState action.
	PresenceState = "PresenceState"

	// PresenceStateList action.
	PresenceStateList = "PresenceStateList"

	// PRIDebugFileSet action.
	PRIDebugFileSet = "PRIDebugFileSet"

	// PRIDebugFileUnset action.
	PRIDebugFileUnset = "PRIDebugFileUnset"

	// PRIDebugSet action.
	PRIDebugSet = "PRIDebugSet"

	// PRIShowSpans action.
	PRIShowSpans = "PRIShowSpans"

	// QueueAdd action.
	QueueAdd = "QueueAdd"

	// QueueChangePriorityCaller action.
	QueueChangePriorityCaller = "QueueChangePriorityCaller"

	// QueueLog action.
	QueueLog = "QueueLog"

	// QueueMemberRingInUse action.
	QueueMemberRingInUse = "QueueMemberRingInUse"

	// QueuePause action.
	QueuePause = "QueuePause"

	// QueuePenalty action.
	QueuePenalty = "QueuePenalty"

	// QueueReload action.
	QueueReload = "QueueReload"

	// QueueRemove action.
	QueueRemove = "QueueRemove"

	// QueueReset action.
	QueueReset = "QueueReset"

	// QueueRule action.
	QueueRule = "QueueRule"

	// Queues action.
	Queues = "Queues"

	// QueueStatus action.
	QueueStatus = "QueueStatus"

	// QueueSummary action.
	QueueSummary = "QueueSummary"

	// Redirect action.
	Redirect = "Redirect"

	// Reload action.
	Reload = "Reload"

	// SendText action.
	SendText = "SendText"

	// Setvar action.
	Setvar = "Setvar"

	// ShowDialPlan action.
	ShowDialPlan = "ShowDialPlan"

	// SIPnotify action.
	SIPnotify = "SIPnotify"

	// SIPpeers action.
	SIPpeers = "SIPpeers"

	// SIPpeerstatus action.
	SIPpeerstatus = "SIPpeerstatus"

	// SIPqualifypeer action.
	SIPqualifypeer = "SIPqualifypeer"

	// SIPshowpeer action.
	SIPshowpeer = "SIPshowpeer"

	// SIPshowregistry action.
	SIPshowregistry = "SIPshowregistry"

	// SKINNYdevices action.
	SKINNYdevices = "SKINNYdevices"

	// SKINNYlines action.
	SKINNYlines = "SKINNYlines"

	// SKINNYshowdevice action.
	SKINNYshowdevice = "SKINNYshowdevice"

	// SKINNYshowline action.
	SKINNYshowline = "SKINNYshowline"

	// SorceryMemoryCacheExpire action.
	SorceryMemoryCacheExpire = "SorceryMemoryCacheExpire"

	// SorceryMemoryCacheExpireObject action.
	SorceryMemoryCacheExpireObject = "SorceryMemoryCacheExpireObject"

	// SorceryMemoryCachePopulate action.
	SorceryMemoryCachePopulate = "SorceryMemoryCachePopulate"

	// SorceryMemoryCacheStale action.
	SorceryMemoryCacheStale = "SorceryMemoryCacheStale"

	// SorceryMemoryCacheStaleObject action.
	SorceryMemoryCacheStaleObject = "SorceryMemoryCacheStaleObject"

	// Status action.
	Status = "Status"

	// StopMixMonitor action.
	StopMixMonitor = "StopMixMonitor"

	// StopMonitor action.
	StopMonitor = "StopMonitor"

	// UnpauseMonitor action.
	UnpauseMonitor = "UnpauseMonitor"

	// UpdateConfig action.
	UpdateConfig = "UpdateConfig"

	// UserEvent action.
	UserEvent = "UserEvent"

	// VoicemailRefresh action.
	VoicemailRefresh = "VoicemailRefresh"

	// VoicemailUsersList action.
	VoicemailUsersList = "VoicemailUsersList"

	// VoicemailUserStatus action.
	VoicemailUserStatus = "VoicemailUserStatus"

	// WaitEvent action.
	WaitEvent = "WaitEvent"
)

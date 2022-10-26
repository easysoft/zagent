package virtualboxsrv

type SettingsVersion string

const (
	SettingsVersionNull SettingsVersion = "Null"

	SettingsVersionV10 SettingsVersion = "v10"

	SettingsVersionV11 SettingsVersion = "v11"

	SettingsVersionV12 SettingsVersion = "v12"

	SettingsVersionV13pre SettingsVersion = "v13pre"

	SettingsVersionV13 SettingsVersion = "v13"

	SettingsVersionV14 SettingsVersion = "v14"

	SettingsVersionV15 SettingsVersion = "v15"

	SettingsVersionV16 SettingsVersion = "v16"

	SettingsVersionV17 SettingsVersion = "v17"

	SettingsVersionV18 SettingsVersion = "v18"

	SettingsVersionV19 SettingsVersion = "v19"

	SettingsVersionV110 SettingsVersion = "v110"

	SettingsVersionV111 SettingsVersion = "v111"

	SettingsVersionV112 SettingsVersion = "v112"

	SettingsVersionV113 SettingsVersion = "v113"

	SettingsVersionV114 SettingsVersion = "v114"

	SettingsVersionV115 SettingsVersion = "v115"

	SettingsVersionFuture SettingsVersion = "Future"
)

type AccessMode string

const (
	AccessModeReadOnly AccessMode = "ReadOnly"

	AccessModeReadWrite AccessMode = "ReadWrite"
)

type MachineState string

const (
	MachineStateNull MachineState = "Null"

	MachineStatePoweredOff MachineState = "PoweredOff"

	MachineStateSaved MachineState = "Saved"

	MachineStateTeleported MachineState = "Teleported"

	MachineStateAborted MachineState = "Aborted"

	MachineStateRunning MachineState = "Running"

	MachineStatePaused MachineState = "Paused"

	MachineStateStuck MachineState = "Stuck"

	MachineStateTeleporting MachineState = "Teleporting"

	MachineStateLiveSnapshotting MachineState = "LiveSnapshotting"

	MachineStateStarting MachineState = "Starting"

	MachineStateStopping MachineState = "Stopping"

	MachineStateSaving MachineState = "Saving"

	MachineStateRestoring MachineState = "Restoring"

	MachineStateTeleportingPausedVM MachineState = "TeleportingPausedVM"

	MachineStateTeleportingIn MachineState = "TeleportingIn"

	MachineStateFaultTolerantSyncing MachineState = "FaultTolerantSyncing"

	MachineStateDeletingSnapshotOnline MachineState = "DeletingSnapshotOnline"

	MachineStateDeletingSnapshotPaused MachineState = "DeletingSnapshotPaused"

	MachineStateOnlineSnapshotting MachineState = "OnlineSnapshotting"

	MachineStateRestoringSnapshot MachineState = "RestoringSnapshot"

	MachineStateDeletingSnapshot MachineState = "DeletingSnapshot"

	MachineStateSettingUp MachineState = "SettingUp"

	MachineStateSnapshotting MachineState = "Snapshotting"

	MachineStateFirstOnline MachineState = "FirstOnline"

	MachineStateLastOnline MachineState = "LastOnline"

	MachineStateFirstTransient MachineState = "FirstTransient"

	MachineStateLastTransient MachineState = "LastTransient"
)

type SessionState string

const (
	SessionStateNull SessionState = "Null"

	SessionStateUnlocked SessionState = "Unlocked"

	SessionStateLocked SessionState = "Locked"

	SessionStateSpawning SessionState = "Spawning"

	SessionStateUnlocking SessionState = "Unlocking"
)

type CPUPropertyType string

const (
	CPUPropertyTypeNull CPUPropertyType = "Null"

	CPUPropertyTypePAE CPUPropertyType = "PAE"

	CPUPropertyTypeLongMode CPUPropertyType = "LongMode"

	CPUPropertyTypeTripleFaultReset CPUPropertyType = "TripleFaultReset"
)

type HWVirtExPropertyType string

const (
	HWVirtExPropertyTypeNull HWVirtExPropertyType = "Null"

	HWVirtExPropertyTypeEnabled HWVirtExPropertyType = "Enabled"

	HWVirtExPropertyTypeVPID HWVirtExPropertyType = "VPID"

	HWVirtExPropertyTypeNestedPaging HWVirtExPropertyType = "NestedPaging"

	HWVirtExPropertyTypeUnrestrictedExecution HWVirtExPropertyType = "UnrestrictedExecution"

	HWVirtExPropertyTypeLargePages HWVirtExPropertyType = "LargePages"

	HWVirtExPropertyTypeForce HWVirtExPropertyType = "Force"
)

type ParavirtProvider string

const (
	ParavirtProviderNone ParavirtProvider = "None"

	ParavirtProviderDefault ParavirtProvider = "Default"

	ParavirtProviderLegacy ParavirtProvider = "Legacy"

	ParavirtProviderMinimal ParavirtProvider = "Minimal"

	ParavirtProviderHyperV ParavirtProvider = "HyperV"

	ParavirtProviderKVM ParavirtProvider = "KVM"
)

type FaultToleranceState string

const (
	FaultToleranceStateInactive FaultToleranceState = "Inactive"

	FaultToleranceStateMaster FaultToleranceState = "Master"

	FaultToleranceStateStandby FaultToleranceState = "Standby"
)

type LockType string

const (
	LockTypeNull LockType = "Null"

	LockTypeShared LockType = "Shared"

	LockTypeWrite LockType = "Write"

	LockTypeVM LockType = "VM"
)

type SessionType string

const (
	SessionTypeNull SessionType = "Null"

	SessionTypeWriteLock SessionType = "WriteLock"

	SessionTypeRemote SessionType = "Remote"

	SessionTypeShared SessionType = "Shared"
)

type DeviceType string

const (
	DeviceTypeNull DeviceType = "Null"

	DeviceTypeFloppy DeviceType = "Floppy"

	DeviceTypeDVD DeviceType = "DVD"

	DeviceTypeHardDisk DeviceType = "HardDisk"

	DeviceTypeNetwork DeviceType = "Network"

	DeviceTypeUSB DeviceType = "USB"

	DeviceTypeSharedFolder DeviceType = "SharedFolder"

	DeviceTypeGraphics3D DeviceType = "Graphics3D"
)

type DeviceActivity string

const (
	DeviceActivityNull DeviceActivity = "Null"

	DeviceActivityIdle DeviceActivity = "Idle"

	DeviceActivityReading DeviceActivity = "Reading"

	DeviceActivityWriting DeviceActivity = "Writing"
)

type ClipboardMode string

const (
	ClipboardModeDisabled ClipboardMode = "Disabled"

	ClipboardModeHostToGuest ClipboardMode = "HostToGuest"

	ClipboardModeGuestToHost ClipboardMode = "GuestToHost"

	ClipboardModeBidirectional ClipboardMode = "Bidirectional"
)

type DnDMode string

const (
	DnDModeDisabled DnDMode = "Disabled"

	DnDModeHostToGuest DnDMode = "HostToGuest"

	DnDModeGuestToHost DnDMode = "GuestToHost"

	DnDModeBidirectional DnDMode = "Bidirectional"
)

type Scope string

const (
	ScopeGlobal Scope = "Global"

	ScopeMachine Scope = "Machine"

	ScopeSession Scope = "Session"
)

type BIOSBootMenuMode string

const (
	BIOSBootMenuModeDisabled BIOSBootMenuMode = "Disabled"

	BIOSBootMenuModeMenuOnly BIOSBootMenuMode = "MenuOnly"

	BIOSBootMenuModeMessageAndMenu BIOSBootMenuMode = "MessageAndMenu"
)

type ProcessorFeature string

const (
	ProcessorFeatureHWVirtEx ProcessorFeature = "HWVirtEx"

	ProcessorFeaturePAE ProcessorFeature = "PAE"

	ProcessorFeatureLongMode ProcessorFeature = "LongMode"

	ProcessorFeatureNestedPaging ProcessorFeature = "NestedPaging"
)

type FirmwareType string

const (
	FirmwareTypeBIOS FirmwareType = "BIOS"

	FirmwareTypeEFI FirmwareType = "EFI"

	FirmwareTypeEFI32 FirmwareType = "EFI32"

	FirmwareTypeEFI64 FirmwareType = "EFI64"

	FirmwareTypeEFIDUAL FirmwareType = "EFIDUAL"
)

type PointingHIDType string

const (
	PointingHIDTypeNone PointingHIDType = "None"

	PointingHIDTypePS2Mouse PointingHIDType = "PS2Mouse"

	PointingHIDTypeUSBMouse PointingHIDType = "USBMouse"

	PointingHIDTypeUSBTablet PointingHIDType = "USBTablet"

	PointingHIDTypeComboMouse PointingHIDType = "ComboMouse"

	PointingHIDTypeUSBMultiTouch PointingHIDType = "USBMultiTouch"
)

type KeyboardHIDType string

const (
	KeyboardHIDTypeNone KeyboardHIDType = "None"

	KeyboardHIDTypePS2Keyboard KeyboardHIDType = "PS2Keyboard"

	KeyboardHIDTypeUSBKeyboard KeyboardHIDType = "USBKeyboard"

	KeyboardHIDTypeComboKeyboard KeyboardHIDType = "ComboKeyboard"
)

type BitmapFormat string

const (
	BitmapFormatOpaque BitmapFormat = "Opaque"

	BitmapFormatBGR BitmapFormat = "BGR"

	BitmapFormatBGR0 BitmapFormat = "BGR0"

	BitmapFormatBGRA BitmapFormat = "BGRA"

	BitmapFormatRGBA BitmapFormat = "RGBA"

	BitmapFormatPNG BitmapFormat = "PNG"

	BitmapFormatJPEG BitmapFormat = "JPEG"
)

type DhcpOpt string

const (
	DhcpOptSubnetMask DhcpOpt = "SubnetMask"

	DhcpOptTimeOffset DhcpOpt = "TimeOffset"

	DhcpOptRouter DhcpOpt = "Router"

	DhcpOptTimeServer DhcpOpt = "TimeServer"

	DhcpOptNameServer DhcpOpt = "NameServer"

	DhcpOptDomainNameServer DhcpOpt = "DomainNameServer"

	DhcpOptLogServer DhcpOpt = "LogServer"

	DhcpOptCookie DhcpOpt = "Cookie"

	DhcpOptLPRServer DhcpOpt = "LPRServer"

	DhcpOptImpressServer DhcpOpt = "ImpressServer"

	DhcpOptResourseLocationServer DhcpOpt = "ResourseLocationServer"

	DhcpOptHostName DhcpOpt = "HostName"

	DhcpOptBootFileSize DhcpOpt = "BootFileSize"

	DhcpOptMeritDumpFile DhcpOpt = "MeritDumpFile"

	DhcpOptDomainName DhcpOpt = "DomainName"

	DhcpOptSwapServer DhcpOpt = "SwapServer"

	DhcpOptRootPath DhcpOpt = "RootPath"

	DhcpOptExtensionPath DhcpOpt = "ExtensionPath"

	DhcpOptIPForwardingEnableDisable DhcpOpt = "IPForwardingEnableDisable"

	DhcpOptNonLocalSourceRoutingEnableDisable DhcpOpt = "NonLocalSourceRoutingEnableDisable"

	DhcpOptPolicyFilter DhcpOpt = "PolicyFilter"

	DhcpOptMaximumDatagramReassemblySize DhcpOpt = "MaximumDatagramReassemblySize"

	DhcpOptDefaultIPTime2Live DhcpOpt = "DefaultIPTime2Live"

	DhcpOptPathMTUAgingTimeout DhcpOpt = "PathMTUAgingTimeout"

	DhcpOptIPLayerParametersPerInterface DhcpOpt = "IPLayerParametersPerInterface"

	DhcpOptInterfaceMTU DhcpOpt = "InterfaceMTU"

	DhcpOptAllSubnetsAreLocal DhcpOpt = "AllSubnetsAreLocal"

	DhcpOptBroadcastAddress DhcpOpt = "BroadcastAddress"

	DhcpOptPerformMaskDiscovery DhcpOpt = "PerformMaskDiscovery"

	DhcpOptMaskSupplier DhcpOpt = "MaskSupplier"

	DhcpOptPerformRouteDiscovery DhcpOpt = "PerformRouteDiscovery"

	DhcpOptRouterSolicitationAddress DhcpOpt = "RouterSolicitationAddress"

	DhcpOptStaticRoute DhcpOpt = "StaticRoute"

	DhcpOptTrailerEncapsulation DhcpOpt = "TrailerEncapsulation"

	DhcpOptARPCacheTimeout DhcpOpt = "ARPCacheTimeout"

	DhcpOptEthernetEncapsulation DhcpOpt = "EthernetEncapsulation"

	DhcpOptTCPDefaultTTL DhcpOpt = "TCPDefaultTTL"

	DhcpOptTCPKeepAliveInterval DhcpOpt = "TCPKeepAliveInterval"

	DhcpOptTCPKeepAliveGarbage DhcpOpt = "TCPKeepAliveGarbage"

	DhcpOptNetworkInformationServiceDomain DhcpOpt = "NetworkInformationServiceDomain"

	DhcpOptNetworkInformationServiceServers DhcpOpt = "NetworkInformationServiceServers"

	DhcpOptNetworkTimeProtocolServers DhcpOpt = "NetworkTimeProtocolServers"

	DhcpOptVendorSpecificInformation DhcpOpt = "VendorSpecificInformation"

	DhcpOptOption44 DhcpOpt = "Option44"

	DhcpOptOption45 DhcpOpt = "Option45"

	DhcpOptOption46 DhcpOpt = "Option46"

	DhcpOptOption47 DhcpOpt = "Option47"

	DhcpOptOption48 DhcpOpt = "Option48"

	DhcpOptOption49 DhcpOpt = "Option49"

	DhcpOptIPAddressLeaseTime DhcpOpt = "IPAddressLeaseTime"

	DhcpOptOption64 DhcpOpt = "Option64"

	DhcpOptOption65 DhcpOpt = "Option65"

	DhcpOptTFTPServerName DhcpOpt = "TFTPServerName"

	DhcpOptBootfileName DhcpOpt = "BootfileName"

	DhcpOptOption68 DhcpOpt = "Option68"

	DhcpOptOption69 DhcpOpt = "Option69"

	DhcpOptOption70 DhcpOpt = "Option70"

	DhcpOptOption71 DhcpOpt = "Option71"

	DhcpOptOption72 DhcpOpt = "Option72"

	DhcpOptOption73 DhcpOpt = "Option73"

	DhcpOptOption74 DhcpOpt = "Option74"

	DhcpOptOption75 DhcpOpt = "Option75"

	DhcpOptOption119 DhcpOpt = "Option119"
)

type DhcpOptEncoding string

const (
	DhcpOptEncodingLegacy DhcpOptEncoding = "Legacy"

	DhcpOptEncodingHex DhcpOptEncoding = "Hex"
)

type VFSType string

const (
	VFSTypeFile VFSType = "File"

	VFSTypeCloud VFSType = "Cloud"

	VFSTypeS3 VFSType = "S3"

	VFSTypeWebDav VFSType = "WebDav"
)

type ImportOptions string

const (
	ImportOptionsKeepAllMACs ImportOptions = "KeepAllMACs"

	ImportOptionsKeepNATMACs ImportOptions = "KeepNATMACs"

	ImportOptionsImportToVDI ImportOptions = "ImportToVDI"
)

type ExportOptions string

const (
	ExportOptionsCreateManifest ExportOptions = "CreateManifest"

	ExportOptionsExportDVDImages ExportOptions = "ExportDVDImages"

	ExportOptionsStripAllMACs ExportOptions = "StripAllMACs"

	ExportOptionsStripAllNonNATMACs ExportOptions = "StripAllNonNATMACs"
)

type VirtualSystemDescriptionType string

const (
	VirtualSystemDescriptionTypeIgnore VirtualSystemDescriptionType = "Ignore"

	VirtualSystemDescriptionTypeOS VirtualSystemDescriptionType = "OS"

	VirtualSystemDescriptionTypeName VirtualSystemDescriptionType = "Name"

	VirtualSystemDescriptionTypeProduct VirtualSystemDescriptionType = "Product"

	VirtualSystemDescriptionTypeVendor VirtualSystemDescriptionType = "Vendor"

	VirtualSystemDescriptionTypeVersion VirtualSystemDescriptionType = "Version"

	VirtualSystemDescriptionTypeProductUrl VirtualSystemDescriptionType = "ProductUrl"

	VirtualSystemDescriptionTypeVendorUrl VirtualSystemDescriptionType = "VendorUrl"

	VirtualSystemDescriptionTypeDescription VirtualSystemDescriptionType = "Description"

	VirtualSystemDescriptionTypeLicense VirtualSystemDescriptionType = "License"

	VirtualSystemDescriptionTypeMiscellaneous VirtualSystemDescriptionType = "Miscellaneous"

	VirtualSystemDescriptionTypeCPU VirtualSystemDescriptionType = "CPU"

	VirtualSystemDescriptionTypeMemory VirtualSystemDescriptionType = "VmMemory"

	VirtualSystemDescriptionTypeHardDiskControllerIDE VirtualSystemDescriptionType = "HardDiskControllerIDE"

	VirtualSystemDescriptionTypeHardDiskControllerSATA VirtualSystemDescriptionType = "HardDiskControllerSATA"

	VirtualSystemDescriptionTypeHardDiskControllerSCSI VirtualSystemDescriptionType = "HardDiskControllerSCSI"

	VirtualSystemDescriptionTypeHardDiskControllerSAS VirtualSystemDescriptionType = "HardDiskControllerSAS"

	VirtualSystemDescriptionTypeHardDiskImage VirtualSystemDescriptionType = "HardDiskImage"

	VirtualSystemDescriptionTypeFloppy VirtualSystemDescriptionType = "Floppy"

	VirtualSystemDescriptionTypeCDROM VirtualSystemDescriptionType = "CDROM"

	VirtualSystemDescriptionTypeNetworkAdapter VirtualSystemDescriptionType = "NetworkAdapter"

	VirtualSystemDescriptionTypeUSBController VirtualSystemDescriptionType = "USBController"

	VirtualSystemDescriptionTypeSoundCard VirtualSystemDescriptionType = "SoundCard"

	VirtualSystemDescriptionTypeSettingsFile VirtualSystemDescriptionType = "SettingsFile"
)

type VirtualSystemDescriptionValueType string

const (
	VirtualSystemDescriptionValueTypeReference VirtualSystemDescriptionValueType = "Reference"

	VirtualSystemDescriptionValueTypeOriginal VirtualSystemDescriptionValueType = "Original"

	VirtualSystemDescriptionValueTypeAuto VirtualSystemDescriptionValueType = "Auto"

	VirtualSystemDescriptionValueTypeExtraConfig VirtualSystemDescriptionValueType = "ExtraConfig"
)

type GraphicsControllerType string

const (
	GraphicsControllerTypeNull GraphicsControllerType = "Null"

	GraphicsControllerTypeVBoxVGA GraphicsControllerType = "VBoxVGA"

	GraphicsControllerTypeVMSVGA GraphicsControllerType = "VMSVGA"
)

type CleanupMode string

const (
	CleanupModeUnregisterOnly CleanupMode = "UnregisterOnly"

	CleanupModeDetachAllReturnNone CleanupMode = "DetachAllReturnNone"

	CleanupModeDetachAllReturnHardDisksOnly CleanupMode = "DetachAllReturnHardDisksOnly"

	CleanupModeFull CleanupMode = "Full"
)

type CloneMode string

const (
	CloneModeMachineState CloneMode = "MachineState"

	CloneModeMachineAndChildStates CloneMode = "MachineAndChildStates"

	CloneModeAllStates CloneMode = "AllStates"
)

type CloneOptions string

const (
	CloneOptionsLink CloneOptions = "Link"

	CloneOptionsKeepAllMACs CloneOptions = "KeepAllMACs"

	CloneOptionsKeepNATMACs CloneOptions = "KeepNATMACs"

	CloneOptionsKeepDiskNames CloneOptions = "KeepDiskNames"
)

type AutostopType string

const (
	AutostopTypeDisabled AutostopType = "Disabled"

	AutostopTypeSaveState AutostopType = "SaveState"

	AutostopTypePowerOff AutostopType = "PowerOff"

	AutostopTypeAcpiShutdown AutostopType = "AcpiShutdown"
)

type HostNetworkInterfaceMediumType string

const (
	HostNetworkInterfaceMediumTypeUnknown HostNetworkInterfaceMediumType = "Unknown"

	HostNetworkInterfaceMediumTypeEthernet HostNetworkInterfaceMediumType = "Ethernet"

	HostNetworkInterfaceMediumTypePPP HostNetworkInterfaceMediumType = "PPP"

	HostNetworkInterfaceMediumTypeSLIP HostNetworkInterfaceMediumType = "SLIP"
)

type HostNetworkInterfaceStatus string

const (
	HostNetworkInterfaceStatusUnknown HostNetworkInterfaceStatus = "Unknown"

	HostNetworkInterfaceStatusUp HostNetworkInterfaceStatus = "Up"

	HostNetworkInterfaceStatusDown HostNetworkInterfaceStatus = "Down"
)

type HostNetworkInterfaceType string

const (
	HostNetworkInterfaceTypeBridged HostNetworkInterfaceType = "Bridged"

	HostNetworkInterfaceTypeHostOnly HostNetworkInterfaceType = "HostOnly"
)

type AdditionsFacilityType string

const (
	AdditionsFacilityTypeNone AdditionsFacilityType = "None"

	AdditionsFacilityTypeVBoxGuestDriver AdditionsFacilityType = "VBoxGuestDriver"

	AdditionsFacilityTypeAutoLogon AdditionsFacilityType = "AutoLogon"

	AdditionsFacilityTypeVBoxService AdditionsFacilityType = "VBoxService"

	AdditionsFacilityTypeVBoxTrayClient AdditionsFacilityType = "VBoxTrayClient"

	AdditionsFacilityTypeSeamless AdditionsFacilityType = "Seamless"

	AdditionsFacilityTypeGraphics AdditionsFacilityType = "Graphics"

	AdditionsFacilityTypeAll AdditionsFacilityType = "All"
)

type AdditionsFacilityClass string

const (
	AdditionsFacilityClassNone AdditionsFacilityClass = "None"

	AdditionsFacilityClassDriver AdditionsFacilityClass = "Driver"

	AdditionsFacilityClassService AdditionsFacilityClass = "Service"

	AdditionsFacilityClassProgram AdditionsFacilityClass = "Program"

	AdditionsFacilityClassFeature AdditionsFacilityClass = "Feature"

	AdditionsFacilityClassThirdParty AdditionsFacilityClass = "ThirdParty"

	AdditionsFacilityClassAll AdditionsFacilityClass = "All"
)

type AdditionsFacilityStatus string

const (
	AdditionsFacilityStatusInactive AdditionsFacilityStatus = "Inactive"

	AdditionsFacilityStatusPaused AdditionsFacilityStatus = "Paused"

	AdditionsFacilityStatusPreInit AdditionsFacilityStatus = "PreInit"

	AdditionsFacilityStatusInit AdditionsFacilityStatus = "InitModels"

	AdditionsFacilityStatusActive AdditionsFacilityStatus = "Active"

	AdditionsFacilityStatusTerminating AdditionsFacilityStatus = "Terminating"

	AdditionsFacilityStatusTerminated AdditionsFacilityStatus = "Terminated"

	AdditionsFacilityStatusFailed AdditionsFacilityStatus = "Failed"

	AdditionsFacilityStatusUnknown AdditionsFacilityStatus = "Unknown"
)

type AdditionsRunLevelType string

const (
	AdditionsRunLevelTypeNone AdditionsRunLevelType = "None"

	AdditionsRunLevelTypeSystem AdditionsRunLevelType = "System"

	AdditionsRunLevelTypeUserland AdditionsRunLevelType = "Userland"

	AdditionsRunLevelTypeDesktop AdditionsRunLevelType = "Desktop"
)

type AdditionsUpdateFlag string

const (
	AdditionsUpdateFlagNone AdditionsUpdateFlag = "None"

	AdditionsUpdateFlagWaitForUpdateStartOnly AdditionsUpdateFlag = "WaitForUpdateStartOnly"
)

type GuestSessionStatus string

const (
	GuestSessionStatusUndefined GuestSessionStatus = "Undefined"

	GuestSessionStatusStarting GuestSessionStatus = "Starting"

	GuestSessionStatusStarted GuestSessionStatus = "Started"

	GuestSessionStatusTerminating GuestSessionStatus = "Terminating"

	GuestSessionStatusTerminated GuestSessionStatus = "Terminated"

	GuestSessionStatusTimedOutKilled GuestSessionStatus = "TimedOutKilled"

	GuestSessionStatusTimedOutAbnormally GuestSessionStatus = "TimedOutAbnormally"

	GuestSessionStatusDown GuestSessionStatus = "Down"

	GuestSessionStatusError GuestSessionStatus = "Error"
)

type GuestSessionWaitForFlag string

const (
	GuestSessionWaitForFlagNone GuestSessionWaitForFlag = "None"

	GuestSessionWaitForFlagStart GuestSessionWaitForFlag = "Start"

	GuestSessionWaitForFlagTerminate GuestSessionWaitForFlag = "TerminateTasks"

	GuestSessionWaitForFlagStatus GuestSessionWaitForFlag = "Status"
)

type GuestSessionWaitResult string

const (
	GuestSessionWaitResultNone GuestSessionWaitResult = "None"

	GuestSessionWaitResultStart GuestSessionWaitResult = "Start"

	GuestSessionWaitResultTerminate GuestSessionWaitResult = "TerminateTasks"

	GuestSessionWaitResultStatus GuestSessionWaitResult = "Status"

	GuestSessionWaitResultError GuestSessionWaitResult = "Error"

	GuestSessionWaitResultTimeout GuestSessionWaitResult = "Timeout"

	GuestSessionWaitResultWaitFlagNotSupported GuestSessionWaitResult = "WaitFlagNotSupported"
)

type GuestUserState string

const (
	GuestUserStateUnknown GuestUserState = "Unknown"

	GuestUserStateLoggedIn GuestUserState = "LoggedIn"

	GuestUserStateLoggedOut GuestUserState = "LoggedOut"

	GuestUserStateLocked GuestUserState = "Locked"

	GuestUserStateUnlocked GuestUserState = "Unlocked"

	GuestUserStateDisabled GuestUserState = "Disabled"

	GuestUserStateIdle GuestUserState = "Idle"

	GuestUserStateInUse GuestUserState = "InUse"

	GuestUserStateCreated GuestUserState = "Created"

	GuestUserStateDeleted GuestUserState = "Deleted"

	GuestUserStateSessionChanged GuestUserState = "SessionChanged"

	GuestUserStateCredentialsChanged GuestUserState = "CredentialsChanged"

	GuestUserStateRoleChanged GuestUserState = "RoleChanged"

	GuestUserStateGroupAdded GuestUserState = "GroupAdded"

	GuestUserStateGroupRemoved GuestUserState = "GroupRemoved"

	GuestUserStateElevated GuestUserState = "Elevated"
)

type FileSeekOrigin string

const (
	FileSeekOriginBegin FileSeekOrigin = "Begin"

	FileSeekOriginCurrent FileSeekOrigin = "Current"

	FileSeekOriginEnd FileSeekOrigin = "End"
)

type ProcessInputFlag string

const (
	ProcessInputFlagNone ProcessInputFlag = "None"

	ProcessInputFlagEndOfFile ProcessInputFlag = "EndOfFile"
)

type ProcessOutputFlag string

const (
	ProcessOutputFlagNone ProcessOutputFlag = "None"

	ProcessOutputFlagStdErr ProcessOutputFlag = "StdErr"
)

type ProcessWaitForFlag string

const (
	ProcessWaitForFlagNone ProcessWaitForFlag = "None"

	ProcessWaitForFlagStart ProcessWaitForFlag = "Start"

	ProcessWaitForFlagTerminate ProcessWaitForFlag = "TerminateTasks"

	ProcessWaitForFlagStdIn ProcessWaitForFlag = "StdIn"

	ProcessWaitForFlagStdOut ProcessWaitForFlag = "StdOut"

	ProcessWaitForFlagStdErr ProcessWaitForFlag = "StdErr"
)

type ProcessWaitResult string

const (
	ProcessWaitResultNone ProcessWaitResult = "None"

	ProcessWaitResultStart ProcessWaitResult = "Start"

	ProcessWaitResultTerminate ProcessWaitResult = "TerminateTasks"

	ProcessWaitResultStatus ProcessWaitResult = "Status"

	ProcessWaitResultError ProcessWaitResult = "Error"

	ProcessWaitResultTimeout ProcessWaitResult = "Timeout"

	ProcessWaitResultStdIn ProcessWaitResult = "StdIn"

	ProcessWaitResultStdOut ProcessWaitResult = "StdOut"

	ProcessWaitResultStdErr ProcessWaitResult = "StdErr"

	ProcessWaitResultWaitFlagNotSupported ProcessWaitResult = "WaitFlagNotSupported"
)

type FileCopyFlag string

const (
	FileCopyFlagNone FileCopyFlag = "None"

	FileCopyFlagNoReplace FileCopyFlag = "NoReplace"

	FileCopyFlagFollowLinks FileCopyFlag = "FollowLinks"

	FileCopyFlagUpdate FileCopyFlag = "Update"
)

type FsObjMoveFlags string

const (
	FsObjMoveFlagsNone FsObjMoveFlags = "None"

	FsObjMoveFlagsReplace FsObjMoveFlags = "Replace"

	FsObjMoveFlagsFollowLinks FsObjMoveFlags = "FollowLinks"

	FsObjMoveFlagsAllowDirectoryMoves FsObjMoveFlags = "AllowDirectoryMoves"
)

type DirectoryCreateFlag string

const (
	DirectoryCreateFlagNone DirectoryCreateFlag = "None"

	DirectoryCreateFlagParents DirectoryCreateFlag = "Parents"
)

type DirectoryCopyFlags string

const (
	DirectoryCopyFlagsNone DirectoryCopyFlags = "None"

	DirectoryCopyFlagsCopyIntoExisting DirectoryCopyFlags = "CopyIntoExisting"
)

type DirectoryRemoveRecFlag string

const (
	DirectoryRemoveRecFlagNone DirectoryRemoveRecFlag = "None"

	DirectoryRemoveRecFlagContentAndDir DirectoryRemoveRecFlag = "ContentAndDir"

	DirectoryRemoveRecFlagContentOnly DirectoryRemoveRecFlag = "ContentOnly"
)

type FsObjRenameFlag string

const (
	FsObjRenameFlagNoReplace FsObjRenameFlag = "NoReplace"

	FsObjRenameFlagReplace FsObjRenameFlag = "Replace"
)

type ProcessCreateFlag string

const (
	ProcessCreateFlagNone ProcessCreateFlag = "None"

	ProcessCreateFlagWaitForProcessStartOnly ProcessCreateFlag = "WaitForProcessStartOnly"

	ProcessCreateFlagIgnoreOrphanedProcesses ProcessCreateFlag = "IgnoreOrphanedProcesses"

	ProcessCreateFlagHidden ProcessCreateFlag = "Hidden"

	ProcessCreateFlagNoProfile ProcessCreateFlag = "NoProfile"

	ProcessCreateFlagWaitForStdOut ProcessCreateFlag = "WaitForStdOut"

	ProcessCreateFlagWaitForStdErr ProcessCreateFlag = "WaitForStdErr"

	ProcessCreateFlagExpandArguments ProcessCreateFlag = "ExpandArguments"

	ProcessCreateFlagUnquotedArguments ProcessCreateFlag = "UnquotedArguments"
)

type ProcessPriority string

const (
	ProcessPriorityInvalid ProcessPriority = "Invalid"

	ProcessPriorityDefault ProcessPriority = "Default"
)

type SymlinkType string

const (
	SymlinkTypeUnknown SymlinkType = "Unknown"

	SymlinkTypeDirectory SymlinkType = "Directory"

	SymlinkTypeFile SymlinkType = "File"
)

type SymlinkReadFlag string

const (
	SymlinkReadFlagNone SymlinkReadFlag = "None"

	SymlinkReadFlagNoSymlinks SymlinkReadFlag = "NoSymlinks"
)

type ProcessStatus string

const (
	ProcessStatusUndefined ProcessStatus = "Undefined"

	ProcessStatusStarting ProcessStatus = "Starting"

	ProcessStatusStarted ProcessStatus = "Started"

	ProcessStatusPaused ProcessStatus = "Paused"

	ProcessStatusTerminating ProcessStatus = "Terminating"

	ProcessStatusTerminatedNormally ProcessStatus = "TerminatedNormally"

	ProcessStatusTerminatedSignal ProcessStatus = "TerminatedSignal"

	ProcessStatusTerminatedAbnormally ProcessStatus = "TerminatedAbnormally"

	ProcessStatusTimedOutKilled ProcessStatus = "TimedOutKilled"

	ProcessStatusTimedOutAbnormally ProcessStatus = "TimedOutAbnormally"

	ProcessStatusDown ProcessStatus = "Down"

	ProcessStatusError ProcessStatus = "Error"
)

type ProcessInputStatus string

const (
	ProcessInputStatusUndefined ProcessInputStatus = "Undefined"

	ProcessInputStatusBroken ProcessInputStatus = "Broken"

	ProcessInputStatusAvailable ProcessInputStatus = "Available"

	ProcessInputStatusWritten ProcessInputStatus = "Written"

	ProcessInputStatusOverflow ProcessInputStatus = "Overflow"
)

type PathStyle string

const (
	PathStyleDOS PathStyle = "DOS"

	PathStyleUNIX PathStyle = "UNIX"

	PathStyleUnknown PathStyle = "Unknown"
)

type FileAccessMode string

const (
	FileAccessModeReadOnly FileAccessMode = "ReadOnly"

	FileAccessModeWriteOnly FileAccessMode = "WriteOnly"

	FileAccessModeReadWrite FileAccessMode = "ReadWrite"

	FileAccessModeAppendOnly FileAccessMode = "AppendOnly"

	FileAccessModeAppendRead FileAccessMode = "AppendRead"
)

type FileOpenAction string

const (
	FileOpenActionOpenExisting FileOpenAction = "OpenExisting"

	FileOpenActionOpenOrCreate FileOpenAction = "OpenOrCreate"

	FileOpenActionCreateNew FileOpenAction = "CreateNew"

	FileOpenActionCreateOrReplace FileOpenAction = "CreateOrReplace"

	FileOpenActionOpenExistingTruncated FileOpenAction = "OpenExistingTruncated"

	FileOpenActionAppendOrCreate FileOpenAction = "AppendOrCreate"
)

type FileSharingMode string

const (
	FileSharingModeRead FileSharingMode = "Read"

	FileSharingModeWrite FileSharingMode = "Write"

	FileSharingModeReadWrite FileSharingMode = "ReadWrite"

	FileSharingModeDelete FileSharingMode = "Delete"

	FileSharingModeReadDelete FileSharingMode = "ReadDelete"

	FileSharingModeWriteDelete FileSharingMode = "WriteDelete"

	FileSharingModeAll FileSharingMode = "All"
)

type FileOpenExFlags string

const (
	FileOpenExFlagsNone FileOpenExFlags = "None"
)

type FileStatus string

const (
	FileStatusUndefined FileStatus = "Undefined"

	FileStatusOpening FileStatus = "Opening"

	FileStatusOpen FileStatus = "Open"

	FileStatusClosing FileStatus = "Closing"

	FileStatusClosed FileStatus = "Closed"

	FileStatusDown FileStatus = "Down"

	FileStatusError FileStatus = "Error"
)

type FsObjType string

const (
	FsObjTypeUnknown FsObjType = "Unknown"

	FsObjTypeFifo FsObjType = "Fifo"

	FsObjTypeDevChar FsObjType = "DevChar"

	FsObjTypeDirectory FsObjType = "Directory"

	FsObjTypeDevBlock FsObjType = "DevBlock"

	FsObjTypeFile FsObjType = "File"

	FsObjTypeSymlink FsObjType = "Symlink"

	FsObjTypeSocket FsObjType = "Socket"

	FsObjTypeWhiteOut FsObjType = "WhiteOut"
)

type DnDAction string

const (
	DnDActionIgnore DnDAction = "Ignore"

	DnDActionCopy DnDAction = "Copy"

	DnDActionMove DnDAction = "Move"

	DnDActionLink DnDAction = "Link"
)

type DirectoryOpenFlag string

const (
	DirectoryOpenFlagNone DirectoryOpenFlag = "None"

	DirectoryOpenFlagNoSymlinks DirectoryOpenFlag = "NoSymlinks"
)

type MediumState string

const (
	MediumStateNotCreated MediumState = "NotCreated"

	MediumStateCreated MediumState = "Created"

	MediumStateLockedRead MediumState = "LockedRead"

	MediumStateLockedWrite MediumState = "LockedWrite"

	MediumStateInaccessible MediumState = "Inaccessible"

	MediumStateCreating MediumState = "Creating"

	MediumStateDeleting MediumState = "Deleting"
)

type MediumType string

const (
	MediumTypeNormal MediumType = "Normal"

	MediumTypeImmutable MediumType = "Immutable"

	MediumTypeWritethrough MediumType = "Writethrough"

	MediumTypeShareable MediumType = "Shareable"

	MediumTypeReadonly MediumType = "Readonly"

	MediumTypeMultiAttach MediumType = "MultiAttach"
)

type MediumVariant string

const (
	MediumVariantStandard MediumVariant = "Standard"

	MediumVariantVmdkSplit2G MediumVariant = "VmdkSplit2G"

	MediumVariantVmdkRawDisk MediumVariant = "VmdkRawDisk"

	MediumVariantVmdkStreamOptimized MediumVariant = "VmdkStreamOptimized"

	MediumVariantVmdkESX MediumVariant = "VmdkESX"

	MediumVariantVdiZeroExpand MediumVariant = "VdiZeroExpand"

	MediumVariantFixed MediumVariant = "Fixed"

	MediumVariantDiff MediumVariant = "Diff"

	MediumVariantNoCreateDir MediumVariant = "NoCreateDir"
)

type DataType string

const (
	DataTypeInt32 DataType = "Int32"

	DataTypeInt8 DataType = "Int8"

	DataTypeString DataType = "String"
)

type DataFlags string

const (
	DataFlagsNone DataFlags = "None"

	DataFlagsMandatory DataFlags = "Mandatory"

	DataFlagsExpert DataFlags = "Expert"

	DataFlagsArray DataFlags = "Array"

	DataFlagsFlagMask DataFlags = "FlagMask"
)

type MediumFormatCapabilities string

const (
	MediumFormatCapabilitiesUuid MediumFormatCapabilities = "Uuid"

	MediumFormatCapabilitiesCreateFixed MediumFormatCapabilities = "CreateFixed"

	MediumFormatCapabilitiesCreateDynamic MediumFormatCapabilities = "CreateDynamic"

	MediumFormatCapabilitiesCreateSplit2G MediumFormatCapabilities = "CreateSplit2G"

	MediumFormatCapabilitiesDifferencing MediumFormatCapabilities = "Differencing"

	MediumFormatCapabilitiesAsynchronous MediumFormatCapabilities = "Asynchronous"

	MediumFormatCapabilitiesFile MediumFormatCapabilities = "File"

	MediumFormatCapabilitiesProperties MediumFormatCapabilities = "Properties"

	MediumFormatCapabilitiesTcpNetworking MediumFormatCapabilities = "TcpNetworking"

	MediumFormatCapabilitiesVFS MediumFormatCapabilities = "VFS"

	MediumFormatCapabilitiesCapabilityMask MediumFormatCapabilities = "CapabilityMask"
)

type KeyboardLED string

const (
	KeyboardLEDNumLock KeyboardLED = "NumLock"

	KeyboardLEDCapsLock KeyboardLED = "CapsLock"

	KeyboardLEDScrollLock KeyboardLED = "ScrollLock"
)

type MouseButtonState string

const (
	MouseButtonStateLeftButton MouseButtonState = "LeftButton"

	MouseButtonStateRightButton MouseButtonState = "RightButton"

	MouseButtonStateMiddleButton MouseButtonState = "MiddleButton"

	MouseButtonStateWheelUp MouseButtonState = "WheelUp"

	MouseButtonStateWheelDown MouseButtonState = "WheelDown"

	MouseButtonStateXButton1 MouseButtonState = "XButton1"

	MouseButtonStateXButton2 MouseButtonState = "XButton2"

	MouseButtonStateMouseStateMask MouseButtonState = "MouseStateMask"
)

type TouchContactState string

const (
	TouchContactStateNone TouchContactState = "None"

	TouchContactStateInContact TouchContactState = "InContact"

	TouchContactStateInRange TouchContactState = "InRange"

	TouchContactStateContactStateMask TouchContactState = "ContactStateMask"
)

type FramebufferCapabilities string

const (
	FramebufferCapabilitiesUpdateImage FramebufferCapabilities = "UpdateImage"

	FramebufferCapabilitiesVHWA FramebufferCapabilities = "VHWA"

	FramebufferCapabilitiesVisibleRegion FramebufferCapabilities = "VisibleRegion"
)

type GuestMonitorStatus string

const (
	GuestMonitorStatusDisabled GuestMonitorStatus = "Disabled"

	GuestMonitorStatusEnabled GuestMonitorStatus = "Enabled"
)

type NetworkAttachmentType string

const (
	NetworkAttachmentTypeNull NetworkAttachmentType = "Null"

	NetworkAttachmentTypeNAT NetworkAttachmentType = "NAT"

	NetworkAttachmentTypeBridged NetworkAttachmentType = "Bridged"

	NetworkAttachmentTypeInternal NetworkAttachmentType = "Internal"

	NetworkAttachmentTypeHostOnly NetworkAttachmentType = "HostOnly"

	NetworkAttachmentTypeGeneric NetworkAttachmentType = "Generic"

	NetworkAttachmentTypeNATNetwork NetworkAttachmentType = "NATNetwork"
)

type NetworkAdapterType string

const (
	NetworkAdapterTypeNull NetworkAdapterType = "Null"

	NetworkAdapterTypeAm79C970A NetworkAdapterType = "Am79C970A"

	NetworkAdapterTypeAm79C973 NetworkAdapterType = "Am79C973"

	NetworkAdapterTypeI82540EM NetworkAdapterType = "I82540EM"

	NetworkAdapterTypeI82543GC NetworkAdapterType = "I82543GC"

	NetworkAdapterTypeI82545EM NetworkAdapterType = "I82545EM"

	NetworkAdapterTypeVirtio NetworkAdapterType = "Virtio"
)

type NetworkAdapterPromiscModePolicy string

const (
	NetworkAdapterPromiscModePolicyDeny NetworkAdapterPromiscModePolicy = "Deny"

	NetworkAdapterPromiscModePolicyAllowNetwork NetworkAdapterPromiscModePolicy = "AllowNetwork"

	NetworkAdapterPromiscModePolicyAllowAll NetworkAdapterPromiscModePolicy = "AllowAll"
)

type PortMode string

const (
	PortModeDisconnected PortMode = "Disconnected"

	PortModeHostPipe PortMode = "HostPipe"

	PortModeHostDevice PortMode = "HostDevice"

	PortModeRawFile PortMode = "RawFile"

	PortModeTCP PortMode = "TCP"
)

type USBControllerType string

const (
	USBControllerTypeNull USBControllerType = "Null"

	USBControllerTypeOHCI USBControllerType = "OHCI"

	USBControllerTypeEHCI USBControllerType = "EHCI"

	USBControllerTypeXHCI USBControllerType = "XHCI"

	USBControllerTypeLast USBControllerType = "Last"
)

type USBConnectionSpeed string

const (
	USBConnectionSpeedNull USBConnectionSpeed = "Null"

	USBConnectionSpeedLow USBConnectionSpeed = "Low"

	USBConnectionSpeedFull USBConnectionSpeed = "Full"

	USBConnectionSpeedHigh USBConnectionSpeed = "High"

	USBConnectionSpeedSuper USBConnectionSpeed = "Super"

	USBConnectionSpeedSuperPlus USBConnectionSpeed = "SuperPlus"
)

type USBDeviceState string

const (
	USBDeviceStateNotSupported USBDeviceState = "NotSupported"

	USBDeviceStateUnavailable USBDeviceState = "Unavailable"

	USBDeviceStateBusy USBDeviceState = "Busy"

	USBDeviceStateAvailable USBDeviceState = "Available"

	USBDeviceStateHeld USBDeviceState = "Held"

	USBDeviceStateCaptured USBDeviceState = "Captured"
)

type USBDeviceFilterAction string

const (
	USBDeviceFilterActionNull USBDeviceFilterAction = "Null"

	USBDeviceFilterActionIgnore USBDeviceFilterAction = "Ignore"

	USBDeviceFilterActionHold USBDeviceFilterAction = "Hold"
)

type AudioDriverType string

const (
	AudioDriverTypeNull AudioDriverType = "Null"

	AudioDriverTypeWinMM AudioDriverType = "WinMM"

	AudioDriverTypeOSS AudioDriverType = "OSS"

	AudioDriverTypeALSA AudioDriverType = "ALSA"

	AudioDriverTypeDirectSound AudioDriverType = "DirectSound"

	AudioDriverTypeCoreAudio AudioDriverType = "CoreAudio"

	AudioDriverTypeMMPM AudioDriverType = "MMPM"

	AudioDriverTypePulse AudioDriverType = "Pulse"

	AudioDriverTypeSolAudio AudioDriverType = "SolAudio"
)

type AudioControllerType string

const (
	AudioControllerTypeAC97 AudioControllerType = "AC97"

	AudioControllerTypeSB16 AudioControllerType = "SB16"

	AudioControllerTypeHDA AudioControllerType = "HDA"
)

type AudioCodecType string

const (
	AudioCodecTypeNull AudioCodecType = "Null"

	AudioCodecTypeSB16 AudioCodecType = "SB16"

	AudioCodecTypeSTAC9700 AudioCodecType = "STAC9700"

	AudioCodecTypeAD1980 AudioCodecType = "AD1980"

	AudioCodecTypeSTAC9221 AudioCodecType = "STAC9221"
)

type AuthType string

const (
	AuthTypeNull AuthType = "Null"

	AuthTypeExternal AuthType = "External"

	AuthTypeGuest AuthType = "Guest"
)

type Reason string

const (
	ReasonUnspecified Reason = "Unspecified"

	ReasonHostSuspend Reason = "HostSuspend"

	ReasonHostResume Reason = "HostResume"

	ReasonHostBatteryLow Reason = "HostBatteryLow"

	ReasonSnapshot Reason = "Snapshot"
)

type StorageBus string

const (
	StorageBusNull StorageBus = "Null"

	StorageBusIDE StorageBus = "IDE"

	StorageBusSATA StorageBus = "SATA"

	StorageBusSCSI StorageBus = "SCSI"

	StorageBusFloppy StorageBus = "Floppy"

	StorageBusSAS StorageBus = "SAS"

	StorageBusUSB StorageBus = "USB"
)

type StorageControllerType string

const (
	StorageControllerTypeNull StorageControllerType = "Null"

	StorageControllerTypeLsiLogic StorageControllerType = "LsiLogic"

	StorageControllerTypeBusLogic StorageControllerType = "BusLogic"

	StorageControllerTypeIntelAhci StorageControllerType = "IntelAhci"

	StorageControllerTypePIIX3 StorageControllerType = "PIIX3"

	StorageControllerTypePIIX4 StorageControllerType = "PIIX4"

	StorageControllerTypeICH6 StorageControllerType = "ICH6"

	StorageControllerTypeI82078 StorageControllerType = "I82078"

	StorageControllerTypeLsiLogicSas StorageControllerType = "LsiLogicSas"

	StorageControllerTypeUSB StorageControllerType = "USB"
)

type ChipsetType string

const (
	ChipsetTypeNull ChipsetType = "Null"

	ChipsetTypePIIX3 ChipsetType = "PIIX3"

	ChipsetTypeICH9 ChipsetType = "ICH9"
)

type NATAliasMode string

const (
	NATAliasModeAliasLog NATAliasMode = "AliasLog"

	NATAliasModeAliasProxyOnly NATAliasMode = "AliasProxyOnly"

	NATAliasModeAliasUseSamePorts NATAliasMode = "AliasUseSamePorts"
)

type NATProtocol string

const (
	NATProtocolUDP NATProtocol = "UDP"

	NATProtocolTCP NATProtocol = "TCP"
)

type BandwidthGroupType string

const (
	BandwidthGroupTypeNull BandwidthGroupType = "Null"

	BandwidthGroupTypeDisk BandwidthGroupType = "Disk"

	BandwidthGroupTypeNetwork BandwidthGroupType = "Network"
)

type VBoxEventType string

const (
	VBoxEventTypeInvalid VBoxEventType = "Invalid"

	VBoxEventTypeAny VBoxEventType = "Any"

	VBoxEventTypeVetoable VBoxEventType = "Vetoable"

	VBoxEventTypeMachineEvent VBoxEventType = "MachineEvent"

	VBoxEventTypeSnapshotEvent VBoxEventType = "SnapshotEvent"

	VBoxEventTypeInputEvent VBoxEventType = "InputEvent"

	VBoxEventTypeLastWildcard VBoxEventType = "LastWildcard"

	VBoxEventTypeOnMachineStateChanged VBoxEventType = "OnMachineStateChanged"

	VBoxEventTypeOnMachineDataChanged VBoxEventType = "OnMachineDataChanged"

	VBoxEventTypeOnExtraDataChanged VBoxEventType = "OnExtraDataChanged"

	VBoxEventTypeOnExtraDataCanChange VBoxEventType = "OnExtraDataCanChange"

	VBoxEventTypeOnMediumRegistered VBoxEventType = "OnMediumRegistered"

	VBoxEventTypeOnMachineRegistered VBoxEventType = "OnMachineRegistered"

	VBoxEventTypeOnSessionStateChanged VBoxEventType = "OnSessionStateChanged"

	VBoxEventTypeOnSnapshotTaken VBoxEventType = "OnSnapshotTaken"

	VBoxEventTypeOnSnapshotDeleted VBoxEventType = "OnSnapshotDeleted"

	VBoxEventTypeOnSnapshotChanged VBoxEventType = "OnSnapshotChanged"

	VBoxEventTypeOnGuestPropertyChanged VBoxEventType = "OnGuestPropertyChanged"

	VBoxEventTypeOnMousePointerShapeChanged VBoxEventType = "OnMousePointerShapeChanged"

	VBoxEventTypeOnMouseCapabilityChanged VBoxEventType = "OnMouseCapabilityChanged"

	VBoxEventTypeOnKeyboardLedsChanged VBoxEventType = "OnKeyboardLedsChanged"

	VBoxEventTypeOnStateChanged VBoxEventType = "OnStateChanged"

	VBoxEventTypeOnAdditionsStateChanged VBoxEventType = "OnAdditionsStateChanged"

	VBoxEventTypeOnNetworkAdapterChanged VBoxEventType = "OnNetworkAdapterChanged"

	VBoxEventTypeOnSerialPortChanged VBoxEventType = "OnSerialPortChanged"

	VBoxEventTypeOnParallelPortChanged VBoxEventType = "OnParallelPortChanged"

	VBoxEventTypeOnStorageControllerChanged VBoxEventType = "OnStorageControllerChanged"

	VBoxEventTypeOnMediumChanged VBoxEventType = "OnMediumChanged"

	VBoxEventTypeOnVRDEServerChanged VBoxEventType = "OnVRDEServerChanged"

	VBoxEventTypeOnUSBControllerChanged VBoxEventType = "OnUSBControllerChanged"

	VBoxEventTypeOnUSBDeviceStateChanged VBoxEventType = "OnUSBDeviceStateChanged"

	VBoxEventTypeOnSharedFolderChanged VBoxEventType = "OnSharedFolderChanged"

	VBoxEventTypeOnRuntimeError VBoxEventType = "OnRuntimeError"

	VBoxEventTypeOnCanShowWindow VBoxEventType = "OnCanShowWindow"

	VBoxEventTypeOnShowWindow VBoxEventType = "OnShowWindow"

	VBoxEventTypeOnCPUChanged VBoxEventType = "OnCPUChanged"

	VBoxEventTypeOnVRDEServerInfoChanged VBoxEventType = "OnVRDEServerInfoChanged"

	VBoxEventTypeOnEventSourceChanged VBoxEventType = "OnEventSourceChanged"

	VBoxEventTypeOnCPUExecutionCapChanged VBoxEventType = "OnCPUExecutionCapChanged"

	VBoxEventTypeOnGuestKeyboard VBoxEventType = "OnGuestKeyboard"

	VBoxEventTypeOnGuestMouse VBoxEventType = "OnGuestMouse"

	VBoxEventTypeOnNATRedirect VBoxEventType = "OnNATRedirect"

	VBoxEventTypeOnHostPCIDevicePlug VBoxEventType = "OnHostPCIDevicePlug"

	VBoxEventTypeOnVBoxSVCAvailabilityChanged VBoxEventType = "OnVBoxSVCAvailabilityChanged"

	VBoxEventTypeOnBandwidthGroupChanged VBoxEventType = "OnBandwidthGroupChanged"

	VBoxEventTypeOnGuestMonitorChanged VBoxEventType = "OnGuestMonitorChanged"

	VBoxEventTypeOnStorageDeviceChanged VBoxEventType = "OnStorageDeviceChanged"

	VBoxEventTypeOnClipboardModeChanged VBoxEventType = "OnClipboardModeChanged"

	VBoxEventTypeOnDnDModeChanged VBoxEventType = "OnDnDModeChanged"

	VBoxEventTypeOnNATNetworkChanged VBoxEventType = "OnNATNetworkChanged"

	VBoxEventTypeOnNATNetworkStartStop VBoxEventType = "OnNATNetworkStartStop"

	VBoxEventTypeOnNATNetworkAlter VBoxEventType = "OnNATNetworkAlter"

	VBoxEventTypeOnNATNetworkCreationDeletion VBoxEventType = "OnNATNetworkCreationDeletion"

	VBoxEventTypeOnNATNetworkSetting VBoxEventType = "OnNATNetworkSetting"

	VBoxEventTypeOnNATNetworkPortForward VBoxEventType = "OnNATNetworkPortForward"

	VBoxEventTypeOnGuestSessionStateChanged VBoxEventType = "OnGuestSessionStateChanged"

	VBoxEventTypeOnGuestSessionRegistered VBoxEventType = "OnGuestSessionRegistered"

	VBoxEventTypeOnGuestProcessRegistered VBoxEventType = "OnGuestProcessRegistered"

	VBoxEventTypeOnGuestProcessStateChanged VBoxEventType = "OnGuestProcessStateChanged"

	VBoxEventTypeOnGuestProcessInputNotify VBoxEventType = "OnGuestProcessInputNotify"

	VBoxEventTypeOnGuestProcessOutput VBoxEventType = "OnGuestProcessOutput"

	VBoxEventTypeOnGuestFileRegistered VBoxEventType = "OnGuestFileRegistered"

	VBoxEventTypeOnGuestFileStateChanged VBoxEventType = "OnGuestFileStateChanged"

	VBoxEventTypeOnGuestFileOffsetChanged VBoxEventType = "OnGuestFileOffsetChanged"

	VBoxEventTypeOnGuestFileRead VBoxEventType = "OnGuestFileRead"

	VBoxEventTypeOnGuestFileWrite VBoxEventType = "OnGuestFileWrite"

	VBoxEventTypeOnVideoCaptureChanged VBoxEventType = "OnVideoCaptureChanged"

	VBoxEventTypeOnGuestUserStateChanged VBoxEventType = "OnGuestUserStateChanged"

	VBoxEventTypeOnGuestMultiTouch VBoxEventType = "OnGuestMultiTouch"

	VBoxEventTypeOnHostNameResolutionConfigurationChange VBoxEventType = "OnHostNameResolutionConfigurationChange"

	VBoxEventTypeOnSnapshotRestored VBoxEventType = "OnSnapshotRestored"

	VBoxEventTypeOnMediumConfigChanged VBoxEventType = "OnMediumConfigChanged"

	VBoxEventTypeLast VBoxEventType = "Last"
)

type GuestMouseEventMode string

const (
	GuestMouseEventModeRelative GuestMouseEventMode = "Relative"

	GuestMouseEventModeAbsolute GuestMouseEventMode = "Absolute"
)

type GuestMonitorChangedEventType string

const (
	GuestMonitorChangedEventTypeEnabled GuestMonitorChangedEventType = "Enabled"

	GuestMonitorChangedEventTypeDisabled GuestMonitorChangedEventType = "Disabled"

	GuestMonitorChangedEventTypeNewOrigin GuestMonitorChangedEventType = "NewOrigin"
)

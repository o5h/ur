package dashboard

//Mode - robot mode
type Mode int8

const (
	ModeUndefined Mode = iota
	ModeNoController
	ModeDisconnected
	ModeConfirmSafety
	ModeBooting
	ModePowerOff
	ModePowerOn
	ModeIdle
	ModeBackdrive
	ModeRunning
)

func (m Mode) String() string {
	switch m {
	case ModeNoController:
		return "NO_CONTROLLER"
	case ModeDisconnected:
		return "DISCONNECTED"
	case ModeConfirmSafety:
		return "CONFIRM_SAFETY"
	case ModeBooting:
		return "BOOTING"
	case ModePowerOff:
		return "POWER_OFF"
	case ModePowerOn:
		return "POWER_ON"
	case ModeIdle:
		return "IDLE"
	case ModeBackdrive:
		return "BACKDRIVE"
	case ModeRunning:
		return "RUNNING"
	default:
		return "UNDEFINED"
	}
}

func modeOf(s string) Mode {
	switch s {
	case "NO_CONTROLLER":
		return ModeNoController
	case "DISCONNECTED":
		return ModeDisconnected
	case "CONFIRM_SAFETY":
		return ModeConfirmSafety
	case "BOOTING":
		return ModeBooting
	case "POWER_OFF":
		return ModePowerOff
	case "POWER_ON":
		return ModePowerOn
	case "IDLE":
		return ModeIdle
	case "BACKDRIVE":
		return ModeBackdrive
	case "RUNNING":
		return ModeRunning
	default:
		return ModeUndefined
	}
}

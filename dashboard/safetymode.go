package dashboard

type SafetyMode uint8

const (
	SafetyModeUndefined SafetyMode = iota
	SafetyModeNormal
	SafetyModeReduced
	SafetyModeProtectiveStop
	SafetyModeRecovery
	SafetyModeSafeGuardStop
	SafetyModeSystemEmergencyStop
	SafetyModeRobotEmergencyStop
	SafetyModeViolation
	SafetyModeFault
)

func (m SafetyMode) String() string {
	switch m {
	case SafetyModeNormal:
		return "NORMAL"
	case SafetyModeReduced:
		return "REDUCED"
	case SafetyModeProtectiveStop:
		return "PROTECTIVE_STOP"
	case SafetyModeRecovery:
		return "RECOVERY"
	case SafetyModeSafeGuardStop:
		return "SAFEGUARD_STOP"
	case SafetyModeSystemEmergencyStop:
		return "SYSTEM_EMERGENCY_STOP"
	case SafetyModeRobotEmergencyStop:
		return "ROBOT_EMERGENCY_STOP"
	case SafetyModeViolation:
		return "VIOLATION"
	case SafetyModeFault:
		return "FAULT"
	default:
		return "UNDEFINED"
	}
}

func safetyMode(s string) SafetyMode {
	switch s {
	case "NORMAL":
		return SafetyModeNormal
	case "REDUCED":
		return SafetyModeReduced
	case "PROTECTIVE_STOP":
		return SafetyModeProtectiveStop
	case "RECOVERY":
		return SafetyModeRecovery
	case "SAFEGUARD_STOP":
		return SafetyModeSafeGuardStop
	case "SYSTEM_EMERGENCY_STOP":
		return SafetyModeSystemEmergencyStop
	case "ROBOT_EMERGENCY_STOP":
		return SafetyModeRobotEmergencyStop
	case "VIOLATION":
		return SafetyModeViolation
	case "FAULT":
		return SafetyModeFault
	default:
		return SafetyModeUndefined
	}
}

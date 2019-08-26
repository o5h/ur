package dashboard

//ProgramState - STOPPED|PLAYING
type ProgramState uint8

const (
	// ProgramStateUndefined - undefined
	ProgramStateUndefined ProgramState = iota
	// ProgramStateStopped - "STOPPED"
	ProgramStateStopped
	// ProgramStatePlaying "PLAYING"
	ProgramStatePlaying
	// ProgramStatePaused - "PAUSED"
	ProgramStatePaused
)

func (s ProgramState) String() string {
	switch s {
	case ProgramStateStopped:
		return "STOPPED"
	case ProgramStatePlaying:
		return "PLAYING"
	case ProgramStatePaused:
		return "PAUSED"
	default:
		return "UNDEFINED"
	}
}
func programState(s string) ProgramState {
	switch s {
	case "STOPPED":
		return ProgramStateStopped
	case "PLAYING":
		return ProgramStatePlaying
	case "PAUSED":
		return ProgramStatePaused
	default:
		return ProgramStateUndefined
	}
}

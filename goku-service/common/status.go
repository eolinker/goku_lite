package common

import "strings"

//InstanceStatus instanceStanceStatus
type InstanceStatus int

const (
	//InstanceRun run
	InstanceRun = iota
	//InstanceDown down
	InstanceDown
	//InstanceChecking check
	InstanceChecking
)

func (status InstanceStatus) String() string {
	switch status {
	case InstanceRun:
		return "run"
	case InstanceDown:
		return "down"
	case InstanceChecking:
		return "checking"
	}
	return "unkown"
}

//ParseStatus parseStatus
func ParseStatus(status string) InstanceStatus {
	s := strings.ToLower(status)

	switch s {
	case "down":
		return InstanceDown
	case "checking":
		return InstanceChecking
	default:
		return InstanceRun
	}
}

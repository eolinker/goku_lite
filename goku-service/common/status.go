package common

import "strings"

type InstanceStatus int

const (
	InstanceRun = iota
	InstanceDown
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

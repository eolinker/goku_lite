package config

import "strings"

//StaticResponseStrategy 静态响应策略
type StaticResponseStrategy int

const (
	//Always always
	Always StaticResponseStrategy = iota
	//Success success
	Success
	//Errored errored
	Errored
	//Incomplete incomplete
	Incomplete
)

func (s StaticResponseStrategy) String() string {
	switch s {
	case Always:
		return "always"
	case Success:
		return "success"
	case Errored:
		return "errored"
	case Incomplete:
		return "incomplete"

	}
	return "unknown"
}

//Parse parse
func Parse(v string) StaticResponseStrategy {
	switch strings.ToLower(v) {
	case "always":
		return Always
	case "success":
		return Success
	case "errored":
		return Errored
	case "incomplete":
		return Incomplete
	}
	return Always
}

//Title title
func (s StaticResponseStrategy) Title() string {
	switch s {
	case Always:
		return "Always - Present in every response"
	case Success:
		return "Success - Present in every non-failed response "
	case Errored:
		return "Errored - Present in every failed response (error not nil)"
	case Incomplete:
		return "Incomplete - Present in incomplete responses"
	}
	return "unknown"
}

package redis

type Error string

func (err Error) Error() string { return string(err) }

type Scanner interface {
	// RedisScan assigns a value from a Redis value. The argument src is one of
	// the reply types listed in the section `Executing Commands`.
	//
	// An error should be returned if the value cannot be stored without
	// loss of information.
	RedisScan(src interface{}) error
}

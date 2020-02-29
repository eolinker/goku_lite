package goku_plugin

import (
	"time"
)

type Cmder interface {
	Name() string
	Args() []interface{}

	Err() error
}

type Cmd interface {
	Val() interface{}
	Result() (interface{}, error)
	String() (string, error)
	Int() (int, error)
	Int64() (int64, error)
	Uint64() (uint64, error)
	//Float32() (float32, error)
	Float64() (float64, error)
	Bool() (bool, error)
}
type SliceCmd interface {
	Val() []interface{}
	Result() ([]interface{}, error)
	String() string
}
type StatusCmd interface {
	Val() string
	Result() (string, error)
	String() string
}
type IntCmd interface {
	Val() int64
	Result() (int64, error)
	String() string
}
type DurationCmd interface {
	Val() time.Duration
	Result() (time.Duration, error)
	String() string
}
type TimeCmd interface {
	Val() time.Time
	Result() (time.Time, error)
	String() string
}
type BoolCmd interface {
	Val() bool
	Result() (bool, error)
	String() string
}
type StringCmd interface {
	Val() string
	Result() (string, error)
	Bytes() ([]byte, error)
	Int() (int, error)
	Int64() (int64, error)
	Uint64() (uint64, error)
	//Float32() (float32, error)
	Float64() (float64, error)
	Scan(val interface{}) error
	String() string
}
type FloatCmd interface {
	Val() float64
	Result() (float64, error)
	String() string
}
type StringSliceCmd interface {
	Val() []string
	Result() ([]string, error)
	String() string
	ScanSlice(container interface{}) error
}
type BoolSliceCmd interface {
	Val() []bool
	Result() ([]bool, error)
	String() string
}
type StringStringMapCmd interface {
	Val() map[string]string
	Result() (map[string]string, error)
	String() string
}
type StringIntMapCmd interface {
	Val() map[string]int64
	Result() (map[string]int64, error)
	String() string
}
type StringStructMapCmd interface {
	Val() map[string]struct{}
	Result() (map[string]struct{}, error)
	String() string
}

//
//type XMessage struct {
//	ID     string
//	Values map[string]interface{}
//}
//
//type XMessageSliceCmd interface {
//	Val() []XMessage
//	Result() ([]XMessage, error)
//	String() string
//}
//type XStream struct {
//	Stream   string
//	Messages []XMessage
//}
//type XStreamSliceCmd interface {
//	Val() []XStream
//	Result() ([]XStream, error)
//	String() string
//}
//type XPending struct {
//	Count     int64
//	Lower     string
//	Higher    string
//	Consumers map[string]int64
//}
//type XPendingCmd interface {
//	Val() *XPending
//	Result() (*XPending, error)
//	String() string
//}
//type XPendingExt struct {
//	Id         string
//	Consumer   string
//	Idle       time.Duration
//	RetryCount int64
//}
//type XPendingExtCmd interface {
//	Val() []XPendingExt
//	Result() ([]XPendingExt, error)
//	String() string
//}
//type ZSliceCmd interface {
//	Val() []redis.Z
//	Result() ([]redis.Z, error)
//	String() string
//}
//type ZWithKeyCmd interface {
//	Val() redis.ZWithKey
//	Result() (redis.ZWithKey, error)
//	String() string
//}
//type ScanCmd interface {
//	Val() (keys []string, cursor uint64)
//	Result() (keys []string, cursor uint64, err error)
//	String() string
//
//	Iterator() *redis.ScanIterator
//}
//type ClusterNode interface {
//}
//type ClusterSlot interface {
//}
//type ClusterSlotsCmd interface {
//	Val() []ClusterSlot
//	Result() ([]ClusterSlot, error)
//	String() string
//}
//
//// GeoLocation is used with GeoAdd to add geospatial location.
//type GeoLocation struct {
//	Name                      string
//	Longitude, Latitude, Dist float64
//	GeoHash                   int64
//}
//
//// GeoRadiusQuery is used with GeoRadius to query geospatial index.
//type GeoRadiusQuery struct {
//	Radius float64
//	// Can be m, km, ft, or mi. Default is km.
//	Unit        string
//	WithCoord   bool
//	WithDist    bool
//	WithGeoHash bool
//	Count       int
//	// Can be ASC or DESC. Default is no sort order.
//	Sort      string
//	Store     string
//	StoreDist string
//}
//
//type GeoLocationCmd interface {
//	Val() []GeoLocation
//	Result() ([]GeoLocation, error)
//	String() string
//}
//type GeoPos struct {
//	Longitude, Latitude float64
//}
//type GeoPosCmd interface {
//	Val() []*GeoPos
//	Result() ([]*GeoPos, error)
//	String() string
//}
//type CommandInfo struct {
//	Name        string
//	Arity       int8
//	Flags       []string
//	FirstKeyPos int8
//	LastKeyPos  int8
//	StepCount   int8
//	ReadOnly    bool
//}
//
//type CommandsInfoCmd interface {
//	Val() map[string]*CommandInfo
//	Result() (map[string]*CommandInfo, error)
//	String() string
//}

package goku_plugin

import "time"

type Cmdable interface {
	Ping() StatusCmd
	Del(keys ...string) IntCmd
	Unlink(keys ...string) IntCmd
	Dump(key string) StringCmd
	Exists(keys ...string) IntCmd
	Expire(key string, expiration time.Duration) BoolCmd
	ExpireAt(key string, tm time.Time) BoolCmd
	Keys(pattern string) StringSliceCmd

	RandomKey() StringCmd
	Rename(key, newkey string) StatusCmd
	RenameNX(key, newkey string) BoolCmd
	Restore(key string, ttl time.Duration, value string) StatusCmd
	RestoreReplace(key string, ttl time.Duration, value string) StatusCmd

	Append(key, value string) IntCmd
	//BitCountS(key string, Start, End int64) IntCmd
	BitOpAnd(destKey string, keys ...string) IntCmd
	BitOpOr(destKey string, keys ...string) IntCmd
	BitOpXor(destKey string, keys ...string) IntCmd
	BitOpNot(destKey string, key string) IntCmd
	BitPos(key string, bit int64, pos ...int64) IntCmd

	Decr(key string) IntCmd
	DecrBy(key string, decrement int64) IntCmd
	Get(key string) StringCmd
	GetBit(key string, offset int64) IntCmd
	GetRange(key string, start, end int64) StringCmd
	GetSet(key string, value interface{}) StringCmd
	Incr(key string) IntCmd
	IncrBy(key string, value int64) IntCmd
	IncrByFloat(key string, value float64) FloatCmd
	MGet(keys ...string) SliceCmd
	MSet(pairs ...interface{}) StatusCmd
	MSetNX(pairs ...interface{}) BoolCmd
	Set(key string, value interface{}, expiration time.Duration) StatusCmd
	SetBit(key string, offset int64, value int) IntCmd
	SetNX(key string, value interface{}, expiration time.Duration) BoolCmd
	SetXX(key string, value interface{}, expiration time.Duration) BoolCmd
	SetRange(key string, offset int64, value string) IntCmd
	StrLen(key string) IntCmd
	HDel(key string, fields ...string) IntCmd
	HExists(key, field string) BoolCmd
	HGet(key, field string) StringCmd
	HGetAll(key string) StringStringMapCmd
	HIncrBy(key, field string, incr int64) IntCmd
	HIncrByFloat(key, field string, incr float64) FloatCmd
	HKeys(key string) StringSliceCmd
	HLen(key string) IntCmd
	HMGet(key string, fields ...string) SliceCmd
	HMSet(key string, fields map[string]interface{}) StatusCmd
	HSet(key, field string, value interface{}) BoolCmd
	HSetNX(key, field string, value interface{}) BoolCmd
	HVals(key string) StringSliceCmd
	BLPop(timeout time.Duration, keys ...string) StringSliceCmd
	BRPop(timeout time.Duration, keys ...string) StringSliceCmd
	BRPopLPush(source, destination string, timeout time.Duration) StringCmd
	LIndex(key string, index int64) StringCmd
	LInsert(key, op string, pivot, value interface{}) IntCmd
	LInsertBefore(key string, pivot, value interface{}) IntCmd
	LInsertAfter(key string, pivot, value interface{}) IntCmd
	LLen(key string) IntCmd
	LPop(key string) StringCmd
	LPush(key string, values ...interface{}) IntCmd
	LPushX(key string, value interface{}) IntCmd
	LRange(key string, start, stop int64) StringSliceCmd
	LRem(key string, count int64, value interface{}) IntCmd
	LSet(key string, index int64, value interface{}) StatusCmd
	LTrim(key string, start, stop int64) StatusCmd
	RPop(key string) StringCmd
	RPopLPush(source, destination string) StringCmd
	RPush(key string, values ...interface{}) IntCmd
	RPushX(key string, value interface{}) IntCmd
	SAdd(key string, members ...interface{}) IntCmd
	SCard(key string) IntCmd
	SDiff(keys ...string) StringSliceCmd
	SDiffStore(destination string, keys ...string) IntCmd
	SInter(keys ...string) StringSliceCmd
	SInterStore(destination string, keys ...string) IntCmd
	SIsMember(key string, member interface{}) BoolCmd
	SMembers(key string) StringSliceCmd
	SMembersMap(key string) StringStructMapCmd
	SMove(source, destination string, member interface{}) BoolCmd
	SPop(key string) StringCmd
	SPopN(key string, count int64) StringSliceCmd
	SRandMember(key string) StringCmd
	SRandMemberN(key string, count int64) StringSliceCmd
	SRem(key string, members ...interface{}) IntCmd
	SUnion(keys ...string) StringSliceCmd
	SUnionStore(destination string, keys ...string) IntCmd
}

type StatefulCmdable interface {
	Cmdable
	Auth(password string) StatusCmd
	Select(index int) StatusCmd
	SwapDB(index1, index2 int) StatusCmd
	ClientSetName(name string) BoolCmd
}

type Pipeliner interface {
	StatefulCmdable
	Do(args ...interface{}) Cmd
	//Process(cmd Cmder) error
	Close() error
	Discard() error
	Exec() ([]Cmder, error)
}

type Redis interface {
	Cmdable
	Pipeline() Pipeliner
	Pipelined(fn func(Pipeliner) error) ([]Cmder, error)
}

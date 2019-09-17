package redis_plugin_proxy

import (
	redis2 "github.com/eolinker/goku-plugin"
	"github.com/go-redis/redis"
	"time"
)

type RedisProxy struct {
	redisClient redis.Cmdable
}

func (p *RedisProxy) Ping() redis2.StatusCmd { return p.redisClient.Ping() }

func (p *RedisProxy) Del(keys ...string) redis2.IntCmd { return p.redisClient.Del(keys...) }

func (p *RedisProxy) Unlink(keys ...string) redis2.IntCmd { return p.redisClient.Unlink(keys...) }

func (p *RedisProxy) Dump(key string) redis2.StringCmd { return p.redisClient.Dump(key) }

func (p *RedisProxy) Exists(keys ...string) redis2.IntCmd { return p.redisClient.Exists(keys...) }

func (p *RedisProxy) Expire(key string, expiration time.Duration) redis2.BoolCmd {
	return p.redisClient.Expire(key, expiration)
}

func (p *RedisProxy) ExpireAt(key string, tm time.Time) redis2.BoolCmd {
	return p.redisClient.ExpireAt(key, tm)
}

func (p *RedisProxy) Keys(pattern string) redis2.StringSliceCmd { return p.redisClient.Keys(pattern) }

func (p *RedisProxy) RandomKey() redis2.StringCmd { return p.redisClient.RandomKey() }

func (p *RedisProxy) Rename(key, newkey string) redis2.StatusCmd {
	return p.redisClient.Rename(key, newkey)
}

func (p *RedisProxy) RenameNX(key, newkey string) redis2.BoolCmd {
	return p.redisClient.RenameNX(key, newkey)
}

func (p *RedisProxy) Restore(key string, ttl time.Duration, value string) redis2.StatusCmd {
	return p.redisClient.Restore(key, ttl, value)
}

func (p *RedisProxy) RestoreReplace(key string, ttl time.Duration, value string) redis2.StatusCmd {
	return p.redisClient.RestoreReplace(key, ttl, value)
}

func (p *RedisProxy) Append(key, value string) redis2.IntCmd { return p.redisClient.Append(key, value) }

//func (p *RedisProxy) BitCountS(key string, Start, End int64) redis2.IntCmd { return p.redisClient.BitCountS(key,Start,End)}

func (p *RedisProxy) BitOpAnd(destKey string, keys ...string) redis2.IntCmd {
	return p.redisClient.BitOpAnd(destKey, keys ...)
}

func (p *RedisProxy) BitOpOr(destKey string, keys ...string) redis2.IntCmd {
	return p.redisClient.BitOpOr(destKey, keys ...)
}

func (p *RedisProxy) BitOpXor(destKey string, keys ...string) redis2.IntCmd {
	return p.redisClient.BitOpXor(destKey, keys ...)
}

func (p *RedisProxy) BitOpNot(destKey string, key string) redis2.IntCmd {
	return p.redisClient.BitOpNot(destKey, key)
}

func (p *RedisProxy) BitPos(key string, bit int64, pos ...int64) redis2.IntCmd {
	return p.redisClient.BitPos(key, bit, pos ...)
}

func (p *RedisProxy) Decr(key string) redis2.IntCmd { return p.redisClient.Decr(key) }

func (p *RedisProxy) DecrBy(key string, decrement int64) redis2.IntCmd {
	return p.redisClient.DecrBy(key, decrement)
}

func (p *RedisProxy) Get(key string) redis2.StringCmd { return p.redisClient.Get(key) }

func (p *RedisProxy) GetBit(key string, offset int64) redis2.IntCmd {
	return p.redisClient.GetBit(key, offset)
}

func (p *RedisProxy) GetRange(key string, start, end int64) redis2.StringCmd {
	return p.redisClient.GetRange(key, start, end)
}

func (p *RedisProxy) GetSet(key string, value interface{}) redis2.StringCmd {
	return p.redisClient.GetSet(key, value)
}

func (p *RedisProxy) Incr(key string) redis2.IntCmd { return p.redisClient.Incr(key) }

func (p *RedisProxy) IncrBy(key string, value int64) redis2.IntCmd {
	return p.redisClient.IncrBy(key, value)
}

func (p *RedisProxy) IncrByFloat(key string, value float64) redis2.FloatCmd {
	return p.redisClient.IncrByFloat(key, value)
}

func (p *RedisProxy) MGet(keys ...string) redis2.SliceCmd { return p.redisClient.MGet(keys...) }

func (p *RedisProxy) MSet(pairs ...interface{}) redis2.StatusCmd { return p.redisClient.MSet(pairs) }

func (p *RedisProxy) MSetNX(pairs ...interface{}) redis2.BoolCmd { return p.redisClient.MSetNX(pairs) }

func (p *RedisProxy) Set(key string, value interface{}, expiration time.Duration) redis2.StatusCmd {
	return p.redisClient.Set(key, value, expiration)
}

func (p *RedisProxy) SetBit(key string, offset int64, value int) redis2.IntCmd {
	return p.redisClient.SetBit(key, offset, value)
}

func (p *RedisProxy) SetNX(key string, value interface{}, expiration time.Duration) redis2.BoolCmd {
	return p.redisClient.SetNX(key, value, expiration)
}

func (p *RedisProxy) SetXX(key string, value interface{}, expiration time.Duration) redis2.BoolCmd {
	return p.redisClient.SetXX(key, value, expiration)
}

func (p *RedisProxy) SetRange(key string, offset int64, value string) redis2.IntCmd {
	return p.redisClient.SetRange(key, offset, value)
}

func (p *RedisProxy) StrLen(key string) redis2.IntCmd { return p.redisClient.StrLen(key) }

func (p *RedisProxy) HDel(key string, fields ...string) redis2.IntCmd {
	return p.redisClient.HDel(key, fields ...)
}

func (p *RedisProxy) HExists(key, field string) redis2.BoolCmd {
	return p.redisClient.HExists(key, field)
}

func (p *RedisProxy) HGet(key, field string) redis2.StringCmd { return p.redisClient.HGet(key, field) }

func (p *RedisProxy) HGetAll(key string) redis2.StringStringMapCmd { return p.redisClient.HGetAll(key) }

func (p *RedisProxy) HIncrBy(key, field string, incr int64) redis2.IntCmd {
	return p.redisClient.HIncrBy(key, field, incr)
}

func (p *RedisProxy) HIncrByFloat(key, field string, incr float64) redis2.FloatCmd {
	return p.redisClient.HIncrByFloat(key, field, incr)
}

func (p *RedisProxy) HKeys(key string) redis2.StringSliceCmd { return p.redisClient.HKeys(key) }

func (p *RedisProxy) HLen(key string) redis2.IntCmd { return p.redisClient.HLen(key) }

func (p *RedisProxy) HMGet(key string, fields ...string) redis2.SliceCmd {
	return p.redisClient.HMGet(key, fields ...)
}

func (p *RedisProxy) HMSet(key string, fields map[string]interface{}) redis2.StatusCmd {
	return p.redisClient.HMSet(key, fields)
}

func (p *RedisProxy) HSet(key, field string, value interface{}) redis2.BoolCmd {
	return p.redisClient.HSet(key, field, value)
}

func (p *RedisProxy) HSetNX(key, field string, value interface{}) redis2.BoolCmd {
	return p.redisClient.HSetNX(key, field, value)
}

func (p *RedisProxy) HVals(key string) redis2.StringSliceCmd { return p.redisClient.HVals(key) }

func (p *RedisProxy) BLPop(timeout time.Duration, keys ...string) redis2.StringSliceCmd {
	return p.redisClient.BLPop(timeout, keys ...)
}

func (p *RedisProxy) BRPop(timeout time.Duration, keys ...string) redis2.StringSliceCmd {
	return p.redisClient.BRPop(timeout, keys ...)
}

func (p *RedisProxy) BRPopLPush(source, destination string, timeout time.Duration) redis2.StringCmd {
	return p.redisClient.BRPopLPush(source, destination, timeout)
}

func (p *RedisProxy) LIndex(key string, index int64) redis2.StringCmd {
	return p.redisClient.LIndex(key, index)
}

func (p *RedisProxy) LInsert(key, op string, pivot, value interface{}) redis2.IntCmd {
	return p.redisClient.LInsert(key, op, pivot, value)
}

func (p *RedisProxy) LInsertBefore(key string, pivot, value interface{}) redis2.IntCmd {
	return p.redisClient.LInsertBefore(key, pivot, value)
}

func (p *RedisProxy) LInsertAfter(key string, pivot, value interface{}) redis2.IntCmd {
	return p.redisClient.LInsertAfter(key, pivot, value)
}

func (p *RedisProxy) LLen(key string) redis2.IntCmd { return p.redisClient.LLen(key) }

func (p *RedisProxy) LPop(key string) redis2.StringCmd { return p.redisClient.LPop(key) }

func (p *RedisProxy) LPush(key string, values ...interface{}) redis2.IntCmd {
	return p.redisClient.LPush(key, values ...)
}

func (p *RedisProxy) LPushX(key string, value interface{}) redis2.IntCmd {
	return p.redisClient.LPushX(key, value)
}

func (p *RedisProxy) LRange(key string, start, stop int64) redis2.StringSliceCmd {
	return p.redisClient.LRange(key, start, stop)
}

func (p *RedisProxy) LRem(key string, count int64, value interface{}) redis2.IntCmd {
	return p.redisClient.LRem(key, count, value)
}

func (p *RedisProxy) LSet(key string, index int64, value interface{}) redis2.StatusCmd {
	return p.redisClient.LSet(key, index, value)
}

func (p *RedisProxy) LTrim(key string, start, stop int64) redis2.StatusCmd {
	return p.redisClient.LTrim(key, start, stop)
}

func (p *RedisProxy) RPop(key string) redis2.StringCmd { return p.redisClient.RPop(key) }

func (p *RedisProxy) RPopLPush(source, destination string) redis2.StringCmd {
	return p.redisClient.RPopLPush(source, destination)
}

func (p *RedisProxy) RPush(key string, values ...interface{}) redis2.IntCmd {
	return p.redisClient.RPush(key, values ...)
}

func (p *RedisProxy) RPushX(key string, value interface{}) redis2.IntCmd {
	return p.redisClient.RPushX(key, value)
}

func (p *RedisProxy) SAdd(key string, members ...interface{}) redis2.IntCmd {
	return p.redisClient.SAdd(key, members ...)
}

func (p *RedisProxy) SCard(key string) redis2.IntCmd { return p.redisClient.SCard(key) }

func (p *RedisProxy) SDiff(keys ...string) redis2.StringSliceCmd { return p.redisClient.SDiff(keys...) }

func (p *RedisProxy) SDiffStore(destination string, keys ...string) redis2.IntCmd {
	return p.redisClient.SDiffStore(destination, keys ...)
}

func (p *RedisProxy) SInter(keys ...string) redis2.StringSliceCmd {
	return p.redisClient.SInter(keys...)
}

func (p *RedisProxy) SInterStore(destination string, keys ...string) redis2.IntCmd {
	return p.redisClient.SInterStore(destination, keys ...)
}

func (p *RedisProxy) SIsMember(key string, member interface{}) redis2.BoolCmd {
	return p.redisClient.SIsMember(key, member)
}

func (p *RedisProxy) SMembers(key string) redis2.StringSliceCmd { return p.redisClient.SMembers(key) }

func (p *RedisProxy) SMembersMap(key string) redis2.StringStructMapCmd {
	return p.redisClient.SMembersMap(key)
}

func (p *RedisProxy) SMove(source, destination string, member interface{}) redis2.BoolCmd {
	return p.redisClient.SMove(source, destination, member)
}

func (p *RedisProxy) SPop(key string) redis2.StringCmd { return p.redisClient.SPop(key) }

func (p *RedisProxy) SPopN(key string, count int64) redis2.StringSliceCmd {
	return p.redisClient.SPopN(key, count)
}

func (p *RedisProxy) SRandMember(key string) redis2.StringCmd { return p.redisClient.SRandMember(key) }

func (p *RedisProxy) SRandMemberN(key string, count int64) redis2.StringSliceCmd {
	return p.redisClient.SRandMemberN(key, count)
}

func (p *RedisProxy) SRem(key string, members ...interface{}) redis2.IntCmd {
	return p.redisClient.SRem(key, members ...)
}

func (p *RedisProxy) SUnion(keys ...string) redis2.StringSliceCmd {
	return p.redisClient.SUnion(keys...)
}

func (p *RedisProxy) SUnionStore(destination string, keys ...string) redis2.IntCmd {
	return p.redisClient.SUnionStore(destination, keys ...)
}

func (p *RedisProxy) Pipeline() redis2.Pipeliner {
	pipe := p.redisClient.Pipeline()
	return &PipelineProxy{
		RedisProxy: RedisProxy{redisClient: pipe},
		pipeliner:  pipe,
	}

}

func (p *RedisProxy) Pipelined(fn func(redis2.Pipeliner) error) ([]redis2.Cmder, error) {
	cmders, e := p.redisClient.Pipelined(func(pipeliner redis.Pipeliner) error {
		pip := &PipelineProxy{
			RedisProxy{pipeliner}, pipeliner,
		}
		return fn(pip)

	})
	if e != nil {
		return nil, e
	}

	cmds := make([]redis2.Cmder, 0, len(cmders))
	for _, c := range cmders {
		cmds = append(cmds, c)
	}
	return cmds, nil
}

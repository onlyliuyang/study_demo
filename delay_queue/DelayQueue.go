package delay_queue

import (
	"context"
	"time"
)

// Msg Msg消息
type Msg struct {
	Topic     string        //消息的主题
	Key       string        //消息的key
	Body      []byte        //消息的body
	Delay     time.Duration //延迟时间
	ReadyTime time.Time     //消息准备好执行的时间
}

const delayQueuePushRedisScript = `
	-- KEYS[1]: topicZSet
	-- KEYS[2]: topicHash
	-- ARGV[1]: 消息的KEY
	-- ARGV[2]: 消息的Body
	-- ARGV[3]: 消息准备好执行的时间

	local topicZSet = KEYS[1]
	local topicHash	= KEYS[2]
	local key	= ARGV[1]
	local body	= ARGV[2]
	local readyTime	= tonumber(ARGV[3])

	-- 添加readyTime到zset
	local count = redis.call("zadd", topicZSet, readyTime, key)
	-- 消息已经存在
	if count == 0 then
		return 0
	end
	-- 添加body到hash
	redis.call("hsetnx", topicHash, key, body)
	return 1
`

type SimpleRedisDelayQueue struct {
}

func (q *SimpleRedisDelayQueue) Push(ctx context.Context, msg *Msg) error {
	//如果设置了ReadyTime，就使用RedisTime
	var readyTime int64
	if !msg.ReadyTime.IsZero() {
		readyTime = msg.ReadyTime.Unix()
	} else {
		//否则使用delay
		readyTime = time.Now().Add(msg.Delay).Unix()
	}

	//success, err := q.Push()
}
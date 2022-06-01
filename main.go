package main

import (
	"fmt"
	"github.com/go-redis/redis"
	ormlog "golang_study/log"
	ormredis "golang_study/redis"
)

func init() {
	ormlog.InitLogger()
	err := ormredis.InitClient()
	if err != nil {
		ormlog.SugarLogger.Error("连接错误")
	} else {
		ormlog.SugarLogger.Infof("连接成功 host:%s,port:%s", "127.0.0.1", "6379")
	}
}

func main() {
	/*
		orm.InitLogger()
		defer orm.SugarLogger.Sync()
		//error
		orm.SugarLogger.Error("这是error")
		orm.SugarLogger.Errorf("这是%s", "error")
		//info
		orm.SugarLogger.Info("这是info")
		orm.SugarLogger.Infof("这是%s", "info")
		orm.SugarLogger.Infow("这是info，key and values", "key", "value")
		//debug
		orm.SugarLogger.Debug("这是debug")
		orm.SugarLogger.Debugf("这是%s", "debug")
		orm.SugarLogger.Debugw("这是debug，key and values", "key", "value")
		//warm
		orm.SugarLogger.Warn("这是warm")
		orm.SugarLogger.Warnf("这是%s", "warm")
		orm.SugarLogger.Warnw("这是warm，key and values", "key", "value")

	*/
	//redisExample2()
	iter := ormredis.Rdb.Scan(0, "s*", 0).Iterator()
	for iter.Next() {
		err := ormredis.Rdb.Del(iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

//set/get示例
func redisExample() {
	err := ormredis.Rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := ormredis.Rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := ormredis.Rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

//zset示例
func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := ormredis.Rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := ormredis.Rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := ormredis.Rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = ormredis.Rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

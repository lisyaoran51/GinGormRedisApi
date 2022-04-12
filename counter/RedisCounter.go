package counter

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v7"
)


type redisCounter struct {
	host string
	db int
	// expires time.Duration

}

const counterKey = "call_count"

func NewRedisCounter(host string, db int) Counter {
	return &redisCounter{
		host: host,
		db: db,
	}
}

func (c *redisCounter) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: c.host,
		Password: "",
		DB: c.db,
	})
}

func (c *redisCounter) String() string {
	return c.getClient().String()
}

func (c *redisCounter) Count() int {

	client := c.getClient()

	value, err := client.Get(counterKey).Result()
	if err != nil {
		fmt.Println("redisCounter.Count(): failed to get value", counterKey)
	}

	var tempCount int
	tempCount, err = strconv.Atoi(value)

	if err != nil {
		tempCount = 0
		fmt.Println("redisCounter.Count(): nil value ", counterKey, value)
	} 

	client.Set(counterKey, tempCount + 1, 0)
	fmt.Println("redisCounter.Count(): add one ", tempCount)

	return 0
}

func (c *redisCounter) GetCount() int {
	client := c.getClient()

	value, err := client.Get(counterKey).Result()
	if err != nil {
		fmt.Println("redisCounter.Count(): failed to get value", counterKey)
	}

	tempCount, _ := strconv.Atoi(value)


	return tempCount
}


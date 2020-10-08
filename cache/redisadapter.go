package cache

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
)

const notStoreErr = "could not store in DB"

func init() {

	connectionList = newConectionList()
}

var (
	connectionList *redisConn
	rOnce          sync.Once
)

type redisConn struct {
	connection *redis.Client
}

func (c *redisConn) setConnection(redisConnection *redis.Client) {

	c.connection = redisConnection

}

func (c *redisConn) getConnection() (*redis.Client, error) {

	redisConnection := c.connection

	// if not cached 1.read redis server url from vault 2. connect to
	//redis 3. cache the connection

	if redisConnection == nil {
		//1. not cached

		redisClient := connectRedis("localhost:6379", "", 0)
		c.setConnection(redisClient)
		return redisClient, nil

	}
	return redisConnection, nil
}

func newConectionList() *redisConn {
	rOnce.Do(func() {
		connectionList = &redisConn{}
	})

	return connectionList
}

func connectRedis(address, password string, db int) *redis.Client {
	return redis.NewClient(loadRedisOptions(address, password,
		db))
}

func loadRedisOptions(address, password string, db int) *redis.Options {
	return &redis.Options{Addr: address, Password: password, DB: db}
}

func HSet(key, fld, val string) (bool, error) {
	c, err := connectionList.getConnection()

	if err != nil {
		log.Fatalf(notStoreErr+":%v", err)
	}

	boolCmd := c.HSet(key, fld, val)

	return boolCmd.Result()

}

func SAdd(set, member string) (int64, error) {
	c, err := connectionList.getConnection()

	if err != nil {
		log.Fatalf(notStoreErr+":%v", err)
	}
	intCmd := c.SAdd(set, member)

	return intCmd.Result()
}

func SIsMember(set, element string) (bool, error) {
	c, err := connectionList.getConnection()

	if err != nil {
		log.Fatalf(notStoreErr+":%v", err)
	}
	boolCmd := c.SIsMember(set, element)

	return boolCmd.Result()

}

func HGet(key, fld string) (string, error) {
	c, err := connectionList.getConnection()
	if err != nil {
		log.Fatalf(notStoreErr+":%v", err)
	}
	strCmd := c.HGet(key, fld)
	return strCmd.Result()
}

func LPush(q, val string) (int64, error) {
	c, err := connectionList.getConnection()

	if err != nil {
		log.Fatalf(notStoreErr+":%v", err)
	}

	intCmd := c.LPush(q, val)

	return intCmd.Result()
}

func RPop(q string) (string, error) {
	c, err := connectionList.getConnection()

	if err != nil {
		log.Fatalf(notStoreErr+":%v", err)
	}

	strCmd := c.RPop(q)

	return strCmd.Result()
}

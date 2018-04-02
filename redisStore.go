package base64Captcha

import (
	"github.com/garyburd/redigo/redis"
	"strings"
	"strconv"
	"log"
)

const MaxPoolSize = 10

var redisProvider = &redisStore{}

type redisStore struct{
	p *redis.Pool
	expire int64
}

func (rs *redisStore)Set(id string, value string){
	conn := rs.p.Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", id, rs.expire, value)
	if err != nil {
		log.Fatalf("redisStore set error: %s", err.Error())
	}
}

func (rs *redisStore)Get(id string, clear bool) string{
	conn := rs.p.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", id))
	if err != nil {
		log.Fatalf("redisStore get error: %s", err.Error())
	}
	if clear{
		_, err = conn.Do("DEL", id)
		if err != nil {
			log.Fatalf("redisStore get error: %s", err.Error())
		}
	}
	return value
}

// extraConfig like redis server addr,pool size,password,dbnum
// e.g. 127.0.0.1:6379,100,astaxie,0
func (rs *redisStore)InitStore(sc storeConfig) error{

	var addr, password string
	var poolsize, dbnum int

	configs := strings.Split(sc.extraConfig, ",")
	if len(configs) > 0 {
		addr = configs[0]
	}
	if len(configs) > 1 {
		poolsize, err := strconv.Atoi(configs[1])
		if err != nil || poolsize < 0 {
			poolsize = MaxPoolSize
		}
	} else {
		poolsize = MaxPoolSize
	}
	if len(configs) > 2 {
		password = configs[2]
	}
	if len(configs) > 3 {
		dbnum, err := strconv.Atoi(configs[3])
		if err != nil || dbnum < 0 {
			dbnum = 0
		}
	} else {
		dbnum = 0
	}

	pool := &redis.Pool{
		MaxIdle:poolsize,
		Dial:func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err = c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			_, err = c.Do("SELECT", dbnum)
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
	rs.p = pool
	rs.expire = sc.expire
	return rs.p.Get().Err()
}

func init (){
	Register("redis", redisProvider)
}
package test

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/goinggo/mapstructure"

	//"golang.org/x/crypto/ssh/agent"
	"testing"
	"time"
)
//var rdb *redis.Client
func TestConnectRedis(t *testing.T){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
	t.Log(err)
	}else{
		t.Log("success")
	}
}
//string操作
func TestSetAndGet(t *testing.T){
	rdb:= redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		t.Log(err)
	}else{
		t.Log("success")
	}
	//set key
	rdb.Set(context.Background(),"today","2022-11-14",40*time.Second)
	res:=rdb.Get(context.Background(),"today1")
	t.Log(res)
	//strlen key
	cmd:=rdb.StrLen(context.Background(),"today")
	t.Log(cmd)
	//mget
	res1:=rdb.MSet(context.Background(),"birthday","2002.09.18","lover","noBody")
	t.Log(res1)
	//mget
	res2:=rdb.MGet(context.Background(),"birthday","lover")
	t.Log(res2)
	//append
	res3:=rdb.Append(context.Background(),"birthday",":01:01:01")
	t.Log(res3)
	//t.Log(rdb.Get(context.Background(),"today"))
}

type student struct {
	Age  string `redis:"age"`
	Gender string	`redis:"gender"`
	Name string	 `redis:"name"`
}
//hash操作(结构化操作)
func TestHashOperation(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		t.Log(err)
	} else {
		t.Log("success")
	}
	//res:=rdb.HGetAll(context.Background(),"runoobkey")
	res1 := rdb.HGetAll(context.Background(), "xiaowei")
	t.Log(res1.Val()["age"])
	for k,v:=range res1.Val(){
		t.Log(k,v)
	}
	 xw :=new(student)
	data,err:=rdb.HGetAll(context.Background(),"xiaowei").Result()
	if err!=nil{
		t.Log(err)
	}
	err=mapstructure.Decode(data,xw)
	if err!=nil{
		t.Log(err)
	}
	t.Log(xw)

}

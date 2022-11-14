package test

import (
	"context"
	//"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis/v9"
	"github.com/mitchellh/mapstructure"
	"testing"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}
type student1 struct {
	Name   string
	Age    int
	Gender string
}

func TestCh2(t *testing.T){

	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"gender":"男",
	}
	var xw student1
	err := mapstructure.Decode(input, &xw)
	if err != nil {
		t.Log(err)
	}
	t.Log(xw)
}
func TestCh(t *testing.T){

	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}
	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		t.Log(err)
	}
	t.Log(result)
}
//获取指定key的field值
func  TestGetByField(t *testing.T){
	rdb:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	_,err:=rdb.Ping(context.Background()).Result()
	if err!=nil{
		t.Log("Connect Redis Server Failed")
		return
	}else{
		t.Log("Connect Redis Server Success")
	}
	name,err:=rdb.HGet(context.Background(),"school","name").Result()
	t.Log(name)
}
//List Operation
func TestListOperation(t *testing.T)  {
	rdb:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	strs,err:=rdb.LRange(context.Background(),"stringlist",0,-1).Result()
	if err!=nil{
		t.Log(err)
	}
	t.Log(strs)
}
package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

//创建客户端
//redis-cli -h 172.27.139.101 -p 25367
func initClient() (*redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr: "172.27.139.101:25367",
		Password: "",
		DB: 1,
		PoolSize: 5,
	})
	pong, err := client.Ping().Result()//检查是否成功连接
	fmt.Println(pong, err)//PONG nil
	return client
}

//String操作
func stringOperaion(client *redis.Client) {
	//第三个参数是过期时间，如果是0，则表示没有过期时间
	//set
	err := client.Set("name", "xys", 0).Err()
	if err != nil {
		panic(err)
	}
	//get
	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", val)

	//设置过期时间
	err = client.Set("age", "20", 1 * time.Second).Err()
	if err != nil {
		panic(err)
	}
	client.Incr("age")//自增
	client.Incr("age")
	client.Decr("age")//自减
	val, err = client.Get("age").Result()//21
	if err != nil {
		panic(err)
	}
	fmt.Println("age: ", val)

	//由于Key"age"的过期时间是1s，1s后key会自动被删除
	time.Sleep(time.Second * 1)
	val, err = client.Get("age").Result()
	if err != nil {
		//由于key"age"已经过期，因此会有一个redis：nil的错误
		fmt.Printf("error:%v\n", err)
	}
	fmt.Println("age", val)
}

//list操作
func listOperation(client *redis.Client) {
	//RPush
	client.RPush("fruit", "apple")//在fruit这个list的尾部加apple
	//LPush
	client.LPush("fruit", "peach")//在fruit在list头添加peach
	//LLen
	length, err := client.LLen("fruit").Result()//返回fruit这个list的长度
	if err != nil {
		panic(err)
	}
	fmt.Println("length:", length)
	//LPop
	value, err := client.LPop("fruit").Result()//返回并删除list的首元素
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit:", value)
	//RPop
	value, err = client.RPop("fruit").Result()//返回并删除list的尾元素
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit:", value)
}

//set操作
func setOperation(client *redis.Client) {
	//SAdd
	client.SAdd("blacklist", "Obama")//向blacklist中添加元素
	client.SAdd("blacklist", "Hillary")
	client.SAdd("blacklist", "the Elder")
	client.SAdd("whitelist", "the Elder")
	//SIsMember
	isMember, err := client.SIsMember("blacklist", "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is Bush in blacklist", isMember)
	//SInter求交集
	names, err := client.SInter("blacklist", "whitelist").Result()
	if err != err {
		panic(err)
	}
	fmt.Println("Inter result:", names)
	//SMembers获取指定集合的所有元素
	all, err := client.SMembers("blacklist").Result()
	if err != err {
		panic(err)
	}
	fmt.Println("blacklist All member:", all)
}

//hash操作
func hashOperation(client *redis.Client) {
	//HSet向名称为key的hash中添加元素field
	client.HSet("user_xyz", "name", "xys")
	client.HSet("user_xyz", "age", "18")
	//HMSet批量的向名称为key的hash中添加元素name和age
	client.HMSet("user_test", map[string]interface{}{"name": "test", "age":"20"})
	//HMGet批量的获取名称为key的hash中的指定字段的值
	fields, err := client.HMGet("user_test", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields in user_test:", fields)//[test 20]
	//HLen获取名为key的hash中的字段个数
	length, err := client.HLen("user_xys").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("field coun in user_xys:", length)//0
	//HDel删除名为key的fields字段
	age, err := client.HDel("user_test", "age").Result()
	if err != nil {
		fmt.Printf("Get user_test age error:%v\n", err)
	} else {
		fmt.Println("user_test age is:", age)//1
	}
}

//开启十个gorutine对设置连接池大小为5的redis读写
func connectPool(client *redis.Client) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xyz%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}
			fmt.Printf("PoolStats, TotalConns:%d\n", client.PoolStats().TotalConns)
		}()
	}
	wg.Wait()
}

func main() {
	client := initClient()
	fmt.Println(client)//Redis<172.27.139.101:25367 db:1>
	//stringOperaion(client)
	//listOperation(client)
	//setOperation(client)
	//hashOperation(client)
	//connectPool(client)
}
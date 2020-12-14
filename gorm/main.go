package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model   //包含ID(默认作为主键)，CreateAt，UpdateAt，DeleteAt
	Name         string
	Age          int8
	Birthday     time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //字段大小为255字节
	MemberNumber *string `gorm:"unique;not null"` //唯一，不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //自增
	//给Address 创建一个名字是  `addr`的索引
	Address  string `gorm:"index:addr"`
	IgnoreMe int    `gorm:"-"` //忽略这个字段
}

type Student struct {
	Id   uint `gorm:"primary_key"`
	Name string
}

func main() {
	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	//为了处理time.Time，需要包括parseTime作为参数。
	db, err := gorm.Open("mysql", "root:XBj0QdYM0P@(172.27.139.101:24299)/"+
		"test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//检查

	//检查模型`Student`是否存在
	exist := db.HasTable(&Student{})
	fmt.Println(exist)
	//检查表`student`是否存在
	exist = db.HasTable("student")
	fmt.Println(exist)

	//创建

	// 如果设置禁用表名复数形式属性为 true，`Student` 的表名将是 `student`
	db.SingularTable(true)
	//为模型`User`创建表
	if !exist {
		db.CreateTable(&Student{})
		//创建表'users'时将"ENGINE = InnoDB"附加到SQL语句
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Student{})
	}

	//创建记录
	var stu Student
	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("name%d", i)
		stu = Student{Id: uint(i), Name: name}
		db.Create(&stu) //Create insert the value into database
		//flag := db.NewRecord(stu) //NewRecord check if value's primary key is blank
		//fmt.Println(flag)
	}

	//查询

	//获取第一条记录，按主键排序
	var query1 Student
	// SELECT * FROM student ORDER BY id LIMIT 1;
	db.First(&query1)
	fmt.Println(query1)

	//获取最后一条记录，按主键排序
	// SELECT * FROM student ORDER BY id DESC LIMIT 1;
	var query2 Student
	db.Last(&query2)
	fmt.Println(query2)

	//获取所有记录
	//SELECT * FROM users;
	var stus []Student
	db.Find(&stus)
	for i, j := range stus {
		fmt.Println(i, j)
	}

	//通过主键进行查询
	//SELECT * FROM student WHERE id = 10;
	var query3 Student
	db.First(&query3, 10)
	fmt.Println(query3)

	//Where

	//获取第一条匹配的记录
	//SELECT * FROM student WHERE name = 'name0' limit 1;
	var query4 Student
	db.Where("name = ?", "name0").First(&query4)
	fmt.Println(query4)

	//获取所有匹配的记录
	//SELECT * FROM users WHERE name = 'name0';
	var query5 []Student
	db.Where("name = ?", "name0").Find(&query5)
	fmt.Println(query5)

	//<> 除了name == xiaoming以外的所有
	var query6 []Student
	db.Where("name <> ?", "xiaoming").Find(&query6)
	fmt.Println(query6)

	//IN
	//SELECT * FROM student WHERE name = 'name0' OR name = 'name2'
	var query7 []Student
	db.Where("name in (?)", []string{"name0", "name2"}).Find(&query7)
	fmt.Println(query7)

	//LIKE
	//SELECT * FROM student WHERE name LIKE 'xiao%'
	//"%" 符号用于在模式的前后定义通配符（默认字母）
	var query8 []Student
	db.Where("name LIKE ?", "xiao%").Find(&query8)
	fmt.Println(query8)

	//AND
	var query9 []Student
	db.Where("name LIKE ? AND id >= ?", "name%", "96").Find(&query9)
	fmt.Println(query9)

	//BETWEEN
	var query10 []Student
	db.Where("id BETWEEN ? AND ?", uint(1), uint(3))
	fmt.Println(query10)

}

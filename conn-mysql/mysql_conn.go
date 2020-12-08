package main

import (
	"database/sql"
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
)

func main() {
	//mysql, 用户名:密码@[连接方式](主机名:端口号)/数据库名
	//mysql -h172.27.139.101 -P24299 -uroot -pXBj0QdYM0P
	db, _ := sql.Open("mysql",
		"root:XBj0QdYM0P@(172.27.139.101:24299)/test")

	err := db.Ping()//连接数据库
	if err != nil {
		fmt.Println("Connect Error!")
	}
	defer db.Close()

	/*
	//操作一：执行数据操作（插入更新删除）语句
	//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	sql := "insert into stu values(2, 'berry')"
	result, _ := db.Exec(sql)//执行SQL语句
	n, _ := result.RowsAffected()//获取受影响的记录数
	fmt.Println("受影响的记录数是", n)

	sqlStr := "insert into stu(id, name) values (?, ?)"
	ret, err := db.Exec(sqlStr, 1, "apple")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theId, _ := ret.LastInsertId()//新插入数据的id
	fmt.Println("insert success, the id is", theId)
	*/

	/*
	//操作二：执行预处理
	//func (db *DB) Prepare(query string) (*Stmt, error)
	//Prepare方法会先将sql语句发送给MySQL服务端，
	//返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询或命令。
	stu := [2][2]string{{"3", "ketty"}, {"4", "rose"}}
	stmt, _ := db.Prepare("insert into stu values (?, ?)")
	for _, s := range stu {
		stmt.Exec(s[0], s[1])//调用预处理语句
	}
	defer stmt.Close()
	*/

	/*
	//操作三：单行查询
	//func (db *DB) QueryRow(query string, args ...interface{}) *Row
	var id, name string
	rows := db.QueryRow("select * from stu where id=4")//获取一行数据
	rows.Scan(&id, &name)//将rows中的数据存到id，name中
	fmt.Println(id, ":", name)
	*/

	/*
	//操作四：多行查询
	//func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	var id, name string
	rows, _ := db.Query("select * from stu")//获取所有数据
	defer rows.Close()//关闭rows，释放持有的数据库连接
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Println(id, ":", name)
	}
	*/
	transactionDemo(db)
}

//事务操作
//开始事务func (db *DB) Begin() (*Tx, error)
//提交事务func (tx *Tx) Commit() error
//回滚事务func (tx *Tx) Rollback() error
func transactionDemo(db *sql.DB) {
	tx, err := db.Begin()//开始事务
	if err != nil {
		if tx != nil {
			tx.Rollback()//回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update stu set name = ? where id = 1"
	_, err = tx.Exec(sqlStr1, "trans1")
	if err != nil {
		tx.Rollback()//回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	sqlStr2 := "Update stu set name = ? where id = 2"
	_, err = tx.Exec(sqlStr2, "trans2")
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit()//提交事务
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
	}
	fmt.Println("exec trans success!")
}

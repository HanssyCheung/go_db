package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//定义一个全局对象DB
var db1 *sql.DB

//定义一个初始化数据库的函数
func initDB1() (err error) {
	//配置Mysql链接参数
	dsn := "root@tcp(127.0.0.1:3306)/database_demo?charset=utf8mb4&parseTime=True"
	//不会校验账号密码是否正确
	//注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db1, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接(校验dsn是否正确)
	err = db1.Ping()
	if err != nil {
		return err
	}
	return nil
}

//插入数据
func insert() {
	s := "insert into users (username,password,create_time) values(?,?,?)"
	r, err := db1.Exec(s, "李四", "ls123", time.Now().UnixMilli())
	if err != nil {
		fmt.Printf("insert failed,err :%v\n", err)
		return
	} else {
		theId, err := r.LastInsertId()
		if err != nil {
			fmt.Printf("get lastinsert Id failed,err: %v\n", err)
			return
		}
		fmt.Printf("insert success,the id is %d.\n", theId)
	}
}

//获得连接
func main() {
	err := initDB1() //调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功")
	}
	insert()

}

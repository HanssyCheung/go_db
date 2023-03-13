package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局对象DB
var db4 *sql.DB

//定义一个初始化数据库的函数
func initDB4() (err error) {
	dsn := "root@tcp(127.0.0.1:3306)/database_demo?charset=utf8mb4&parseTime=True"
	//不会校验账号密码是否正确
	//注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db4, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接(校验dsn是否正确)
	err = db4.Ping()
	if err != nil {
		return err
	}
	return nil
}

//查询数据

type User struct {
	id       int
	username string
	password string
	time     int
}

/*
	单行查询db.QueryRow()执行一次查询，并期望返回最多一行的结果（即Row）。QueryRow总是返回非nil的值，
直到返回值的scan方法被调用时，才会返回被延迟的错误
*/
func queryOneRow() {
	s := "select * from users where id = ?"
	var u User
	err := db4.QueryRow(s, 1).Scan(&u.id, &u.username, &u.password, &u.time)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("u: %v\n", u)
	}
}

//查询多行
/*
	多行查询db.Query()执行一次查询，返回多行结果（即Rows），一般用于select命令。参数args表示query中的占位参数
*/

//查询多条数据实例
func queryMultiRow() {
	sqlStr := "select id, username, password, create_time from users where id > ?"
	rows, err := db4.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed,err: %v\n", err)
		return
	}
	//非常重要，关闭rows释放持有的数据库链接
	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.username, &u.password, &u.time)
		if err != nil {
			fmt.Printf("scan failed,err: %v\n", err)
			return
		} else {
			fmt.Printf("id: %d username: %s password:%s time:%v\n", u.id, u.username, u.password, u.time)
		}
	}
}

//获得连接
func main() {
	err := initDB4() //调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功")
	}
	queryMultiRow()
}

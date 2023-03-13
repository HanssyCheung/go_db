package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局对象DB
var db3 *sql.DB

//定义一个初始化数据库的函数
func initDB3() (err error) {
	//配置Mysql链接参数
	dsn := "root@tcp(127.0.0.1:3306)/database_demo?charset=utf8mb4&parseTime=True"
	//不会校验账号密码是否正确
	//注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db3, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接(校验dsn是否正确)
	err = db3.Ping()
	if err != nil {
		return err
	}
	return nil
}

//删除数据
func deleteData() {
	s := "delete from users where  id=?"
	r, err := db3.Exec(s, 1)
	if err != nil {
		fmt.Printf("删除失败，err = : %v\n", err)
		return
	} else {
		i, _ := r.RowsAffected()
		if err != nil {
			fmt.Printf("更新行 失败，err = : %v\n", err)
			return
		}
		fmt.Printf("更新成功，更新行数：%d.\n", i)
	}
}

//获得连接
func main() {
	err := initDB3() //调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功")
	}
	deleteData()
}

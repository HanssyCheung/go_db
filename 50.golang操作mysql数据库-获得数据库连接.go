package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局对象DB
var db *sql.DB

//定义一个初始化数据库的函数
func initDB() (err error) {
	//配置Mysql链接参数
	//username := "root"        //账号
	//password := ""            //密码
	//host := "127.0.0.1"       //数据库地址，可以是IP或者域名
	//port := 3306              //数据库端口
	//Dbname := "database_demo" //数据库名
	//
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	dsn := "root@tcp(127.0.0.1:3306)/database_demo?charset=utf8mb4&parseTime=True"
	//不会校验账号密码是否正确
	//注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接(校验dsn是否正确)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//获得连接
func main() {
	////初始化连接
	////Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库得连接。如果要检查数据源得名称是否真实有效，应该调用Ping方法
	////返回得DB对象可以安全地被多个goroutine并发使用，并且维护其自己地空闲连接池。因此，Open函数应该仅被调用一次，很少需要关闭这个DB对象
	//db, err = sql.Open("mysql", "root:data_base@/go.db")
	//if err != nil {
	//	panic(err)
	//}
	//print(db)
	////最大连接时长
	//db.SetConnMaxIdleTime(time.Minute * 3)
	//
	////最大连接数
	//db.SetMaxOpenConns(10)
	//
	////空闲连接数
	//db.SetMaxIdleConns(10)
	err := initDB() //调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功")
	}
}

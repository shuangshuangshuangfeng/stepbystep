package dao

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //这个引用是必不可少的，因为需要调用driver.go文件里的init方法来提供一个数据库驱动程序
	"github.com/spf13/viper"
)

var DB *sql.DB     //全局变量，这样可以在别处调用

func Init() error {

	var err error

	//这行代码的作用就是初始化一个sql.DB对象
	DB ,err = sql.Open("mysql", viper.GetString("mysql.source_name"))
	if nil != err {
		return err
	}

	//设置最大超时时间
	DB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	//建立链接
	err = DB.Ping()
	if nil != err{
		return err
	}else{
		log.Println("Mysql Startup Normal!")
	}
	return nil
}
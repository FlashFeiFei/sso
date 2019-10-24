package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"os/signal"
	"runtime"
)

//读写连接
var readContent *gorm.DB
var writeContent *gorm.DB

//初始化连接数据库的信息
func init() {
	_, path, _, _ := runtime.Caller(0)
	log.Println(path)
	//读连接
	if readContent == nil {
		//只读连接
		readContent, err := gorm.Open(
			"mysql",
			"user:password@/dbname?charset=utf8&parseTime=True&rejectReadOnly=true")
		if err != nil {
			readContent.Close()
			log.Fatalln("读的数据库连接失败:", err)
		}
		// SetMaxIdleCons 设置连接池中的最大闲置连接数。
		readContent.DB().SetMaxIdleConns(10)
		// SetMaxOpenCons 设置数据库的最大连接数量。
		readContent.DB().SetMaxOpenConns(100)
	}

	//写连接
	if writeContent == nil {
		writeContent, err := gorm.Open(
			"mysql",
			"user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			writeContent.Close()
			log.Fatalln("读的数据库连接失败:", err)
		}
		// SetMaxIdleCons 设置连接池中的最大闲置连接数。
		writeContent.DB().SetMaxIdleConns(5)
		// SetMaxOpenCons 设置数据库的最大连接数量。
		writeContent.DB().SetMaxOpenConns(50)
	}

	//异步的去监听程序是否收到中断信号，从而释放连接资源
	if readContent != nil && writeContent != nil {
		go safeOut()
	}

}

//获取读连接
func GetReadContent() *gorm.DB {
	return readContent
}

//获取写的连接
func GetWriteContent() *gorm.DB {
	return writeContent
}

//安全退出
func safeOut() {
	//有缓冲通道,不会阻塞
	c := make(chan os.Signal, 1)
	//监听所有中断安全退出的信号
	signal.Notify(c, os.Interrupt, os.Kill)

	s := <-c
	log.Println("收到终端发来的中断信号:", s)

	//释放数据库连接资源
	readContent.Close()
	writeContent.Close()
}

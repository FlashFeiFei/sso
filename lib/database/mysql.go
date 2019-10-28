package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"os"
	"os/signal"
)

//读写连接
var readContent []*gorm.DB
var writeContent []*gorm.DB

//初始化连接数据库的信息
func init() {
	//解析mysql配置文件
	db_info := ParseConfig("")

	if len(db_info.WriteInfo) <= 0 {
		log.Fatalln("不能没有写连接的mysql")
	}

	if writeContent == nil {
		writeContent = make([]*gorm.DB, 0)
		//初始化写连接的mysql
		initContent(writeContent, db_info.WriteInfo)
	}

	if readContent == nil {
		readContent = make([]*gorm.DB, 0)
		//初始化读连接的mysql
		initContent(readContent, db_info.ReadInfo)
	}

	//异步的去监听程序是否收到中断信号，从而释放连接资源
	if len(writeContent) > 0 {
		go safeOut()
	}
}

//初始化连接
func initContent(contentGroup []*gorm.DB, contentInfo []ContentInfo) {
	for _, content := range contentInfo {
		db, err := gorm.Open("mysql", content.GetContentInfo())
		// SetMaxOpenCons 设置数据库的最大连接数量。
		db.DB().SetMaxOpenConns(10)
		// SetMaxIdleCons 设置连接池中的最大闲置连接数。
		db.DB().SetMaxIdleConns(100)
		if err != nil {
			log.Fatalln("mysql连接出错", err)
		}
		contentGroup = append(contentGroup, db)
	}
}

//获取读连接
func GetReadContent() *gorm.DB {
	content_number := len(readContent)
	//如果读取连接为空，则读写为同一个连接
	if content_number == 0 {
		return GetWriteContent()
	}
	return readContent[rand.Intn(content_number)]
}

//获取写的连接
func GetWriteContent() *gorm.DB {
	content_number := len(writeContent)
	if content_number == 0 {
		panic("写连接为空")
	}
	return writeContent[rand.Intn(content_number)]
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
	for _, writeDB := range writeContent {
		writeDB.Close()
	}

	for _, readDB := range readContent {
		readDB.Close()
	}

}

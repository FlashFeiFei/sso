package database

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
)

//解析配置文件
func ParseConfig(db_config_path string) (dbInfo DBInfostruct) {

	if db_config_path == "" {
		//获取当前执行文件的绝对路径
		_, path, _, _ := runtime.Caller(0)
		application_path := filepath.Dir(filepath.Dir(filepath.Dir(path)))
		log.Println(application_path)
		db_config_path := filepath.Join(application_path, "config", "db.json")
		log.Println("db路径:", db_config_path)
	}

	file, err := os.Open(db_config_path)
	if err != nil {
		log.Fatalln("解析数据库文件失败", err)
	}
	defer file.Close()
	json_decode := json.NewDecoder(file)
	json_decode.Decode(&dbInfo)
	return
}

type DBInfostruct struct {
	WriteInfo []ContentInfo `json:"Write"`
	ReadInfo  []ContentInfo `json:"Read"`
}

//连接信息
type ContentInfo struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
	//数据库连接账号
	User string `json:"User"`
	//密码
	Password string `json:"Password"`
	//连接的数据库
	DBName string `json:"DBName"`
	//连接信息参数
	Param1 url.Values `json:"-"`
}

func (ci ContentInfo) GetContentInfo() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", ci.User, ci.Password, ci.IP, ci.Port, ci.DBName, ci.Param1.Encode())
}

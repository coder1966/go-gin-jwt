package util

import (
	"fmt"
	"github.com/go-ini/ini"
	"time"
)

var (
	AppMode      string
	HttpPort     string
	JwtKey       string // 盐
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	DB         string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	ImgUrl    string
	AccessKey string
	SecretKey string
	Bucket    string

	LogPath string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	LoadServer(file)
	LoadData(file)
	LoadImgServer(file)
	LoadPath(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("KUY7D9P3se")
	ReadTimeOut = file.Section("server").Key("ReadTimeOut").MustDuration(600 * time.Minute)
	WriteTimeOut = file.Section("server").Key("WriteTimeOut").MustDuration(60 * time.Minute)
}

func LoadData(file *ini.File) {
	DB = file.Section("database").Key("DB").MustString("mysql")
	DBHost = file.Section("database").Key("DBHost").MustString("localhost")
	DBPort = file.Section("database").Key("DBPort").MustString("3417")
	DBUser = file.Section("database").Key("DBUser").MustString("gotest")
	DBPassword = file.Section("database").Key("DBPassword").MustString("gotest123456")
	DBName = file.Section("database").Key("DBName").MustString("go_test_db")
}

func LoadImgServer(file *ini.File) {
	ImgUrl = file.Section("imgServer").Key("ImgUrl").String()
	AccessKey = file.Section("imgServer").Key("AccessKey").String()
	SecretKey = file.Section("imgServer").Key("SecretKey").String()
	Bucket = file.Section("imgServer").Key("Bucket").String()
}

func LoadPath(file *ini.File) {
	LogPath = file.Section("path").Key("LogPath").MustString("log/")
}

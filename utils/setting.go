package utils

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	Dbname     string
	JwtKey     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置错误")
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("buisw678")

}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("123456qq")
	Dbname = file.Section("database").Key("Dbname").MustString("ginblog")

}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("GHa7ug7KtGzYS52A_Hn-GmdAop80SwIeVCjsKeTK")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("GGwa6ELilRYQyUITcny3JMWtZnC3AN_whhNKLZQR")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("haolinfeng")
	QiniuServer = file.Section("qiniu").Key("QiniuServer").MustString("http://rio8rsdtv.hd-bkt.clouddn.com/")

}

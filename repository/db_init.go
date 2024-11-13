package repository

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 声明的全局变量
var db *gorm.DB

// InitConfig 读取配置文件
func InitConfig() {
	viper.SetConfigName("config") // 配置文件名称（不带扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

// Init 初始化数据库
func Init() error {
	InitConfig() // 加载配置文件

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.dbname"),
		viper.GetString("database.charset"),
		viper.GetBool("database.parseTime"),
		viper.GetString("database.loc"),
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

// 无配置文件的初始
func Init0() error {
	var err error
	dsn := "root:1234@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

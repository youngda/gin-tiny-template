package modle

import (
	"api/logger"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

// DB 数据库句柄
var DB *sqlx.DB

// InitMysql 初始化MySQL连接
func InitMysql() (err error){

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
		true,
		"Local",
		)
	// sqlx.Connect() 底层做了Open和Ping
	if DB, err = sqlx.Connect("mysql", config); err != nil {
		logger.Log.Fatal("sqlx.Connect failed, err:",zap.Error(err))

		return
	}

	// 设置MySQL相关配置

	//超时设置
	DB.SetConnMaxLifetime(time.Second * 500)
	//空闲最大连接数
	DB.SetMaxIdleConns(1)
	//数据库的最大连接数
	DB.SetMaxOpenConns(20)
	logger.Log.Info("mysql connect success......")
	return
}


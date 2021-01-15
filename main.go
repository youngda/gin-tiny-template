package main

import (
	"api/logger"
	"api/modle"
	"errors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"

	"api/config"
	"api/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "goapi config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {

	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}


	if err:=modle.InitMysql();err!=nil{
		logger.Log.Error("setting db failed,err:",zap.Error(err))
		return
	}


	// Set gin mode.
		gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			logger.Log.Info(err.Error())
		}
		logger.Log.Info("The router has been deployed successfully.")

	}()
	logger.Log.Info("The router has been deployed successfully.")
	logger.Log.Info("ok",zap.String("Start to listening the incoming requests on http address:",viper.GetString("addr")), )
	logger.Log.Error("error",zap.String("value",http.ListenAndServe(viper.GetString("addr"), g).Error()))

}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

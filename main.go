package main

import (
	"github.com/JumpSama/aug-blog/config"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/pkg/logger"
	"github.com/JumpSama/aug-blog/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
)

var cfg = pflag.StringP("config", "c", "", "aug-blog config file path.")

func main() {
	pflag.Parse()

	// init log
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init database
	model.DB.Init()
	defer model.DB.Close()

	// init logger
	logger.Logger.Init()
	defer logger.Logger.Close()

	// set gin mode
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	router.Load(g)

	port := viper.GetString("addr")

	logger.Logger.Sugar.Infof("Start to listening the incoming requests on http address: localhost%s", port)

	logger.Logger.Sugar.Info(http.ListenAndServe(port, g).Error())
}

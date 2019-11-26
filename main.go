package main

import (
	"errors"
	"github.com/JumpSama/aug-blog/config"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var cfg = pflag.StringP("config", "c", "", "aug-blog config file path.")

func pingServer() error {
	url := viper.GetString("url")
	maxCount := viper.GetInt("max_ping_count")

	for i := 0; i < maxCount; i++ {
		resp, err := http.Get(url + "/sd/health")

		if err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}

func main() {
	pflag.Parse()

	// init log
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init database
	model.DB.Init()
	defer model.DB.Close()

	// set gin mode
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(g, middlewares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	port := viper.GetString("addr")

	log.Infof("Start to listening the incoming requests on http address: %s", port)

	log.Info(http.ListenAndServe(port, g).Error())
}

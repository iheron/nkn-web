package main

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
	"nkn-web/src/config"
	//"nkn-web/src/routes/home"
	"os"
)

var _VERSION_ = "unknown"

func main() {
	cmd := cli.NewApp()
	cmd.Name = "nkn-web"
	cmd.Version = config.Parameters.Version
	cmd.HelpName = "nkn-web"
	cmd.Usage = "command line tool for blockchain"
	cmd.UsageText = "nknc [global options] command [command options] [args]"
	cmd.HideHelp = false
	cmd.HideVersion = false
	cmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config",
			Usage:       "config path",
			Value:       "./config.json",
			Destination: &config.ConfigPath,
		},
	}

	cmd.Action = func(c *cli.Context) error {

		config.Init()

		return nil
	}

	cmd.Run(os.Args)
	f, _ := os.Create(config.Parameters.LogFile)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	app := gin.Default()

	//app.Static("/assets", "./assets")
	app.StaticFS("/web", http.Dir("dist"))
	//_ = flag.String("config", "./config/config.json","config file")
	//flag.Parse()

	//app.Use(routerv1.Sign())
	//	app.Use(routerv1.Router(app.Group("/api/v1")))
	//(&home.HomeRouter{}).RouterByGroup(app.Group("/api/v1"))
	app.GET("/ping", func(c *gin.Context) {
		log.Println("http server start at : " + config.Parameters.Port)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	app.Use(func(context *gin.Context) {
		//err := context.MustGet("error")

		//if err != nil {
		//	context.JSON(http.StatusInternalServerError, "server error")
		//	return
		//}

		context.JSON(http.StatusNotFound, "not found")
	})

	app.Run(":" + config.Parameters.Port) // listen and serve on 0.0.0.0:8080

}

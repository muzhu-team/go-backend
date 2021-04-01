package main

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/routers"
	_ "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	_ "github.com/gorilla/websocket"
	_ "golang.org/x/net/proxy"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}
func addNum() func(num int) int {
	add := func(num int) int {
		i := 1
		for a := 1; a <= num; a++ {
			i += a
		}
		return i
	}
	return add
}

func main() {
	addNumber := addNum()
	fmt.Println(addNumber(100), "add number")

	fmt.Println(addNumber(200), "add number")
	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}

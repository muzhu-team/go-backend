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
	"github.com/shirou/gopsutil/cpu"
	_ "golang.org/x/net/proxy"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func ticker(tick time.Duration) float64 {
	percent, _ := cpu.Percent(tick, false)
	return percent[0]
}

func main() {
	//
	//tick := 1 * time.Second
	//
	//// You can generate a Token from the "Tokens Tab" in the UI
	//const token = "k7eP_n6yW4JylHbKQUrxQ4nXOhRjmGYBm9zK4pE_DFBJaePoBLn4BuZiSOMtlr62yjLrHaPAGsRsm1i9jEMMNA=="
	//const bucket = "second"
	//const org = "qingdaowangningsoftware"
	//
	//client := influxdb2.NewClient("http://localhost:8086", token)
	//// always close client at the end
	////defer client.Close()
	//
	//defer client.Close()
	//// get non-blocking write client
	//writeAPI := client.WriteAPI(org, bucket)
	//
	//errorsCh := writeAPI.Errors()
	//// Create go proc for reading and logging errors
	//go func() {
	//	for err := range errorsCh {
	//		fmt.Printf("write error: %s\n", err.Error())
	//	}
	//}()
	//go func() {
	//	writeAPI := client.WriteAPI(org, bucket)
	//	for range time.Tick(1 * time.Second) {
	//		cpuPercent := ticker(tick)
	//
	//		p := influxdb2.NewPointWithMeasurement("cpu").
	//			AddTag("unit", "cputemp010000000000").
	//			AddField("cpuPercent", cpuPercent).
	//			SetTime(time.Now())
	//		fmt.Println(cpuPercent, "cputemp010000000000")
	//
	//		// write point asynchronously
	//		writeAPI.WritePoint(p)
	//		// Flush writes
	//		writeAPI.Flush()
	//	}
	//}()
	//go func() {
	//	writeAPI := client.WriteAPI(org, bucket)
	//	for range time.Tick(1 * time.Second) {
	//		cpuPercent := ticker(tick)
	//
	//		p := influxdb2.NewPointWithMeasurement("cpu").
	//			AddTag("unit", "cputemp11111").
	//			AddField("cpuPercent", cpuPercent).
	//			SetTime(time.Now())
	//		fmt.Println(cpuPercent, "cputemp11111")
	//
	//		// write point asynchronously
	//		writeAPI.WritePoint(p)
	//		// Flush writes
	//		writeAPI.Flush()
	//	}
	//}()
	//
	//go func() {
	//	writeAPI := client.WriteAPI(org, bucket)
	//	for range time.Tick(1 * time.Second) {
	//		cpuPercent := ticker(tick)
	//
	//		p := influxdb2.NewPointWithMeasurement("cpu").
	//			AddTag("unit", "cpuPercent2222222").
	//			AddField("cpuPercent", cpuPercent).
	//			SetTime(time.Now())
	//		fmt.Println(cpuPercent, "cpuPercent2222222")
	//
	//		// write point asynchronously
	//		writeAPI.WritePoint(p)
	//		// Flush writes
	//		writeAPI.Flush()
	//	}
	//}()
	//
	//go func() {
	//	writeAPI := client.WriteAPI(org, bucket)
	//	for range time.Tick(1 * time.Second) {
	//		cpuPercent := ticker(tick)
	//
	//		p := influxdb2.NewPointWithMeasurement("cpu").
	//			AddTag("unit", "cputemp333333").
	//			AddField("cpuPercent", cpuPercent).
	//			SetTime(time.Now())
	//		fmt.Println(cpuPercent, "cputemp333333")
	//
	//		// write point asynchronously
	//		writeAPI.WritePoint(p)
	//		// Flush writes
	//		writeAPI.Flush()
	//	}
	//}()
	//
	//
	//
	//
	//
	//query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h) |> filter(fn: (r) => r._measurement == \"stat\")", bucket)
	//// Get query client
	//queryAPI := client.QueryAPI(org)
	//// get QueryTableResult
	//result, err := queryAPI.Query(context.Background(), query)
	//if err == nil {
	//	// Iterate over query response
	//	for result.Next() {
	//		// Notice when group key has changed
	//		if result.TableChanged() {
	//			fmt.Printf("table: %s\n", result.TableMetadata().String())
	//		}
	//		// Access data
	//		fmt.Printf("value: %v\n", result.Record().Value())
	//	}
	//	// check for an error
	//	if result.Err() != nil {
	//		fmt.Printf("query parsing error: %\n", result.Err().Error())
	//	}
	//} else {
	//	panic(err)
	//}
	//
	//time.Sleep(time.Hour*1)

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

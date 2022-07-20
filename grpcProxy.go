package main

//
//import (
//	"flag"
//	"fmt"
//	"go-common/library/naming/discovery"
//	"google.golang.org/grpc/encoding"
//	"net/http"
//	"strconv"
//	"strings"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//)
//
//var (
//	color string
//)
//
//// Reply ...
//type Reply struct {
//	data []byte
//}
//
//// JSON is impl of encoding.Codec
//type JSON struct {
//}
//
//// Name is name of JSON
//func (j JSON) Name() string {
//	return "json"
//}
//
//// Marshal is json marshal
//func (j JSON) Marshal(v interface{}) (out []byte, err error) {
//	return v.([]byte), nil
//}
//
//// Unmarshal is json unmarshal
//func (j JSON) Unmarshal(data []byte, v interface{}) (err error) {
//	v.(*Reply).data = data
//	return nil
//}
//
//func ipFromDiscovery(appid string, color string) (ip string, err error) {
//	d := discovery.New(nil)
//	defer d.Close()
//	b := d.Build(appid)
//	defer b.Close()
//	ch := b.Watch()
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
//	defer cancel()
//	select {
//	case <-ch:
//	case <-ctx.Done():
//		err = fmt.Errorf("查找节点超时，请检查appid是否填写正确")
//		return
//	}
//	ins, ok := b.Fetch(context.Background())
//	if !ok {
//		err = fmt.Errorf("discovery 拉取 %s 失败", appid)
//		return
//	}
//	for _, vs := range ins {
//		for _, v := range vs {
//			if v.Metadata["color"] == color {
//				for _, addr := range v.Addrs {
//					if strings.Contains(addr, "grpc://") && v.Zone == "sh001" {
//						ip = strings.Replace(addr, "grpc://", "", -1)
//						return
//					}
//				}
//			}
//		}
//	}
//	err = fmt.Errorf("discovery 找不到服务节点: %s color: %s", appid, color)
//	return
//}
//
//func main() {
//	encoding.RegisterCodec(JSON{})
//	gin.SetMode(gin.ReleaseMode)
//	router := gin.Default()
//	router.GET("/grpc_proxy", handle)
//	router.Run(":1001")
//}
//
//func handle(c *gin.Context) {
//	env, _ := strconv.Atoi(c.Query("env"))
//	service := c.Query("service")
//	path := c.Query("path")
//	data := c.Query("data")
//	// fmt.Println("data: ", data)
//	flag.Set("appid", "live.grpc-proxy")
//	flag.Set("deploy.env", "uat")
//	switch env {
//	case 0:
//		color = ""
//	case 6:
//		color = "white-live-fat6"
//	case 7:
//		color = "black-live-fat7"
//	case 8:
//		color = "brown-live-fat8"
//	case 9:
//		color = "cool-live-fat9"
//	case 10:
//		color = "awesome-live-fat10"
//	case 11:
//		color = "coffee-live-fat11"
//	case 12:
//		color = "kole-live-fat12"
//	case 13:
//		color = "fenda-live-fat13"
//	case 14:
//		color = "suda-live-fat14"
//	case 15:
//		color = "suntory-live-fat15"
//	case 16:
//		color = "live-live-fat16"
//	case 17:
//		color = "live-live-fat17"
//	case 18:
//		color = "live-live-fat18"
//	case 21:
//		color = "live-test-1"
//	case 22:
//		color = "live-bachi"
//	case 23:
//		color = "live-didi"
//	case 24:
//		color = "live-kesheng"
//	case 25:
//		color = "live-jiaqi"
//	case 26:
//		color = "live-andy"
//	case 27:
//		color = "live-gigi"
//	case 28:
//		color = "live-cooper"
//	case 29:
//		color = "live-meiying"
//	case 30:
//		color = "live-test-10"
//	case 31:
//		color = "live-test-11"
//	case 32:
//		color = "live-test-12"
//	}
//	if color != "" {
//		flag.Set("deploy.color", color)
//	}
//	flag.Parse()
//	addr, err := ipFromDiscovery(service, color)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	conn, dialErr := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.CallContentSubtype("json")))
//	defer func() {
//		conn.Close()
//		if err := recover(); err != nil {
//			fmt.Printf("recover: %v", err)
//		}
//	}()
//	if dialErr != nil {
//		c.JSON(200, gin.H{
//			"error": dialErr.Error(),
//		})
//		return
//	}
//	var resp Reply
//	rpcErr := conn.Invoke(c, path, []byte(data), &resp)
//	if rpcErr != nil {
//		//e := make(map[string]string)
//		//e["error"] = rpcErr.Error()
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code":  -1,
//			"error": rpcErr.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, string(resp.data))
//		fmt.Println(string(resp.data))
//	}
//}

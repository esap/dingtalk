package main

import (
	"fmt"
	"time"

	"github.com/esap/dingtalk"
)

var (
	token   = ""
	appid   = ""
	appSec  = "14OuUUzupJnd9qYVfqwTdZ5K09_ruDW4Oe2NexXjEB_dZIYpS994HnRH7Iex1QlR"
	appKey  = "dingtbpdimiyyucwqnyg"
	agentid = 268326543
	webhook = "https://oapi.dingtalk.com/robot/send?access_token=e1cdf5bb20ca6ed9fb078adccfe1d3f15a68217580bc8b068aefe816e28b01a3"
)

func main() {

	dingtalk.Debug = true
	app := dingtalk.New(token, appKey, appSec, "", agentid)
	// fmt.Println(app.GetAccessToken())
	fmt.Println("simplelist")
	fmt.Println(app.GetUserSimpleList())
	time.Sleep(1 * time.Second)

}

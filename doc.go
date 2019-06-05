/*
Package dingtalk provide dingtalk-sdk for go

5行代码，开启微信API示例:

	package main

	import (
		"net/http"
		"github.com/esap/dingtalk" // 微信SDK包
	)

	func main() {
		dingtalk.Debug = true
		dingtalk.Set("yourToken", "yourAppID", "yourSecret", "yourEncodingAesKey")
		http.HandleFunc("/", WxHandler)
		http.ListenAndServe(":9090", nil)
	}

	func WxHandler(w http.ResponseWriter, r *http.Request) {
		dingtalk.VerifyURL(w, r).NewText("客服消息1").Send().NewText("客服消息2").Send().NewText("查询结果...").Reply()
	}

More info: https://github.com/esap/dingtalk

*/
package dingtalk

package main

import (
	"bytes"
	"fmt"
	gateway "github.com/rpcxio/rpcx-gateway"
	"github.com/smallnest/rpcx/codec"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	cc := &codec.MsgpackCodec{}

	args := &Args{
		A: 100,
		B: 200,
	}

	data, _ := cc.Encode(args)
	// request
	url := fmt.Sprintf("http://127.0.0.1:%d/", 8086)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		log.Printf("failed to create request: %s", err.Error())
		return
	}

	// 设置header
	h := req.Header
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")
	h.Set(gateway.XServicePath, "Arith")
	h.Set(gateway.XServiceMethod, "Mul")

	// 发送http请求
	//  http请求===>rpcx请求===>rpcx服务===>返回rpcx结果===>转换为http的response===>输出到client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to call: %s", err.Error())
	}
	defer res.Body.Close()
	// 获取结果
	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read response: %s", err.Error())
	}
	// 解码
	reply := &Reply{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Printf("failed to decode reply: %s", err.Error())
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}

package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var G_lineBotclient *linebot.Client

func main() {
	var(
		line_channel_secret string
		line_channel_token string
		httpClient *http.Client
		err error
	)

	log.Println("server starthe")

	line_channel_secret = "c31afe4242e7247a5a0c59a9e8027655"
	line_channel_token = "fO18hoQ45hl6YSvHCI/NY8uuEE3OBRvWRmoi+h6+FOOfNtJcLTv+OA49er/GYC0pZyS4fleFfgz87xR5ZFbyDDNn/9g2oHYfrMXeGOy706B3/zd4K43v6Bf+h/3VpIsOlLZU0tih7lRF5AusGMLPYwdB04t89/1O/w1cDnyilFU="


	httpClient = &http.Client{}
	if G_lineBotclient, err = linebot.New(line_channel_secret, line_channel_token, linebot.WithHTTPClient(httpClient)); err != nil{
		log.Fatalf("failed to initial line bot: %v", err)
		return
	}

	InitHttpServer()
}

func handleEvent(resp http.ResponseWriter, req *http.Request){
	var(
		events []*linebot.Event
		event *linebot.Event
		err error
	)
	fmt.Println("handleEvent launched:", req)

	events = make([]*linebot.Event,0)

	if events, err = G_lineBotclient.ParseRequest(req); err!=nil{
		log.Fatalf("failed to parse request: %v", err)
	}

	for _, event = range events {
		if event.Type == linebot.EventTypeMessage{
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				G_lineBotclient.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
			}
		}
	}

	resp.Write([]byte("OK"))
}

func InitHttpServer(){
	var(
		err error
		listener net.Listener
		mux *http.ServeMux
		httpSvr *http.Server
		port string
	)

	port = os.Getenv("PORT")
	if port == ""{
		port = "80"
	}

	if listener, err = net.Listen("tcp",":"+port); err!=nil{
		return
	}

	fmt.Println("port:", port)
	mux = http.NewServeMux()
	mux.HandleFunc("/callback",handleEvent)

	httpSvr = &http.Server{
		ReadTimeout: 5000 * time.Millisecond,
		WriteTimeout: 5000 * time.Millisecond,
		Handler: mux,
	}
	httpSvr.Serve(listener)
}
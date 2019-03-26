package events

import (
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/weizhe0422/TrafficInfoHelper/dataBean"
	"github.com/weizhe0422/TrafficInfoHelper/httpClient"
	"github.com/weizhe0422/TrafficInfoHelper/resource"
	"log"
	"strings"
)

func EventTHSR(client *linebot.Client, event *linebot.Event) {
	var (
		stationInfo *Station
		stations    []dataBean.THSRStationInfo
		err         error
		replyMsg    strings.Builder
	)
	log.Println("enter into EventTHSR")

	stationInfo = InitStation()
	if stations, err = stationInfo.GetStations(); err != nil {
		return
	}
	log.Println("stations:", stations)

	for idx, station := range stations {
		replyMsg.WriteString(station[idx].StationName.ZhTw)
	}
	client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg.String())).Do()
}

type Station struct {
	StationName string
	StationID   int
}

func InitStation() *Station {
	return new(Station)
}

func (s *Station) GetStations() ([]dataBean.THSRStationInfo, error) {
	var (
		client      *httpClient.Client
		respContent []byte
		err         error
		jsonResult  []dataBean.THSRStationInfo
	)

	client = httpClient.InitHttpClient(resource.URLTHSRStationInfo)
	if respContent, err = client.GetHttpResp(); err != nil {
		log.Printf("failed to get response from %v, error is %v", resource.URLTHSRStationInfo, err)
		return nil, err
	}

	if err = json.Unmarshal(respContent, &jsonResult); err != nil {
		log.Printf("failed to unmarshall response: %v", err)
		return nil, err
	}

	return jsonResult, nil
}

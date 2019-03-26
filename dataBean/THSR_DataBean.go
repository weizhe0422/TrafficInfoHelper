package dataBean

import "time"

type THSRStationInfo struct {
	StationUID  string `json:"StationUID"`
	StationID   string `json:"StationID"`
	StationName struct {
		ZhTw string `json:"Zh_tw"`
		En   string `json:"En"`
	} `json:"StationName"`
	StationPosition struct {
		PositionLat float64 `json:"PositionLat"`
		PositionLon float64 `json:"PositionLon"`
	} `json:"StationPosition"`
	StationAddress string    `json:"StationAddress"`
	OperatorID     string    `json:"OperatorID"`
	UpdateTime     time.Time `json:"UpdateTime"`
	VersionID      int       `json:"VersionID"`
}

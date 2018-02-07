package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type requestInfo struct {
	Nonce string
	AppId string
	Sign  string
}

type CarInfo struct {
	CarNo     string
	MaxVolume float64
	MaxWeight float64
}

type Waybill struct {
	WaybillNumber  string
	ActualWeight   float64
	ActualVolume   float64
	FreightCharges float64
	PackageNumber  int
	Necessary      string
	UnderStowed    string
	Split          string
	OtherInfo      string
	Result         string
}

type ReqCalculate struct {
	Info      requestInfo
	OrderNo   string
	CalType   string
	NotifyUrl string
	Cars      []CarInfo
	GoodsList []Waybill
}

const (
	APPID     = "1234567890"
	APPKEY    = "62324db02faefc973dcf976f317bbacf"
	NOTIFYURL = "http://abc.com"
)

func gen_req_info() requestInfo {
	reqInfo := requestInfo{
		AppId: APPID,
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
	}
	signStr := "AppId=" + reqInfo.AppId +
		"&Nonce=" + reqInfo.Nonce +
		"AppKey=" + APPKEY
	reqInfo.Sign = strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(signStr))))
	return reqInfo
}

func main() {
	reqBody := ReqCalculate{
		Info:      gen_req_info(),
		OrderNo:   "",
		CalType:   "moneyOpt",
		NotifyUrl: NOTIFYURL,
		Cars:      nil,
		GoodsList: nil,
	}
	var body string
	byteBody, e := json.Marshal(reqBody)
	if e != nil {
		fmt.Println(e)
	} else {
		body = string(byteBody)
		fmt.Println(body)
	}
	return
}

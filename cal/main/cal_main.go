package main

import (
	"encoding/json"
	"fmt"
	"note_code/cal"
	"time"
)

func main() {
	//fmt.Println(time.Now(),"operation begin")
	//for {
	//	Operation()
	//	time.Sleep(time.Second*3)
	//}
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		cal.Produce("123456789")
	}
}
func Operation() {
	req := cal.ConsumerDto{}
	resp := cal.ProducerDto{}
	body := cal.Consume()
	json.Unmarshal([]byte(body), &req)
	if len(req.Cars) == 0 || len(req.GoodsList) == 0 {
		return
	}
	for _, v := range req.Cars {
		resp.CarSummary = append(resp.CarSummary, cal.MQCarSummary{
			CarNo:       v.CarNo,
			TotalMoney:  0,
			TotalWeight: 0,
			TotalVolume: 0,
		})
	}
	car := resp.CarSummary[0]
	for _, v := range req.GoodsList {
		if v.Ne == "true" || v.Uns != "" {
			resp.Result = append(resp.Result, cal.MQCalResult{
				Id:     v.Id,
				Result: car.CarNo,
			})
			car.TotalVolume += v.Av
			car.TotalWeight += v.Aw
			car.TotalMoney += v.Fc
		} else {
			resp.Result = append(resp.Result, cal.MQCalResult{
				Id:     v.Id,
				Result: "",
			})
		}
	}
	resp.CarSummary[0] = car
	resp.CalTimes = req.CalTimes
	resp.UsingId = req.UsingId
	resp.ErrorCode = 0

	js, err := json.Marshal(&resp)
	if err != nil {
		fmt.Println("resp json marshal 出错")
	}
	cal.Produce(string(js))
}

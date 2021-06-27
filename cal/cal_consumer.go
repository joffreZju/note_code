package cal

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

func Consume() string{
	Topic:="gotest"
	URL:="http://publictest-rest.ons.aliyun.com"
	Ak:="ak"
	Sk:="sk"
	ConsumerID:="CID_goo"
	newline := "\n"

	for{
		date := fmt.Sprintf("%d",time.Now().UnixNano())[0:13]
		signStr := Topic+newline+ConsumerID+newline+date
		sign := calSigh(signStr,Sk)
		client := &http.Client{}
		req, err := http.NewRequest("GET",URL+"/message/?topic="+Topic+"&time="+date+"&num=2",nil)
		if err != nil {
			fmt.Println(time.Now(),"cal_consumer:get req构造出错")
		}

		req.Header.Set("Signature",sign)
		req.Header.Set("AccessKey",Ak)
		req.Header.Set("ConsumerID",ConsumerID)

		resp, err := client.Do(req)
		if err!= nil {
			fmt.Println(time.Now(),"cal_consumer:get req出错")
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		var msgs []Msg
		//fmt.Println(string(body))
		json.Unmarshal(body,&msgs)

		for _,v := range msgs{
			//fmt.Println(v.Body)
			date_ := fmt.Sprintf("%d",time.Now().UnixNano())[0:13]
			delUrl := URL + "/message/?msgHandle="+ v.MsgHandle + "&topic="+Topic+"&time="+ date_
			signStr := Topic+newline + ConsumerID +newline + v.MsgHandle +newline + date_
			sign := calSigh(signStr,Sk)
			req,err := http.NewRequest(http.MethodDelete,delUrl,nil)
			if err!=nil{
				fmt.Println(time.Now(),"cal_consumer:delete请求构造出错")
			}

			req.Header.Set("Signature",sign)
			req.Header.Set("AccessKey",Ak)
			req.Header.Set("ConsumerID",ConsumerID)

			resp,err := client.Do(req)
			if err!=nil{
				fmt.Println(time.Now(),"cal_consumer:delete请求出错")
			}

			fmt.Println(time.Now(),"cal_consumer:",resp.Status,"删除成功")
			return v.Body
		}
	}
}


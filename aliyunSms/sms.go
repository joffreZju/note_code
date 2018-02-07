package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-sms-go-sdk/sms"
)

// modify it to yours
const (
	ENDPOINT  = "https://sms.aliyuncs.com/"
	ACCESSID  = "LTAIwxFn7egYfvra"
	ACCESSKEY = "nBfpqo4StRZv9JreRsLQpFaZKKUT1h"
)

func main() {
	sms.HttpDebugEnable = true
	c := sms.New(ACCESSID, ACCESSKEY)
	// send to one person
	for i := 0; i < 5; i++ {
		e, err := c.SendOne("13777367114", "壹算科技", "SMS_43055008", `{"customer":"车燃燃"}`)
		if err != nil {
			fmt.Println("send sms failed", err, e.Error())
			//os.Exit(0)
			return
		}
	}
	// send to more than one person
	//e, err = c.SendMulti([]string{"13735544671", "13777367114"}, "壹算科技", "SMS_43055008", `{"customer":"wangjunfu"}`)
	//if err != nil {
	//	fmt.Println("send sms failed", err, e.Error())
	//	os.Exit(0)
	//}
	//fmt.Println("send sms succeed", e.GetRequestId())
}

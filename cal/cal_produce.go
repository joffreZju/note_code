package cal

import (
	"fmt"
	"net/http"
	"strings"
	"crypto/md5"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"time"
	"io/ioutil"
)
//define for MQ

// md5
func contentMD5(content string) string{
	byteContent := []byte(content)
	mdContent := md5.Sum(byteContent)
	s := fmt.Sprintf("%x",mdContent)
	return s
}

//calSign
func calSigh(signStr, sk string)string {
	mac := hmac.New(sha1.New, []byte(sk))
	mac.Write([]byte(signStr))
	s := base64.StdEncoding.EncodeToString([]byte(mac.Sum(nil)))
	strings.TrimRight(s, " ")
	return s
}

type Msg struct {
	Body      string
	MsgHandle string
}

func Produce(body string){
	Topic:="stowage_test"
	URL:="http://publictest-rest.ons.aliyun.com"
	Ak:="ak"
	Sk:="sk"
	ProducerID:="PID_calculation"

	for i:=0;i<1;i++{
		newline := "\n"
		content := contentMD5(body)
		date := fmt.Sprintf("%d",time.Now().UnixNano())[0:13]

		signStr := Topic+newline+ProducerID+newline+content+newline+date
		sign := calSigh(signStr,Sk)

		client := &http.Client{}

		req, err := http.NewRequest("POST",URL+"/message/?topic="+Topic+"&time="+date+"&tag=http"+"&key=http", strings.NewReader(body))
		if err != nil {
			fmt.Println(time.Now(),"cal_producer: req构造出错")
		}

		req.Header.Set("Signature",sign)
		req.Header.Set("AccessKey",Ak)
		req.Header.Set("ProducerID",ProducerID)
		req.Header.Set("Content-Type","text/html;charset=UTF-8")

		resp, err := client.Do(req)
		if err!= nil {
			fmt.Println(time.Now(),"cal_producer: req出错")
		}

		fmt.Println(time.Now(),"cal_producer: producer状态",resp.Status)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(time.Now(),"cal_producer: ",string(body))
	}

}




package main

import (
	"fmt"

	"encoding/json"
)

func main() {
	str := `{
  "Tpl":{
    "WaybillNumber": "aaaa",
      "ActualWeight": "bbbb",
      "ActualVolume": "vvvvv",
      "FreightCharges": "ffff",
      "PackageNumber": "eeeee"
  }
}`
	m := make(map[string]interface{})
	e := json.Unmarshal([]byte(str), &m)
	fmt.Println(e, m)
	fmt.Println(str)

	//nonce := fmt.Sprintf("%d",time.Now().Unix())
	//str := "AppId=1234567890&Nonce=" + nonce + "&AppKey=62324db02faefc973dcf976f317bbacf"
	//s:=fmt.Sprintf("%x",sha256.Sum256([]byte(str)))
	//fmt.Println(nonce, strings.ToUpper(s))

	//s64 := base64.StdEncoding.EncodeToString([]byte("123"))
	//fmt.Println(s64)
	//d64, _ := base64.StdEncoding.DecodeString(s64)
	//fmt.Println(fmt.Sprintf("%s",d64))

	//cacheObject, err := cache.NewCache("memory",`{"interval":3}`)
	//if err != nil {
	//	beego.Critical(err)
	//}
	//cacheObject.Put("wjf",1,time.Second*5)
	//cacheObject.Put("wjf",2,time.Second*5)
	////time.Sleep(4*time.Second)
	//beego.Debug(cacheObject.Get("wjf"))

	//验证随机数会不会重复
	//set :=mapset.NewSet()
	//var sl []int64
	//start := time.Now()
	//fmt.Println(start)
	//
	//for i:=0;i<10000;i++{
	//	//time.Sleep(time.Second*3)
	//	a := time.Now().UnixNano()%999983
	//	if a < 100000{
	//		fmt.Println("----",a,"------")
	//		a += 100000
	//	}
	//	sl = append(sl,a)
	//	set.Add(a)
	//}
	//fmt.Println(sl)
	//fmt.Println(set.Cardinality(),len(sl))

}

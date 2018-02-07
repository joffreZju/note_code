package main

import (
	"github.com/astaxie/beego"
)

type DemoController struct {
	beego.Controller
}
type User struct {
	Id      int
	Name    string
	Sex     string
	ClassId int
}

type Car struct {
	CarNo  string
	Cubage int
	Load   int
}
type Goods struct {
	Code            string
	Actual_Weight   int
	Actual_Volume   int
	Freight_Charges int
	Send_Site_Name  string
	Package_Number  int
	Package_Type    string
	Goods_Name      string
	Customer_Name   string
	Necessary       string
	Understowed     string
}
type CalFirstReq struct {
	Key       string
	CalType   string
	OrderNo   string
	CarInfo   []Car
	GoodsList []Goods
	NotifyUrl string
}
type Hstore struct {
	Id   int
	Data string
}

func (c *DemoController) Cook() {
	//beego.BeeLogger
	//beego.Debug(utility.Instance().Get("wjf"))
	c.Ctx.SetCookie("wjf", "wjftestcookie", 1200, "/", 1, 1, true)
	h := Hstore{
		Id:   1,
		Data: "hahah",
	}
	//fmt.Println(c.Ctx.Request.Cookie("wjf"))
	//fmt.Println(c.Ctx.Input.Cookie("wjf"))
	//fmt.Println(c.Ctx.GetCookie("wjf"))
	c.Data["json"] = &h
	c.ServeJSON()
}

func main() {

	//测试newSetFromSlice() ，必须声明一个[]interface{}
	//h1 := Hstore{
	//	Id:   1,
	//	Data: "111",
	//}
	//h2 := Hstore{
	//	Id:   2,
	//	Data: "222",
	//}
	//var sh []interface{}
	//sh = append(sh, h1)
	//sh = append(sh, h2)
	//set := mapset.NewSetFromSlice(sh)
	//logs.Info(set)

	//mapset中存储结构体直接判断contains
	// set := mapset.NewSet()
	// set.Add(Hstore{
	//     Id:1,
	//     Data:"11111",
	// })
	//     set.Add(Hstore{
	//     Id:2,
	//     Data:"22222",
	// })
	// logs.Info(set.Contains(Hstore{
	//     Id:1,
	//     Data:"11111",
	// }))

	//测试set de add() Cardinality()
	// set.Add("hahhahahahahha")
	// set.Add("hahhaha")
	// set.Add("hahaha")
	// set.Add("hhaha")
	// logs.Info("--------------",set.Cardinality(),"1","2","3","4","5","6","7")

	//测试returning id
	// o := orm.NewOrm()
	// var id int = 100
	// o.Raw("insert into public.class (name) values ('哈哈哈哈') RETURNING ID").QueryRow(&id)
	// fmt.Println(id)

	// 多表查询的语法，以及重名字段的处理，返回到map里
	// var maps []orm.Params
	// sql := `SELECT u.id, u.name as uname, u.sex, c.name as cname
	// FROM public.user as u inner join public.class as c
	// on u.class_id = c.id;`
	// o.Raw(sql).Values(&maps)
	// logs.Info(maps)

	//查询不到或者数据库出错时 err!=nil
	// o:= orm.NewOrm()
	// sql := "SELECT * from public.user where id = ?"
	// var u User
	// err := o.Raw(sql,1).QueryRow(&u)
	// //b, _ := json.Marshal(u)
	//if err!=nil {
	//	logs.Info("haha")
	//}
	// logs.Info(err)

	//Hstore插入数据
	// o:= orm.NewOrm()
	//s:="7"+"=>"+fmt.Sprintf("%d",11111)
	//str := "测试"
	//str += "hstore"
	// sql := "insert into public.hstore_test (data,str) values (?::hstore,?) returning id"
	////不可以返回插入id
	//p,_ := o.Raw(sql).Prepare()
	//if r,err:=p.Exec(s,str);err!=nil {
	//	logs.Info(err)
	//}else {
	//	id,_ := r.LastInsertId()
	//	logs.Info(id)
	//}
	////如下方式可以获取插入数据库时生成的id
	////var id int
	////_ = o.Raw(sql,s,str).QueryRow(&id)
	////logs.Info(id)

	//hstore的 select
	//o := orm.NewOrm()
	//sql := "SELECT id,data->'1' as one from public.hstore_test where id = 1"
	//var maps []orm.Params
	//o.Raw(sql).Values(&maps)
	//var i int
	//i,_ = strconv.Atoi(fmt.Sprintf("%s",maps[0]["id"]))
	//logs.Info(reflect.TypeOf(i))

	//先select 到结构体，再for循环到 []interface{} 中，然后建立mapset
	//o := orm.NewOrm()
	//var h []Hstore
	//sql := "SELECT id,data-> ? as data from public.hstore_test where id = 1"
	//str := fmt.Sprintf("%d",1)
	//o.Raw(sql,str).QueryRows(&h)
	//var i []interface{}
	//for _,v := range h{
	//	i = append(i,v)
	//}
	//_ = sql
	//hset := mapset.NewSetFromSlice(i)
	//fmt.Println(hset)

	//测试临时结构体
	//o:= orm.NewOrm()
	//h := struct {
	//	Id int
	//	Str_str string
	//	Hah int
	//}{}
	//sql := `select id,str_str from public.hstore_test where id = ?`
	//_ = o.Raw(sql,
	//4).QueryRow(&h)
	//fmt.Println(h)
	//fmt.Println(h.Str_str,"haha")

	//测试 sql count，测试一模一样的sql语句是否会有缓存导致数据不真实
	//o:= orm.NewOrm()
	//var a int
	//sql := `select count(*) from public.hstore_test`
	//_ = o.Raw(sql).QueryRow(&a)
	//fmt.Println(a)
	//time.Sleep(3*time.Second)
	//_ = o.Raw(sql).QueryRow(&a)
	//fmt.Println(a)

	//hstore的 update
	//o:=orm.NewOrm()
	//_ = o.Begin()
	//s := "7=>"
	//con := "444"
	//sql := `update public.hstore_test set data=data||(?) where id = ?`
	//_,err := o.Raw(sql,s+con,4).Exec()
	//con = "555"
	//_,err = o.Raw(sql,s+con,5).Exec()
	//_ = o.Commit()
	//if err==nil{
	//	logs.Info("true")
	//}else{
	//	logs.Info("false")
	//}

	//测试json list，
	// unmarshal：json字符串里面有而结构体没有的字段，会忽略；
	//marshal：结构体里面int string 等字段会有默认值
	//js := []byte(`[
	//         {
	//             "carNo": null,
	//             "cubage": 135,
	//             "load": null,
	//             "haha":123
	//         },
	//         {
	//             "carNo": "赣C24985",
	//             "cubage": 115,
	//             "load": 25
	//         }
	//     ]`)
	//var cars []Car
	//json.Unmarshal(js,&cars)
	//fmt.Println(cars)
	//for _,v := range cars{
	//	fmt.Println(v.CarNo,"haha","hhh")
	//}

	//结构体 数组字段会赋值null
	//data := CalFirstReq{
	//	Key:"124",
	//	OrderNo:"46",
	//}
	//byt,_ := json.Marshal(&data)
	//fmt.Println(string(byt))

	//js := []byte(`{}`)
	//var ca Car
	//_ = json.Unmarshal(js,&ca)
	//if ca.CarNo == nil{
	//	fmt.Println("nil")
	//}

	//json.Marshal() 会给 int 类型的字段赋默认值
	//h := Hstore{
	//	Id:1,
	//	//Data:"123",
	//}
	//h1 := Hstore{
	//	Data:"123",
	//}
	//var hs []Hstore
	//hs = append(hs,h)
	//hs = append(hs,h1)
	//byt,_ := json.Marshal(hs)
	//fmt.Println(string(byt))

	//切片赋值
	//var ps []person
	//p1 := person{
	//	id :1,
	//	name:"wang",
	//}
	//p2:=person{
	//	name:"liu",
	//}
	//ps=append(ps,p1)
	//ps=append(ps,p2)
	//var p []person
	//p= ps
	//fmt.Println(p)
	//修改结构体切片内容
	//for k,_ := range ps{
	//	ps[k].id = 3
	//}
	//fmt.Println(ps)

	//测试strings.Contains()
	//var s1,s2 string
	//s1 = "wangjun"
	//s2 = "jun"
	//fmt.Println(s1,s2)
	//fmt.Println(strings.Contains(s1,s2))

}

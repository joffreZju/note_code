package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/deckarep/golang-set"
	_ "github.com/lib/pq"
	"time"
)

//Model Struct
type SpzCoupon struct {
	Id        int    `orm:"column(id);auto"`
	Number    int    `orm:"column(number)"`
	SecretKey string `orm:"column(secret_key)"`
	Status    int    `orm:"column(status)"`
}

func init() {
	orm.RegisterDataBase("default", "postgres", "user=user01 password=allsum123 host=rm-uf6q1kk0byn74g70zo.pg.rds.aliyuncs.com port=3432 dbname=db_test sslmode=disable")

	orm.RegisterModel(new(SpzCoupon))

	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	str := "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"
	skSet := mapset.NewSet()
	var skSlice []string
	i := 0
	for ; ; i++ {
		sk := ""
		for j := 0; j < 6; j++ {
			sk += string(str[time.Now().UnixNano()%33])
		}
		if !skSet.Contains(sk) {
			skSet.Add(sk)
			skSlice = append(skSlice, sk)
		}
		if skSet.Cardinality() == 1000000 {
			break
		}
	}
	fmt.Println(skSet.Cardinality(), len(skSlice))

	fmt.Println(time.Now())

	coupons := []SpzCoupon{}
	for i = 0; i < len(skSlice); i++ {
		coupons = append(coupons, SpzCoupon{
			Number:    i + 10000001,
			SecretKey: skSlice[i],
			Status:    0,
		})
	}
	fmt.Println(len(coupons))
	if o.Begin() == nil {
		o.InsertMulti(len(coupons)/100, coupons)
	}
	fmt.Println(o.Commit())

	fmt.Println(time.Now())
}

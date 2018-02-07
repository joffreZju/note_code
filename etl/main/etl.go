package main

import (
	"encoding/hex"
	"note_code/etl/model2"
	"note_code/etl/model3"
	"math/rand"
	"time"
	
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(5)
	
	var err error
	err = orm.RegisterDataBase("default", "postgres", "user=joffre password=allsumjoffre host=rm-uf6p5m3r3sat7tdxxo.pg.rds.aliyuncs.com port=3432 dbname=webv2pro sslmode=disable")
	logs.Debug(err)
	//orm.RegisterModel(new(model2.SpzUser), new(model2.SpzTemplate), new(model2.SpzCalRecord), new(model2.SpzCalLib), new(model2.SpzCalCarsummary))
	//err = orm.RunSyncdb("default", false, true)
	//logs.Debug(err)
	
	err = orm.RegisterDataBase("db3", "postgres", "user=allsum password=stowage@allsum,./ host=rm-uf6q1kk0byn74g70zo.pg.rds.aliyuncs.com port=3432 dbname=stowage sslmode=disable")
	logs.Debug(err)
	
	orm.RegisterModel(new(model3.User), new(model3.Account), new(model3.CalTemplate), new(model3.CalRecord), new(model3.CalGoods), new(model3.CarSummary))
	err = orm.RunSyncdb("db3", false, true)
	logs.Debug(err)
}

func randomByte16() string {
	var code = make([]byte, 16)
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		code[i] = byte(r.Intn(255))
	}
	return hex.EncodeToString(code)
}

func handleEveryUser(o orm.Ormer, u2 *model2.SpzUser) (e error) {
	//转换user数据
	u3 := &model3.User{
		Tel:        u2.Mobile,
		Password:   u2.Password,
		UserName:   u2.UserName,
		CreateTime: u2.Ctt,
		UserType:   u2.Role,
	}
	e = model3.CreateUser(o, u3)
	if e != nil {
		logs.Debug("insert user", e, u3.Tel)
		o.Rollback()
		return e
	}
	//同时建立账户
	acc3 := &model3.Account{
		AccountNo: randomByte16(),
		Userid:    u3.Id,
		UserType:  u3.UserType,
		Status:    1,
		Banlance:  20000,
	}
	e = model3.AddAccount(o, acc3)
	if e != nil {
		logs.Debug("insert account:", e, u3.Tel)
		o.Rollback()
		return e
	}
	//转移用户匹配模板
	tpls2 := model2.GetTplsOfUser(u2.Id)
	tpls3 := []*model3.CalTemplate{}
	for _, v := range tpls2 {
		tpls3 = append(tpls3, &model3.CalTemplate{
			UserId:         u3.Id,
			WaybillNumber:  v.WaybillNumber,
			ActualVolume:   v.ActualVolume,
			ActualWeight:   v.ActualWeight,
			PackageNumber:  v.PackageNumber,
			FreightCharges: v.FreightCharges,
			Ctt:            v.Ctt,
		})
	}
	if len(tpls3) != 0 {
		_, e = o.InsertMulti(len(tpls3), tpls3)
		if e != nil {
			logs.Debug("insert tpls", e, u3.Tel)
			o.Rollback()
			return e
		}
	}
	//转移使用记录
	crs2 := model2.GetCalRecordsOfUser(u2.UserCode)
	for _, v := range crs2 {
		if e = handleEveryCalRecord(o, v, u3, acc3); e != nil {
			return e
		}
	}
	return
}
func handleEveryCalRecord(o orm.Ormer, cr2 *model2.SpzCalRecord, u3 *model3.User, acc3 *model3.Account) (e error) {
	cr3 := &model3.CalRecord{
		UserId:     u3.Id,
		AccountId:  acc3.Id,
		PayStatus:  4,
		CalNo:      cr2.OrderNo,
		CalTimes:   cr2.CalTimes,
		LastResult: cr2.LastResult,
		CalType:    cr2.CalType,
		Ctt:        cr2.Ctt,
		Ltt:        cr2.Ctt,
		Utt:        cr2.Utt,
	}
	e = model3.InsertCalRecord(o, cr3)
	if e != nil {
		logs.Debug("insert calrecord", e, u3.Tel)
		o.Rollback()
		return e
	}
	//处理goods
	goods2 := model2.GetGoodsOfCalRecord(cr2.Id)
	goods3 := []*model3.CalGoods{}
	for _, v := range goods2 {
		goods3 = append(goods3, &model3.CalGoods{
			CalRecordId:    cr3.Id,
			CalTimes:       v.CalTimes,
			WaybillNumber:  v.WaybillNumber,
			ActualVolume:   v.ActualVolume,
			ActualWeight:   v.ActualWeight,
			FreightCharges: v.FreightCharges,
			PackageNumber:  v.PackageNumber,
			Necessary:      v.Necessary,
			Understowed:    v.Understowed,
			OtherInfo:      v.OtherInfo,
			Split:          v.Split,
			SplitInfo:      v.SplitInfo,
			CalResult:      v.CalResult,
			Ctt:            v.Ctt,
			Utt:            v.Utt,
		})
	}
	if len(goods3) != 0 {
		_, e = o.InsertMulti(len(goods3), goods3)
		if e != nil {
			logs.Debug("insert goods", e, u3.Tel)
			o.Rollback()
			return e
		}
	}
	//处理car summary
	cars2 := model2.GetCarsOfCalRecord(cr2.Id)
	cars3 := []*model3.CarSummary{}
	for _, v := range cars2 {
		cars3 = append(cars3, &model3.CarSummary{
			CalRecordId:  cr3.Id,
			CalTimes:     v.CalTimes,
			UserId:       u3.Id,
			CarNo:        v.CarNo,
			MaxVolume:    v.MaxVolume,
			MaxWeight:    v.MaxWeight,
			TotalMoney:   v.TotalMoney,
			TotalVolume:  v.TotalVolume,
			TotalWeight:  v.TotalWeight,
			StowageRatio: v.StowageRatio,
			Ctt:          v.Ctt,
			Utt:          v.Utt,
		})
	}
	if len(cars3) != 0 {
		_, e = o.InsertMulti(len(cars3), cars3)
		if e != nil {
			logs.Debug("insert cars:", e, u3.Tel)
			o.Rollback()
			return e
		}
	}
	return
}

func readUsers() []*model2.SpzUser {
	o := orm.NewOrm()
	var users []*model2.SpzUser
	_, e := o.Raw(`select * from spz_user`).QueryRows(&users)
	if e != nil {
		return nil
	}
	return users
}

func main() {
	users := readUsers()
	o := orm.NewOrm()
	logs.Debug(o.Using("db3"))
	logs.Debug(o.Begin())
	for _, v := range users {
		if e := handleEveryUser(o, v); e != nil {
			logs.Debug("main:", e, v.Mobile)
			logs.Debug("main:", o.Rollback())
			return
		}
		logs.Debug("etl success", v.Mobile)
	}
	logs.Debug("db commit:", o.Commit())
}

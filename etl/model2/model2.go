package model2

import (
	"time"
	"github.com/astaxie/beego/orm"

)

type SpzUser struct {
	Id         int       `orm:"column(id);auto"`
	UserCode   string    `orm:"column(user_code);null"`
	Password   string    `orm:"column(password);null"`
	Mobile     string    `orm:"column(mobile);null"`
	UserName   string    `orm:"column(user_name);null"`
	AgentCode  string    `orm:"column(agent_code);null"`
	Ctt        time.Time `orm:"column(ctt);type(timestamp with time zone);null"`
	Utt        time.Time `orm:"column(utt);type(timestamp with time zone);null"`
	Role       int       `orm:"column(role);null"`
	UserStatus int       `orm:"column(user_status);null"`
}

type SpzTemplate struct {
	Id             int       `orm:"column(id);auto"`
	UserId         int       `orm:"column(user_id);null"`
	WaybillNumber  string    `orm:"column(waybill_number);null"`
	ActualWeight   string    `orm:"column(actual_weight);null"`
	ActualVolume   string    `orm:"column(actual_volume);null"`
	FreightCharges string    `orm:"column(freight_charges);null"`
	PackageNumber  string    `orm:"column(package_number);null"`
	Ctt            time.Time `orm:"column(ctt);type(timestamp with time zone);null"`
}

func GetTplsOfUser(uid int)(tpls []*SpzTemplate){
	tpls = []*SpzTemplate{}
	sql := `select * from spz_template where user_id = ?`
	_, e := orm.NewOrm().Raw(sql,uid).QueryRows(&tpls)
	if e != nil {
		return nil
	}
	return tpls
}

type SpzCalRecord struct {
	Id         int       `orm:"column(id);auto"`
	AccountId  int64     `orm:"column(account_id);null"`
	OrderNo    string    `orm:"column(order_no);null"`
	UsingCost  float64   `orm:"column(using_cost);null"`
	UserCode   string    `orm:"column(user_code);null"`
	CalTimes   int       `orm:"column(cal_times);null"`
	LastResult int       `orm:"column(last_result);null"`
	UserType   string    `orm:"column(user_type);null"`
	CalType    string    `orm:"column(cal_type);null"`
	NotifyUrl  string    `orm:"column(notify_url);null"`
	Ctt        time.Time `orm:"column(ctt);type(timestamp with time zone);null"`
	Utt        time.Time `orm:"column(utt);type(timestamp with time zone);null"`
}
func GetCalRecordsOfUser(uc string)(crs []*SpzCalRecord){
	crs = []*SpzCalRecord{}
	sql := `select * from spz_cal_record where user_code = ?`
	_, e := orm.NewOrm().Raw(sql,uc).QueryRows(&crs)
	if e != nil {
		return nil
	}
	return crs
}

type SpzCalLib struct {
	Id             int       `orm:"column(id);auto"`
	UsingId        int       `orm:"column(using_id);null"`
	WaybillNumber  string    `orm:"column(waybill_number);null"`
	ActualWeight   float64   `orm:"column(actual_weight);null"`
	ActualVolume   float64   `orm:"column(actual_volume);null"`
	FreightCharges float64   `orm:"column(freight_charges);null"`
	PackageNumber  int       `orm:"column(package_number);null"`
	Necessary      string    `orm:"column(necessary);null"`
	CalResult      string    `orm:"column(cal_result);null"`
	Understowed    string    `orm:"column(understowed);null"`
	CalTimes       int       `orm:"column(cal_times);null"`
	OtherInfo      string    `orm:"column(other_info);null"`
	Split          string    `orm:"column(split);null"`
	SplitInfo      string    `orm:"column(split_info);null"`
	Ctt            time.Time `orm:"column(ctt);type(timestamp with time zone);null"`
	Utt            time.Time `orm:"column(utt);type(timestamp with time zone);null"`
}

func GetGoodsOfCalRecord(usingId int)(goods []*SpzCalLib){
	goods = []*SpzCalLib{}
	sql := `select * from spz_cal_lib where using_id = ?`
	_, e := orm.NewOrm().Raw(sql, usingId).QueryRows(&goods)
	if e != nil {
		return nil
	}
	return goods
}

type SpzCalCarsummary struct {
	Id           int       `orm:"column(id);auto"`
	UsingId      int       `orm:"column(using_id);null"`
	CalTimes     int       `orm:"column(cal_times);null"`
	CarNo        string    `orm:"column(car_no);null"`
	TotalMoney   float64   `orm:"column(total_money);null"`
	TotalWeight  float64   `orm:"column(total_weight);null"`
	TotalVolume  float64   `orm:"column(total_volume);null"`
	Ctt          time.Time `orm:"column(ctt);type(timestamp with time zone);null"`
	MaxVolume    float64   `orm:"column(max_volume);null"`
	MaxWeight    float64   `orm:"column(max_weight);null"`
	Utt          time.Time `orm:"column(utt);type(timestamp with time zone);null"`
	UserCode     string    `orm:"column(user_code);null"`
	StowageRatio float64   `orm:"column(stowage_ratio);null"`
}
func GetCarsOfCalRecord(usingId int)(cars []*SpzCalCarsummary){
	cars = []*SpzCalCarsummary{}
	sql := `select * from spz_cal_carsummary where using_id = ?`
	_, e := orm.NewOrm().Raw(sql, usingId).QueryRows(&cars)
	if e != nil {
		return nil
	}
	return cars
}

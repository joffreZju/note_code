package model3

import "github.com/astaxie/beego/orm"

type Account struct {
	Id        int    `orm:"pk;auto"`
	AccountNo string `orm:"unique;size(64)" json:"-"` //
	Userid    int    `orm:"unique"`                   //个人和企业id
	UserType  int    //1 个人，2代理商，3企业
	Banlance  int64  //账户余额
	Topup     int64  //充值金额,消费者
	Spend     int64  //消费金额
	Withdraw  int64  //代理商提现
	Status    int    //1正常，2冻结
}

func AddAccount(o orm.Ormer, a *Account) (err error) {
	id, err := o.Insert(a)
	if err != nil {
		return
	}
	a.Id = int(id)
	return
}
//
//func CheckAccountExist(uid int, tp int) (b bool) {
//	count, err := orm.NewOrm().QueryTable("Account").Filter("Userid", uid).Filter("UserType", tp).Count()
//	if err == nil && count > 0 {
//		b = true
//	}
//	b = false
//	return
//}
//
//func GetAccountById(id int) (a *Account, err error) {
//	a = &Account{Id: id}
//	err = orm.NewOrm().Read(a)
//	return
//}
//
//func GetAccountByAccountNo(no string) (a *Account, err error) {
//	a = &Account{AccountNo: no}
//	err = orm.NewOrm().Read(a, "AccountNo")
//	return
//}
//
//func GetAccountByUserId(id int) (a *Account, err error) {
//	a = new(Account)
//	err = orm.NewOrm().QueryTable("Account").Filter("Userid", id).One(a)
//	return
//}
//
//func UpdateAccount(a *Account, fields ...string) (err error) {
//	_, err = orm.NewOrm().Update(a, fields...)
//	return
//}

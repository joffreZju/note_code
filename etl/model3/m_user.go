package model3

import (
	"time"
	"github.com/astaxie/beego/orm"
)
const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

type User struct {
	Id         int       `orm:"auto;pk;column(id)" ` // 用户ID，表内自增
	Unum       string    `orm:"null"  json:"-"`
	Tel        string    `orm:"unique;size(15)" json:",omitempty"`
	Password   string    `json:"-"`                              // 密码
	UserName   string    `orm:"null;size(64)" json:",omitempty"` // 用户名
	Icon       string    `orm:"null;size(64)" json:",omitempty"`
	Descp      string    `orm:"null" json:",omitempty"`
	Gender     int8      `orm:"null" json:",omitempty"`
	Address    string    `orm:"null;size(64)" json:",omitempty"`
	LoginTime  time.Time `orm:"type(datetime);null" json:"-"`                  //登录时间
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:",omitempty"` //
	Mail       string    `orm:"null;size(64)" json:",omitempty"`
	UserType   int       `orm:"default(1)"` //1 普通用户,2 代理商
	Referer    string    `orm:"null;size(16)" json:",omitempty"`
	AgentUid   int       `orm:"null" json:"-"`
	RegisterID string    `orm:"null;size(32)" json:",omitempty"` // 用于给用户推送消息
	//Groups     []*Group  `orm:"-" json:",omitempty"` // 用户的所在组织
}

func (u *User) TableName() string {
	return "allsum_user"
}

func CreateUser(o orm.Ormer, u *User) (err error) {
	id, err := o.Insert(u)
	if err != nil {
		return
	}
	u.Id = int(id)
	return
}
//
//func CreateOrUpdateAuser(u *User) (err error) {
//	u.CreateTime = time.Now()
//	o := orm.NewOrm()
//	us := new(User)
//	err = o.QueryTable("allsum_user").Filter("Tel", u.Tel).One(us)
//	if err == orm.ErrNoRows {
//		//insert
//		var id int64
//		id, err = o.Insert(u)
//		if err != nil {
//			return
//		}
//		u.Id = int(id)
//	} else if err == nil {
//		//update
//		us.UserType = 2
//		u.Id = us.Id
//		_, err = o.Update(us)
//	}
//	return
//	//_, id, err := orm.NewOrm().ReadOrCreate(u, "tel")
//	//if err != nil {
//	//	return
//	//}
//	//u.Id = int(id)
//	//return
//}

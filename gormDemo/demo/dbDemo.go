package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"note_code/gormDemo"
	"time"
)

type Student struct {
	Id    int
	SName string //`gorm:"default:'wjf'"`
	SId   int
	Ctime time.Time
	Slice gormDemo.IntSlice    `gorm:"type:int[]"`
	Js    gormDemo.MyInterface `gorm:"default:null;type:jsonb"`
}

func (Student) TableName() string {
	return "students"
}

func main() {
	dbuser := "allsum"
	dbpwd := "stowage@allsum,./"
	dbhost := "rm-uf6q1kk0byn74g70zo.pg.rds.aliyuncs.com"
	dbport := "3432"
	dbname := "asaccount"
	constr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbuser, dbpwd, dbhost, dbport, dbname)
	db, err := gorm.Open("postgres", constr)
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.LogMode(true)
	// 自动迁移模式bash
	schema := "group."
	//schema := "C0607145711618."
	fmt.Println(db.Table(schema + Student{}.TableName()).AutoMigrate(&Student{}).Error)

	//struct 找不到err=ErrRecordNotFound, slice 找不到err=nil
	//stu := Student{}
	//e := db.Table(schema+Student{}.TableName()).Find(&stu, "id=?", 50).Error
	//fmt.Println(e, stu)
	//stus := []Student{}
	//e = db.Table(schema+Student{}.TableName()).Find(&stus, "id>?", 50).Error
	//fmt.Println(e, stus)

	//limit order select 放在前面
	//stu := []Student{}
	//e := db.Table(schema+Student{}.TableName()).Select([]string{"s_id", "s_name"}).Limit(3).Order("id desc").
	//	Where(&Student{SName:"joffre"}).Where("id in (?)", []int{12, 13, 14, 15, 20}).Find(&stu).Error
	//fmt.Println(e)
	//for _, v := range stu {
	//	fmt.Println(v)
	//}

	//链接多个table函数最后一个有效,查询单列
	//tx := db.Table(schema+"group").Begin()
	//var ids []int
	//e := tx.Table(schema+"user").Where("id>?",2).Pluck("id", &ids).Error
	//fmt.Println(e, ids)

	//使用myInterface实现json存储
	//stu := &Student{}
	//stu.Js = gormDemo.MyInterface{
	//	"name": "wangjungfu",
	//	"score":[]int{1,2,3},
	//}
	//e := db.Table(schema+Student{}.TableName()).Create(stu).Error
	//if e != nil{
	//	fmt.Println(e)
	//}
	//stu := Student{}
	//e := db.Table(schema + stu.TableName()).First(&stu, 23).Error
	//if e != nil {
	//	fmt.Println(e)
	//}
	//fmt.Println(stu)

	//查询单列存到数组，可以使用table + where
	//ids := []int{}
	//sql := fmt.Sprint(`select distinct(s_id) from "group".students`)
	//db.Raw(sql).Pluck("id", &ids)
	//fmt.Println(ids)

	//HasTable()只对public下的表有效
	//b := db.HasTable(fmt.Sprintf(`"%s"."%s"`,schema, Student{}.TableName()))
	//fmt.Println(b)

	// 创建
	//m := map[string]interface{}{"name":"wjf","no":"12345"}
	//js, _ := json.Marshal(m)
	//stu := &Student{
	//	SId:21621179,
	//	SName:"joffre",
	//	Slice:gormDemo.NewIntSlice(1,2,3),
	//	Js:string(js),
	//}
	//fmt.Println(db.NewRecord(stu))
	//fmt.Println(stu.Id)
	//db.Table(schema+ Student{}.TableName()).Create(&stu)
	//fmt.Println(db.NewRecord(stu))
	//fmt.Println(stu)

	// 查询
	//us := []Student{}
	//db.Where("s_name <> ?","123").Find(&us)
	//fmt.Println(us)

	// 读取
	//a := 0
	//e := db.Table(schema+Student{}.TableName()).Where("id > ?",2 ).Count(&a).Error
	//fmt.Println(e,a )

	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//fmt.Println(product)

	//stu := Student{}
	//db.Table(schema+Student{}.TableName()).FirstOrInit(&stu, &Student{SName:"haha"}) //初始化结构体
	//db.Table(schema+Student{}.TableName()).FirstOrCreate(&stu, &Student{SName:"joffre"}) //直接插入数据库
	//fmt.Println(stu)

	//更新
	//stu := &Student{
	//	SId:2222,
	//	Slice:gormDemo.NewIntSlice(4,5,6,7),
	//}
	//tx := db.Begin()
	//e := tx.Table(schema + Student{}.TableName()).Model(&Student{Id:1}).
	//	Updates(stu).Error //Model()必须在Table()后面,Model里面必须用主键
	//fmt.Println(e)
	//tx.Commit()
	//fmt.Println(stu)

	//删除
	//tx := db.Table(schema + Student{}.TableName()).Begin()
	//rows := tx.Delete(&Student{},"id = 11 and s_name = ?","joffre").RowsAffected
	//if rows == 2{
	//	tx.Commit()
	//}else {
	//	tx.Rollback()
	//}

	//json 解析空数组
	//str := `null`
	//str := `[]`
	//sl := []int{}
	//e := json.Unmarshal([]byte(str),&sl)
	//fmt.Println(e,sl)

}

func Insert(slice, insertion []int, index int) []int {
	result := make([]int, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

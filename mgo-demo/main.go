// This program provides a sample application for using MongoDB with
// the mgo driver.
package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sync"
	"time"
	"runtime"
	"net/http"
)

const (
	MongoDBHosts = "10.214.224.142:20000"
	AuthDatabase = "kg_test"
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "hzds"
)

type (
	SaleRecord struct {
		ID          bson.ObjectId `bson:"_id,omitempty"`
		CustomerId  bson.ObjectId `bson:"customer_id"`
		ProductId   int           `bson:"product_id"`
		Price       int           `bson:"price"`
		SaleId      int           `bson:"sale_id"`
		Discount    float64       `bson:"discount"`
		Amount      int           `bson:"amount"`
		BuyTime     time.Time     `bson:"buy_time"`
		Payment     int           `bson:"payment"`
		StoreId     int           `bson:"store_id"`
		Year        int           `bson:"year"`
		Month       int           `bson:"month"`
		Day         int           `bson:"day"`
		Hour        int           `bson:"hour"`
		Minute      int           `bson:"minute"`
		UpdateCount int           `bson:"update_count"`
	}
)

var mgoSession *mgo.Session

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  180 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}
	var err error
	mgoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	
	mgoSession.SetMode(mgo.Monotonic, true)
	mgoSession.SetPoolLimit(500)
	
	//SerialFindAndUpdate(mongoSession)
	//ConcurrentFindAndUpdate(mongoSession)
	http.HandleFunc("/fast", ConcurrentFindAndUpdate)
	http.HandleFunc("/slow", SerialFindAndUpdate)
	log.Println(http.ListenAndServe(":9002", nil))
}

func SerialFindAndUpdate(w http.ResponseWriter, r *http.Request) {
	sessionCopy := mgoSession.Copy()
	defer sessionCopy.Close()
	col := sessionCopy.DB(TestDatabase).C("hzds_sale_test")
	
	month := 8
	s_hour, e_hour := 16, 16
	minute := 59
	for h := s_hour; h <= e_hour; h++ {
		for m := 0; m <= minute; m++ {
			sales := []*SaleRecord{}
			err := col.Find(bson.M{
				"month":  month,
				"hour":   h,
				"minute": m,
			}).All(&sales)
			if err != nil {
				log.Println(err)
				return
			}
			for _, v := range sales {
				v.UpdateCount ++
				err = col.Update(bson.M{"_id": v.ID}, v)
				if err != nil {
					log.Println(err)
					return
				}
			}
			log.Printf("hour:%d, minute:%d, sales_count:%d 已完成", h, m, len(sales))
		}
	}
}

func ConcurrentFindAndUpdate(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	month := 8
	s_hour, e_hour := 16, 16
	minute := 59
	
	sessionCopy := mgoSession.Copy()
	var waitGroup sync.WaitGroup
	limit := make(chan bool, 100)
	
	for h := s_hour; h <= e_hour; h++ {
		for m := 0; m <= minute; m++ {
			limit <- true
			waitGroup.Add(1)
			go concurrentFind(month, h, m, &waitGroup, sessionCopy, limit)
		}
	}
	waitGroup.Wait()
	log.Println("测试结束,总时间:", time.Now().Sub(start).String())
}

func concurrentFind(month, hour, minute int, waitGroup *sync.WaitGroup, mongoSession *mgo.Session, limit chan bool) {
	defer waitGroup.Done()
	
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	
	col := sessionCopy.DB(TestDatabase).C("hzds_sale_test")
	
	sales := []*SaleRecord{}
	err := col.Find(bson.M{
		"month":  month,
		"hour":   hour,
		"minute": minute,
	}).All(&sales)
	if err != nil {
		log.Printf("Find : ERROR : %s\n", err)
		return
	}
	for _, v := range sales {
		v.UpdateCount ++
		err = col.Update(bson.M{"_id": v.ID}, v)
		if err != nil {
			log.Println(err)
			return
		}
	}
	//log.Printf("hour:%d, minute:%d, sales_count:%d 已完成", hour, minute, len(sales))
	<-limit
}

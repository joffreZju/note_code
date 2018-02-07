package cal


type CarInfo struct {
	CarNo     string
	MaxVolume int
	MaxWeight int
}

type ConsumerDto struct {
	CalType   string
	OrderNo   string
	UsingId   int
	CalTimes  int
	Cars      []CarInfo
	GoodsList []MQWaybill
}

type MQWaybill struct {
	Id int
	Aw int
	Av int
	Fc int
	Ne string
	Uns string
}

//接收MQ计算结果Dto
type ProducerDto struct {
	ErrorCode  int
	OrderNo    string
	UsingId    int
	CalTimes   int
	CarSummary []MQCarSummary
	Result      []MQCalResult
}
type MQCarSummary struct{
	CarNo       string
	TotalWeight int
	TotalVolume int
	TotalMoney  int
}
type MQCalResult struct {
	Id int
	Result string
}


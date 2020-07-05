package model

type Stock struct {
	Id              int       `json:"Id" gorm:"primary_key;auto_increment" description:"主键id"`
	Code            string    `json:"code" gorm:"size:64;unique_index;not null" description:"股票代码"`
	Name            string    `json:"name" gorm:"size:64;not null" description:"股票名"`
	NorthFlow       string    `json:"northFlow" gorm:"size:64`
	Performance     string    `json:"performance" gorm:"size:32`
	Type            string    `json:"stockType" gorm:"size:32`
	TriggerPrice    float64   `json:"triggerPrice"`
	TriggerOrgBuy   float64   `json:"triggerOrgBuy"`
	TriggerDate     LocalTime `json:"triggerDate"`
	TopPrice        float64   `json:"topPrice"`
	TopDate         LocalTime `json:"topDate"`
	TriggerTopDay   int       `json:"triggerTopDay"`
	FloorPrice      float64   `json:"floorPrice"`
	FloorDate       LocalTime `json:"floorDate"`
	TriggerFloorDay int       `json:"triggerFloorDay"`
	OrgBuyTime      int       `json:"orgBuyTime"`
	OrgBuyMoney     float64   `json:"orgBuyMoney"`
	OrgSellTime     int       `json:"orgSellTime"`
	OrgSellMoney    float64   `json:"orgSellMoney"`
	CurPrice        float64   `json:"curPrice"`
	Industry        string    `json:"industry" gorm:"size:128`
	CreatedAt       LocalTime `json:"createdAt" description:"创建时间"`
	UpdatedAt       LocalTime `json:"ureatedAt" description:"更新时间"`
}

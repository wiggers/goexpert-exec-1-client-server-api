package entity

type Exchange struct {
	ID            uint   `gorm:"primaryKey"`
	Code          string `json:"code"`
	CodeIn        string `json:"codein"`
	Name          string `json:"name"`
	High          string `json:"high"`
	Low           string `json:"low"`
	VarBid        string `json:"varBid"`
	PctChange     string `json:"pctChange"`
	Bid           string `json:"bid"`
	Ask           string `json:"ask"`
	DateTimestamp string `json:"timestamp"`
	CreateDate    string `json:"create_date"`
}

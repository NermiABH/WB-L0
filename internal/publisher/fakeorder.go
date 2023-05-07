package publisher

import (
	"Wb-L0/internal/model"
	"github.com/brianvoe/gofakeit/v6"
	"strconv"
	"time"
)

func NewFakeOrder() model.OrderJsonStruct {
	return model.OrderJsonStruct{
		OrderUid:    gofakeit.UUID(),
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: model.Delivery{
			Name:    gofakeit.Name(),
			Phone:   gofakeit.Phone(),
			Zip:     gofakeit.Zip(),
			City:    gofakeit.City(),
			Address: gofakeit.Street(),
			Region:  gofakeit.State(),
			Email:   gofakeit.Email(),
		},
		Payment: model.Payment{
			Transaction: gofakeit.UUID(),
			RequestId:   "",
			Currency:    gofakeit.Currency().Short,
			Provider:    gofakeit.Word(),
			Amount:      gofakeit.Number(1000, 5000),
			PaymentDt: int(gofakeit.DateRange(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 12, 31, 23, 59, 59, 999999999, time.UTC)).Unix()),
			Bank:         gofakeit.Word(),
			DeliveryCost: gofakeit.Number(500, 2000),
			GoodsTotal:   gofakeit.Number(100, 1000),
			CustomFee:    gofakeit.Number(0, 100),
		},
		Items: model.Items{
			{
				ChrtId:      gofakeit.Number(1000000, 9999999),
				TrackNumber: "WBILMTESTTRACK",
				Price:       gofakeit.Number(100, 1000),
				Rid:         gofakeit.UUID(),
				Name:        gofakeit.Word(),
				Sale:        gofakeit.Number(0, 50),
				Size:        strconv.Itoa(gofakeit.Number(0, 10)),
				TotalPrice:  gofakeit.Number(1, 1000),
				NmId:        gofakeit.Number(1000000, 9999999),
				Brand:       gofakeit.Company(),
				Status:      gofakeit.Number(100, 400),
			},
		},
		Locale:            gofakeit.LanguageAbbreviation(),
		InternalSignature: "",
		CustomerId:        gofakeit.Username(),
		DeliveryService:   gofakeit.Word(),
		Shardkey:          strconv.Itoa(gofakeit.Number(0, 10)),
		SmId:              gofakeit.Number(0, 100),
		DateCreated: gofakeit.DateRange(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2022, 12, 31, 23, 59, 59, 999999999, time.UTC)),
		OofShard: strconv.Itoa(gofakeit.Number(0, 50)),
	}
}

//type FakeOrder struct {
//	OrderUid    string `json:"order_uid" fake:"{uuid}"`
//	TrackNumber string `json:"track_number" fake:""`
//	Entry       string `json:"entry"`
//	Delivery    struct {
//		Name    string `json:"name" fake:"{name}"`
//		Phone   string `json:"phone" fake:"{phone}"`
//		Zip     string `json:"zip" fake:"{number:1000000,999999}"`
//		City    string `json:"city" fake:"{city}"`
//		Address string `json:"address" fake:"{address}"`
//		Region  string `json:"region"`
//		Email   string `json:"email" fake:"{email}"`
//	} `json:"delivery"`
//	Payment struct {
//		Transaction  string `json:"transaction"`
//		RequestId    string `json:"request_id" fake:"{uuid}"`
//		Currency     string `json:"currency" fake:"{currency}"`
//		Provider     string `json:"provider"`
//		Amount       int    `json:"amount"`
//		PaymentDt    int    `json:"payment_dt"`
//		Bank         string `json:"bank"`
//		DeliveryCost int    `json:"delivery_cost"`
//		GoodsTotal   int    `json:"goods_total"`
//		CustomFee    int    `json:"custom_fee"`
//	} `json:"payment"`
//	Items []struct {
//		ChrtId      int    `json:"chrt_id"`
//		TrackNumber string `json:"track_number"`
//		Price       int    `json:"price"`
//		Rid         string `json:"rid"`
//		Name        string `json:"name"`
//		Sale        int    `json:"sale"`
//		Size        string `json:"size"`
//		TotalPrice  int    `json:"total_price"`
//		NmId        int    `json:"nm_id"`
//		Brand       string `json:"brand"`
//		Status      int    `json:"status"`
//	} `json:"items"`
//	Locale            string    `json:"locale"`
//	InternalSignature string    `json:"internal_signature"`
//	CustomerId        string    `json:"customer_id"`
//	DeliveryService   string    `json:"delivery_service"`
//	Shardkey          string    `json:"shardkey"`
//	SmId              int       `json:"sm_id"`
//	DateCreated       time.Time `json:"date_created"`
//	OofShard          string    `json:"oof_shard"`
//}

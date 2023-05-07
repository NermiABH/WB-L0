package model

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"time"
)

var validate = validator.New()

type Order struct {
	UUID      string `json:"order_uid"`
	OrderJson string
}

type OrderJsonStruct struct {
	OrderUid          string `json:"order_uid" validate:"required"`
	TrackNumber       string `json:"track_number" validate:"required"`
	Entry             string `json:"entry" validate:"required"`
	Delivery          `json:"delivery" validate:"required"`
	Payment           `json:"payment" validate:"required"`
	Items             `json:"items" validate:"required"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature" validate:"required"`
	CustomerId        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	SmId              int       `json:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}

type Delivery struct {
	Name    string `json:"name" validate:"required,min=3,max=50"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required,alpha"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required,alpha"`
	Email   string `json:"email" validate:"required"`
}

type Payment struct {
	Transaction  string `json:"transaction,uuid" validate:"required"`
	RequestId    string `json:"request_id" validate:"required"`
	Currency     string `json:"currency,numeric,len" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int    `json:"amount" validate:"required"`
	PaymentDt    int    `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"required"`
	GoodsTotal   int    `json:"goods_total" validate:"required"`
	CustomFee    int    `json:"custom_fee" validate:"required"`
}
type Items []struct {
	ChrtId      int    `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int    `json:"sale" validate:"required"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  int    `json:"total_price" validate:"required"`
	NmId        int    `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      int    `json:"status" validate:"required"`
}

func JsonToOrder(orderJson []byte) (*Order, error) {
	var orderJsonStruct OrderJsonStruct
	if err := json.Unmarshal(orderJson, &orderJsonStruct); err != nil {
		return nil, err
	}
	if err := validate.Struct(&orderJsonStruct); err != nil {
		return nil, err
	}
	order := Order{orderJsonStruct.OrderUid, string(orderJson)}
	return &order, nil
}

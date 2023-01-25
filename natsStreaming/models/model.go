package model

import (
	"sync"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type InMemoryCache struct {
	sync.RWMutex
	store map[string]interface{}
}

func NewCache() *InMemoryCache {
	return &InMemoryCache{store: make(map[string]interface{})}
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	var value interface{}
	data, ok := c.store[key]
	if ok {
		value = data
	}
	return value, ok
}

func (c *InMemoryCache) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	_, ok := c.store[key]
	if !ok {
		c.store[key] = value
	}
}

type Order struct {
	OrderUid          string   `json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Delivery          Delivery `json:"delivery"`
	Payment           Payment  `json:"payment"`
	Items             []Items  `json:"items"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerId        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	Shardkey          string   `json:"shardkey"`
	SmId              int      `json:"sm_id"`
	DateCreated       string   `json:"date_created"`
	OofShard          string   `json:"oof_shard"`
}
type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Items struct {
	ChrtId      int64  `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

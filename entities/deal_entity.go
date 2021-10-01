package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type DealEntity struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	DealId           int                `json:"id" bson:"deal_id,omitempty"`
	ProductId        int                `json:"product_id" bson:"product_id,omitempty"`
	ProductName      string             `json:"product_name" bson:"product_name,omitempty"`
	Sku              string             `json:"sku" bson:"sku,omitempty"`
	SellerId         int                `json:"seller_id" bson:"seller_id,omitempty"`
	Price            float32            `json:"price" bson:"price,omitempty"`
	SpecialPrice     float32            `json:"special_price" bson:"special_price,omitempty"`
	SpecialFromDate  string             `json:"special_from_date" bson:"special_from_date,omitempty"`
	SpecialToDate    string             `json:"special_to_date" bson:"special_to_date,omitempty"`
	QtyMax           int                `json:"qty_max" bson:"qty_max,omitempty"`
	QtyOrdered       int                `json:"qty_ordered" bson:"qty_ordered,omitempty"`
	QtyLimit         int                `json:"qty_limit" bson:"qty_limit,omitempty"`
	IsHotDeal        uint8              `json:"is_hot_deal" bson:"is_hot_deal,omitempty"`
	AreaId           int                `json:"area_id" bson:"area_id,omitempty"`
	SortOrder        int                `json:"sort_order" bson:"sort_order,omitempty"`
	IsActive         int8               `json:"is_active" bson:"is_active,omitempty"`
	Url              string             `json:"url" bson:"url,omitempty"`
	MinPrice         int                `json:"min_price" bson:"min_price,omitempty"`
	StepPrice        int                `json:"step_price" bson:"step_price,omitempty"`
	DealSpecialPrice int                `json:"deal_special_price" bson:"deal_special_price,omitempty"`
	IsChild          int                `json:"is_child" bson:"is_child,omitempty"`
	DiscountPercent  float32            `json:"discount_percent" bson:"discount_percent"`
	Tags             string             `json:"tags" bson:"tags,omitempty" default:""`
	CampaignId       int                `json:"campaign_id" bson:"campaign_id,omitempty"`
	Priority         int                `json:"priority" bson:"priority,omitempty"`
}

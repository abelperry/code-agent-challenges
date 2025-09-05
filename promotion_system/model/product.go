package model

import "time"

// Product 商品信息
type Product struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Category   string    `json:"category"` // 商品分类: "electronics", "clothing", "books"
	Brand      string    `json:"brand"`
	CreateTime time.Time `json:"create_time"`
}

// User 用户信息
type User struct {
	ID           int       `json:"id"`
	Level        string    `json:"level"` // 用户等级: "bronze", "silver", "gold", "platinum"
	RegisterTime time.Time `json:"register_time"`
}

// Order 订单信息
type Order struct {
	Products    []Product `json:"products"`
	User        User      `json:"user"`
	TotalAmount float64   `json:"total_amount"`
	CreateTime  time.Time `json:"create_time"`
}

// PromotionResult 促销计算结果
type PromotionResult struct {
	OriginalAmount float64            `json:"original_amount"`
	DiscountAmount float64            `json:"discount_amount"`
	FinalAmount    float64            `json:"final_amount"`
	AppliedRules   []string           `json:"applied_rules"`
	Details        map[string]float64 `json:"details"`
}

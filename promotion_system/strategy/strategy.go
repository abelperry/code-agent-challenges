package strategy

import "code-agent-challenges/promotion_system/model"

type PromotionStrategy interface {
	Calculate(order model.Order) model.PromotionResult
}

func NewPromotionStrategy(strategyType string) PromotionStrategy {
	return nil
}

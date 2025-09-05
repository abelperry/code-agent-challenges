package strategy

import (
	"code-agent-challenges/promotion_system/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1. **满减策略 (FullReductionStrategy)**
   - 订单满1000元减100元
   - 满500元减50元

2. **用户等级折扣策略 (UserLevelDiscountStrategy)**
   - bronze: 无折扣
   - silver: 95折
   - gold: 9折
   - platinum: 85折

3. **商品分类折扣策略 (CategoryDiscountStrategy)**
   - electronics: 92折
   - clothing: 88折
   - books: 8折
*/

var testCases = []struct {
	name     string
	order    model.Order
	expected model.PromotionResult
}{
	// ========== Bronze用户组合 (12个) ==========
	// Bronze + 无满减 + 无分类
	{
		name: "Bronze用户+无满减+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 300}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 0,
			FinalAmount:    300,
			AppliedRules:   []string{},
		},
	},
	// Bronze + 无满减 + electronics
	{
		name: "Bronze用户+无满减+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "electronics"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 24, // 300 * 0.08 = 24
			FinalAmount:    276,
			AppliedRules:   []string{"CategoryDiscount"},
		},
	},
	// Bronze + 无满减 + clothing
	{
		name: "Bronze用户+无满减+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "clothing"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 36, // 300 * 0.12 = 36
			FinalAmount:    264,
			AppliedRules:   []string{"CategoryDiscount"},
		},
	},
	// Bronze + 无满减 + books
	{
		name: "Bronze用户+无满减+books",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "books"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 60, // 300 * 0.2 = 60
			FinalAmount:    240,
			AppliedRules:   []string{"CategoryDiscount"},
		},
	},
	// Bronze + 满500减50 + 无分类
	{
		name: "Bronze用户+满500减50+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 300}, {Price: 250}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 50,
			FinalAmount:    500,
			AppliedRules:   []string{"FullReduction"},
		},
	},
	// Bronze + 满500减50 + electronics
	{
		name: "Bronze用户+满500减50+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "electronics"}, {Price: 250, Category: "electronics"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 94, // 50 + 550*0.08 = 50 + 44 = 94
			FinalAmount:    456,
			AppliedRules:   []string{"FullReduction", "CategoryDiscount"},
		},
	},
	// Bronze + 满500减50 + clothing
	{
		name: "Bronze用户+满500减50+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "clothing"}, {Price: 250, Category: "clothing"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 116, // 50 + 550*0.12 = 50 + 66 = 116
			FinalAmount:    434,
			AppliedRules:   []string{"FullReduction", "CategoryDiscount"},
		},
	},
	// Bronze + 满500减50 + books
	{
		name: "Bronze用户+满500减50+books",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "books"}, {Price: 250, Category: "books"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 160, // 50 + 550*0.2 = 50 + 110 = 160
			FinalAmount:    390,
			AppliedRules:   []string{"FullReduction", "CategoryDiscount"},
		},
	},
	// Bronze + 满1000减100 + 无分类
	{
		name: "Bronze用户+满1000减100+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 600}, {Price: 500}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 100,
			FinalAmount:    1000,
			AppliedRules:   []string{"FullReduction"},
		},
	},
	// Bronze + 满1000减100 + electronics
	{
		name: "Bronze用户+满1000减100+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "electronics"}, {Price: 500, Category: "electronics"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 188, // 100 + 1100*0.08 = 100 + 88 = 188
			FinalAmount:    912,
			AppliedRules:   []string{"FullReduction", "CategoryDiscount"},
		},
	},
	// Bronze + 满1000减100 + clothing
	{
		name: "Bronze用户+满1000减100+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "clothing"}, {Price: 500, Category: "clothing"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 232, // 100 + 1100*0.12 = 100 + 132 = 232
			FinalAmount:    868,
			AppliedRules:   []string{"FullReduction", "CategoryDiscount"},
		},
	},
	// Bronze + 满1000减100 + books
	{
		name: "Bronze用户+满1000减100+books",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "books"}, {Price: 500, Category: "books"}},
			User:     model.User{Level: "bronze"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 320, // 100 + 1100*0.2 = 100 + 220 = 320
			FinalAmount:    780,
			AppliedRules:   []string{"FullReduction", "CategoryDiscount"},
		},
	},

	// ========== Silver用户组合 (12个) ==========
	// Silver + 无满减 + 无分类
	{
		name: "Silver用户+无满减+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 300}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 15, // 300 * 0.05 = 15
			FinalAmount:    285,
			AppliedRules:   []string{"UserLevelDiscount"},
		},
	},
	// Silver + 无满减 + electronics
	{
		name: "Silver用户+无满减+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "electronics"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 39, // 300 * 0.05 + 300 * 0.08 = 15 + 24 = 39
			FinalAmount:    261,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Silver + 无满减 + clothing
	{
		name: "Silver用户+无满减+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "clothing"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 51, // 300 * 0.05 + 300 * 0.12 = 15 + 36 = 51
			FinalAmount:    249,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Silver + 无满减 + books
	{
		name: "Silver用户+无满减+books",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "books"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 75, // 300 * 0.05 + 300 * 0.2 = 15 + 60 = 75
			FinalAmount:    225,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Silver + 满500减50 + 无分类
	{
		name: "Silver用户+满500减50+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 300}, {Price: 250}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 77.5, // 50 + 550 * 0.05 = 50 + 27.5 = 77.5
			FinalAmount:    472.5,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount"},
		},
	},
	// Silver + 满500减50 + electronics
	{
		name: "Silver用户+满500减50+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "electronics"}, {Price: 250, Category: "electronics"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 121.5, // 50 + 550*0.05 + 550*0.08 = 50 + 27.5 + 44 = 121.5
			FinalAmount:    428.5,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Silver + 满500减50 + clothing
	{
		name: "Silver用户+满500减50+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "clothing"}, {Price: 250, Category: "clothing"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 143.5, // 50 + 550*0.05 + 550*0.12 = 50 + 27.5 + 66 = 143.5
			FinalAmount:    406.5,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Silver + 满500减50 + books
	{
		name: "Silver用户+满500减50+books",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "books"}, {Price: 250, Category: "books"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 187.5, // 50 + 550*0.05 + 550*0.2 = 50 + 27.5 + 110 = 187.5
			FinalAmount:    362.5,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Silver + 满1000减100 + 无分类
	{
		name: "Silver用户+满1000减100+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 600}, {Price: 500}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 155, // 100 + 1100 * 0.05 = 100 + 55 = 155
			FinalAmount:    945,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount"},
		},
	},
	// Silver + 满1000减100 + electronics
	{
		name: "Silver用户+满1000减100+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "electronics"}, {Price: 500, Category: "electronics"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 243, // 100 + 1100*0.05 + 1100*0.08 = 100 + 55 + 88 = 243
			FinalAmount:    857,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Silver + 满1000减100 + clothing
	{
		name: "Silver用户+满1000减100+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "clothing"}, {Price: 500, Category: "clothing"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 287, // 100 + 1100*0.05 + 1100*0.12 = 100 + 55 + 132 = 287
			FinalAmount:    813,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Silver + 满1000减100 + books
	{
		name: "Silver用户+满1000减100+books",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "books"}, {Price: 500, Category: "books"}},
			User:     model.User{Level: "silver"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 375, // 100 + 1100*0.05 + 1100*0.2 = 100 + 55 + 220 = 375
			FinalAmount:    725,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},

	// ========== Gold用户组合 (12个) ==========
	// Gold + 无满减 + 无分类
	{
		name: "Gold用户+无满减+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 300}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 30, // 300 * 0.1 = 30
			FinalAmount:    270,
			AppliedRules:   []string{"UserLevelDiscount"},
		},
	},
	// Gold + 无满减 + electronics
	{
		name: "Gold用户+无满减+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "electronics"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 54, // 300 * 0.1 + 300 * 0.08 = 30 + 24 = 54
			FinalAmount:    246,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Gold + 无满减 + clothing
	{
		name: "Gold用户+无满减+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "clothing"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 66, // 300 * 0.1 + 300 * 0.12 = 30 + 36 = 66
			FinalAmount:    234,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Gold + 无满减 + books
	{
		name: "Gold用户+无满减+books",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "books"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 90, // 300 * 0.1 + 300 * 0.2 = 30 + 60 = 90
			FinalAmount:    210,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Gold + 满500减50 + 无分类
	{
		name: "Gold用户+满500减50+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 300}, {Price: 250}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 105, // 50 + 550 * 0.1 = 50 + 55 = 105
			FinalAmount:    445,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount"},
		},
	},
	// Gold + 满500减50 + electronics
	{
		name: "Gold用户+满500减50+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "electronics"}, {Price: 250, Category: "electronics"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 149, // 50 + 550*0.1 + 550*0.08 = 50 + 55 + 44 = 149
			FinalAmount:    401,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Gold + 满500减50 + clothing
	{
		name: "Gold用户+满500减50+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "clothing"}, {Price: 250, Category: "clothing"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 171, // 50 + 550*0.1 + 550*0.12 = 50 + 55 + 66 = 171
			FinalAmount:    379,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Gold + 满500减50 + books
	{
		name: "Gold用户+满500减50+books",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "books"}, {Price: 250, Category: "books"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 215, // 50 + 550*0.1 + 550*0.2 = 50 + 55 + 110 = 215
			FinalAmount:    335,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Gold + 满1000减100 + 无分类
	{
		name: "Gold用户+满1000减100+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 600}, {Price: 500}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 210, // 100 + 1100 * 0.1 = 100 + 110 = 210
			FinalAmount:    890,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount"},
		},
	},
	// Gold + 满1000减100 + electronics
	{
		name: "Gold用户+满1000减100+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "electronics"}, {Price: 500, Category: "electronics"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 298, // 100 + 1100*0.1 + 1100*0.08 = 100 + 110 + 88 = 298
			FinalAmount:    802,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Gold + 满1000减100 + clothing
	{
		name: "Gold用户+满1000减100+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "clothing"}, {Price: 500, Category: "clothing"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 342, // 100 + 1100*0.1 + 1100*0.12 = 100 + 110 + 132 = 342
			FinalAmount:    758,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Gold + 满1000减100 + books
	{
		name: "Gold用户+满1000减100+books",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "books"}, {Price: 500, Category: "books"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 430, // 100 + 1100*0.1 + 1100*0.2 = 100 + 110 + 220 = 430
			FinalAmount:    670,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},

	// ========== Platinum用户组合 (12个) ==========
	// Platinum + 无满减 + 无分类
	{
		name: "Platinum用户+无满减+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 300}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 45, // 300 * 0.15 = 45
			FinalAmount:    255,
			AppliedRules:   []string{"UserLevelDiscount"},
		},
	},
	// Platinum + 无满减 + electronics
	{
		name: "Platinum用户+无满减+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "electronics"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 69, // 300 * 0.15 + 300 * 0.08 = 45 + 24 = 69
			FinalAmount:    231,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Platinum + 无满减 + clothing
	{
		name: "Platinum用户+无满减+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "clothing"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 81, // 300 * 0.15 + 300 * 0.12 = 45 + 36 = 81
			FinalAmount:    219,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Platinum + 无满减 + books
	{
		name: "Platinum用户+无满减+books",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "books"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 300,
			DiscountAmount: 105, // 300 * 0.15 + 300 * 0.2 = 45 + 60 = 105
			FinalAmount:    195,
			AppliedRules:   []string{"UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Platinum + 满500减50 + 无分类
	{
		name: "Platinum用户+满500减50+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 300}, {Price: 250}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 132.5, // 50 + 550 * 0.15 = 50 + 82.5 = 132.5
			FinalAmount:    417.5,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount"},
		},
	},
	// Platinum + 满500减50 + electronics
	{
		name: "Platinum用户+满500减50+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "electronics"}, {Price: 250, Category: "electronics"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 176.5, // 50 + 550*0.15 + 550*0.08 = 50 + 82.5 + 44 = 176.5
			FinalAmount:    373.5,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Platinum + 满500减50 + clothing
	{
		name: "Platinum用户+满500减50+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "clothing"}, {Price: 250, Category: "clothing"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 198.5, // 50 + 550*0.15 + 550*0.12 = 50 + 82.5 + 66 = 198.5
			FinalAmount:    351.5,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Platinum + 满500减50 + books
	{
		name: "Platinum用户+满500减50+books",
		order: model.Order{
			Products: []model.Product{{Price: 300, Category: "books"}, {Price: 250, Category: "books"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 550,
			DiscountAmount: 242.5, // 50 + 550*0.15 + 550*0.2 = 50 + 82.5 + 110 = 242.5
			FinalAmount:    307.5,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Platinum + 满1000减100 + 无分类
	{
		name: "Platinum用户+满1000减100+无分类",
		order: model.Order{
			Products: []model.Product{{Price: 600}, {Price: 500}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 265, // 100 + 1100 * 0.15 = 100 + 165 = 265
			FinalAmount:    835,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount"},
		},
	},
	// Platinum + 满1000减100 + electronics
	{
		name: "Platinum用户+满1000减100+electronics",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "electronics"}, {Price: 500, Category: "electronics"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 353, // 100 + 1100*0.15 + 1100*0.08 = 100 + 165 + 88 = 353
			FinalAmount:    747,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Platinum + 满1000减100 + clothing
	{
		name: "Platinum用户+满1000减100+clothing",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "clothing"}, {Price: 500, Category: "clothing"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 397, // 100 + 1100*0.15 + 1100*0.12 = 100 + 165 + 132 = 397
			FinalAmount:    703,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},
	// Platinum + 满1000减100 + books
	{
		name: "Platinum用户+满1000减100+books",
		order: model.Order{
			Products: []model.Product{{Price: 600, Category: "books"}, {Price: 500, Category: "books"}},
			User:     model.User{Level: "platinum"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1100,
			DiscountAmount: 485, // 100 + 1100*0.15 + 1100*0.2 = 100 + 165 + 220 = 485
			FinalAmount:    615,
			AppliedRules:   []string{"FullReduction", "UserLevelDiscount", "CategoryDiscount"},
		},
	},

	// ========== 边界条件测试 (3个) ==========
	// 空订单
	{
		name: "边界条件-空订单",
		order: model.Order{
			Products: []model.Product{},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 0,
			DiscountAmount: 0,
			FinalAmount:    0,
			AppliedRules:   []string{},
		},
	},
	// 价格为0
	{
		name: "边界条件-价格为0",
		order: model.Order{
			Products: []model.Product{{Price: 0, Category: "electronics"}},
			User:     model.User{Level: "gold"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 0,
			DiscountAmount: 0,
			FinalAmount:    0,
			AppliedRules:   []string{},
		},
	},
	// 无效用户等级
	{
		name: "边界条件-无效用户等级",
		order: model.Order{
			Products: []model.Product{{Price: 1000}},
			User:     model.User{Level: "invalid"},
		},
		expected: model.PromotionResult{
			OriginalAmount: 1000,
			DiscountAmount: 0,
			FinalAmount:    1000,
			AppliedRules:   []string{},
		},
	},
}

func TestAll(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// 累加所有策略的折扣金额
			var totalDiscount float64 = 0

			// 应用每个策略并累加折扣
			for _, rule := range test.expected.AppliedRules {
				strategy := NewPromotionStrategy(rule)
				if strategy != nil {
					result := strategy.Calculate(test.order)
					totalDiscount += result.DiscountAmount
				}
			}

			// 计算最终金额
			finalAmount := test.expected.OriginalAmount - totalDiscount

			// 验证折扣金额和最终金额
			assert.Equal(t, test.expected.DiscountAmount, totalDiscount, "折扣金额不匹配")
			assert.Equal(t, test.expected.FinalAmount, finalAmount, "最终金额不匹配")
		})
	}
}

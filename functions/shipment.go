package functions

import "github.com/yourusername/logistics-api/models"

func ScoreOption(opt *models.ShippingOption, urgency string) {
	factor := 1.0
	if urgency == "high" {
		factor = 1.5
	}
	opt.Score = factor*float64(100-opt.EstimatedTime) - opt.Cost
}
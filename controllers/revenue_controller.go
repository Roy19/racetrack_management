package controllers

import "github.com/Roy19/racetrack-management/interfaces"

type RevenueController struct {
	revenueService interfaces.IRevenueService
}

func (rc *RevenueController) CalculateRevenue() (int, int) {
	return rc.revenueService.CalculateRevenue()
}

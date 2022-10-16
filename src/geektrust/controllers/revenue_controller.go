package controllers

import "geektrust/interfaces"

type RevenueController struct {
	RevenueService interfaces.IRevenueService
}

func (rc *RevenueController) CalculateRevenue() (int, int) {
	return rc.RevenueService.CalculateRevenue()
}

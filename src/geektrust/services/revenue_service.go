package services

import (
	"math"
	"time"

	"geektrust/models"
)

type RevenueService struct {
	RaceTrackManagement *models.RaceTrackManagement
}

func (rs *RevenueService) CalculateRevenue() (int, int) {
	regularTrackRevenue := calculateRevenueForTrack(rs.RaceTrackManagement, models.REGULAR)
	vipTrackRevenue := calculateRevenueForTrack(rs.RaceTrackManagement, models.VIP)
	return regularTrackRevenue, vipTrackRevenue
}

func calculateRevenueForTrack(raceTracks *models.RaceTrackManagement,
	trackType models.RacetrackType) int {
	totalRevenue := 0
	for _, v := range raceTracks.RaceTracks {
		if v.RaceTrackType == trackType {
			for _, slot := range v.BookedSlots {
				t := slot.StartTime.Add(time.Duration(3)*time.Hour +
					time.Duration(15)*time.Minute)
				diff := slot.EndTime.Sub(t)
				if diff.Hours() <= 0 {
					totalRevenue += 3 *
						models.GetChargeGivenTrackAndVehicleType(v.AllowedVehicleType,
							v.RaceTrackType)
				} else {
					extraHours := math.Ceil(diff.Hours())
					totalRevenue += (3 + int(extraHours)) *
						models.GetChargeGivenTrackAndVehicleType(v.AllowedVehicleType, v.RaceTrackType)
				}
			}
		}
	}
	return totalRevenue
}

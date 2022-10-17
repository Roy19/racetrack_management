package services

import (
	"math"
	"time"

	"geektrust/models"
)

const (
	defaultHours   int = 3
	defaultMinutes int = 15
)

const (
	bikeTrackRegularCostPerHour int = 60
	carTrackRegularCostPerHour  int = 120
	carTrackVipCostPerHour      int = 250
	suvTrackRegularCostPerHour  int = 200
	suvTrackVipCostPerHour      int = 300
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
				t := slot.StartTime.Add(time.Duration(defaultHours)*time.Hour +
					time.Duration(defaultMinutes)*time.Minute)
				diff := slot.EndTime.Sub(t)
				totalRevenue += (defaultHours + int(math.Ceil(math.Max(0, diff.Hours())))) *
					getChargeGivenTrackAndVehicleType(v.AllowedVehicleType, v.RaceTrackType)
			}
		}
	}
	return totalRevenue
}

func getChargeGivenTrackAndVehicleType(vehicleType models.VehicleType, trackType models.RacetrackType) int {
	if trackType == models.REGULAR {
		if vehicleType == models.BIKE {
			return bikeTrackRegularCostPerHour
		}
		if vehicleType == models.CAR {
			return carTrackRegularCostPerHour
		}
		if vehicleType == models.SUV {
			return suvTrackRegularCostPerHour
		}
	} else if trackType == models.VIP {
		if vehicleType == models.CAR {
			return carTrackVipCostPerHour
		}
		if vehicleType == models.SUV {
			return suvTrackVipCostPerHour
		}
	}
	return 0
}

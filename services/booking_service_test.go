package services

import (
	"testing"
	"time"

	"github.com/Roy19/racetrack-management/interfaces"
	"github.com/Roy19/racetrack-management/models"
)

const error_format_string = "Failed in method %v. Input: %v, Got: %v, Expected: %v\n"

func TestTryBookingSlot(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	bookingService := BookingService{
		raceTrackManagement: &mockRaceTrackManagement,
	}
	t1, _ := time.Parse("15:04:05", "13:00:00")
	t2 := t1.Add(time.Hour * time.Duration(3))
	slot1 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.BIKE,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	got := bookingService.TryBookingSlot(&slot1)
	expected := true
	if got != expected {
		t.Errorf(error_format_string, "TestTryBookingSlot", slot1, got, expected)
	}
	// should also check if slot is appended in place
}

func TestCantBookSameSlotIfNotAvailable(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	bookingService := BookingService{
		raceTrackManagement: &mockRaceTrackManagement,
	}
	t1, _ := time.Parse("15:04:05", "13:00:00")
	t2 := t1.Add(time.Hour * time.Duration(3))
	slot1 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.CAR,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	bookingService.TryBookingSlot(&slot1)
	bookingService.TryBookingSlot(&slot1)
	bookingService.TryBookingSlot(&slot1)
	got := bookingService.TryBookingSlot(&slot1)
	expected := false
	if got != expected {
		t.Errorf(error_format_string, "TestTryBookingSlot", slot1, got, expected)
	}
}

func initMockRaceTrackManagement() models.RaceTrackManagement {
	mock := interfaces.RaceTrackManagementBuilder{
		RaceTracks: make([]*models.RaceTrack, 0),
	}
	mock.AddRacetrackForVechicleAndRacetrackType(models.BIKE, models.REGULAR, 4)
	mock.AddRacetrackForVechicleAndRacetrackType(models.CAR, models.REGULAR, 2)
	mock.AddRacetrackForVechicleAndRacetrackType(models.CAR, models.VIP, 1)
	mock.AddRacetrackForVechicleAndRacetrackType(models.SUV, models.REGULAR, 2)
	mock.AddRacetrackForVechicleAndRacetrackType(models.SUV, models.VIP, 1)
	return mock.BuildRacetrack()
}

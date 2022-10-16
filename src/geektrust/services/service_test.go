package services

import (
	"testing"
	"time"

	"geektrust/builders"
	"geektrust/models"
)

const error_format_string = "Failed in method %v. Input: %v, Got: %v, Expected: %v\n"

func TestTryBookingSlot(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	bookingService := BookingService{
		RaceTrackManagement: &mockRaceTrackManagement,
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
	if !checkBookedSlotsCount(&mockRaceTrackManagement, 1) {
		t.Errorf(error_format_string, "TestTryBookingSlot", mockRaceTrackManagement,
			mockRaceTrackManagement.GetTotalSlotsBooked(), 1)
	}
}

func TestCantBookSameSlotIfNotAvailable(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	bookingService := BookingService{
		RaceTrackManagement: &mockRaceTrackManagement,
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
		t.Errorf(error_format_string, "TestCantBookSameSlotIfNotAvailable",
			slot1,
			got,
			expected)
	}
	if !checkBookedSlotsCount(&mockRaceTrackManagement, 3) {
		t.Errorf(error_format_string, "TestCantBookSameSlotIfNotAvailable",
			mockRaceTrackManagement,
			mockRaceTrackManagement.GetTotalSlotsBooked(),
			3)
	}
}

func TestTryBookingDifferentSlots(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	bookingService := BookingService{
		RaceTrackManagement: &mockRaceTrackManagement,
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
	t1, _ = time.Parse("15:04:05", "14:00:00")
	t2 = t1.Add(time.Hour * time.Duration(3))
	slot2 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.CAR,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "15:00:00")
	t2 = t1.Add(time.Hour * time.Duration(3))
	slot3 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.CAR,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "16:00:05")
	t2 = t1.Add(time.Hour * time.Duration(3))
	slot4 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.CAR,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "18:00:00")
	t2 = t1.Add(time.Hour * time.Duration(3))
	slot5 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.CAR,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	got := bookingService.TryBookingSlot(&slot1)
	got = got && bookingService.TryBookingSlot(&slot2)
	got = got && bookingService.TryBookingSlot(&slot3)
	got = got && bookingService.TryBookingSlot(&slot4)
	got = got && bookingService.TryBookingSlot(&slot5)
	expected := true
	if got != expected {
		t.Errorf(error_format_string, "TestTryBookingDifferentSlots",
			slot1,
			got,
			expected)
	}
	if !checkBookedSlotsCount(&mockRaceTrackManagement, 5) {
		t.Errorf(error_format_string, "TestTryBookingDifferentSlots",
			mockRaceTrackManagement,
			mockRaceTrackManagement.GetTotalSlotsBooked(),
			5)
	}
}

func checkBookedSlotsCount(raceTrackManagement *models.RaceTrackManagement,
	expected int) bool {
	return raceTrackManagement.GetTotalSlotsBooked() == expected
}

func TestCalculateRevenueWithMultipleSlotsBooked(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	revenueService := RevenueService{
		RaceTrackManagement: &mockRaceTrackManagement,
	}
	bookingService := BookingService{
		RaceTrackManagement: &mockRaceTrackManagement,
	}
	duration := 3
	vehicle := models.CAR
	t1, _ := time.Parse("15:04:05", "13:00:00")
	t2 := t1.Add(time.Hour * time.Duration(duration))
	slot1 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "14:00:00")
	t2 = t1.Add(time.Hour * time.Duration(duration))
	slot2 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "15:00:00")
	t2 = t1.Add(time.Hour * time.Duration(duration))
	slot3 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "16:00:05")
	t2 = t1.Add(time.Hour * time.Duration(duration))
	slot4 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "18:00:00")
	t2 = t1.Add(time.Hour * time.Duration(duration))
	slot5 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "16:00:05")
	t2 = t1.Add(time.Hour * time.Duration(duration))
	slot6 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.SUV,
			IdentificationNumber: "XY2",
		},
		StartTime: t1,
		EndTime:   t2,
	}

	bookingService.TryBookingSlot(&slot1)
	bookingService.TryBookingSlot(&slot2)
	bookingService.TryBookingSlot(&slot3)
	bookingService.TryBookingSlot(&slot4)
	bookingService.TryBookingSlot(&slot5)
	bookingService.TryBookingSlot(&slot6)

	regularRevenue, vipRevenue := revenueService.CalculateRevenue()
	expected1 := 2040
	expected2 := 750
	if regularRevenue != expected1 || vipRevenue != expected2 {
		t.Errorf(error_format_string, "TestCalculateRevenueWithMultipleSlotsBooked",
			mockRaceTrackManagement, regularRevenue, expected1)
		t.Errorf(error_format_string, "TestCalculateRevenueWithMultipleSlotsBooked",
			mockRaceTrackManagement, vipRevenue, expected2)
	}
	got := mockRaceTrackManagement.GetTotalSlotsBooked()
	want := 6
	if got != want {
		t.Errorf(error_format_string, "TestCalculateRevenueWithMultipleSlotsBooked",
			mockRaceTrackManagement, got, want)
	}
}

func TestCalculateRevenueWithMultipleSlotsBookedNoExtraCharge(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	revenueService := RevenueService{
		RaceTrackManagement: &mockRaceTrackManagement,
	}
	bookingService := BookingService{
		RaceTrackManagement: &mockRaceTrackManagement,
	}
	duration := 3
	vehicle := models.CAR
	t1, _ := time.Parse("15:04:05", "13:00:00")
	t2 := t1.Add(time.Hour*time.Duration(duration) + time.Minute*time.Duration(10))
	slot1 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "14:00:00")
	t2 = t1.Add(time.Hour * time.Duration(duration))
	slot2 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "15:00:00")
	t2 = t1.Add(time.Hour*time.Duration(duration) + time.Minute*time.Duration(15))
	slot3 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "16:11:00")
	t2 = t1.Add(time.Hour * time.Duration(duration))
	slot4 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "18:00:00")
	t2 = t1.Add(time.Hour * time.Duration(duration))
	slot5 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicle,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	t1, _ = time.Parse("15:04:05", "16:00:05")
	t2 = t1.Add(time.Hour*time.Duration(duration) + time.Minute*time.Duration(5))
	slot6 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.SUV,
			IdentificationNumber: "XY2",
		},
		StartTime: t1,
		EndTime:   t2,
	}

	bookingService.TryBookingSlot(&slot1)
	bookingService.TryBookingSlot(&slot2)
	bookingService.TryBookingSlot(&slot3)
	bookingService.TryBookingSlot(&slot4)
	bookingService.TryBookingSlot(&slot5)
	bookingService.TryBookingSlot(&slot6)

	regularRevenue, vipRevenue := revenueService.CalculateRevenue()
	expected1 := 2040
	expected2 := 750
	if regularRevenue != expected1 || vipRevenue != expected2 {
		t.Errorf(error_format_string, "TestCalculateRevenueWithMultipleSlotsBookedNoExtraCharge",
			mockRaceTrackManagement, regularRevenue, expected1)
		t.Errorf(error_format_string, "TestCalculateRevenueWithMultipleSlotsBookedNoExtraCharge",
			mockRaceTrackManagement, vipRevenue, expected2)
	}

	got := mockRaceTrackManagement.GetTotalSlotsBooked()
	want := 6
	if got != want {
		t.Errorf(error_format_string, "TestCalculateRevenueWithMultipleSlotsBookedNoExtraCharge",
			mockRaceTrackManagement, got, want)
	}
}

func TestCalculateRevenue(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	revenueService := RevenueService{
		RaceTrackManagement: &mockRaceTrackManagement,
	}
	bookingService := BookingService{
		RaceTrackManagement: &mockRaceTrackManagement,
	}
	duration := 3
	vehicle := models.BIKE
	t1, _ := time.Parse("15:04:05", "13:00:00")
	t2 := t1.Add(time.Hour*time.Duration(duration) +
		time.Minute*time.Duration(15))
	slot1 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.BIKE,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	bookingService.TryBookingSlot(&slot1)
	regularRevenue, vipRevenue := revenueService.CalculateRevenue()
	expected1 := duration * models.GetChargeGivenTrackAndVehicleType(vehicle, models.REGULAR)
	expected2 := 0
	if regularRevenue != expected1 || vipRevenue != expected2 {
		t.Errorf(error_format_string, "TestCalculateRevenue",
			slot1, regularRevenue, expected1)
		t.Errorf(error_format_string, "TestCalculateRevenue",
			slot1, vipRevenue, expected2)
	}
}

func TestCalculateRevenueOverNormalTimeCharedNextHour(t *testing.T) {
	mockRaceTrackManagement := initMockRaceTrackManagement()
	revenueService := RevenueService{
		RaceTrackManagement: &mockRaceTrackManagement,
	}
	bookingService := BookingService{
		RaceTrackManagement: &mockRaceTrackManagement,
	}
	duration := 3
	vehicle := models.BIKE
	t1, _ := time.Parse("15:04:05", "13:00:00")
	t2 := t1.Add(time.Hour*time.Duration(duration) +
		time.Minute*time.Duration(20))
	slot1 := models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          models.BIKE,
			IdentificationNumber: "XY1",
		},
		StartTime: t1,
		EndTime:   t2,
	}
	bookingService.TryBookingSlot(&slot1)
	regularRevenue, vipRevenue := revenueService.CalculateRevenue()
	expected1 := (duration + 1) * models.GetChargeGivenTrackAndVehicleType(vehicle, models.REGULAR)
	expected2 := 0
	if regularRevenue != expected1 || vipRevenue != expected2 {
		t.Errorf(error_format_string, "TestCalculateRevenueOverNormalTimeCharedNextHour",
			slot1, regularRevenue, expected1)
		t.Errorf(error_format_string, "TestCalculateRevenueOverNormalTimeCharedNextHour",
			slot1, vipRevenue, expected2)
	}
}

func initMockRaceTrackManagement() models.RaceTrackManagement {
	mock := builders.RaceTrackManagementBuilder{
		RaceTracks: make([]*models.RaceTrack, 0),
	}
	mock.AddRacetrackForVechicleAndRacetrackType(models.BIKE, models.REGULAR, 4)
	mock.AddRacetrackForVechicleAndRacetrackType(models.CAR, models.REGULAR, 2)
	mock.AddRacetrackForVechicleAndRacetrackType(models.CAR, models.VIP, 1)
	mock.AddRacetrackForVechicleAndRacetrackType(models.SUV, models.REGULAR, 2)
	mock.AddRacetrackForVechicleAndRacetrackType(models.SUV, models.VIP, 1)
	return mock.BuildRacetrack()
}

package builders

import (
	"geektrust/models"
	"testing"
)

func TestRacetrackManagementBuilder(t *testing.T) {
	raceTrackManagementBuilder := RaceTrackManagementBuilder{
		RaceTracks: make([]*models.RaceTrack, 0),
	}
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(
		models.BIKE, models.REGULAR, 2,
	)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(
		models.CAR, models.VIP, 1,
	)
	raceTrackManagement := raceTrackManagementBuilder.BuildRacetrack()
	got := len(raceTrackManagement.RaceTracks)
	want := 3
	if got != want {
		t.Errorf("Failed in TestRacetrackManagementBuilder. Got: %v, Want: %v", got, want)
	}
}

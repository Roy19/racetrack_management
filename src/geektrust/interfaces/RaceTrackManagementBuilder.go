package interfaces

import "geektrust/models"

type IRaceTrackManagementBuilder interface {
	AddRacetrackForVechicleAndRacetrackType(
		vechicleType models.VehicleType,
		racetrackType models.RacetrackType,
		times int) IRaceTrackManagementBuilder
	BuildRacetrack() models.RaceTrackManagement
}

type RaceTrackManagementBuilder struct {
	RaceTracks []*models.RaceTrack
}

func (rcm *RaceTrackManagementBuilder) AddRacetrackForVechicleAndRacetrackType(
	vechicleType models.VehicleType,
	racetrackType models.RacetrackType,
	times int) IRaceTrackManagementBuilder {
	for idx := 1; idx <= times; idx++ {
		raceTrackForVehicleType := &models.RaceTrack{
			AllowedVehicleType: vechicleType,
			RaceTrackType:      racetrackType,
			BookedSlots:        make([]*models.BookedSlot, 0),
		}
		rcm.RaceTracks = append(rcm.RaceTracks, raceTrackForVehicleType)
	}
	return rcm
}

func (rcm *RaceTrackManagementBuilder) BuildRacetrack() models.RaceTrackManagement {
	raceTrackManagement := models.RaceTrackManagement{}
	raceTrackManagement.RaceTracks = rcm.RaceTracks
	return raceTrackManagement
}

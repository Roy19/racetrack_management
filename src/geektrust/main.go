package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"geektrust/builders"
	"geektrust/commands"
	"geektrust/controllers"
	"geektrust/models"
	"geektrust/services"
)

const (
	regularBikeTracks int = 4
	regularCarTracks  int = 2
	regularSuvTracks  int = 2

	vipCarTracks int = 1
	vipSuvTracks int = 1
)

func initServiceContainer() commands.CommandExecutor {
	raceTrackManagementBuilder := builders.RaceTrackManagementBuilder{
		RaceTracks: make([]*models.RaceTrack, 0),
	}
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.BIKE, models.REGULAR, regularBikeTracks)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.CAR, models.REGULAR, regularCarTracks)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.CAR, models.VIP, vipCarTracks)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.SUV, models.REGULAR, regularSuvTracks)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.SUV, models.VIP, vipSuvTracks)
	raceTrackManagement := raceTrackManagementBuilder.BuildRacetrack()
	bookingService := services.BookingService{
		RaceTrackManagement: &raceTrackManagement,
	}
	revenueService := services.RevenueService{
		RaceTrackManagement: &raceTrackManagement,
	}
	bookingController := controllers.BookingController{
		BookingService: &bookingService,
	}
	revenueController := controllers.RevenueController{
		RevenueService: &revenueService,
	}
	commandExecutor := commands.CommandExecutor{
		BookingController: &bookingController,
		RevenueController: &revenueController,
	}
	return commandExecutor
}

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")
		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	commandExecutor := initServiceContainer()
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		commandExecutor.ExecutorCommand(tokens)
	}
}

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

func initServiceContainer() commands.CommandExecutor {
	raceTrackManagementBuilder := builders.RaceTrackManagementBuilder{
		RaceTracks: make([]*models.RaceTrack, 0),
	}
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.BIKE, models.REGULAR, 4)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.CAR, models.REGULAR, 2)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.CAR, models.VIP, 1)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.SUV, models.REGULAR, 2)
	raceTrackManagementBuilder.AddRacetrackForVechicleAndRacetrackType(models.SUV, models.VIP, 1)
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
		//fmt.Println(tokens)
		commandExecutor.ExecutorCommand(tokens)
	}
}

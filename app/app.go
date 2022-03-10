package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"taxi-fares/app/model"
	"taxi-fares/app/util"
)

func Run() {
	fmt.Println("Application Start")
	defer fmt.Println("Application Stop")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input number of line: ")
	scanner.Scan()
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || n < 1 {
		fmt.Println("Input number not valid")
		return
	}

	data := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("Scan stdin error", scanner.Err())
			return
		}
		input := scanner.Text()
		data[i] = input
	}

	fare, err := calculateFares(data)
	if err != nil {
		fmt.Println("Failed calculate taxi fares", err)
	}
	fmt.Println("Taxi Fare: ", fare)
}

func calculateFares(data []string) (fare int, err error) {

	// when there are less than two lines of data
	if len(data) < 2 {
		return 0, errors.New("less than two lines of data")
	}

	prevDistanceMeter := model.DistanceMeter{}
	totalMilage := float64(0)
	for i, dt := range data {
		dt = strings.TrimSpace(dt)
		distanceMeter, err := util.StringToDistanceMeter(dt)
		if err != nil {
			return 0, err
		}

		// check When the past time has been sent
		if i > 0 {
			if distanceMeter.ElapsedTime < prevDistanceMeter.ElapsedTime {
				return 0, fmt.Errorf("past time has been sent. %s < previous %s", distanceMeter.ElapsedTimeStr, prevDistanceMeter.ElapsedTimeStr)
			}

			if (distanceMeter.ElapsedTime - prevDistanceMeter.ElapsedTime).Minutes() > 5 {
				return 0, fmt.Errorf("the interval between records is more than 5 minutes apart. previous %s, current %s", prevDistanceMeter.ElapsedTimeStr, distanceMeter.ElapsedTimeStr)
			}
		}

		if distanceMeter.Milage < 0 {
			return 0, errors.New("milage can't be less than 0")
		}

		totalMilage = distanceMeter.Milage

		prevDistanceMeter = distanceMeter
	}

	if totalMilage == 0 {
		return fare, errors.New("total milage is 0.0m")
	}

	fare = util.CalculateFare(totalMilage)

	return fare, nil
}

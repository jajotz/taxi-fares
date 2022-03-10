package util

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"taxi-fares/app/model"
	"time"
)

const (
	elapsedTimePattern = "^[0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{3}$" // 00:00:00.000 to 99:99:99.999
)

var (
	//1. The base fare is 400 yen for up to 1 km.
	//2. Up to 10 km, 40 yen is added every 400 meters.
	//3. Over 10km, 40 yen is added every 350 meters
	fareTiers = []model.FareTier{
		{
			Start: 0,
			End:   1000,
			Cycle: 0,
			Fare:  400,
		},
		{
			Start: 1000,
			End:   10000,
			Cycle: 400,
			Fare:  40,
		},
		{
			Start: 10000,
			End:   -1, // unlimited
			Cycle: 350,
			Fare:  40,
		},
	}
)

// StringToDistanceMeter - convert string to DistanceMeter model
func StringToDistanceMeter(str string) (distanceMeter model.DistanceMeter, err error) {
	split := strings.Split(str, " ")
	if len(split) != 2 {
		return distanceMeter, errors.New("input not under the format of: hh:mm:ss.fff<SPACE>xxxxxxxx.f<LF>")
	}

	matched, _ := regexp.MatchString(elapsedTimePattern, split[0])
	if !matched {
		return distanceMeter, fmt.Errorf("input first half (elapsed-time) %s is not valid. should between 00:00:00.000 to 99:99:99.999", split[0])
	}
	elapsedSplit := strings.Split(split[0], ".")
	hms := strings.Split(elapsedSplit[0], ":")
	h, _ := strconv.Atoi(hms[0])
	m, _ := strconv.Atoi(hms[1])
	s, _ := strconv.Atoi(hms[2])
	ms, _ := strconv.Atoi(elapsedSplit[1])

	milage, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return distanceMeter, fmt.Errorf("input second half (milage) %s is not valid", split[1])
	}

	distanceMeter = model.DistanceMeter{
		ElapsedTimeStr: split[0],
		ElapsedTime:    time.Duration(h*int(time.Hour)) + time.Duration(m*int(time.Minute)) + time.Duration(s*int(time.Second)) + time.Duration(ms*int(time.Millisecond)),
		Milage:         milage,
	}
	return distanceMeter, nil
}

// CalculateFare based on FareTier
func CalculateFare(milage float64) (fare int) {

	milageLeft := milage
	for _, tier := range fareTiers {
		tierRange := tier.End - tier.Start
		if tierRange < 0 {
			tierRange = tier.Start
		}

		if tier.Cycle == 0 {
			fare += tier.Fare
		} else {
			multiplier := 0

			if milageLeft-tierRange < 0 {
				multiplier = int(math.Ceil(milageLeft / tier.Cycle))
			} else {
				multiplier = int(math.Ceil(tierRange / tier.Cycle))
			}
			fare += multiplier * tier.Fare
		}

		milageLeft -= tierRange
		if milageLeft < 0 {
			break
		}
	}

	return fare
}

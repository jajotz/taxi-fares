package model

import "time"

type (
	DistanceMeter struct {
		ElapsedTimeStr string
		ElapsedTime    time.Duration
		Milage         float64
	}

	FareTier struct {
		Start float64
		End   float64
		Cycle float64
		Fare  int
	}
)

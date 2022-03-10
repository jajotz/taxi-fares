package util

import (
	"errors"
	"taxi-fares/app/model"
	"testing"
	"time"
)

func TestStringToDistanceMeter(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want model.DistanceMeter
		err  error
	}{
		{
			args: args{str: "00:00:00.000 0.0"},
			want: model.DistanceMeter{
				ElapsedTimeStr: "00:00:00.000",
				ElapsedTime:    time.Duration(0),
				Milage:         0,
			},
		},
		{
			args: args{str: "00:01:00.123 480.9"},
			want: model.DistanceMeter{
				ElapsedTimeStr: "00:01:00.123",
				ElapsedTime:    time.Duration(1*time.Minute) + time.Duration(123*time.Millisecond),
				Milage:         480.9,
			},
		},
		{
			args: args{str: "abc"},
			want: model.DistanceMeter{},
			err:  errors.New("input not under the format of: hh:mm:ss.fff<SPACE>xxxxxxxx.f<LF>"),
		},
		{
			args: args{str: "abc 123"},
			want: model.DistanceMeter{},
			err:  errors.New("input first half (elapsed-time) abc is not valid. should between 00:00:00.000 to 99:99:99.999"),
		},
		{
			args: args{str: "00:01:00.123 abc"},
			want: model.DistanceMeter{},
			err:  errors.New("input second half (milage) abc is not valid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToDistanceMeter(tt.args.str)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Error(err)
				}
			}

			if got != tt.want {
				t.Errorf("StringToDistanceMeter() = %v, want %v", got, tt.want)

			}
		})
	}
}
func TestCalculateFare(t *testing.T) {
	type args struct {
		milage float64
	}
	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{
			args: args{milage: 0},
			want: 400,
		},
		{
			args: args{milage: 1000},
			want: 400,
		},
		{
			args: args{milage: 1001},
			want: 440,
		},
		{
			args: args{milage: 1400},
			want: 440,
		},
		{
			args: args{milage: 1401},
			want: 480,
		},
		{
			args: args{milage: 1800.8},
			want: 520,
		},
		{
			args: args{milage: 9999},
			want: 1320,
		},
		{
			args: args{milage: 10000},
			want: 1320,
		},
		{
			args: args{milage: 10001},
			want: 1360,
		},
		{
			args: args{milage: 20000},
			want: 2480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateFare(tt.args.milage)
			if got != tt.want {
				t.Errorf("CalculateFare() = %v, want %v", got, tt.want)

			}
		})
	}
}

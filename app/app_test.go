package app

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:01:00.123 480.9",
				"00:02:00.125 1141.2",
				"00:03:00.100 1800.8"}},
			want: 520,
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0"}},
			want: 0,
			err:  errors.New("less than two lines of data"),
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:02:00.000 100.0",
				"00:01:00.000 200.0"}},
			want: 0,
			err:  errors.New("past time has been sent. 00:01:00.000 < previous 00:02:00.000"),
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:02:00.000 100.0",
				"00:07:00.001 200.0"}},
			want: 0,
			err:  errors.New("the interval between records is more than 5 minutes apart. previous 00:02:00.000, current 00:07:00.001"),
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:01:00.123 0.0"}},
			want: 0,
			err:  errors.New("total milage is 0.0m"),
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:01:00.123 -10.0"}},
			want: 0,
			err:  errors.New("milage can't be less than 0"),
		},
		{
			args: args{data: []string{
				"00:00:00.abc 0.0",
				"00:01:00.123 -10.0"}},
			want: 0,
			err:  errors.New("input first half (elapsed-time) 00:00:00.abc is not valid. should between 00:00:00.000 to 99:99:99.999"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run()
		})
	}
}

func TestCalculateFares(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:01:00.123 480.9",
				"00:02:00.125 1141.2",
				"00:03:00.100 1800.8"}},
			want: 520,
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0"}},
			want: 0,
			err:  errors.New("less than two lines of data"),
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:02:00.000 100.0",
				"00:01:00.000 200.0"}},
			want: 0,
			err:  errors.New("past time has been sent. 00:01:00.000 < previous 00:02:00.000"),
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:02:00.000 100.0",
				"00:07:00.001 200.0"}},
			want: 0,
			err:  errors.New("the interval between records is more than 5 minutes apart. previous 00:02:00.000, current 00:07:00.001"),
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:01:00.123 0.0"}},
			want: 0,
			err:  errors.New("total milage is 0.0m"),
		},
		{
			args: args{data: []string{
				"00:00:00.000 0.0",
				"00:01:00.123 -10.0"}},
			want: 0,
			err:  errors.New("milage can't be less than 0"),
		},
		{
			args: args{data: []string{
				"00:00:00.abc 0.0",
				"00:01:00.123 -10.0"}},
			want: 0,
			err:  errors.New("input first half (elapsed-time) 00:00:00.abc is not valid. should between 00:00:00.000 to 99:99:99.999"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateFares(tt.args.data)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Error(err)
				}
			}

			if got != tt.want {
				t.Errorf("calculateFares() = %v, want %v", got, tt.want)

			}
		})
	}
}

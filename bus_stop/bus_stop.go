package bus_stop

import (
	"errors"
	"time"
)

func Time(hours int, minutes int, stop_number int) [][]int {
	estimates := [][]int{}

	bus_stops := 10
	routes := 3

	estimates = append(estimates, []int{0, 1})
	estimates = append(estimates, []int{2, 3})
	estimates = append(estimates, []int{4, 5})

	each_stop_serviced_every_mins_per_route := 15
	service_hours_per_day = 24
	route_start_running_mins_after_previous_one := 2
	stop_mins_away_previous_one = 2

	chicago, err := time.LoadLocation("America/Chicago")

	if err != nil {
		fmt.Println(err)
		return
	}

	now := time.Now().In(chicago)
	year, month, day := now.Date()

	
	route_start_time := 

	return estimates
}

func EstimateArrival(stop_number int) ([][]int, error) {
	bus_stops := 10

	if stop_number < 1 || stop_number > bus_stops {
		return [][]int{}, errors.New("invalid stop number")
	}

	hours, minutes, _ := time.Now().Clock()

	next_two_arrival_times_per_route := Time(hours, minutes, 0)

	return next_two_arrival_times_per_route, nil
}

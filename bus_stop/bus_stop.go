package bus_stop

import (
	"errors"
	"fmt"
	"time"
)

func Time(hours int, minutes int, stop_number int) [][]int {
	estimates := [][]int{}

	routes := 3
	bus_stops := 10

	// each_stop_serviced_every_mins_per_route := 15
	// service_hours_per_day = 24
	route_start_running_mins_after_previous_one := 2
	stop_mins_away_previous_one := 2

	location, err := time.LoadLocation("America/New_York")

	if err != nil {
		fmt.Println(err)
		return estimates
	}

	now := time.Now().In(location)
	year, month, day := now.Date()

	today_at_midnight := time.Date(year, month, day, 0, 0, 0, 0, location)

	start_time := today_at_midnight

	duration_since_start := time.Since(start_time)
	minutes_since_start := int(duration_since_start.Minutes())
	arrival_time_1 := minutes_since_start
	arrival_time_2 := minutes_since_start

	for route_index := 0; route_index < routes; route_index++ {
		arrival_time_1 = arrival_time_1 + route_index*route_start_running_mins_after_previous_one
		arrival_time_2 = arrival_time_2 + route_index*route_start_running_mins_after_previous_one

		for bus_stop_number := 1; bus_stop_number <= bus_stops; bus_stop_number++ {
			// arrival_time_1 = arrival_time_1 + (bus_stop_number-1)*(minutes_since_start/stop_mins_away_previous_one)
			arrival_time_2 = arrival_time_2 + (bus_stop_number-1)*stop_mins_away_previous_one //each_stop_serviced_every_mins_per_route
		}

		estimates = append(estimates, []int{arrival_time_1 - minutes_since_start, arrival_time_2 - minutes_since_start})
	}

	return estimates
}

func EstimateArrival(stop_number int) ([][]int, error) {
	bus_stops := 10

	if stop_number < 1 || stop_number > bus_stops {
		return [][]int{}, errors.New("invalid stop number")
	}

	hours, minutes, _ := time.Now().Clock()

	next_two_arrival_times_per_route := Time(hours, minutes, stop_number)

	return next_two_arrival_times_per_route, nil
}

package bus_stop

import (
	"testing"
)

func TestTime(t *testing.T) {
	got := Time(0, 0, 1)

	route_1 := got[0]
	route_2 := got[1]
	route_3 := got[2]

	route_1_arrival_time_1 := route_1[0]
	route_2_arrival_time_1 := route_2[0]
	route_3_arrival_time_1 := route_3[0]

	route_1_arrival_time_2 := route_1[1]
	route_2_arrival_time_2 := route_2[1]
	route_3_arrival_time_2 := route_3[1]

	want_route_1_arrival_time_1 := 18
	want_route_1_arrival_time_2 := 1090

	want_route_2_arrival_time_1 := 18
	want_route_2_arrival_time_2 := 1090

	want_route_3_arrival_time_1 := 18
	want_route_3_arrival_time_2 := 1090

	if route_1_arrival_time_1 != want_route_1_arrival_time_1 {
		t.Errorf("Time(0, 0, 1) = %d; want route_1_arrival_time_1 = %d", route_1_arrival_time_1, want_route_1_arrival_time_1)
	}
	if route_2_arrival_time_1 != want_route_2_arrival_time_1 {
		t.Errorf("Time(0, 0, 1) = %d; want route_2_arrival_time_1 = %d", route_2_arrival_time_1, want_route_2_arrival_time_1)
	}
	if route_3_arrival_time_1 != want_route_3_arrival_time_1 {
		t.Errorf("Time(0, 0, 1) = %d; want route_3_arrival_time_1 = %d", route_3_arrival_time_1, want_route_3_arrival_time_1)
	}
	if route_1_arrival_time_2 != want_route_1_arrival_time_2 {
		t.Errorf("Time(0, 0, 1) = %d; want route_1_arrival_time_2 = %d", route_1_arrival_time_2, want_route_1_arrival_time_2)
	}
	if route_2_arrival_time_2 != want_route_2_arrival_time_2 {
		t.Errorf("Time(0, 0, 1) = %d; want route_2_arrival_time_2 = %d", route_2_arrival_time_2, want_route_2_arrival_time_2)
	}
	if route_3_arrival_time_2 != want_route_3_arrival_time_2 {
		t.Errorf("Time(0, 0, 1) = %d; want route_3_arrival_time_2 = %d", route_3_arrival_time_2, want_route_3_arrival_time_2)
	}
}

// func TestEstimateArrival(t *testing.T) {
// 	got, _ := EstimateArrival(1)
// 	if got[0][0] != 0 {
// 		t.Errorf("EstimateArrival(1) = %d; want 0", got)
// 	}
// }

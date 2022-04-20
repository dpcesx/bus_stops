package bus_stop

import (
	"testing"
)

func TestTime(t *testing.T) {
	got := Time(0, 0, 1)
	if got[0][0] != 0 {
		t.Errorf("Time(0, 0, 1) = %d; want 0", got)
	}
}

func TestEstimateArrival(t *testing.T) {
	got, _ := EstimateArrival(1)
	if got[0][0] != 0 {
		t.Errorf("EstimateArrival(1) = %d; want 0", got)
	}
}

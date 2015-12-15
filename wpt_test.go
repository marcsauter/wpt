package wpt

import "testing"

func TestTakeOff(t *testing.T) {
	wps, err := NewWaypoints("./Waypoints_Startplatz.gpx")
	if err != nil {
		t.Error(err)
	}
	name, _ := wps.Find(46.70996411, 7.77341942)
	if name != "Niederhorn" {
		t.Errorf("should be \"Niederhorn\" - got %s\n", name)
	}
}

func TestLanding(t *testing.T) {
	wps, err := NewWaypoints("./Waypoints_Landeplatz.gpx")
	if err != nil {
		t.Error(err)
	}
	name, _ := wps.Find(46.6810178, 7.82421503)
	if name != "Unterseen Lehn" {
		t.Errorf("should be \"Unterseen Lehn\" - got %s\n", name)
	}
}

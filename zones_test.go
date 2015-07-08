package timezone

import "testing"

var expectedZones = []Timezone{
	Timezone{"Europe/Helsinki", "FI"},
}

func TestZones(t *testing.T) {
	if Locations[0] != expectedZones[0] {
		t.Errorf("ZonesTest returned %s", Locations[0])
	}

	if len(Locations) != 416 {
		t.Errorf("Length of zones was not 3, actual %d", len(Locations))
	}
}

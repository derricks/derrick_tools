package cmd

import (
	"testing"
)

func TestConstellationsWithStars(t *testing.T) {
	sampleConstellations := []constellation{
		constellation{"Andromeda", "Princess of Ethiopia", []string{"Alpheratz", "Mirach", "Almach", "Sadiradra", "Nembus", "Titawin", "Keff al Salsalat", "Adhil", "Veritate"}},
		constellation{"Antlia", "Air Pump", []string{"Macondo"}},
		constellation{"Apus", "Bird of Paradise", []string{}},
		constellation{"Aquarius", "Water Bearer", []string{"Some star"}},
		constellation{"Aquila", "Eagle", []string{}},
	}

	withStars := constellationsWithStars(sampleConstellations)
	if len(withStars) != 3 {
		t.Fatalf("Constellations with stars should have 3 elements but has %d", len(withStars))
	}
}

package cmd

import (
	"testing"
)

func TestConstellationsWithStars(t *testing.T) {
	sampleConstellations := []constellation{
		constellation{1, "Andromeda", "Princess of Ethiopia", []string{"Alpheratz", "Mirach", "Almach", "Sadiradra", "Nembus", "Titawin", "Keff al Salsalat", "Adhil", "Veritate"}},
		constellation{2, "Antlia", "Air Pump", []string{"Macondo"}},
		constellation{3, "Apus", "Bird of Paradise", []string{}},
		constellation{4, "Aquarius", "Water Bearer", []string{"Some star"}},
		constellation{5, "Aquila", "Eagle", []string{}},
	}

	withStars := constellationsWithStars(sampleConstellations)
	if len(withStars) != 3 {
		t.Fatalf("Constellations with stars should have 3 elements but has %d", len(withStars))
	}
}

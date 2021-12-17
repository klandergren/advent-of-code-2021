package day09

import "testing"

var testLines = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
}

func Test_AdjacentHeights_00(t *testing.T) {
	hm, err := NewHeightMap(testLines)

	if err != nil {
		t.Error(err)
	}

	aHts := hm.AdjacentHeights(0, 0)

	if len(aHts) != 2 {
		t.Errorf("adjacent heights for (0,0) should be 2, got: %d", len(aHts))
	}

	if south := aHts[0]; south != 3 {
		t.Errorf("first adjacent height for (0,0) should be 3, got: %d", south)
	}

	if east := aHts[1]; east != 1 {
		t.Errorf("second adjacent height for (0,0) should be 1, got: %d", east)
	}

}

func Test_AdjacentHeights_33(t *testing.T) {
	hm, err := NewHeightMap(testLines)

	if err != nil {
		t.Error(err)
	}

	aHts := hm.AdjacentHeights(3, 3)

	if len(aHts) != 4 {
		t.Errorf("adjacent heights for (3,3) should be 4, got: %d", len(aHts))
	}

	if north := aHts[0]; north != 6 {
		t.Errorf("first adjacent height for (3,3) should be 6, got: %d", north)
	}

	if west := aHts[1]; west != 6 {
		t.Errorf("second adjacent height for (3,3) should be 6, got: %d", west)
	}

	if south := aHts[2]; south != 9 {
		t.Errorf("third adjacent height for (3,3) should be 9, got: %d", south)
	}

	if east := aHts[3]; east != 8 {
		t.Errorf("fourth adjacent height for (3,3) should be 8, got: %d", east)
	}

}

func Test_AdjacentHeights_50(t *testing.T) {
	hm, err := NewHeightMap(testLines)

	if err != nil {
		t.Error(err)
	}

	aHts := hm.AdjacentHeights(5, 0)

	if len(aHts) != 3 {
		t.Errorf("adjacent heights for (5,0) should be 3, got: %d", len(aHts))
	}

	if west := aHts[0]; west != 9 {
		t.Errorf("first adjacent height for (3,3) should be 9, got: %d", west)
	}

	if south := aHts[1]; south != 9 {
		t.Errorf("second adjacent height for (3,3) should be 9, got: %d", south)
	}

	if east := aHts[2]; east != 3 {
		t.Errorf("third adjacent height for (3,3) should be 3, got: %d", east)
	}

}

func Test_LowPointHeights(t *testing.T) {
	hm, err := NewHeightMap(testLines)

	if err != nil {
		t.Error(err)
	}

	lPtHts := hm.LowPointHeights()

	if len(lPtHts) != 4 {
		t.Errorf("low point heights for test height map should be 4, got: %d", len(lPtHts))
	}
}

func Test_SumRisk(t *testing.T) {
	hm, err := NewHeightMap(testLines)

	if err != nil {
		t.Error(err)
	}

	sumRisk := hm.SumRisk()

	if sumRisk != 15 {
		t.Errorf("sum risk for test height map should be 15, got: %d", sumRisk)
	}
}

func Test_BasinSizes(t *testing.T) {
	hm, err := NewHeightMap(testLines)

	if err != nil {
		t.Error(err)
	}

	bs := hm.BasinSizes()

	if len(bs) != 4 {
		t.Errorf("basin sizes length should be 4, got: %d", len(bs))
	}
}

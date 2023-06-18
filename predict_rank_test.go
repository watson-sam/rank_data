package rank_data_test

import (
	"reflect"
	"testing"

	"rank_data"
)

func TestRankAverage(t *testing.T) {
	// Test case 1
	data1 := []float64{4, 2, 7, 2, 9, 5, 1}
	expectedRanks1 := []float64{4, 2.5, 6, 2.5, 7, 5, 1}
	ranks1 := rank_data.RankAverage(data1)
	if !reflect.DeepEqual(ranks1, expectedRanks1) {
		t.Errorf(
			"RankAverage failed for test case 1. Expected %v, but got %v",
			expectedRanks1,
			ranks1,
		)
	}

	// Test case 2: Negative values and duplicates
	data2 := []float64{-2, 4, 7, 2, 4, -2, 9}
	expectedRanks2 := []float64{1.5, 4.5, 6, 3, 4.5, 1.5, 7}
	ranks2 := rank_data.RankAverage(data2)
	if !reflect.DeepEqual(ranks2, expectedRanks2) {
		t.Errorf(
			"RankAverage failed for test case 2. Expected %v, but got %v",
			expectedRanks2,
			ranks2,
		)
	}

	// Test case 3: Test with empty data
	data3 := []float64{}
	expectedRanks3 := []float64{}
	ranks3 := rank_data.RankAverage(data3)
	if !reflect.DeepEqual(ranks3, expectedRanks3) {
		t.Errorf(
			"RankAverage failed for test case 3. Expected %v, but got %v",
			expectedRanks3,
			ranks3,
		)
	}

	// Test case 4: Test with identical values
	data4 := []float64{5, 5, 5, 5, 5}
	expectedRanks4 := []float64{3, 3, 3, 3, 3}
	ranks4 := rank_data.RankAverage(data4)
	if !reflect.DeepEqual(ranks4, expectedRanks4) {
		t.Errorf(
			"RankAverage failed for test case 4. Expected %v, but got %v",
			expectedRanks4,
			ranks4,
		)
	}

	// Test case 5
	data5 := []float64{9, 2, 8, 6, 3, 1, 4, 7, 5, 10, 4, 5}
	expectedRanks5 := []float64{11, 2, 10, 8, 3, 1, 4.5, 9, 6.5, 12, 4.5, 6.5}
	ranks5 := rank_data.RankAverage(data5)
	if !reflect.DeepEqual(ranks5, expectedRanks5) {
		t.Errorf(
			"RankAverage failed for test case 5. Expected %v, but got %v",
			expectedRanks5,
			ranks5,
		)
	}
}

func TestRankMin(t *testing.T) {
	// Test case 1
	data1 := []float64{4, 2, 7, 2, 9, 5, 1}
	expectedRanks1 := []float64{4, 2, 6, 2, 7, 5, 1}
	ranks1 := rank_data.RankMin(data1)
	if !reflect.DeepEqual(ranks1, expectedRanks1) {
		t.Errorf("RankMin failed for test case 1. Expected %v, but got %v", expectedRanks1, ranks1)
	}

	// Test case 2: Negative values and duplicates
	data2 := []float64{-2, 4, 7, 2, 4, -2, 9}
	expectedRanks2 := []float64{1, 4, 6, 3, 4, 1, 7}
	ranks2 := rank_data.RankMin(data2)
	if !reflect.DeepEqual(ranks2, expectedRanks2) {
		t.Errorf("RankMin failed for test case 2. Expected %v, but got %v", expectedRanks2, ranks2)
	}

	// Test case 3: Test with empty data
	data3 := []float64{}
	expectedRanks3 := []float64{}
	ranks3 := rank_data.RankMin(data3)
	if !reflect.DeepEqual(ranks3, expectedRanks3) {
		t.Errorf("RankMin failed for test case 3. Expected %v, but got %v", expectedRanks3, ranks3)
	}

	// Test case 4: Test with identical values
	data4 := []float64{5, 5, 5, 5, 5}
	expectedRanks4 := []float64{1, 1, 1, 1, 1}
	ranks4 := rank_data.RankMin(data4)
	if !reflect.DeepEqual(ranks4, expectedRanks4) {
		t.Errorf("RankMin failed for test case 4. Expected %v, but got %v", expectedRanks4, ranks4)
	}

	// Test case 5
	data5 := []float64{9, 2, 8, 6, 3, 1, 4, 7, 5, 10, 4, 5}
	expectedRanks5 := []float64{11, 2, 10, 8, 3, 1, 4, 9, 6, 12, 4, 6}
	ranks5 := rank_data.RankMin(data5)
	if !reflect.DeepEqual(ranks5, expectedRanks5) {
		t.Errorf("RankMin failed for test case 5. Expected %v, but got %v", expectedRanks5, ranks5)
	}
}

func TestRankMax(t *testing.T) {
	// Test case 1
	data1 := []float64{4, 2, 7, 2, 9, 5, 1}
	expectedRanks1 := []float64{4, 3, 6, 3, 7, 5, 1}
	ranks1 := rank_data.RankMax(data1)
	if !reflect.DeepEqual(ranks1, expectedRanks1) {
		t.Errorf("RankMax failed for test case 1. Expected %v, but got %v", expectedRanks1, ranks1)
	}

	// Test case 2: Negative values and duplicates
	data2 := []float64{-2, 4, 7, 2, 4, -2, 9}
	expectedRanks2 := []float64{2, 5, 6, 3, 5, 2, 7}
	ranks2 := rank_data.RankMax(data2)
	if !reflect.DeepEqual(ranks2, expectedRanks2) {
		t.Errorf("RankMax failed for test case 2. Expected %v, but got %v", expectedRanks2, ranks2)
	}

	// Test case 3: Test with empty data
	data3 := []float64{}
	expectedRanks3 := []float64{}
	ranks3 := rank_data.RankMax(data3)
	if !reflect.DeepEqual(ranks3, expectedRanks3) {
		t.Errorf("RankMax failed for test case 3. Expected %v, but got %v", expectedRanks3, ranks3)
	}

	// Test case 4: Test with identical values
	data4 := []float64{5, 5, 5, 5, 5}
	expectedRanks4 := []float64{5, 5, 5, 5, 5}
	ranks4 := rank_data.RankMax(data4)
	if !reflect.DeepEqual(ranks4, expectedRanks4) {
		t.Errorf("RankMax failed for test case 4. Expected %v, but got %v", expectedRanks4, ranks4)
	}

	// Test case 5
	data5 := []float64{9, 2, 8, 6, 3, 1, 4, 7, 5, 10, 4, 5}
	expectedRanks5 := []float64{11, 2, 10, 8, 3, 1, 5, 9, 7, 12, 5, 7}
	ranks5 := rank_data.RankMax(data5)
	if !reflect.DeepEqual(ranks5, expectedRanks5) {
		t.Errorf("RankMax failed for test case 5. Expected %v, but got %v", expectedRanks5, ranks5)
	}
}
func TestRankDence(t *testing.T) {
	// Test case 1
	data1 := []float64{4, 2, 7, 2, 9, 5, 1}
	expectedRanks1 := []float64{3, 2, 5, 2, 6, 4, 1}
	ranks1 := rank_data.RankDense(data1)
	if !reflect.DeepEqual(ranks1, expectedRanks1) {
		t.Errorf(
			"RankDense failed for test case 1. Expected %v, but got %v",
			expectedRanks1,
			ranks1,
		)
	}

	// Test case 2: Negative values and duplicates
	data2 := []float64{-2, 4, 7, 2, 4, -2, 9}
	expectedRanks2 := []float64{1, 3, 4, 2, 3, 1, 5}
	ranks2 := rank_data.RankDense(data2)
	if !reflect.DeepEqual(ranks2, expectedRanks2) {
		t.Errorf(
			"RankDense failed for test case 2. Expected %v, but got %v",
			expectedRanks2,
			ranks2,
		)
	}

	// Test case 3: Test with empty data
	data3 := []float64{}
	expectedRanks3 := []float64{}
	ranks3 := rank_data.RankDense(data3)
	if !reflect.DeepEqual(ranks3, expectedRanks3) {
		t.Errorf(
			"RankDense failed for test case 3. Expected %v, but got %v",
			expectedRanks3,
			ranks3,
		)
	}

	// Test case 4: Test with identical values
	data4 := []float64{5, 5, 5, 5, 5}
	expectedRanks4 := []float64{1, 1, 1, 1, 1}
	ranks4 := rank_data.RankDense(data4)
	if !reflect.DeepEqual(ranks4, expectedRanks4) {
		t.Errorf(
			"RankDense failed for test case 4. Expected %v, but got %v",
			expectedRanks4,
			ranks4,
		)
	}

	// Test case 5
	data5 := []float64{9, 2, 8, 6, 3, 1, 4, 7, 5, 10, 4, 5}
	expectedRanks5 := []float64{9, 2, 8, 6, 3, 1, 4, 7, 5, 10, 4, 5}
	ranks5 := rank_data.RankDense(data5)
	if !reflect.DeepEqual(ranks5, expectedRanks5) {
		t.Errorf(
			"RankDense failed for test case 5. Expected %v, but got %v",
			expectedRanks5,
			ranks5,
		)
	}
}

func TestRankOrdinal(t *testing.T) {
	// Test case 1
	data1 := []float64{4, 2, 7, 2, 9, 5, 1}
	expectedRanks1 := []float64{4, 2, 6, 3, 7, 5, 1}
	ranks1 := rank_data.RankOrdinal(data1)
	if !reflect.DeepEqual(ranks1, expectedRanks1) {
		t.Errorf(
			"RankOrdinal failed for test case 1. Expected %v, but got %v",
			expectedRanks1,
			ranks1,
		)
	}

	// Test case 2: Negative values and duplicates
	data2 := []float64{-2, 4, 7, 2, 4, -2, 9}
	expectedRanks2 := []float64{1, 4, 6, 3, 5, 2, 7}
	ranks2 := rank_data.RankOrdinal(data2)
	if !reflect.DeepEqual(ranks2, expectedRanks2) {
		t.Errorf(
			"RankOrdinal failed for test case 2. Expected %v, but got %v",
			expectedRanks2,
			ranks2,
		)
	}

	// Test case 3: Test with empty data
	data3 := []float64{}
	expectedRanks3 := []float64{}
	ranks3 := rank_data.RankOrdinal(data3)
	if !reflect.DeepEqual(ranks3, expectedRanks3) {
		t.Errorf(
			"RankOrdinal failed for test case 3. Expected %v, but got %v",
			expectedRanks3,
			ranks3,
		)
	}

	// Test case 4: Test with identical values
	data4 := []float64{5, 5, 5, 5, 5}
	expectedRanks4 := []float64{1, 2, 3, 4, 5}
	ranks4 := rank_data.RankOrdinal(data4)
	if !reflect.DeepEqual(ranks4, expectedRanks4) {
		t.Errorf(
			"RankOrdinal failed for test case 4. Expected %v, but got %v",
			expectedRanks4,
			ranks4,
		)
	}

	// Test case 5
	data5 := []float64{9, 2, 8, 6, 3, 1, 4, 7, 5, 10, 4, 5}
	expectedRanks5 := []float64{11, 2, 10, 8, 3, 1, 4, 9, 6, 12, 5, 7}
	ranks5 := rank_data.RankOrdinal(data5)
	if !reflect.DeepEqual(ranks5, expectedRanks5) {
		t.Errorf(
			"RankOrdinal failed for test case 5. Expected %v, but got %v",
			expectedRanks5,
			ranks5,
		)
	}
}

type case_ struct {
	Values   []float64
	Method   string
	Expected []float64
}

func TestRankData(t *testing.T) {
	cases := []case_{
		{[]float64{}, "average", []float64{}},
		{[]float64{}, "min", []float64{}},
		{[]float64{}, "max", []float64{}},
		{[]float64{}, "dense", []float64{}},
		{[]float64{}, "ordinal", []float64{}},
		//
		{[]float64{100}, "average", []float64{1}},
		{[]float64{100}, "min", []float64{1}},
		{[]float64{100}, "max", []float64{1}},
		{[]float64{100}, "dense", []float64{1}},
		{[]float64{100}, "ordinal", []float64{1}},
		//
		{[]float64{100, 100, 100}, "average", []float64{2.0, 2.0, 2.0}},
		{[]float64{100, 100, 100}, "min", []float64{1.0, 1.0, 1.0}},
		{[]float64{100, 100, 100}, "max", []float64{3.0, 3.0, 3.0}},
		{[]float64{100, 100, 100}, "dense", []float64{1.0, 1.0, 1.0}},
		{[]float64{100, 100, 100}, "ordinal", []float64{1.0, 2.0, 3.0}},
		//
		{[]float64{100, 300, 200}, "average", []float64{1.0, 3.0, 2.0}},
		{[]float64{100, 300, 200}, "min", []float64{1.0, 3.0, 2.0}},
		{[]float64{100, 300, 200}, "max", []float64{1.0, 3.0, 2.0}},
		{[]float64{100, 300, 200}, "dense", []float64{1.0, 3.0, 2.0}},
		{[]float64{100, 300, 200}, "ordinal", []float64{1.0, 3.0, 2.0}},
		//
		{[]float64{100, 200, 300, 200}, "average", []float64{1.0, 2.5, 4.0, 2.5}},
		{[]float64{100, 200, 300, 200}, "min", []float64{1.0, 2.0, 4.0, 2.0}},
		{[]float64{100, 200, 300, 200}, "max", []float64{1.0, 3.0, 4.0, 3.0}},
		{[]float64{100, 200, 300, 200}, "dense", []float64{1.0, 2.0, 3.0, 2.0}},
		{[]float64{100, 200, 300, 200}, "ordinal", []float64{1.0, 2.0, 4.0, 3.0}},
		//
		{[]float64{100, 200, 300, 200, 100}, "average", []float64{1.5, 3.5, 5.0, 3.5, 1.5}},
		{[]float64{100, 200, 300, 200, 100}, "min", []float64{1.0, 3.0, 5.0, 3.0, 1.0}},
		{[]float64{100, 200, 300, 200, 100}, "max", []float64{2.0, 4.0, 5.0, 4.0, 2.0}},
		{[]float64{100, 200, 300, 200, 100}, "dense", []float64{1.0, 2.0, 3.0, 2.0, 1.0}},
		{[]float64{100, 200, 300, 200, 100}, "ordinal", []float64{1.0, 3.0, 5.0, 4.0, 2.0}},
	}

	for i, c := range cases {
		ranks := rank_data.RankData(c.Values, c.Method)
		if !reflect.DeepEqual(ranks, c.Expected) {
			t.Errorf(
				"RankData failed for test case %v. Expected %v, but got %v",
				i,
				c.Expected,
				ranks,
			)
		}
	}
}

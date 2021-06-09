package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestThreeLineStrike(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{98, 90, 82, 62}, series.Float, "Open"),
		series.New([]float64{100, 95, 86, 103}, series.Float, "High"),
		series.New([]float64{80, 73, 67, 101}, series.Float, "Close"),
		series.New([]float64{77, 68, 65, 59}, series.Float, "Low"),
	)
	result_bear := ThreeLineStrike(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestThreeLineStrike Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{70, 80, 95, 125}, series.Float, "Open"),
		series.New([]float64{100, 110, 120, 130}, series.Float, "High"),
		series.New([]float64{90, 105, 115, 55}, series.Float, "Close"),
		series.New([]float64{60, 75, 93, 50}, series.Float, "Low"),
	)

	result_bull := ThreeLineStrike(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestThreeLineStrike Bull unexpected result: got %v want %v", result_bull, expected_bull)
	}

}

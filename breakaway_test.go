package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestBreakAway(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{70, 115, 120, 125, 125}, series.Float, "Open"),
		series.New([]float64{85, 125, 130, 135, 130}, series.Float, "High"),
		series.New([]float64{80, 120, 125, 130, 100}, series.Float, "Close"),
		series.New([]float64{60, 110, 115, 120, 95}, series.Float, "Low"),
	)
	result_bear := BreakAway(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestBreakAway Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{130, 100, 95, 90, 85}, series.Float, "Open"),
		series.New([]float64{135, 105, 100, 95, 115}, series.Float, "High"),
		series.New([]float64{120, 90, 85, 80, 110}, series.Float, "Close"),
		series.New([]float64{115, 85, 80, 75, 80}, series.Float, "Low"),
	)

	result_bull := BreakAway(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestBreakAway Bull unexpected result: got %v want %v", result_bull, expected_bull)
	}

}

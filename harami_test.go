package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestHarami(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{100, 110}, series.Float, "Open"),
		series.New([]float64{125, 115}, series.Float, "High"),
		series.New([]float64{120, 105}, series.Float, "Close"),
		series.New([]float64{95, 100}, series.Float, "Low"),
	)
	result_bear := Harami(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestHarami Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{120, 105}, series.Float, "Open"),
		series.New([]float64{125, 115}, series.Float, "High"),
		series.New([]float64{100, 110}, series.Float, "Close"),
		series.New([]float64{95, 100}, series.Float, "Low"),
	)
	result_bull := Harami(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestHarami Bull unexpected result: got %v want %v", result_bear, expected_bear)
	}

}

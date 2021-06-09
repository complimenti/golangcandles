package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestAbandonedBaby(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{90, 125, 110}, series.Float, "Open"),
		series.New([]float64{105, 135, 115}, series.Float, "High"),
		series.New([]float64{100, 130, 95}, series.Float, "Close"),
		series.New([]float64{85, 120, 92}, series.Float, "Low"),
	)
	result_bear := AbandonedBaby(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestAbandonedBaby Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{100, 50, 80}, series.Float, "Open"),
		series.New([]float64{105, 60, 100}, series.Float, "High"),
		series.New([]float64{90, 55, 95}, series.Float, "Close"),
		series.New([]float64{85, 45, 75}, series.Float, "Low"),
	)
	result_bull := AbandonedBaby(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestAbandonedBaby Bull unexpected result: got %v want %v", result_bear, expected_bear)
	}

}

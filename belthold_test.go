package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestBeltHold(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{60, 100}, series.Float, "Open"),
		series.New([]float64{75, 100}, series.Float, "High"),
		series.New([]float64{70, 80}, series.Float, "Close"),
		series.New([]float64{55, 75}, series.Float, "Low"),
	)
	result_bear := BeltHold(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestBeltHold Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{120, 70}, series.Float, "Open"),
		series.New([]float64{125, 95}, series.Float, "High"),
		series.New([]float64{105, 90}, series.Float, "Close"),
		series.New([]float64{100, 70}, series.Float, "Low"),
	)

	result_bull := BeltHold(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestBeltHold Bull unexpected result: got %v want %v", result_bull, expected_bull)
	}

}

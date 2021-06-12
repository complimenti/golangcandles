package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestKicking(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{100, 90}, series.Float, "Open"),
		series.New([]float64{120, 90}, series.Float, "High"),
		series.New([]float64{120, 70}, series.Float, "Close"),
		series.New([]float64{100, 70}, series.Float, "Low"),
	)
	result_bear := Kicking(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestKicking Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{90, 100}, series.Float, "Open"),
		series.New([]float64{90, 120}, series.Float, "High"),
		series.New([]float64{70, 120}, series.Float, "Close"),
		series.New([]float64{70, 100}, series.Float, "Low"),
	)
	result_bull := Kicking(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestKicking Bull unexpected result: got %v want %v", result_bear, expected_bear)
	}

}

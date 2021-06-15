package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestLongLeggedDoji(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{85, 95, 105, 170}, series.Float, "Open"),
		series.New([]float64{95, 105, 115, 210}, series.Float, "High"),
		series.New([]float64{90, 100, 110, 160}, series.Float, "Close"),
		series.New([]float64{80, 90, 100, 120}, series.Float, "Low"),
	)
	result_bear := LongLeggedDoji(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestLongLeggedDoji Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{90, 100, 110, 160}, series.Float, "Open"),
		series.New([]float64{95, 105, 115, 210}, series.Float, "High"),
		series.New([]float64{85, 95, 105, 170}, series.Float, "Close"),
		series.New([]float64{80, 90, 100, 120}, series.Float, "Low"),
	)
	result_bull := LongLeggedDoji(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestLongLeggedDoji Bull unexpected result: got %v want %v", result_bear, expected_bear)
	}
}

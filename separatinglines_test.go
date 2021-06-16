package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestSeparatingLines(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{200, 150, 152}, series.Float, "Open"),
		series.New([]float64{210, 190, 152}, series.Float, "High"),
		series.New([]float64{170, 180, 130}, series.Float, "Close"),
		series.New([]float64{160, 140, 120}, series.Float, "Low"),
	)
	result_bear := SeparatingLines(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestSeparatingLines Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{50, 100, 102}, series.Float, "Open"),
		series.New([]float64{80, 110, 130}, series.Float, "High"),
		series.New([]float64{70, 80, 120}, series.Float, "Close"),
		series.New([]float64{40, 70, 102}, series.Float, "Low"),
	)
	result_bull := SeparatingLines(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestSeparatingLines Bull unexpected result: got %v want %v", result_bear, expected_bear)
	}
}

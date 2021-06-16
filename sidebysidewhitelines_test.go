package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestSideBySideWhiteLines(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{130, 110, 50, 50}, series.Float, "Open"),
		series.New([]float64{140, 115, 75, 77}, series.Float, "High"),
		series.New([]float64{120, 90, 70, 68}, series.Float, "Close"),
		series.New([]float64{110, 85, 45, 42}, series.Float, "Low"),
	)
	result_bear := SideBySideWhiteLines(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestSideBySideWhiteLines Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{70, 100, 150, 150}, series.Float, "Open"),
		series.New([]float64{90, 130, 210, 200}, series.Float, "High"),
		series.New([]float64{80, 120, 185, 190}, series.Float, "Close"),
		series.New([]float64{60, 90, 140, 135}, series.Float, "Low"),
	)
	result_bull := SideBySideWhiteLines(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestSideBySideWhiteLines Bull unexpected result: got %v want %v", result_bear, expected_bear)
	}
}

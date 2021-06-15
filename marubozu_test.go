package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestMarubozu(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{200}, series.Float, "Open"),
		series.New([]float64{200}, series.Float, "High"),
		series.New([]float64{100}, series.Float, "Close"),
		series.New([]float64{100}, series.Float, "Low"),
	)
	result_bear := Marubozu(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestMarubozu Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{100}, series.Float, "Open"),
		series.New([]float64{200}, series.Float, "High"),
		series.New([]float64{200}, series.Float, "Close"),
		series.New([]float64{100}, series.Float, "Low"),
	)
	result_bull := Marubozu(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestMarubozu Bull unexpected result: got %v want %v", result_bear, expected_bear)
	}

}

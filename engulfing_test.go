package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestEngulfing(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{80, 100}, series.Float, "Open"),
		series.New([]float64{95, 105}, series.Float, "High"),
		series.New([]float64{90, 70}, series.Float, "Close"),
		series.New([]float64{75, 65}, series.Float, "Low"),
	)
	result_bear := Engulfing(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestEngulfing Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{90, 70}, series.Float, "Open"),
		series.New([]float64{95, 105}, series.Float, "High"),
		series.New([]float64{80, 100}, series.Float, "Close"),
		series.New([]float64{75, 65}, series.Float, "Low"),
	)

	result_bull := Engulfing(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestEngulfing Bull unexpected result: got %v want %v", result_bull, expected_bull)
	}

}

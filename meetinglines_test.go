package golangcandles

import (
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func TestMeetingLines(t *testing.T) {
	bear_df := dataframe.New(
		series.New([]float64{85, 95, 130}, series.Float, "Open"),
		series.New([]float64{100, 120, 140}, series.Float, "High"),
		series.New([]float64{95, 115, 110}, series.Float, "Close"),
		series.New([]float64{85, 90, 105}, series.Float, "Low"),
	)
	result_bear := MeetingLines(bear_df)
	expected_bear := -100
	if result_bear != expected_bear {
		t.Errorf("TestMeetingLines Bear unexpected result: got %v want %v", result_bear, expected_bear)
	}

	bull_df := dataframe.New(
		series.New([]float64{200, 180, 110}, series.Float, "Open"),
		series.New([]float64{210, 195, 160}, series.Float, "High"),
		series.New([]float64{190, 150, 155}, series.Float, "Close"),
		series.New([]float64{180, 140, 105}, series.Float, "Low"),
	)
	result_bull := MeetingLines(bull_df)
	expected_bull := 100
	if result_bull != expected_bull {
		t.Errorf("TestMeetingLines Bull unexpected result: got %v want %v", result_bear, expected_bear)
	}
}

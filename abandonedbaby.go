package golangcandles

import (
	"math"

	"github.com/go-gota/gota/dataframe"
)

func AbandonedBaby(df dataframe.DataFrame) (output int) {

	ohlc_3 := GenerateOhcl(df, 3)
	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_1 := GenerateOhcl(df, 1)

	// BEAR
	if ohlc_3.Open < ohlc_3.Close {
		if ohlc_3.High < ohlc_2.Low {
			if math.Abs((ohlc_2.Close-ohlc_2.Open)/ohlc_2.Open) < 0.1 {
				if ohlc_1.Open < ohlc_2.Low && ohlc_1.Close < ohlc_1.Open {
					output = -100
				}
			}
		}
	}

	// BULL
	if ohlc_3.Open > ohlc_3.Close {
		if ohlc_3.Low > ohlc_2.High {
			if math.Abs((ohlc_2.Close-ohlc_2.Open)/ohlc_2.Open) <= 0.1 {
				if ohlc_1.Open > ohlc_2.High && ohlc_1.Close > ohlc_1.Open {
					output = 100
				}
			}
		}
	}

	return
}

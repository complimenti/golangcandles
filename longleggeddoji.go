package golangcandles

import (
	"math"

	"github.com/go-gota/gota/dataframe"
)

func LongLeggedDoji(df dataframe.DataFrame) (output int) {

	ohlc_4 := GenerateOhcl(df, 4)
	ohlc_3 := GenerateOhcl(df, 3)
	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_1 := GenerateOhcl(df, 1)

	// BEAR
	if ohlc_4.Close > ohlc_4.Open {
		if ohlc_3.Close > ohlc_3.Open {
			if ohlc_2.Close > ohlc_2.Open {
				if math.Abs((ohlc_1.Close-ohlc_1.Open)/ohlc_1.Open) < 0.1 {
					if math.Abs((ohlc_1.High-ohlc_1.Open)/ohlc_1.Open) > 0.2 {
						if math.Abs((ohlc_1.Close-ohlc_1.Low)/ohlc_1.Low) > 0.2 {
							output = -100
						}
					}
				}
			}
		}
	}

	// BULL
	if ohlc_4.Close < ohlc_4.Open {
		if ohlc_3.Close < ohlc_3.Open {
			if ohlc_2.Close < ohlc_2.Open {
				if math.Abs((ohlc_1.Open-ohlc_1.Close)/ohlc_1.Close) < 0.1 {
					if math.Abs((ohlc_1.Low-ohlc_1.Close)/ohlc_1.Close) > 0.2 {
						if math.Abs((ohlc_1.Open-ohlc_1.High)/ohlc_1.High) > 0.2 {
							output = 100
						}
					}
				}
			}
		}
	}

	return
}

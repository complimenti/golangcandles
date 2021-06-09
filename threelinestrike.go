package golangcandles

import (
	"github.com/go-gota/gota/dataframe"
)

func ThreeLineStrike(df dataframe.DataFrame) (output int) {

	ohlc_4 := GenerateOhcl(df, 4)
	ohlc_3 := GenerateOhcl(df, 3)
	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_1 := GenerateOhcl(df, 1)

	// BEAR
	if ohlc_3.Close < ohlc_4.Close {
		if ohlc_2.Close < ohlc_3.Close {
			if (ohlc_4.Close < ohlc_3.Open) && (ohlc_3.Open < ohlc_4.Open) {
				if (ohlc_3.Close < ohlc_2.Open) && (ohlc_2.Open < ohlc_3.Open) {
					if ohlc_1.Open < ohlc_2.Close {
						if ohlc_1.Close > ohlc_4.Open {
							output = -100
						}
					}
				}
			}

		}
	}

	// BULL
	if ohlc_3.Close > ohlc_4.Close {
		if ohlc_2.Close > ohlc_3.Close {
			if (ohlc_4.Close > ohlc_3.Open) && (ohlc_3.Open > ohlc_4.Open) {
				if (ohlc_3.Close > ohlc_2.Open) && (ohlc_2.Open > ohlc_3.Open) {
					if ohlc_1.Open > ohlc_2.Close {
						if ohlc_1.Close < ohlc_4.Open {
							output = 100
						}
					}
				}
			}

		}
	}

	return
}

package golangcandles

import (
	"math"

	"github.com/go-gota/gota/dataframe"
)

func SideBySideWhiteLines(df dataframe.DataFrame) (output int) {

	ohlc_1 := GenerateOhcl(df, 1)
	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_3 := GenerateOhcl(df, 3)
	ohlc_4 := GenerateOhcl(df, 4)

	// BEAR
	if ohlc_4.Open > ohlc_4.Close {
		if ohlc_3.Open > ohlc_3.Close {
			if ohlc_3.Low > ohlc_2.High {
				if ohlc_2.Open < ohlc_2.Close {
					if ohlc_1.Open < ohlc_1.Close {
						if math.Abs((ohlc_2.Open-ohlc_1.Open)/ohlc_1.Open) < 0.1 {
							if math.Abs((ohlc_2.Close-ohlc_1.Close)/ohlc_1.Close) < 0.1 {
								if ohlc_2.Open == ohlc_1.Open {
									output = -100
								}
							}
						}
					}
				}
			}
		}
	}

	// BULL
	if ohlc_4.Open < ohlc_4.Close {
		if ohlc_3.Open < ohlc_3.Close {
			if ohlc_3.Low < ohlc_2.High {
				if ohlc_2.Open < ohlc_2.Close {
					if ohlc_1.Open < ohlc_1.Close {
						if math.Abs((ohlc_2.Open-ohlc_1.Open)/ohlc_1.Open) < 0.1 {
							if math.Abs((ohlc_2.Close-ohlc_1.Close)/ohlc_1.Close) < 0.1 {
								if ohlc_2.Open == ohlc_1.Open {
									output = 100
								}
							}
						}
					}
				}
			}
		}
	}

	return
}

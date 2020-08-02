package utils

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

const delta = 0.000001

func Compare(x, y float64) bool {
	return x < y+delta && x > y-delta
}

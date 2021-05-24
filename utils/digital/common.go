package digital

import "math"

// 截取小数点后n位，非四舍五入
func TruncatFloat(num float64, n int) float64 {
	factor := math.Pow10(n)
	res := math.Trunc(num*factor) / factor
	return res
}

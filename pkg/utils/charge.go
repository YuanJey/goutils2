package utils

func CalculateFee(userCount int) float64 {
	var fee float64

	if userCount > 1000 {
		fee = float64(userCount) * 0.1 * 30
	} else {
		fee = 3000
	}
	return fee
}

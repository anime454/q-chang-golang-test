package currency

func BahtToSatang(baht float64) int {
	return int(baht * 100)
}

func SatangToBaht(satang int) float64 {
	return float64(satang) / 100
}

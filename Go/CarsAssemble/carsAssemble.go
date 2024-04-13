package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := float64(productionRate) * (successRate / 100)
	return int(carsPerHour / 60)
}

func CalculateCost(carsCount int) uint {
	groupCars := carsCount / 10
	individualCars := carsCount % 10
	return uint(groupCars*95000 + individualCars*10000)
}

func max(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	// enough fuel
	if startFuel >= target {
		return 0
	}

	distancesWithMinRefuels := make([]int, len(stations)+1)
	distancesWithMinRefuels[0] = startFuel
	for i := 0; i < len(stations); i++ {
		for time := i; time >= 0; time-- {
			if distancesWithMinRefuels[time] >= stations[i][0] {
				distancesWithMinRefuels[time+1] = max(distancesWithMinRefuels[time+1], distancesWithMinRefuels[time]+stations[i][1])
			}
		}
	}

	for i := 0; i < len(distancesWithMinRefuels); i++ {
		if distancesWithMinRefuels[i] >= target {
			return i
		}
	}
	return -1
}

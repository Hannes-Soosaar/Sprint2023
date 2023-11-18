package sprint

func Payout(amount int, denominations []int) (payout []int) {

	if amount < 0 {
		return []int{}
	}
	// bubble sort the denomination to be in decending order
	swappedValue := true
	for swappedValue {
		swappedValue = false
		for i := 0; i < len(denominations)-1; i++ {
			for j := 0; j < len(denominations)-i-1; j++ {
				if denominations[j] < denominations[j+1] {
					denominations[j], denominations[j+1] = denominations[j+1], denominations[j]
				}
			}
		}
		payout = []int{}
		for _, denomination := range denominations {
			for amount >= denomination {
				payout = append(payout, denomination)
				amount -= denomination
			}
		}
	}
	if amount == 0 {
		return payout
	}
	return []int{}
}

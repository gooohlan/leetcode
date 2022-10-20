package Array

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for _, booking := range bookings {
		diff[booking[0]-1] += booking[2]
		if booking[1] < n {
			diff[booking[1]-1] -= booking[2]
		}
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff
}

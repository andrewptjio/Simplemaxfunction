package max

func Max(number []int) int {
	largestNumber := 0

	for a := 0; a < len(number); a++ {
		if number[a] == number[len(number)-1] {
			var temp int
			indexawal := a
			indexakhir := len(number) - 1

			for indexawal < indexakhir {

				if number[indexawal] == number[indexakhir] {
					temp = number[indexawal]
					indexawal++
					indexakhir--
				} else {
					break
				}

				if temp > largestNumber {
					largestNumber = temp
				}
			}
		}
	}
	return largestNumber
}

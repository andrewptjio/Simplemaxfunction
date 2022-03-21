package max

func Max(Number []int) int {
	largestNumber := 0

	for a := 0; a < len(Number); a++ {
		if Number[a] == Number[len(Number)-1] {
			var temp int
			indexAwal := a
			indexAkhir := len(Number) - 1

			for indexAwal < indexAkhir {

				if Number[indexAwal] == Number[indexAkhir] {
					temp = Number[indexAwal]
					indexAwal++
					indexAkhir--
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

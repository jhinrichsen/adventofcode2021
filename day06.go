package adventofcode2021

/* naive approach blows mem for part 2 around 240 days
 We need to group fishes by age.
	for day := 0; day < days; day++ {
		var babies int
		for i := range fishes {
			fishes[i]--
			if fishes[i] < 0 {
				fishes[i] = 6
				babies++
			}
		}
		for i := 0; i < babies; i++ {
			fishes = append(fishes, 8)
		}
	}
	return len(fishes), nil
*/

func Day06(fishes []uint, days uint) (uint, error) {
	const groups = 8
	var ages [groups + 1]uint // 0..7 plus babies in 8
	for i := 0; i < len(fishes); i++ {
		ages[fishes[i]%groups]++
	}

	for day := uint(0); day < days; day++ {
		babies := ages[0]
		for age := 0; age < groups; age++ {
			ages[age] = ages[age+1]
		}
		ages[6] += babies
		ages[8] = babies
	}

	var sum uint
	for i := range ages {
		sum += ages[i]
	}
	return sum, nil
}

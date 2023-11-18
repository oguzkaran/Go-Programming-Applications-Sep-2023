package numeric

func IsPrime(val int) bool {
	if val <= 1 {
		return false
	}

	if val%2 == 0 {
		return val == 2
	}

	if val%3 == 0 {
		return val == 3
	}

	if val%5 == 0 {
		return val == 5
	}

	if val%7 == 0 {
		return val == 7
	}

	for i := 11; i*i <= val; i += 2 {
		if val%i == 0 {
			return false
		}
	}

	return true
}

func Prime(n int) int {
	count := 0
	val := 2

	for {
		if IsPrime(val) {
			count++
		}

		if count == n {
			return val
		}

		val++
	}
}

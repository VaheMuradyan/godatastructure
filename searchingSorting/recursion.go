package searchingSorting

func CountDigits(num int) int {
	if num == 0 {
		return 0
	} else {
		// Recursive case: reduce the number by dividing by 10 and count the rest
		// TODO: Count only the even numbers
		if (num%10)%2 != 0 {
			return CountDigits(num / 10)
		} else {
			return num%10 + CountDigits(num/10)
		}
	}
}

package util

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func stringToMasInt(str string) []int {
	var res []int
	str = reverse(str)
	for i := 0; i < len(str); i++ {
		res = append(res, int(str[i])-int('0'))
	}
	return res
}

func LuhnAlgorithm(card string) bool {
	cardNumber := stringToMasInt(card)
	for pos := range cardNumber {
		if (pos+1)%2 == 0 {
			cardNumber[pos] *= 2
			if cardNumber[pos] > 9 {
				tmp := cardNumber[pos]
				cardNumber[pos] = 0
				for tmp != 0 {
					cardNumber[pos] += tmp % 10
					tmp /= 10
				}
			}
		}
	}
	control_sum := 0
	for _, val := range cardNumber {
		control_sum += val
	}
	if control_sum%10 == 0 {
		return true
	}
	return false
}

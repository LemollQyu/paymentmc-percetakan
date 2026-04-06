package utils

import "strconv"

func FormatRupiah(amount float64) string {
	str := strconv.FormatInt(int64(amount), 10)

	var result []byte
	for i, c := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result = append(result, '.')
		}
		result = append(result, byte(c))
	}

	return "Rp " + string(result)
}

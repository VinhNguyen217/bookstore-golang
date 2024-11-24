package utils

import "strconv"

func ConvertToVND(price int) string {
	strValue := strconv.Itoa(price)

	// Initialize an empty string for the result
	var result string
	n := len(strValue)

	// Loop through the string and add dots (.) as thousands separators
	for i, r := range strValue {
		// Add a dot every 3 digits from the right (except at the beginning)
		if (n-i)%3 == 0 && i != 0 {
			result += "."
		}
		result += string(r)
	}

	// Append "₫" symbol for Vietnamese đồng
	result += " ₫"

	return result
}

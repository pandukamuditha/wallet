package common

import "strconv"

// Convert string to 64 bit integer of base 10
func StringToInt64(valstr string) (int64, error) {
	return strconv.ParseInt(valstr, 10, 64)
}

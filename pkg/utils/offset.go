package utils

func ValidateOffset(offset int64) int64 {
	if offset < 0 {
		return 0
	}

	return offset
}

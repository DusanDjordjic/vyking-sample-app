package utils

import "app/pkg/config"

func ValidateLimit(limit int64) int64 {
	if limit < 0 {
		return 0
	}

	if limit > config.MAX_LIMIT {
		return config.MAX_LIMIT
	}

	return limit
}

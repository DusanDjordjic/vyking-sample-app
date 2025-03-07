package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func GetStringPathParam(url *url.URL, n uint8) (string, error) {
	if n == 0 {
		panic("path parameters start from 1 but 0 is passed as n")
	}

	trimmed := strings.Trim(url.Path, "/")
	splits := strings.Split(trimmed, "/")
	if len(splits) < int(n) {
		return "", fmt.Errorf("%d. path parameter doesn't exist", n)
	}

	return splits[n-1], nil
}

func GetInt64PathParameter(url *url.URL, n uint8) (int64, error) {
	sValue, err := GetStringPathParam(url, n)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(sValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %d. path param as int64, %s", n, err)
	}

	return v, nil
}

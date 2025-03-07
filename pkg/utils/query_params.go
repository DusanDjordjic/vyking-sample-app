package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetStringQueryParam(r *http.Request, name string) (string, error) {
	v := r.URL.Query().Get(name)
	if len(v) == 0 {
		return "", fmt.Errorf("%s query param doesn't exist", name)
	}

	return v, nil
}

func GetStringQueryParamWithDefault(r *http.Request, name string, def string) string {
	v := r.URL.Query().Get(name)
	if len(v) == 0 {
		return def
	}

	return v
}

func GetInt64QueryParam(r *http.Request, name string) (int64, error) {
	sValue, err := GetStringQueryParam(r, name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(sValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s query param as int64, %s", name, err)
	}

	return v, nil
}

func GetInt64QueryParamWithDefault(r *http.Request, name string, def int64) int64 {
	sValue, err := GetStringQueryParam(r, name)
	if err != nil {
		return def
	}

	v, err := strconv.ParseInt(sValue, 10, 64)
	if err != nil {
		return def
	}

	return v
}

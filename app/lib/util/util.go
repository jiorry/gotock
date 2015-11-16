package util

import (
	"strconv"
	"strings"
)

func ParsePercent(str string) float64 {
	str = strings.TrimSpace(str)
	if str == "--" {
		return float64(0)
	}

	n := len(str)
	val := str
	isPer := false
	if str[n-1:] == "%" {
		val = str[:n-1]
		isPer = true
	}

	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0.0
	}

	if isPer {
		return f / 100
	} else {
		return f
	}
}

func ParseMoney(str string) float64 {
	str = strings.Replace(str, ",", "", -1)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return f
}

func ParseMoneyCN(str string) float64 {
	str = strings.TrimSpace(str)
	if str == "--" {
		return float64(0)
	}

	n := len(str)

	if n < 4 {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0.0
		}
		return f
	}

	ext := str[n-3:]
	val := str[:n-3]
	if n > 6 && str[n-6:] == "万亿" {
		ext = str[n-6:]
		val = str[:n-6]
	} else {

	}

	switch ext {
	case "万亿":
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0.0
		}
		return f * 1000000000000
	case "亿":
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0.0
		}
		return f * 100000000
	case "万":
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0.0
		}
		return f * 10000
	default:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0.0
		}
		return f
	}

}

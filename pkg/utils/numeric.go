// Package utils provides general utility functionality.
package utils

func IsUnit(num, threshold float64) bool {
	return num >= -threshold && num <= 1+threshold
}

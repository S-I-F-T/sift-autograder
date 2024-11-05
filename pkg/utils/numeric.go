// Package utils provides general utility functionality.
package utils

/* Checks whether nums is approximately 
within the unit interval [0,1] but with
a tolerance threshold. */
func IsUnit(num, threshold float64) bool {
	return num >= -threshold && num <= 1+threshold
}

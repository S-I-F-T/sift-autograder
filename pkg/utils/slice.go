package utils

import "slices"

/* Verifies that slice contains all the elements in elems 
by looping through each elements in elems and checking if
it is contained in slice. */
func ContainsAll[S interface{ ~[]E }, E interface{comparable}](slice S, elems []E) bool {
    for _, elem := range elems {
        if !slices.Contains(slice, elem) {
            return false
        }
    }
    return true
}


package utils

import "slices"

func ContainsAll[S interface{ ~[]E }, E interface{comparable}](slice S, elems []E) bool {
    for _, elem := range elems {
        if !slices.Contains(slice, elem) {
            return false
        }
    }
    return true
}


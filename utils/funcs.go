// Package utils defines custom and helper functions to enhance the inventory CLI output
package utils

import "strings"

func Pluralize(count int, singular string) string {
	if count == 1 {
		return singular
	}
	return singular + "s"
}

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

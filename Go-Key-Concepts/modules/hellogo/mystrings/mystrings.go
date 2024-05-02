package mystrings

// Reverse reverses a string left to right
// Notice that we need to capitalize the
// first letter of the function
// If we don't then we won't be able access
// this function outside of the
// mystrings package
func Reverse(s string) string {
	result := ""
	for _, val := range s {
		result = string(val) + result
	}
	return result
}

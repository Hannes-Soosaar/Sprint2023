package sprint

func StrLen(s string) []int {

return []int{len([]rune(s)),len(s)} // must convert the string to rune array to get the number of characters len counts the number of chars
									// len(s)	returns the numebr of bytes in a String.
}
package sprint

func NRune(s string, i int) rune {
return rune(s[i]) // the +1 was not needed as i will call the right position
}
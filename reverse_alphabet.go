package gocrypt

func Reverse_Alphabet(mess string) string {
	var key = map[string]string{
		"z": "a", "J": "Q", "e": "v", "h": "s", "q": "j", "r": "i", "a": "z", "j": "q",
		"M": "N", "R": "I", "E": "V", "U": "F", "Z": "A", "k": "p", "v": "e", "w": "d",
		"y": "b", "N": "M", "X": "C", "i": "r", "n": "m", "x": "c", "C": "X", "m": "n",
		"F": "U", "H": "S", "K": "P", "b": "y", "f": "u", "g": "t", "l": "o", "W": "D",
		"T": "G", "Y": "B", "o": "l", "u": "f", "A": "Z", "O": "L", "p": "k", "s": "h",
		"B": "Y", "L": "O", "G": "T", "I": "R", "P": "K", "Q": "J", "c": "x", "d": "w",
		"t": "g", "D": "W", "S": "H", "V": "E",
	}
	out := ""
	for i := 0; i < len(mess); i++ {
		m, isset := key[string(mess[i])]
		if isset {
			out += m
		} else {
			out += string(mess[i])
		}
	}
	return out
}

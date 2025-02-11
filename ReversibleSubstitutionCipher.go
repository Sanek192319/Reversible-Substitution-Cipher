package main

import (
	"fmt"
	"strings"
)

// Простая таблица подстановки
var substitutionTable = map[rune]rune{
	'A': 'M', 'B': 'N', 'C': 'O', 'D': 'P', 'E': 'Q', 'F': 'R', 'G': 'S', 'H': 'T', 'I': 'U', 'J': 'V',
	'K': 'W', 'L': 'X', 'M': 'Y', 'N': 'Z', 'O': 'A', 'P': 'B', 'Q': 'C', 'R': 'D', 'S': 'E', 'T': 'F',
	'U': 'G', 'V': 'H', 'W': 'I', 'X': 'J', 'Y': 'K', 'Z': 'L',

	'a': 'm', 'b': 'n', 'c': 'o', 'd': 'p', 'e': 'q', 'f': 'r', 'g': 's', 'h': 't', 'i': 'u', 'j': 'v',
	'k': 'w', 'l': 'x', 'm': 'y', 'n': 'z', 'o': 'a', 'p': 'b', 'q': 'c', 'r': 'd', 's': 'e', 't': 'f',
	'u': 'g', 'v': 'h', 'w': 'i', 'x': 'j', 'y': 'k', 'z': 'l',
}

// Инвертирование таблицы для дешифрования
func invertMap(m map[rune]rune) map[rune]rune {
	inverted := make(map[rune]rune)
	for key, value := range m {
		inverted[value] = key
	}
	return inverted
}

// Функция шифрования
func encrypt(text string) string {
	var encrypted strings.Builder
	for _, char := range text {
		if sub, exists := substitutionTable[char]; exists {
			encrypted.WriteRune(sub)
		} else {
			encrypted.WriteRune(char)
		}
	}
	return encrypted.String()
}

// Функция дешифрования
func decrypt(text string) string {
	decryptionTable := invertMap(substitutionTable)
	var decrypted strings.Builder
	for _, char := range text {
		if sub, exists := decryptionTable[char]; exists {
			decrypted.WriteRune(sub)
		} else {
			decrypted.WriteRune(char)
		}
	}
	return decrypted.String()
}

func main() {
	var input string
	fmt.Print("Enter text: ")
	fmt.Scanln(&input)

	encrypted := encrypt(input)
	fmt.Println("Encrypted:", encrypted)

	decrypted := decrypt(encrypted)
	fmt.Println("Decrypted:", decrypted)
}

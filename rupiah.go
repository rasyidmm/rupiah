package rupiah

import (
	"strconv"
	"strings"
	"unicode"
)

var (
	Bilangan = []string{"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan", "Sepuluh", "Sebelas"}
	Besar    = []string{"", "Ribu", "Juta", "Milyar", "Triliun", "Quadriliun"}
)

func Rupiah(rupiah int64) string {
	var (
		result string
	)

	decimal := strconv.FormatInt(rupiah, 10)
	temp := len(CheckDecimal(decimal)) - 1
	for _, x := range CheckDecimal(decimal) {
		hasil := Penyebut(x)
		result += hasil
		if hasil != "" {
			result += Besar[temp]
		}
		temp--
	}

	return GiveSpace(result + "Rupiah")
}
func Penyebut(text string) string {
	var result string
	temp := false
	tert := true
	for i := 0; i < len(text); i++ {
		if string(text[i]) != "0" || temp {
			dec, _ := strconv.ParseInt(string(text[i]), 10, 64)
			if i == 0 {
				if text[i] == '1' {
					result += "Seratus"
				} else {
					result += Bilangan[dec] + "Ratus"
				}

			} else if i == 1 {
				if text[i] == '1' && text[i+1] == '0' {
					result += "se" + "Puluh"
					tert = false
				} else if text[i] == '1' && text[i+1] != '0' {
					dec, _ := strconv.ParseInt(string(text[i+1]), 10, 64)
					tert = false
					result += Bilangan[dec] + "Belas"
				} else {
					result += Bilangan[dec] + "Puluh"
				}
			} else if tert && i == 2 {
				result += Bilangan[dec]
			}

		}
	}
	return result
}

func CheckDecimal(decimal string) []string {
	for {
		if len(decimal)%3 == 0 {
			break
		}
		decimal = "0" + decimal
	}

	return SplitCharacter(decimal)
}

func SplitCharacter(decimal string) []string {
	var (
		result []string
		temp   string
	)

	for i := 0; i < len(decimal); i++ {
		temp = temp + string(decimal[i])
		if len(temp) == 3 {
			result = append(result, temp)
			temp = ""
		}

	}

	return result
}

func GiveSpace(text string) string {
	newName := ""
	for _, c := range text {

		if unicode.IsUpper(c) {
			newName += " "
		}
		newName += string(c)
	}
	newName = strings.TrimSpace(newName)

	return newName
}

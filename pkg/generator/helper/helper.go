package helper

import (
	"math/rand"
	"slices"
	"strings"
	"unicode"
)

type IntDictionary []int
type StringDictionary []string
type BoolDictionary []bool

func (d IntDictionary) Contain(v int) bool {
	return slices.Contains(d, v)
}

func (d IntDictionary) Add(v int) IntDictionary {
	return append(d, v)
}

func (d StringDictionary) Contain(v string) bool {
	return slices.Contains(d, v)
}

func (d StringDictionary) Add(v string) StringDictionary {
	return append(d, v)
}

func (d StringDictionary) Insert(v string, i int) StringDictionary {
	return slices.Insert(d, i, v)
}

func Swap(d, beginString, targetString string, begin, target int) string {
	arrayTemp := []string{}
	result := ""
	for _, i := range d {
		arrayTemp = append(arrayTemp, string(i))
	}
	arrayTemp[begin] = targetString
	arrayTemp[target] = beginString
	for _, i := range arrayTemp {
		result += i
	}
	return result
}

func (d BoolDictionary) Contain(v bool) bool {
	return slices.Contains(d, v)
}

func (d BoolDictionary) Add(v bool) BoolDictionary {
	return append(d, v)
}

var SpacialChars = []string{"@", "!", "$", "%", "^", ">", "<", "*", "(", ")", "[", "]", "{", "}", "+", "-", "/", "|", "~", "?"}

func RandomSpacialChar(d string) string {
	arrayTemp := []string{}
	result := ""
	limitSpacial := rand.Intn(4) + 1
	totalSpacial := 0
	usedSpacial := StringDictionary{}
	usedSpacialChannel := make(chan StringDictionary)
	for _, i := range d {
		arrayTemp = append(arrayTemp, string(i))
	}
	for i := range len(arrayTemp) {
		random := rand.Intn(100)
		if i%2 == 0 && random > 40 && totalSpacial <= limitSpacial {
			randomSpacial := rand.Intn(len(SpacialChars))
			for usedSpacial.Contain(SpacialChars[randomSpacial]) {
				randomSpacial = rand.Intn(len(SpacialChars))
			}
			selectSpacial := SpacialChars[randomSpacial]
			arrayTemp[i] = selectSpacial
			go func() {
				usedSpacialChannel <- usedSpacial.Add(selectSpacial)
			}()
			add := <-usedSpacialChannel
			usedSpacial = add
			totalSpacial++
		}
	}

	for _, i := range arrayTemp {
		result += i
	}

	return result
}

func RandomUpper(str string) string {
	result := ""
	trackUpper := 0
	totalChar := 0
	for i := range len(str) {
		if !unicode.IsDigit(rune(str[i])) {
			totalChar++
		}
	}

	for i := range len(str) {
		if !unicode.IsDigit(rune(str[i])) {
			random := rand.Intn(100)
			if random > 50 && trackUpper < totalChar/2 {
				result += string(strings.ToUpper(string(str[i])))
				trackUpper++
				continue
			}

			result += string(strings.ToLower(string(str[i])))
			continue
		}

		result += string(str[i])
	}
	return result
}

func CountInt(str string) int {
	currentNumberLength := 0
	for _, i := range str {
		if unicode.IsDigit(i) {
			currentNumberLength++
		}
	}
	return currentNumberLength
}

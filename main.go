package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := os.Args[1]
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	words := strings.Split(string(file), " ")
	for i := 1; i < len(words); i++ {
		switch words[i] {
		case "(cap)":
			words[i-1] = cap(words[i-1])
			words[i] = ""
		case "(up)":
			fmt.Printf("'%s'\n", words[i-1])
			words[i-1] = up(words[i-1])
			words[i] = ""
		case "(low)":
			words[i-1] = low(words[i-1])
			words[i] = ""
		case "(bin)":
			words[i-1] = bin(words[i-1])
			words[i] = ""
		case "(hex)":
			words[i-1] = hex(words[i-1])
			words[i] = ""
		}
	}
	words = delete_empty(words)
	regex := regexp.MustCompile(`\s+`)
	result := regex.ReplaceAllString(strings.Join(words, " "), " $1")
	result = startsWithVowel(result)

	fileOutput, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer fileOutput.Close()

	w := bufio.NewWriter(fileOutput)
	_, er := w.WriteString(result)
	if er != nil {
		os.Exit(1)
	}

	w.Flush()
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func hex(word string) string {
	value, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		fmt.Println("Invalid Hex")
		return ""
	}
	return strconv.Itoa(int(value))
}

func bin(digit string) string {
	binary, err := strconv.ParseInt(digit, 2, 64)
	if err != nil {
		fmt.Println("Invalid binary")
		return ""
	}
	return strconv.FormatInt(binary, 10)
}

func up(w string) string {
	upp := strings.ToUpper(w)
	return upp
}

func low(s string) string {
	loww := strings.ToLower(s)
	return loww
}

func cap(maj string) string {
	loww := strings.ToLower(maj)
	w := strings.Title(loww)
	return w
}

//func num(a string) int {
//  app := func(a rune) bool {
//  return !unicode.strings(app)
//}
//suii :=
//}

func punct(text string) string {
	regex := regexp.MustCompile(`\s*([.,!?;:]+)\s*`)
	transformed := regex.ReplaceAllString(text, "$1")
	return transformed
}

func groupofpunct(text string) string {
	regex := regexp.MustCompile(`\s*([.,]{2,}|[!?]+)(?:\s+|$)`)
	transformed := regex.ReplaceAllString(text, "$1 ")
	return transformed
}

func punctofquotes(text string) string {
	regex := regexp.MustCompile(`'\s*(\w+)\s*'`)
	transformed := regex.ReplaceAllStringFunc(text, func(match string) string {
		word := strings.Trim(match, "' ")
		return "'" + word + "'"
	})
	return transformed
}

func morethanoneword(text string) string {
	regex := regexp.MustCompile(`'\s*(.*?)\s*'`)
	transformed := regex.ReplaceAllStringFunc(text, func(match string) string {
		word := strings.Trim(match, "' ")
		return "'" + word + "'"
	})
	return transformed
}

func startsWithVowel(t string) string {
	regex := regexp.MustCompile(`\b([aA])\s+([aeiouh]\w*)`)
	return regex.ReplaceAllString(t, "${1}n $2")
}

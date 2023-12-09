package main
	
	import (
	"fmt"
	
	"strconv"
	
	"regexp"
	
	"strings"
	
	"io/ioutil"
	
	"os"
	)
	
	func main() {
	if len(os.Args) < 2 {
	fmt.Println("File name missing")
	return
	} else if len(os.Args) > 3 {
	fmt.Println("Too many arguments")
	return
	} else if len(os.Args) != 3 {
	return
	}
	fileName := os.Args[1]
	file, err := os.ReadFile(fileName)
	if err != nil {
	fmt.Println(err.Error())
	}
	sample := string(file)
	if err := ioutil.WriteFile(fileName, []byte(sample), 0644); err != nil {
	fmt.Println(err.Error())
	return
	}
	err = nonenglishchar(sample)
	if err != nil {
	fmt.Println(err)
	return
	}
	
	regex := regexp.MustCompile(`\s+`)
	err = brackets(sample)
	if err != nil {
	fmt.Println(err)
	return
	}
	words := strings.Split(string(file), " ")
	if strings.Contains(words[0], "(cap") ||
	strings.Contains(words[0], "(low") ||
	strings.Contains(words[0], "(hex") ||
	strings.Contains(words[0], "(bin") ||
	strings.Contains(words[0], "(up") ||
	strings.Contains(words[0], "(low") {
	fmt.Println("Error: First word is a function")
	return
	}
	
	for i := 1; i < len(words); i++ {
	if strings.Contains(words[i], "(cap)") {
	if i+1 < len(words) && (strings.Contains(words[i+1], "(cap") || strings.Contains(words[i+1], "(low") || strings.Contains(words[i+1], "(hex") || strings.Contains(words[i+1], "(bin") || strings.Contains(words[i+1], "(up") || strings.Contains(words[i+1], "(low")) {
	fmt.Println("Error: there is more than one function")
	return
	}
	words[i-1] = cap(words[i-1])
	words[i] = ""
	}
	
	if strings.Contains(words[i], "(up)") {
	if i+1 < len(words) && (strings.Contains(words[i+1], "(cap") || strings.Contains(words[i+1], "(low") || strings.Contains(words[i+1], "(hex") || strings.Contains(words[i+1], "(bin") || strings.Contains(words[i+1], "(up") || strings.Contains(words[i+1], "(low")) {
	fmt.Println("Error: there is more than one function")
	return
	}
	words[i-1] = up(words[i-1])
	words[i] = ""
	}
	
	if strings.Contains(words[i], "(low)") {
	if i+1 < len(words) && (strings.Contains(words[i+1], "(cap") || strings.Contains(words[i+1], "(low") || strings.Contains(words[i+1], "(hex") || strings.Contains(words[i+1], "(bin") || strings.Contains(words[i+1], "(up") || strings.Contains(words[i+1], "(low")) {
	fmt.Println("Error: there is more than one function")
	return
	}
	words[i-1] = low(words[i-1])
	words[i] = ""
	}
	
	if strings.Contains(words[i], "(bin)") {
	if i+1 < len(words) && (strings.Contains(words[i+1], "(cap") || strings.Contains(words[i+1], "(low") || strings.Contains(words[i+1], "(hex") || strings.Contains(words[i+1], "(bin") || strings.Contains(words[i+1], "(up") || strings.Contains(words[i+1], "(low")) {
	fmt.Println("Error: there is more than one function")
	return
	}
	words[i-1] = bin(words[i-1])
	words[i] = ""
	}
	
	if strings.Contains(words[i], "(hex)") {
	if i+1 < len(words) && (strings.Contains(words[i+1], "(cap") || strings.Contains(words[i+1], "(low") || strings.Contains(words[i+1], "(hex") || strings.Contains(words[i+1], "(bin") || strings.Contains(words[i+1], "(up") || strings.Contains(words[i+1], "(low")) {
	fmt.Println("Error: there is more than one function")
	return
	}
	words[i-1] = hex(words[i-1])
	words[i] = ""
	}
	
	if strings.HasPrefix(words[i], "(low,") {
	if i+2 < len(words) && (strings.Contains(words[i+2], "(cap") || strings.Contains(words[i+2], "(low") || strings.Contains(words[i+2], "(hex") || strings.Contains(words[i+2], "(bin") || strings.Contains(words[i+2], "(up") || strings.Contains(words[i+2], "(low")) {
	fmt.Println("Error: there is more than one function")
	return
	}
	if i+1 >= len(words) {
	fmt.Println("Error: No space in between")
	return
	}
	num, err := strconv.Atoi(strings.TrimSuffix(words[i+1][0:], ")"))
	if err != nil {
	fmt.Println("Error: Invalid Number")
	return
	}
	if err == nil && i-num >= 0 {
	startIndex := i - num
	for j := startIndex; j < i; j++ {
	words[j] = low(words[j])
	}
	words[i] = ""
	words[i+1] = ""
	} else if i-num < 0 {
	fmt.Println("Error: No enough words")
	return
	}
	} else if strings.HasPrefix(words[i], "(up,") {
	if i+2 < len(words) && (strings.Contains(words[i+2], "(cap") || strings.Contains(words[i+2], "(low") || strings.Contains(words[i+2], "(hex") || strings.Contains(words[i+2], "(bin") || strings.Contains(words[i+2], "(up") || strings.Contains(words[i+2], "(low")) {
	fmt.Println("Error: there is more than one function")
	return
	}
	if i+1 >= len(words) {
	fmt.Println("Error: No space in between")
	return
	}
	num, err := strconv.Atoi(strings.TrimSuffix(words[i+1][0:], ")"))
	if err != nil {
	fmt.Println("Error: Invalid Number")
	return
	}
	if err == nil && i-num >= 0 {
	startIndex := i - num
	for j := startIndex; j < i; j++ {
	words[j] = up(words[j])
	}
	words[i] = ""
	words[i+1] = ""
	} else if i-num < 0 {
	fmt.Println("Error: No enough words")
	return
	}
	} else if strings.HasPrefix(words[i], "(cap,") {
	if i+2 < len(words) && (strings.Contains(words[i+2], "(cap") || strings.Contains(words[i+2], "(low") || strings.Contains(words[i+2], "(hex") || strings.Contains(words[i+2], "(bin") || strings.Contains(words[i+2], "(up") || strings.Contains(words[i+2], "(low")) {
	fmt.Println("Error: there is more than one function")
	return
	}
	if i+1 >= len(words) {
	fmt.Println("Error: No space in between")
	return
	}
	num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
	if err != nil {
	fmt.Println("Error: Invalid Number")
	return
	}
	
	if err == nil && i-num >= 0 {
	startIndex := i - num
	for j := startIndex; j < i; j++ {
	words[j] = cap(words[j])
	}
	words[i] = ""
	words[i+1] = ""
	} else if i-num < 0 {
	fmt.Println("Error: No enough words")
	return
	}
	}
	}
	
	regex = regexp.MustCompile(`\s+`)
	result := regex.ReplaceAllString(strings.Join(words, " "), " $1")
	result = punct(result)
	result = quotes(result)
	result = doublequotes(result)
	result = replaceVowelsWithAn(result)
	output := os.Args[2]
	if err := ioutil.WriteFile(output, []byte(result), 0644); err != nil {
	fmt.Println(err.Error())
	return
	}
	
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
	
	func punct(text string) string {
	regex := regexp.MustCompile(`\s*([.,!?;:]+)\s*`)
	transformed := regex.ReplaceAllString(text, "$1")
	re := regexp.MustCompile(`([.,!?;:]+)\s*`)
	suii := re.ReplaceAllString(transformed, "$1 ")
	return suii
	}
	
	func quotes(text string) string {
	regex := regexp.MustCompile(`'\s*(.*?)\s*'`)
	transformed := regex.ReplaceAllStringFunc(text, func(match string) string {
	word := strings.Trim(match, "' ")
	return "'" + word + "'"
	})
	return transformed
	}
	
	func doublequotes(text string) string {
	regex := regexp.MustCompile(`"\s*(.*?)\s*"`)
	transformed := regex.ReplaceAllStringFunc(text, func(match string) string {
	word := strings.Trim(match, "\" ")
	return "\"" + word + "\""
	})
	return transformed
	}
	
	func replaceVowelsWithAn(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
	word := words[i]
	nextWord := words[i+1]
	if (word == "a" || word == "A") && startsWithVowelOrH(nextWord) {
	if word == "a" {
	words[i] = "an"
	} else {
	words[i] = "An"
	}
	}
	}
	replaced := strings.Join(words, " ")
	return replaced
	}
	
	func startsWithVowelOrH(word string) bool {
	if len(word) == 0 {
	return false
	}
	firstChar := word[:1]
	return strings.ContainsAny(strings.ToLower(firstChar), "aeiouh")
	}
	
	func nonenglishchar(str string) error {
	regex := regexp.MustCompile(`[^\x00-\x7F]`)
	if regex.MatchString(str) {
	return fmt.Errorf("Error: The string contains non english characters")
	}
	return nil
	}
	
	func brackets(x string) error {
	regex := regexp.MustCompile(`\s+\)`)
	if regex.MatchString(x) {
	return fmt.Errorf("Error: There is multiple spaces after the number")
	}
	return nil
	} 
	
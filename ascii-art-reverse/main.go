package main
import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"
)
func printUsage() {
	fmt.Println("Usage: go run . [OPTION]")
	fmt.Println("EX: go run . --reverse=<fileName>")
	os.Exit(0)
}
func pretty(x ...any) {
	for _, v := range x {
		b, _ := json.MarshalIndent(v, "", "  ")
		fmt.Println(string(b))
	}
}
func reverse(filepath string) string {
	fonts := "standard.txt"
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v\n", err)
		os.Exit(1)
	}
	splits := strings.Split(string(input), "\n")
	if len(splits) != 8 {
		fmt.Println("The ascii art to reverse must be 8 lines of height")
		os.Exit(1)
	}
	spaces := findSpace(splits)
	userInput := splitUserInput(splits, spaces)
	userInputMap := userInputMapping(userInput)
	// pretty(userInputMap)
	asciiGraphic := getASCIIgraphicFont(fonts)
	output := mapUserInputWithASCIIgraphicFont(userInputMap, asciiGraphic)
	return output
}
func findSpace(splits []string) []int {
	var emptyColumns []int
	for column := 0; column < len(splits[0]); column++ {
		count := 0
		for row := 0; row < len(splits)-1; row++ {
			if splits[row][column] == ' ' {
				count++
			} else {
				count = 0
				break
			}
			if count == len(splits)-1 {
				emptyColumns = append(emptyColumns, column)
				count = 0
			}
		}
	}
	count := 5
	var indexToRem []int
	for i := range emptyColumns {
		if count == 0 {
			count = 5
			continue
		}
		if i > 0 && emptyColumns[i] == emptyColumns[i-1]+1 {
			indexToRem = append(indexToRem, i)
			count--
		}
	}
	for i := len(indexToRem) - 1; i >= 0; i-- {
		emptyColumns = removeIndex(emptyColumns, indexToRem[i])
	}
	return emptyColumns
}
func removeIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}
func splitUserInput(matrix []string, emptyColumns []int) string {
	var result strings.Builder
	start := 0
	end := 0
	for _, column := range emptyColumns {
		if end < len(matrix[0]) {
			end = column
			for _, characters := range matrix {
				if len(characters) > 0 {
					columns := characters[start:end]
					result.WriteString(columns)
					result.WriteString(" ")
				}
				result.WriteString("\n")
			}
			start = end + 1
		}
	}
	return result.String()
}
func userInputMapping(result string) map[int][]string {
	strSlice := strings.Split(result, "\n")
	var buf []string
	charNb := 0
	graphicInput := make(map[int][]string)
	for li, l := range strSlice {
		if li != 0 && li%8 == 0 {
			graphicInput[charNb] = buf
			charNb++
			buf = []string{}
		}
		buf = append(buf, l)
	}
	return graphicInput
}
func getASCIIgraphicFont(fonts string) map[int][]string {
	readFile, err := ioutil.ReadFile(fonts)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
		return nil
	}
	slice := strings.Split(string(readFile), "\n")
	ascii := make(map[int][]string)
	i := 31
	for _, ch := range slice {
		if ch == "" {
			i++
		} else {
			ascii[i] = append(ascii[i], ch)
		}
	}
	return ascii
}
func mapUserInputWithASCIIgraphicFont(graphicInput, ascii map[int][]string) string {
	var keys []int
	for k := range graphicInput {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var output string
	var sliceOfBytes []byte
	for _, value := range keys {
		graphicValue := graphicInput[value]
		for asciiKey, asciiValue := range ascii {
			if reflect.DeepEqual(asciiValue, graphicValue) {
				sliceOfBytes = append(sliceOfBytes, byte(asciiKey))
			}
		}
		output = string(sliceOfBytes)
	}
	return output
}
func main() {
	var filepath string
	flag.StringVar(&filepath, "reverse", "example.txt", "read file from flag")
	flag.Parse()
	if strings.TrimSpace(filepath) == "" {
		printUsage()
	}
	out := reverse(filepath)
	fmt.Println(out)
}



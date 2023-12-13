package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 3 {
		e := os.Remove("test.txt")
		if e != nil {
			return
		}
		file2, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0o600)
		defer file2.Close() // on ferme automatiquement à la fin de notre programme

		line := strings.Split(os.Args[1], "\\n")
		content := getTable()
		if len(getTable()) == 0 {
			fmt.Println("impossible de trouver le fichier")

			return
		}
		for i := 0; i < len(line); i++ {
			if len(line[i]) > 0 {
				chars := []rune(line[i])
				for n := 0; n < 8; n++ {
					for v := 0; v < len(chars); v++ {
						char := (int(chars[v]) - 32) * 9
						charLine := content[char+1+n]
						fmt.Print(charLine)
						_, err := file2.WriteString(string(charLine)) // écrire dans le fichier
						if err != nil {
							panic(err)
						}
					}
					fmt.Print(string(rune('\n')))
					_, err := file2.WriteString("\n") // écrire dans le fichier
					if err != nil {
						panic(err)
					}
				}

			} else {
				fmt.Print(string(rune('\n')))
			}
		}
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . something standard")
	}
}

func getTable() []string {
	var table []string
	args := os.Args[1:]
	if len(args) == 2 {
		if string(args[1]) == "standard" || string(args[1]) == "standard.txt" || string(args[1]) == "Standard" || string(args[1]) == "Standard.txt" {
			file, err := os.Open("standard.txt") // lire le fichier text.txt
			content, _ := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			table = strings.Split(string(content), "\n")
		} else if string(args[1]) == "Shadow" || string(args[1]) == "Shadow.txt" || string(args[1]) == "shadow" || string(args[1]) == "shadow.txt" {
			file, err := os.Open("Shadows.txt") // lire le fichier text.txtls
			content, _ := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			table = strings.Split(string(content), "\n")
		} else if string(args[1]) == "Thinkertoy.txt" || string(args[1]) == "Thinkertoy" || string(args[1]) == "thinkertoy.txt" || string(args[1]) == "thinkertoy" {
			file, err := os.Open("Tinkertoy.txt") // lire le fichier text.txt
			content, _ := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			table = strings.Split(string(content), "\n")
		}
	}
	return table
}

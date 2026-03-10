package main
import (
	"fmt"
"strings"
"strconv"
// "slices"
"os"
)

func capitalize(str string) string{
	capSlice := []byte(str)
	capSlice[0] = capSlice[0] - 32
	str = string(capSlice)
	// fmt.Println(str)
	return str
}

func uppercase(str string) string{
	str = strings.ToUpper(str)
	return str
}

func lowercase(str string) string{
	str = strings.ToLower(str)
	return str
}

func formatBasedOnNumb(key, numb string, sent []string, index int) {
	// replace ',' in the keyword with ')'
	modKey := strings.Replace(key, ",", ")",1)

	// convert string to number and replace ')' with empty space
	modNumb,_ := strconv.Atoi(strings.Replace(numb, ")", "", 1))

	if modKey == "(cap)" {
		for i:=0; i<modNumb; i++{
		   sent[index - 1]	= capitalize(sent[index - 1 ])
			index = index - 1
		}
	} else if modKey == "(low)" {
		for i:=0; i<modNumb; i++{
		   sent[index - 1]	= lowercase(sent[index - 1 ])
			index = index - 1
		}
	} else if modKey == "(up)" {
		for i:=0; i<modNumb; i++ {
			sent[index - 1] = uppercase(sent[index - 1])
			index = index - 1
		}
	}

}

func formatSentence(s string) string{
	words := strings.Split(s, " ")
	// var output string
	formatedWords := []string{}
	quoteIndex := []int{}

	for idx, _ := range words {

		// Capitalize
		if words[idx] == "(cap)" {
			words[idx-1] = capitalize(words[idx-1])
			words[idx] = ""
			// words = append(words[:idx], words[idx + 1:])
			

			// To Uppercase
		} else if words[idx] == "(up)" {
			words[idx-1] = uppercase(words[idx-1])
			words[idx] = ""

			// To LowerCase
		} else if words[idx] == "(low)" {
			words[idx - 1] = lowercase(words[idx-1])
			words[idx] = ""
		} else if words[idx] == "(hex)" {
			// convert from base 16 to base 10
			s, _ := strconv.ParseInt(words[idx -1], 16, 32)

			// convert from int to string
			val := strconv.FormatInt(s,10)
			words[idx -1] = val
			words[idx] = ""
		}  else if words[idx] == "(bin)" {
			// convert binary to decimal
			s, _ := strconv.ParseInt(words[idx -1], 2, 32)

			// convert int to string
			val := strconv.FormatInt(s, 10)
			words[idx - 1] = val
			
			words[idx] = ""

			// To execute keywords based on number
		} else if strings.Contains(words[idx],"(") {
			// fmt.Println(words[idx])
			slice := words[idx:idx+2]
			formatBasedOnNumb(slice[0], slice[1], words, idx)
			words[idx] = ""
			words[idx + 1] = ""

			// Changing A to An
		}  else if words[idx] == "A" || words[idx] == "a" {
			slice := []byte(words[idx + 1])
			if slice[0] == 'a' || slice[0] == 'e' || slice[0] == 'i' || slice[0] == 'o' || slice[0] == 'u' || slice[0] == 'h' {
				words[idx] = words[idx] + "n"
			}

			// To remove space infront of punctuation
		}  else if (strings.Contains(words[idx],"." ) || strings.Contains(words[idx],"," ) || strings.Contains(words[idx],"!" ) || strings.Contains(words[idx],"?" ) || strings.Contains(words[idx],":" ) || strings.Contains(words[idx],";" ) ){
			slice := []byte(words[idx])
			
			// fmt.Println(string(slice))

			if (len(slice) > 1) && (slice[0] < 48 && slice[1] < 48) {
				words[idx - 1] = words[idx - 1] + words[idx]
				
				// To print the remaining alphabet if the following words is not a symbol
				if slice[1] > 64 {
					words[idx] = string(slice[1:])
				} else {
					words[idx] = ""
				}
				
			} else if slice[0] < 48 || (slice[0] > 57 && slice[0] < 65){
				if words[idx - 1] == "" {
					words[idx - 2] = words[idx - 2] + string(slice[0])
				} else {
					words[idx - 1] = words[idx - 1] + string(slice[0])
				}
					words[idx] = string(slice[1:])
			} 
		
			// Save index of single quotes
	 	} else if strings.Contains(words[idx], "'") {
			quoteIndex = append(quoteIndex, idx)
		} 
	
	}

		// fmt.Println(quoteIndex)


	
					// To remove spaces in quotes

	for index,_ := range words {
		 if strings.Contains(words[index], "'"){
		words[quoteIndex[0] + 1] = words[quoteIndex[0]] + words[quoteIndex[0] + 1]
		words[quoteIndex[0]] = ""

		words[quoteIndex[1] - 1]  = words[quoteIndex[1] - 1] + words[quoteIndex[1]]
		words[quoteIndex[1]]  = ""

		} 
	}
	
	
	// To remove spaces
	for index,_:= range words {
		if words[index] == "" {
			continue

			}  else {
			formatedWords = append(formatedWords, words[index])
		} 
	}

	
	output := strings.Replace(strings.Join(formatedWords," "), " ,", ",", -1)
	// output = strings.TrimSpace(output)

	// output:= strings.Join(formatedWords," ")
	// fmt.Println(output)
	return output
}

func main() {
	// // Getting the input from .txt
	inputPath := os.Args[1]
	content, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("File Error: %v", err)
	}	

	// defectSent := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	// formatSentence(defectSent)

	// defectSent2 := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	// formatSentence(defectSent2)

	// formatSentence("Punctuation tests are ... kinda boring ,what do you think ?")
	// formatSentence("There it was. A amazing rock!")
	// formatSentence("There is no greater agony than bearing a untold story inside you.")

	// Saving the Output
	outputPath := os.Args[2]
	data := []byte(formatSentence(string(content)))
	err2 := os.WriteFile(outputPath, data, 0644)
	if err2 != nil {
		fmt.Printf("Error: %v", err2)
	}
	
}

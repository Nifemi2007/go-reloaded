package main
import ("fmt"
"strings"
"strconv"
// "slices"
// "os"
)



func formatSentence(s string) string{
	words := strings.Split(s, " ")

	for idx, _ := range words {

		// Capitalize
		if words[idx] == "(cap)" {
			capSlice := []byte(words[idx -1])
			capSlice[0] = capSlice[0] - 32
			words[idx - 1] = string(capSlice)
			words[idx] = ""
			// To Uppercase
		} else if words[idx] == "(up)" {
			words[idx - 1] = strings.ToUpper(words[idx - 1])
			words[idx] = ""

			// To LowerCase
		} else if words[idx] == "(low)" {
			words[idx - 1] = strings.ToLower(words[idx -1])
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
			words[idx -1] = val
			
			words[idx] = ""

			// To execute keywords based on number
		} else if strings.Contains(words[idx],"(") {
			slice := words[idx:idx+2]
			fmt.Println(slice[0])
		}
	}
	// output := strings.Join(words, " ")
	// fmt.Println(output)
	return s
}

func main() {
	// // Getting the input from .txt
	// inputPath := os.Args[1]
	// content, err := os.ReadFile(inputPath)
	// if err != nil {
	// 	fmt.Printf("File Error: %v", err)
	// }	

	defectSent := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	formatSentence(defectSent)

	defectSent2 := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	formatSentence(defectSent2)


	// // Saving the Output
	// outputPath := os.Args[2]
	// data := []byte(formatSentence(string(content)))
	// err2 := os.WriteFile(outputPath, data, 0644)
	// if err2 != nil {
	// 	fmt.Printf("Error: %v", err2)
	// }
	
}
// // func returnDecimal() {

// // }

// func getKeyword(words string) {
// 	// keywords := []string{}
// 	// if words[i] = '('  get the word until you see ')'
// 	var a,b int;
// 	keywordSlice := []string{}
// 	sliceIndex := []int{}
// 	defectWord := []string{}
// 	for i, val := range words {
// 		// if(string(val) == " ") {
// 		// 	fmt.Println(val)
// 		// }

// 		// Get the keyword
// 		if (string(val) == "(") {
// 			a = i

// 			defectWord = append(defectWord, words[:(a-1)])
// 		} else if(string(val) == ")") {
// 			b = i
// 			keywordSlice = append(keywordSlice, words[(a+1):b])
// 				// break;
// 			}  

		
// 	}

// 	// Edit word based on keyword
	
	
// 	// splitKeyword := strings.Split(keywordSlice, ",")
// 	fmt.Println(keywordSlice[2])
// 	fmt.Println(sliceIndex)
// 	// fmt.Println(defectWord)
// 	// fmt.Println(splitKeyword)
// }

// func main() {
// 	// result := getKeyword("hey")
// 	testSentence := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
// 	getKeyword(testSentence)

// 	fmt.Println(strings.ToUpper("heyIU"))
// 	// arr1 := [2]string{"man", "Bola"}
// 	// fmt.Println(arr1.split(","))
// 	// fmt.Println("hey")
// }
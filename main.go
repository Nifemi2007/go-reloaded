package main
import ("fmt"
// "strings"
"os"
)

func formatSentence(s string) string{
	
	return s
}

func main() {
	// Getting the input from .txt
	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if(err != nil) {
		fmt.Printf("File Error: %v", err)
		return
	}





	// Saving the Output
	outputPath := os.Args[2]
	data := []byte(formatSentence(string(content)))

	err2 := os.WriteFile(outputPath, data, 0644)
	if(err2 != nil) {
		fmt.Printf("Error Value: %v", err2)
		return
	}
	
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
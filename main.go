package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var words []string

func main() {
	fmt.Println("Suggested Start: adieu")
	readFile()
	play()
}

func readFile() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 5 {
			words = append(words, word)
		}
	}

	msg := fmt.Sprintf("Loaded %d words", len(words))
	fmt.Println(msg)
}

func play() {
	reader := bufio.NewReader(os.Stdin)
	guessCount := 1

	fmt.Println("New eliminated letters can be entered as a single string i.e. adie")
	fmt.Println("Known letters should be entered in sets of 5, with asterisks denoting unkown characters")
	fmt.Println("If you know the location of a character is correct, enter it uppercase")
	fmt.Println("If you know the location of a character is incorrect, enter it lowercase")
	fmt.Println("Example, for the word Humor, *Uo** would indicate the letters U and O are known, that U must be in second position, and that O is in the string but NOT in third position")

	for {
		fmt.Println("Enter new eliminated letters")
		excluded, _ := reader.ReadString('\n')
		excluded = strings.TrimSpace(excluded)

		fmt.Println("Enter known letters")
		known, _ := reader.ReadString('\n')
		known = strings.TrimSpace(known)

		processEliminatedLetters(excluded)
		processKnownLetters(known)
		guessCount++

		//fmt.Printf("Suggested Guess: %s\n", words[0])
		fmt.Printf("Guesses Remaining: %d\n", 6-guessCount)

		if len(words) <= 10 {
			fmt.Printf("Possible Solutions: %s\n", words)
		}
	}
}

func processEliminatedLetters(letters string) {
	fmt.Println("=== Filtering out words that contain eliminated letters ===")

	eliminatedCount := 0

	for _, letter := range letters {
		var keepWords []string

		for _, word := range words {
			if !strings.Contains(word, string(letter)) {
				keepWords = append(keepWords, word)
			} else {
				eliminatedCount++
			}
		}

		words = keepWords
	}

	fmt.Printf("Eliminated: %d Remaining: %d\n", eliminatedCount, len(words))
}

func processKnownLetters(letters string) {
	findBySubstring(letters)
	findByPosition(letters)
	eliminateByPosition(letters)
	findBestGuess(letters)
}

func findBySubstring(letters string) {
	letters = strings.ReplaceAll(letters, "*", "")
	letters = strings.ToLower(letters)

	fmt.Println("=== Filtering out words that don't contain known letters ===")

	eliminatedCount := 0

	for _, letter := range letters {
		var keepWords []string

		for _, word := range words {
			if strings.Contains(word, string(letter)) {
				keepWords = append(keepWords, word)
			} else {
				eliminatedCount++
			}
		}

		words = keepWords
	}

	fmt.Printf("Eliminated: %d Remaining: %d\n", eliminatedCount, len(words))
}

func findByPosition(letters string) {
	fmt.Println("=== Filtering out words that don't contain specified letters at specified positions ===")

	eliminatedCount := 0

	for i, letter := range letters {
		if letter == '*' || unicode.IsLower(letter) {
			continue
		}

		lower := strings.ToLower(string(letter))

		var keepWords []string

		for _, word := range words {
			if rune(word[i]) == rune(lower[0]) {
				keepWords = append(keepWords, word)
			} else {
				eliminatedCount++
			}
		}

		words = keepWords
	}

	fmt.Printf("Eliminated: %d Remaining: %d\n", eliminatedCount, len(words))
}

func eliminateByPosition(letters string) {
	fmt.Println("=== Filtering out words that don't contain specified letters at known wrong position ===")

	eliminatedCount := 0

	for i, letter := range letters {
		if letter == '*' || unicode.IsUpper(letter) {
			continue
		}

		var keepWords []string

		for _, word := range words {
			if rune(word[i]) != letter {
				keepWords = append(keepWords, word)
			} else {
				eliminatedCount++
			}
		}

		words = keepWords
	}

	fmt.Printf("Eliminated: %d Remaining: %d\n", eliminatedCount, len(words))
}

func findBestGuess(letters string) {
	letters = strings.ReplaceAll(letters, "*", "")
	letters = strings.ToLower(letters)

	guess := ""
	mostUnique := 0

	skips := []rune(letters)
	var alternatives []string

	for _, word := range words {
		var uniqueLetters = make(map[rune]int)

		for _, letter := range word {
			if contains(skips, letter) {
				continue
			}
			uniqueLetters[letter] = 1
		}

		currentUnique := len(uniqueLetters)

		if currentUnique > mostUnique {
			guess = word
			mostUnique = currentUnique
			alternatives = []string{}
		} else if currentUnique == mostUnique {
			alternatives = append(alternatives, word)
		}
	}

	if len(alternatives) > 10 {
		alternatives = alternatives[0:10]
	}

	fmt.Printf("Suggested Guess: %s\n", guess)
	fmt.Printf("Alternatives: %s\n", alternatives)
}

func contains(s []rune, r rune) bool {
	for _, a := range s {
		if a == r {
			return true
		}
	}
	return false
}

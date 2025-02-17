package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type thirdRotor struct {
	dictionary []struct {
		in  int
		out int
	}
	codePosition   int
	decodePosition int
}

func newThirdRotor() *thirdRotor {
	return &thirdRotor{
		dictionary: []struct {
			in  int
			out int
		}{
			{1, 3}, {5, 20}, {17, 2}, {2, 15}, {14, 18}, {8, 1}, {12, 13}, {18, 10}, {9, 5}, {13, 24}, {3, 19}, {23, 6}, {6, 17}, {16, 12}, {25, 21}, {15, 26}, {4, 8}, {10, 25}, {21, 9}, {26, 14}, {22, 11}, {19, 16}, {11, 23}, {24, 7}, {7, 4}, {20, 22},
		},
		codePosition:   0,
		decodePosition: 0,
	}
}

func (r *thirdRotor) code(input int, direction int) int {
	r.codePosition++
	if r.codePosition > 26 {
		r.codePosition -= 26
	}
	search := r.cycleSearch(input + r.codePosition)
	result := 0
	for i := 0; i < 26; i++ {
		if r.dictionary[i].in == search && switchDirection(direction) == 0 || r.dictionary[i].out == search && switchDirection(direction) == 1 {
			if switchDirection(direction) == 0 {
				result = r.dictionary[i].out
			} else {
				result = r.dictionary[i].in
			}
			break
		}
	}
	return result
}

func (r *thirdRotor) decode(input int, direction int) int {
	r.decodePosition++
	if r.decodePosition > 26 {
		r.decodePosition -= 26
	}
	realPosition := r.decodePosition
	if direction == 0 {
		realPosition++
	} else {
		realPosition--
	}
	result := 0
	for i := 0; i < 26; i++ {
		if r.dictionary[i].in == input && switchDirection(direction) == 0 || r.dictionary[i].out == input && switchDirection(direction) == 1 {
			if switchDirection(direction) == 0 {
				result = r.dictionary[i].out - realPosition
			} else {
				result = r.dictionary[i].in - realPosition
			}

			for result <= 0 {
				result += 26
			}
			break
		}
	}
	return r.cycleSearch(result)
}

func (r *thirdRotor) cycleSearch(search int) int {
	for search > 26 {
		search -= 26
	}
	return search
}

type stator struct {
	dictionary []struct {
		in  int
		out int
	}
}

func newStator() *stator {
	return &stator{
		dictionary: []struct {
			in  int
			out int
		}{
			{5, 24}, {14, 19}, {9, 23}, {26, 13}, {22, 7}, {20, 12}, {17, 1}, {21, 25}, {16, 2}, {18, 3}, {11, 8}, {10, 6}, {4, 15}, {24, 5}, {19, 14}, {23, 9}, {13, 26}, {7, 22}, {12, 20}, {1, 17}, {25, 21}, {2, 16}, {3, 18}, {8, 11}, {6, 10}, {15, 4},
		},
	}
}

func (s *stator) find(input int) int {
	for i := 0; i < 26; i++ {
		if s.dictionary[i].in == input {
			return s.dictionary[i].out
		}
	}
	return 0 // Should not reach here in normal operation
}

type secondRotor struct {
	dictionary []struct {
		in  int
		out int
	}
	codePosition   int
	decodePosition int
}

func newSecondRotor() *secondRotor {
	return &secondRotor{
		dictionary: []struct {
			in  int
			out int
		}{
			{18, 8}, {4, 6}, {13, 5}, {23, 24}, {12, 25}, {19, 21}, {2, 14}, {24, 16}, {3, 19}, {7, 1}, {10, 23}, {6, 2}, {20, 13}, {15, 12}, {14, 10}, {5, 20}, {17, 3}, {16, 17}, {9, 22}, {25, 15}, {11, 18}, {8, 26}, {1, 7}, {21, 9}, {22, 11}, {26, 4},
		},
		codePosition:   0,
		decodePosition: 0,
	}
}

func (r *secondRotor) code(input int, direction int) int {
	r.codePosition++
	realPosition := r.codePosition

	search := r.cycleSearch(input + r.cycleItterator(realPosition))
	result := 0
	for i := 0; i < 26; i++ {
		if r.dictionary[i].in == search && switchDirection(direction) == 0 || r.dictionary[i].out == search && switchDirection(direction) == 1 {
			if switchDirection(direction) == 0 {
				result = r.dictionary[i].out
			} else {
				result = r.dictionary[i].in
			}
			break
		}
	}
	return result
}

func (r *secondRotor) decode(input int, direction int) int {
	r.decodePosition++

	realPosition := r.decodePosition
	if direction == 0 {
		realPosition++
	} else {
		realPosition--
	}

	result := 0
	for i := 0; i < 26; i++ {
		if r.dictionary[i].in == input && switchDirection(direction) == 0 || r.dictionary[i].out == input && switchDirection(direction) == 1 {
			if switchDirection(direction) == 0 {
				result = r.dictionary[i].out - r.cycleItterator(realPosition)
			} else {
				result = r.dictionary[i].in - r.cycleItterator(realPosition)
			}
			for result <= 0 {
				result += 26
			}
			break
		}
	}
	return r.cycleSearch(result)
}

func (r *secondRotor) cycleItterator(itterator int) int {
	for itterator > 26 {
		itterator /= 26
	}
	return itterator
}

func (r *secondRotor) cycleSearch(search int) int {
	for search > 26 {
		search -= 26
	}
	return search
}

type firstRotor struct {
	dictionary []struct {
		in  int
		out int
	}
	codePosition   int
	decodePosition int
}

func newFirstRotor() *firstRotor {
	return &firstRotor{
		dictionary: []struct {
			in  int
			out int
		}{
			{24, 19}, {19, 20}, {5, 13}, {6, 7}, {2, 5}, {15, 25}, {8, 17}, {9, 4}, {20, 10}, {12, 11}, {25, 23}, {11, 6}, {18, 8}, {23, 2}, {21, 16}, {16, 26}, {26, 1}, {7, 21}, {3, 3}, {4, 12}, {14, 15}, {22, 9}, {17, 24}, {13, 22}, {1, 14}, {10, 18},
		},
		codePosition:   0,
		decodePosition: 0,
	}
}

func (r *firstRotor) code(input int, direction int) int {
	r.codePosition++
	realPosition := r.codePosition

	search := r.cycleSearch(input + r.cycleItteratorSqrt(realPosition))
	result := 0
	for i := 0; i < 26; i++ {
		if r.dictionary[i].in == search && switchDirection(direction) == 0 || r.dictionary[i].out == search && switchDirection(direction) == 1 {
			if switchDirection(direction) == 0 {
				result = r.dictionary[i].out
			} else {
				result = r.dictionary[i].in
			}
			break
		}
	}
	return result
}

func (r *firstRotor) decode(input int, direction int) int {
	r.decodePosition++

	realPosition := r.decodePosition
	if direction == 0 {
		realPosition++
	} else {
		realPosition--
	}
	result := 0
	for i := 0; i < 26; i++ {
		if r.dictionary[i].in == input && switchDirection(direction) == 0 || r.dictionary[i].out == input && switchDirection(direction) == 1 {
			if switchDirection(direction) == 0 {
				result = r.dictionary[i].out - r.cycleItteratorSqrt(realPosition)
			} else {
				result = r.dictionary[i].in - r.cycleItteratorSqrt(realPosition)
			}

			for result <= 0 {
				result += 26
			}
			break
		}
	}
	return r.cycleSearch(result)
}

func (r *firstRotor) cycleItteratorSqrt(itterator int) int {
	if itterator/(26*26) > 1 {
		for itterator > (26 * 26) {
			itterator /= (26 * 26)
		}
	} else {
		itterator /= (26 * 26)
	}
	return itterator
}

func (r *firstRotor) cycleSearch(search int) int {
	for search > 26 {
		search -= 26
	}
	return search
}

var (
	rotor3 = newThirdRotor()
	rotor2 = newSecondRotor()
	rotor1 = newFirstRotor()
	stat   = newStator()
	abc    = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func cryptWord(word string) []rune {
	result := make([]rune, 0)
	for _, letter := range word {
		result = append(result, codeByRotors(letter))
	}
	return result
}

func decryptWord(word []rune) []rune {
	result := make([]rune, 0)
	for _, letter := range word {
		result = append(result, decodeByRotors(letter))
	}
	return result
}

func codeByRotors(letter rune) rune {
	number := getLetterIndex(letter) + 1
	thirdRtl := rotor3.code(number, 0)
	secondRtl := rotor2.code(thirdRtl, 0)
	firstRtl := rotor1.code(secondRtl, 0)
	statorRtl := stat.find(firstRtl)
	firstLtr := rotor1.code(statorRtl, 1)
	secondLtr := rotor2.code(firstLtr, 1)
	thirdLtr := rotor3.code(secondLtr, 1)
	result := thirdLtr

	return abc[result-1]
}

func decodeByRotors(letter rune) rune {
	number := getLetterIndex(letter) + 1
	thirdRtl := rotor3.decode(number, 0)
	secondRtl := rotor2.decode(thirdRtl, 0)
	firstRtl := rotor1.decode(secondRtl, 0)
	statorRtl := stat.find(firstRtl)
	firstLtr := rotor1.decode(statorRtl, 1)
	secondLtr := rotor2.decode(firstLtr, 1)
	thirdLtr := rotor3.decode(secondLtr, 1)
	result := thirdLtr

	return abc[result-1]
}

func switchDirection(direction int) int {
	if direction == 1 {
		return 0
	}
	return 1
}

func configureMachine(action string, frtRrStart, sndRrStart, trdRrStart int) {
	if action == "1" {
		rotor1.codePosition = frtRrStart
		rotor2.codePosition = sndRrStart
		rotor3.codePosition = trdRrStart
	} else if action == "2" {
		rotor1.decodePosition = frtRrStart
		rotor2.decodePosition = sndRrStart
		rotor3.decodePosition = trdRrStart
	} else {
		fmt.Println("Error: Unknown action")
		os.Exit(1)
	}
}

func runMachine(action string, inputString string) string {
	result := ""
	stringByWords := strings.Split(inputString, " ")
	if action == "1" {
		for _, word := range stringByWords {
			crypted := cryptWord(word)
			result += string(crypted) + " " // Corrected: Directly convert rune slice to string
		}
	} else if action == "2" {
		for _, word := range stringByWords {
			uncr := decryptWord([]rune(word)) // Convert word to rune slice for decryptWord
			result += string(uncr) + " "      // Corrected: Directly convert rune slice to string
		}
	} else {
		fmt.Println("Error: Unknown action")
		os.Exit(1)
	}
	return result
}

func getLetterIndex(letter rune) int {
	for i, l := range abc {
		if l == letter {
			return i
		}
	}
	return -1 // Should not happen for uppercase letters
}

func main() {
	fmt.Println("Enigma machine")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select input file")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	fmt.Println("Select output file")
	filePath1, _ := reader.ReadString('\n')
	filePath1 = strings.TrimSpace(filePath1)

	file1, err1 := os.Create(filePath1)
	if err1 != nil {
		log.Fatalf("Ошибка при создании файла: %v", err1)
	}
	defer file1.Close()

	// Чтение содержимого файла
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	// Преобразуем данные в строку и выводим
	inputString := string(data)

	fmt.Print("Select the option (1 - Encrypt string, 2 - Decrypt string): ")
	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(action)

	fmt.Print("Set first rotor start positions (1-26): ")
	frtRrStartStr, _ := reader.ReadString('\n')
	frtRrStartStr = strings.TrimSpace(frtRrStartStr)
	frtRrStart, err := strconv.Atoi(frtRrStartStr)
	if err != nil || frtRrStart < 1 || frtRrStart > 26 {
		fmt.Println("Error: Invalid first rotor start position")
		os.Exit(1)
	}

	fmt.Print("Set second rotor start positions (1-26): ")
	sndRrStartStr, _ := reader.ReadString('\n')
	sndRrStartStr = strings.TrimSpace(sndRrStartStr)
	sndRrStart, err := strconv.Atoi(sndRrStartStr)
	if err != nil || sndRrStart < 1 || sndRrStart > 26 {
		fmt.Println("Error: Invalid second rotor start position")
		os.Exit(1)
	}

	fmt.Print("Set third rotor start positions (1-26): ")
	trdRrStartStr, _ := reader.ReadString('\n')
	trdRrStartStr = strings.TrimSpace(trdRrStartStr)
	trdRrStart, err := strconv.Atoi(trdRrStartStr)
	if err != nil || trdRrStart < 1 || trdRrStart > 26 {
		fmt.Println("Error: Invalid third rotor start position")
		os.Exit(1)
	}
	start := time.Now()
	inputString = strings.ToUpper(strings.TrimSpace(inputString))
	reg := regexp.MustCompile("[^a-zA-Z]")
	inputString = reg.ReplaceAllString(inputString, " ")

	if action != "" && frtRrStart >= 1 && sndRrStart >= 1 && trdRrStart >= 1 {
		configureMachine(action, frtRrStart, sndRrStart, trdRrStart)
		result := runMachine(action, inputString)
		fmt.Println("\nResult:")
		fmt.Println(result)
		// Записываем строку в файл
		_, err1 = file1.WriteString(result)
		if err1 != nil {
			log.Fatalf("Ошибка при записи в файл: %v", err1)
		}
	} else {
		fmt.Println("Error: Some settings are missing or failed")
		os.Exit(1)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

}

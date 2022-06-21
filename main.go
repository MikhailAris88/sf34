package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		return
	}
	_ = os.Remove("./output.txt")
	output, err := os.OpenFile("./output.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	writer := bufio.NewWriter(output)
	for reader.Scan() {
		m, err := regexp.MatchString(`\d+[+-]{1}\d+[=]{1}[?]{1}`, reader.Text())
		if err != nil {
			return
		}
		if !m {
			continue
		}

		pattern := regexp.MustCompile(`([\-+]\d+|\d+)`)
		numbers := pattern.FindAllString(reader.Text(), -1)
		var result int
		for _, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				return
			}
			result += n
		}

		replace := regexp.MustCompile(`[?]`)
		_, err = writer.WriteString(replace.ReplaceAllLiteralString(reader.Text(), strconv.Itoa(result)))
		if err != nil {
			return
		}
		writer.WriteString("\n")

	}
	err = writer.Flush()
	if err != nil {
		return
	}
	return
}

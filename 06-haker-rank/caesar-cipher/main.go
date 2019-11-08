package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	lenAlphabet = 26
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	var out strings.Builder
	for _, l := range s {
		switch {
		case unicode.IsLower(l):
			l = rotate(l, k, true)
		case unicode.IsUpper(l):
			l = rotate(l, k, false)
		}
		out.WriteString(string(l))
	}
	return out.String()
}

func rotate(letter rune, k int32, isLower bool) rune {
	lower := []rune("abcdefghijklmnopqrstuvwxyz")
	upper := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var alphabet []rune
	if isLower {
		alphabet = lower
	} else {
		alphabet = upper
	}
	shift := k % 26
	// if we're wrapping around the end with the change k
	if letter+shift > alphabet[lenAlphabet-1] {
		letter = letter - lenAlphabet + shift
	} else {
		letter = letter + shift
	}
	return letter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	_, err = strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

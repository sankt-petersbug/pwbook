package password

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	// LowerLetters is the list of allowed lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of allowed uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of allowed digits.
	Digits = "0123456789"

	// Symbols is the list of allowed symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

// DefaultOptions is the options that will be used if no Options is provided
var DefaultOptions = &Options{
	lowerLetters: true,
	upperLetters: true,
	digits:       true,
	symbols:      true,
}

// Options contains information about the type of generated password
type Options struct {
	lowerLetters, upperLetters, digits, symbols bool
}

func makeCategories(opt *Options) [][]rune {
	var categories [][]rune

	if opt.lowerLetters {
		categories = append(categories, []rune(LowerLetters))
	}
	if opt.upperLetters {
		categories = append(categories, []rune(UpperLetters))
	}
	if opt.digits {
		categories = append(categories, []rune(Digits))
	}
	if opt.symbols {
		categories = append(categories, []rune(Symbols))
	}

	return categories
}

// Generate a password string with given length and options
func Generate(length int, opt *Options) string {
	var buf bytes.Buffer

	if opt == nil {
		opt = DefaultOptions
	}

	categories := makeCategories(opt)
	if len(categories) == 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		category := categories[rand.Intn(len(categories))]
		char := category[rand.Intn(len(category))]
		buf.WriteRune(char)
	}

	return buf.String()
}

func isStrong(s string) bool {
	if !strings.ContainsAny(s, LowerLetters) {
		return false
	}
	if !strings.ContainsAny(s, UpperLetters) {
		return false
	}
	if !strings.ContainsAny(s, Digits) {
		return false
	}
	if !strings.ContainsAny(s, Symbols) {
		return false
	}

	return true
}

// GenerateStrong generate a password that satisfy default password policy
func GenerateStrong() (string, error) {
	const limit = 10000

	for i := 0; i < limit; i++ {
		s := Generate(10, nil)

		if isStrong(s) {
			return s, nil
		}
	}

	msg := fmt.Sprintf("Failed to generate a strong password during %d tries", limit)

	return "", errors.New(msg)
}

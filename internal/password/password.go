package password

import (
    "bytes"
    "math/rand"
    "time"
)

const (
    LowerLetters = "abcdefghijklmnopqrstuvwxyz"
    UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    Digits = "0123456789"
    Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var DefaultOptions = &Options{
    lowerLetters: true,
    upperLetters: true,
    digits: true,
    symbols: true,
}

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
package rand

import (
	"github.com/ismdeep/rand/base"
	"github.com/ismdeep/rand/word"
	"math/rand"
	"strings"
)

func StrWithBase(base string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		j := Intn(len(base))
		s += base[j : j+1]
	}
	return s
}

func Str(n int) string {
	return StrWithBase(base.Base, n)
}

func Password(charSize int, digitSize int, symbolSize int) string {
	s := StrWithBase(base.Char, charSize) + StrWithBase(base.Digit, digitSize) + StrWithBase(base.Symbol, symbolSize)
	items := strings.Split(s, "")
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	return strings.Join(items, "")
}

func PasswordEasyToRemember(n int) string {
	items := make([]string, 0)
	for i := 0; i < n; i++ {
		items = append(items, word.Agwordlist[Intn(len(word.Agwordlist))])
	}
	return strings.Join(items, "-")
}

func PIN(n int) string {
	return Password(0, n, 0)
}

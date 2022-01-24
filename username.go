package rand

import (
	"fmt"
	"github.com/ismdeep/rand/word"
)

func Username() string {
	return fmt.Sprintf("%v-%v", word.Agwordlist[Intn(len(word.Agwordlist))], word.Agwordlist[Intn(len(word.Agwordlist))])
}

package proxy

import (
	"fmt"
	"strings"
	"unicode"
)

type Stringable struct {
	value string
}

func MakeStringable(value string) *Stringable {
	return &Stringable{
		value: value,
	}
}

func (s Stringable) ToSnakeCase() string {
	var result strings.Builder

	for i, r := range s.value {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteByte('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

func (s Stringable) Dd(abstract any) {
	fmt.Println(fmt.Sprintf("dd: %+v", abstract))
}

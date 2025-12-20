package service

import (
	"strings"

	"github.com/Ruginis-AV/finish-six.sprint/pkg/morse"
)

type service struct {
	Converter morse.Converter
}

func New(Converter morse.Converter) *service {
	return &service{
		Converter: morse.DefaultConverter,
	}
}

func (s *service) ConvertString(input string) string {

	mrs := strings.ContainsAny(input, ". -")

	if mrs {
		for _, char := range input {
			if !strings.ContainsRune(". -", char) {
				mrs = false
				break
			}
		}
	}
	if mrs {
		text := s.Converter.ToText(input)
		return text
	} else {
		morse := s.Converter.ToMorse(input)
		return morse
	}
}

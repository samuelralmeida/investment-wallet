package box

import (
	"fmt"
	"slices"
	"strings"
)

var options = map[string][]string{
	"estabildiade":    {"baunilha", "pimenta"},
	"antifragilidade": {"ouro", "dólar"},
}

var optionsList []string

func init() {
	for key, values := range options {
		for _, value := range values {
			optionsList = append(optionsList, fmt.Sprintf("%s:%s", key, value))
		}
	}
}

func OptionsList() []string {
	return optionsList
}

// ValidateOption returns if optins is valid and, if valid, the names of box and flavor
func ValidateOption(option string) (bool, string, string) {
	parts := strings.Split(option, ":")
	if len(parts) != 2 {
		return false, "", ""
	}

	val, ok := options[parts[0]]
	if !ok {
		return false, "", ""
	}

	return slices.Contains(val, parts[1]), parts[0], parts[1]
}

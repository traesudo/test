package test

import (
	"fmt"
	"regexp"
	"strings"
)

func Testphone() {
	var phone = "222213123"
	if strings.Count(phone, " ") > 1 {
		sum := 0
		phone = strings.Map(
			func(r rune) rune {
				println(sum)
				if r == ' ' {
					if sum != 0 {
						r = 'w'
					}
					sum++
				}
				return r
			}, phone)
		phone = strings.ReplaceAll(phone, "w", "")
	}
	digitalPattern := regexp.MustCompile(`\d+`)
	telData := digitalPattern.FindAllString(phone, 2)
	fmt.Println(strings.Count(phone, " "))
	fmt.Println(digitalPattern)
	fmt.Println(telData)
	fmt.Println(len(telData))
}

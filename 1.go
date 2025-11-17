package main

import (
	ru "github.com/Styans/rune"
)

func PrintCombo() {
	first := true

	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for k := j + 1; k <= 9; k++ {

				if !first {
					ru.PrintRune(',')
					ru.PrintRune(' ')
				}
				first = false

				ru.PrintRune(rune(i + '0'))
				ru.PrintRune(rune(j + '0'))
				ru.PrintRune(rune(k + '0'))
			}
		}
	}
}

func string_and_numbers() {
	s := "abcdefghijklmnopqrstuvwxyz"
	a := "0123456789"


	for _, r := range s {
		ru.PrintRune(r)
	}

	ru.PrintRune('\n')

	for _, r := range a {
		ru.PrintRune(r)
	}
}

func string_integars() []int{
  p := "0123456789"
  var nums[]int

  for _, r := range p {
	nums = append(nums, int(r - '0'))
  }
  return nums
  
}

func main() {
    nums := string_integars()
	for _, n := range nums {
		ru.PrintRune(rune(n + '0'))
		// ru.PrintRune(' ')
	}
	ru.PrintRune('\n')
	string_and_numbers()
	ru.PrintRune('\n')
	PrintCombo()
}

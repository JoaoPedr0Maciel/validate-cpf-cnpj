package validator

import (
	"errors"
	"strings"
	"sync"
)

// ValidateCPF validates a CPF number
// @checkReturn
// @param cpf string cpf formatted "00000000000" or "000.000.000-00"
// @return bool true if CPF is valid, false if is invalid
// @return error error during the validation or cpf invalid
func ValidateCPF(cpf string) (bool, error) {
	var wg sync.WaitGroup

	multiplicateNumbers := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	if cpf == "" {
		panic("cpf must be a valid string (000.000.000-00) or (00000000000)")
	}
	cpfWithoutDots := strings.ReplaceAll(cpf, ".", "")
	cleanedCpf := strings.ReplaceAll(cpfWithoutDots, "-", "")

	ch1 := make(chan int)
	ch2 := make(chan int)

	firstDigitUser := int(cleanedCpf[9] - '0')
	secondDigitUser := int(cleanedCpf[10] - '0')

	wg.Add(2)

	go calculatedFirstDigit(cleanedCpf, multiplicateNumbers, ch1, &wg)

	go calculatedSecondDigit(cleanedCpf, multiplicateNumbers, ch2, &wg)

	firstDigitCalculated := <-ch1
	secondDigitCalculated := <-ch2

	wg.Wait()

	if firstDigitCalculated != firstDigitUser && secondDigitCalculated != secondDigitUser {
		return false, errors.New("this CPF is invalid")
	}

	return true, nil
}

func calculatedFirstDigit(cleanedCpf string, multiplicateNumbers []int, ch chan int, wg *sync.WaitGroup) {
	var sum int
	for i, num := range cleanedCpf[0:9] {
		numInt := int(num - '0')
		result := numInt * multiplicateNumbers[1:][i]
		sum += result
	}

	defer wg.Done()

	var fisrtDigitCalculated int

	rest := sum % 11

	if rest <= 2 {
		fisrtDigitCalculated = 0
	} else {
		fisrtDigitCalculated = 11 - rest
	}

	ch <- fisrtDigitCalculated

}

func calculatedSecondDigit(cleanedCpf string, multiplicateNumbers []int, ch chan int, wg *sync.WaitGroup) {
	var sum int
	for i, num := range cleanedCpf[0:10] {
		numInt := int(num - '0')
		result := numInt * multiplicateNumbers[i]
		sum += result
	}

	defer wg.Done()

	var secondDigitCalculated int
	rest := sum % 11

	if rest <= 2 {
		secondDigitCalculated = 0
	} else {
		secondDigitCalculated = 11 - rest
	}

	ch <- secondDigitCalculated
}

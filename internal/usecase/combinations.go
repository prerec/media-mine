package usecase

import (
	"errors"
	"sort"
)

var (
	NoComboErr           = errors.New("no combo")
	NegativeBanknotesErr = errors.New("dont use negative banknotes value")
	NegativeAmountErr    = errors.New("dont use negative amount value")
	ZeroAmountErr        = errors.New("dont use zero amount value")
	ZeroBanknotesErr     = errors.New("dont use zero banknotes value")
)

// CombinationsFinder реализует логику работы по вычислению вариантов размена для указанной суммы денег
func CombinationsFinder(amount int, banknotes []int) ([][]int, error) {

	// Валидируем входные данные
	if err := validateInput(amount, banknotes); err != nil {
		return nil, err
	}

	// Удаляем дубликаты для корректной работы алгоритма поиска комбинаций
	banknotes = removeDuplicates(banknotes)

	// Сортируем массив банкнот для оптимизации поиска комбинаций
	sort.Sort(sort.Reverse(sort.IntSlice(banknotes)))

	// Исключаем ситуации, когда минимальный номинал банкноты больше требуемой суммы
	if len(banknotes) > 0 && banknotes[len(banknotes)-1] > amount {
		return nil, NoComboErr
	}

	var result [][]int

	// find - функция для поиска комбинаций. Принимает на вход текущую сумму 'amount',
	// индекс текущей банкноты 'index', текущую комбинацию банкнот 'current'
	var find func(amount int, index int, current []int)

	find = func(amount int, index int, current []int) {
		// Если текущая сумма = 0, то это значит, что найдена комбинация
		if amount == 0 {
			combo := make([]int, len(current))
			copy(combo, current)
			result = append(result, combo)
			return
		}
		for i := index; i < len(banknotes); i++ {
			// Если номинал банкноты <= суммы,
			if banknotes[i] <= amount {
				// то рекурсивно вызываем find для следующей банкноты, запоминая текущую которая подходит
				find(amount-banknotes[i], i, append(current, banknotes[i]))
			}
		}
	}

	find(amount, 0, []int{})

	if result == nil {
		return nil, NoComboErr
	}
	return result, nil
}

// removeDuplicates удаляет дубликаты номиналов банкнот
func removeDuplicates(slice []int) []int {

	encountered := make(map[int]bool)
	result := []int{}

	for _, v := range slice {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}

// validateNegativeValue исключает работу с отрицательными значениями
func validateNegativeValue(amount int, banknotes []int) error {

	if amount < 0 {
		return NegativeAmountErr
	}

	for _, banknote := range banknotes {
		if banknote < 0 {
			return NegativeBanknotesErr
		}
	}

	return nil
}

// validateZeroValue исключает работу с нулевыми значениями
func validateZeroValue(amount int, banknotes []int) error {

	if amount == 0 {
		return ZeroAmountErr
	}

	for _, banknote := range banknotes {
		if banknote == 0 {
			return ZeroBanknotesErr
		}
	}

	return nil
}

// validateInput объединяет в себе валидирующие функции
func validateInput(amount int, banknotes []int) error {

	if err := validateNegativeValue(amount, banknotes); err != nil {
		return err
	}

	if err := validateZeroValue(amount, banknotes); err != nil {
		return err
	}

	return nil
}

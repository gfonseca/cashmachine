package entity

import "fmt"

// Account represents the main busines rules of the system responsible for make operations on the user's accounts.
type Account struct {
	ID      int
	Balance float32
}

// Withdraw catch some value from users account and return the proper amount of bills
func (acc *Account) Withdraw(value int) (bills []Bill, err error) {
	if value < 1 {
		err = fmt.Errorf("The withdraw value must be at least 1")
		return
	}

	existentValuesStep := 0
	currentBillValue := existentValues[existentValuesStep]

	for i := value; i > 0; {
		if i < currentBillValue {
			existentValuesStep++
			currentBillValue = existentValues[existentValuesStep]
			continue
		}

		b := Bill{currentBillValue}
		bills = append(bills, b)
		i -= currentBillValue
	}

	acc.Balance -= float32(value)

	return
}

// VerifyBalance check if withdraw is possible
func (acc Account) VerifyBalance(value float32) bool {
	return acc.Balance >= value
}

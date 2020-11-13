package entity

import "fmt"

var existentValues = []int{50, 10, 5, 1}

// Bill represents an monetary unit with a value associated to it
type Bill struct {
	Value int
}

func billExists(value int) (exists bool) {
	exists = false
	for i := range existentValues {
		if existentValues[i] == value {
			exists = true
		}
	}
	return
}

// NewBill Return a new bill with current value
func NewBill(value int) (*Bill, error) {
	if valueExists := billExists(value); !valueExists {
		return nil, fmt.Errorf("[Error] Invalid bill value")
	}

	return &Bill{Value: value}, nil
}

//String return an string representation for a Bill
func (b Bill) String() string {
	return fmt.Sprintf("<Bill $%d>", b.Value)
}

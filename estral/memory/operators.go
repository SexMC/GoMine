package memory

import "github.com/google/uuid"

var operators []uuid.UUID

func AddOperator(uuid uuid.UUID) {
	operators = append(operators, uuid)
}

func OperatorExists(uuid uuid.UUID) bool {
	for _, id := range operators {
		if uuid == id {
			return true
		}
	}

	return false
}

func RemoveOperator(uuid uuid.UUID) {
	for index, id := range operators {
		if uuid == id {
			operators[index] = operators[len(operators)-1]
			operators = operators[:len(operators)-1]
		}
	}
}

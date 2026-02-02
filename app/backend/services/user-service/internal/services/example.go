package services

import (
	"rmn-backend/services/user-service/internal/logic/example"
)

var exampleLogic *example.ExampleLogic

func RegisterExample(logic *example.ExampleLogic) {
	exampleLogic = logic
}

func GetExampleLogic() *example.ExampleLogic {
	return exampleLogic
}

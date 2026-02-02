package services

import (
	"mtsgn-access-user-svc/internal/logic/example"
)

var exampleLogic *example.ExampleLogic

func RegisterExample(logic *example.ExampleLogic) {
	exampleLogic = logic
}

func GetExampleLogic() *example.ExampleLogic {
	return exampleLogic
}

package services

type ExampleService struct{}

func (es *ExampleService) PerformOperation() (string, error) {
	// Perform some operation on a data model
	return "Operation result", nil
}

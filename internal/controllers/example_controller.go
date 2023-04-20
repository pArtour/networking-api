package controllers

type ExampleController struct{}

func (ec *ExampleController) GetExample() (string, error) {
	// Your business logic here
	return "Hello, world!", nil
}

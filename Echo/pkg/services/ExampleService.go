package services

type exampleService struct {
	// This is where we would store and dependancies we want to call, i.e repositories.
}

func NewExampleService( /*we would pass dependencies here as part of the constructor*/ ) ExampleService {
	return &exampleService{
		/*link the constructor variable to the struct*/
	}
}

type ExampleService interface {
	ExposedFunction() (int, any)
}

func (s *exampleService) ExposedFunction() (int, any) {
	return 200,
		struct {
			Status string `json:"Status"`
		}{"Hello, World. - Echo"}
}

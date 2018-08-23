package blockchain

type GetDemoInService struct {
	Test string
}
type GetDemoOutService struct {
	Status bool
}

func (this GetDemoInService) GetDemo() GetDemoOutService {
	return GetDemoOutService{true}
}

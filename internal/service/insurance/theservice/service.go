package theservice

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []TheService {
	return allEntities
}

func (s *Service) Get(idx int) (*TheService, error) {
	return &allEntities[idx], nil
}

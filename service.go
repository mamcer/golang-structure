package main

type Service struct {
	Repository repository
}

func NewService(r repository) *Service {
	return &Service{r}
}

func (s *Service) GetMessage(id int) string {
	return s.Repository.GetMessageByID(id)
}

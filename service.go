package main

type Service struct {
	Repository repository
}

func NewService(r repository) *Service {
	return &Service{r}
}

func (s *Service) GetMessage(id int) string {
	m, e := s.Repository.GetMessageByID(id)
	if e == nil {
		return m
	}

	return ""
}

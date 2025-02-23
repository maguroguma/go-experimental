package dummy

func NewService(client Client) Service {
	return &service{client: client}
}

type service struct {
	client Client
}

func (s *service) GetItem() (string, error) {
	item, err := s.client.Get()
	if err != nil {
		return "", err
	}
	return item, nil
}

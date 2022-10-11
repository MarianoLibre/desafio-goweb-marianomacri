package tickets

import "github.com/MarianoLibre/desafio-goweb-marianomacri/internal/domain"

type Service interface {
    GetAll() ([]domain.Ticket, error)
    GetTicketByDestination(destination string) ([]domain.Ticket, error)
}


type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Ticket, error) {

    return nil, nil
}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

    return nil, nil
}

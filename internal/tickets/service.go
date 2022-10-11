package tickets

import (
	"context"

	"github.com/MarianoLibre/desafio-goweb-marianomacri/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetMostVisited(ctx context.Context) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	ts, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	ts, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (s *service) GetMostVisited(ctx context.Context) (string, error) {
	return s.repository.GetMostVisited(ctx)
}

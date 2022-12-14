package handler

import (
	"net/http"

	"github.com/MarianoLibre/desafio-goweb-marianomacri/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		tickets, err := s.service.GetAll(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTicketByDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) MostVisited() gin.HandlerFunc {
	return func(c *gin.Context) {
        mv, err := s.service.GetMostVisited(c)
        if err != nil {
            c.String(http.StatusInternalServerError, err.Error(), nil)
        }

        c.JSON(200, gin.H{"most visited country": mv})
	}
}

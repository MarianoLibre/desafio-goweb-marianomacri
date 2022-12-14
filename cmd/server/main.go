package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/MarianoLibre/desafio-goweb-marianomacri/cmd/server/handler"
	"github.com/MarianoLibre/desafio-goweb-marianomacri/internal/domain"
	"github.com/MarianoLibre/desafio-goweb-marianomacri/internal/tickets"
	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	db, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	repo := tickets.NewRepository(db)
	service := tickets.NewService(repo)
	h := handler.NewService(service)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// Rutas a desarollar:

	ts := r.Group("/tickets")
	{
		ts.GET("/all", h.GetAll())
		ts.GET("/bycountry/:dest", h.GetTicketsByCountry())
		ts.GET("/mostvisited", h.MostVisited())
	}
	// GET - “/ticket/getByCountry/:dest”
	// GET - “/ticket/getAverage/:dest”
	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}

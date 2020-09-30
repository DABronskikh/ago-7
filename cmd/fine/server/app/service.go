package app

import (
	"context"
	"errors"
	fineV1Pb "github.com/DABronskikh/ago-7/pkg/fine/v1"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	pool      *pgxpool.Pool
	poolServ2 *pgxpool.Pool
	poolServ3 *pgxpool.Pool
}

var (
	ErrDB = errors.New("error db")
)

func NewService(pool *pgxpool.Pool, poolServ2 *pgxpool.Pool, poolServ3 *pgxpool.Pool) *Service {
	return &Service{pool: pool, poolServ2: poolServ2, poolServ3: poolServ3}
}

// поиск в сервите №1
func (s *Service) GetFlightsServ1(ctx context.Context, data *fineV1Pb.SearchData) (*fineV1Pb.FlightResponse, error) {
	var flightResponseDB []*fineV1Pb.Flight
	rows, err := s.pool.Query(ctx, `
		SELECT id, cost, fromiata, toiata 
		FROM Flights
		WHERE fromiata = $1 AND toiata = $2
		LIMIT 50
	`, data.FromIATA, data.ToIATA)
	defer rows.Close()

	for rows.Next() {
		flightEl := &fineV1Pb.Flight{}
		err = rows.Scan(&flightEl.Id, &flightEl.Cost, &flightEl.FromIATA, &flightEl.ToIATA)
		if err != nil {
			return nil, ErrDB
		}
		flightResponseDB = append(flightResponseDB, flightEl)
	}

	if err != nil {
		return nil, ErrDB
	}

	return &fineV1Pb.FlightResponse{
		Items: flightResponseDB,
	}, nil
}

// поиск в сервите №2
func (s *Service) GetFlightsServ2(ctx context.Context, data *fineV1Pb.SearchData) (*fineV1Pb.FlightResponse, error) {
	var flightResponseDB []*fineV1Pb.Flight
	rows, err := s.poolServ2.Query(ctx, `
		SELECT id, cost, fromiata, toiata 
		FROM Flights
		WHERE fromiata = $1 AND toiata = $2
		LIMIT 50
	`, data.FromIATA, data.ToIATA)
	defer rows.Close()

	for rows.Next() {
		flightEl := &fineV1Pb.Flight{}
		err = rows.Scan(&flightEl.Id, &flightEl.Cost, &flightEl.FromIATA, &flightEl.ToIATA)
		if err != nil {
			return nil, ErrDB
		}
		flightResponseDB = append(flightResponseDB, flightEl)
	}

	if err != nil {
		return nil, ErrDB
	}

	return &fineV1Pb.FlightResponse{
		Items: flightResponseDB,
	}, nil
}

// поиск в сервите №3
func (s *Service) GetFlightsServ3(ctx context.Context, data *fineV1Pb.SearchData) (*fineV1Pb.FlightResponse, error) {
	var flightResponseDB []*fineV1Pb.Flight
	rows, err := s.poolServ3.Query(ctx, `
		SELECT id, cost, fromiata, toiata 
		FROM Flights
		WHERE fromiata = $1 AND toiata = $2
		LIMIT 50
	`, data.FromIATA, data.ToIATA)
	defer rows.Close()

	for rows.Next() {
		flightEl := &fineV1Pb.Flight{}
		err = rows.Scan(&flightEl.Id, &flightEl.Cost, &flightEl.FromIATA, &flightEl.ToIATA)
		if err != nil {
			return nil, ErrDB
		}
		flightResponseDB = append(flightResponseDB, flightEl)
	}

	if err != nil {
		return nil, ErrDB
	}

	return &fineV1Pb.FlightResponse{
		Items: flightResponseDB,
	}, nil
}

package app

import (
	"context"
	fineV1Pb "github.com/DABronskikh/ago-7/pkg/fine/v1"
	"log"
	"sync"
)

type Server struct {
	businessSvc *Service
}

func NewServer(businessSvc *Service) *Server {
	return &Server{businessSvc}
}

func (s *Server) SearchFlights(
	request *fineV1Pb.SearchData,
	server fineV1Pb.AggregatorService_SearchFlightsServer,
) error {
	log.Println(request)

	countServices := 3
	wg := sync.WaitGroup{}
	wg.Add(countServices)

	go func() {
		log.Println("go serv - 1")
		arrItems, err := s.businessSvc.GetFlightsServ1(context.Background(), request)
		if err != nil {
			log.Println(err)
			return
		}
		if err := server.Send(arrItems); err != nil {
			log.Println(err)
			return
		}
		wg.Done()
	}()

	go func() {
		log.Println("go serv - 2")
		arrItems, err := s.businessSvc.GetFlightsServ2(context.Background(), request)
		if err != nil {
			log.Println(err)
			return
		}
		if err := server.Send(arrItems); err != nil {
			log.Println(err)
			return
		}
		wg.Done()
	}()

	go func() {
		log.Println("go serv - 3")
		arrItems, err := s.businessSvc.GetFlightsServ3(context.Background(), request)
		if err != nil {
			log.Println(err)
			return
		}
		if err := server.Send(arrItems); err != nil {
			log.Println(err)
			return
		}
		wg.Done()
	}()

	wg.Wait()

	return nil
}

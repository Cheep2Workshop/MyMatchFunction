package mmf

import (
	"log"

	"open-match.dev/open-match/pkg/matchfunction"
	"open-match.dev/open-match/pkg/pb"
)

const (
	matchName = "basic-matchfunction"
)

func (s *MatchFunctionService) Run(req *pb.RunRequest, stream pb.MatchFunction_RunServer) error {
	log.Printf("Generating proposals for function %v", req.GetProfile().GetName())

	poolTickets, err := matchfunction.QueryPools(stream.Context(), s.queryServiceClient, req.GetProfile().GetPools())
	if err != nil {
		log.Printf("Failed to query tickets for the give pools, got %s", err.Error())
		return err
	}

	proposals, err := makeMatches(req.GetProfile(), poolTickets)
	if err != nil {
		log.Printf("Failed to generate matches, got %s", err.Error())
		return err
	}

	log.Printf("Streaming %v proposals to Open Match", len(proposals))
	for _, proposal := range proposals {
		if err := stream.Send(&pb.RunResponse{Proposal: proposal}); err != nil {
			log.Printf("Failed to stream proposals to Open Match, got %s", err.Error())
			return err
		}
	}

	return nil
}

func makeMatches(p *pb.MatchProfile, poolTickets map[string][]*pb.Ticket) ([]*pb.Match, error) {
	return nil, nil
}

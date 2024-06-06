package service

import (
	"context"
	"log"
	pb "vote/genproto"
	"vote/storage/postgres"
)

type PartyService struct {
	stg postgres.Storage
	// stg storage.StorageI
	pb.UnimplementedPartyServiceServer
}

func NewPartyService(stg *postgres.Storage) *PartyService {
	return &PartyService{stg: *stg}
}

func (ps *PartyService) CreateParty(ctx context.Context, party *pb.CreatePartyRequest) (*pb.VoidPartyResponse, error) {
	err := ps.stg.Party().Create(party)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (ps *PartyService) GetPartyById(ctx context.Context, party *pb.ByIdPartyRequest) (*pb.GetPartyResponse, error) {
	res, err := ps.stg.Party().GetById((*pb.ByIdPartyRequest)(party))
	if err != nil {
		log.Fatal(err.Error())
	}

	return res, err
}

func (ps *PartyService) GetAllParties(ctx context.Context, party *pb.FilterPartyRequest) (*pb.GetAllPartyResponse, error) {
	res, err := ps.stg.Party().GetAll((*pb.FilterPartyRequest)(party))
	if err != nil {
		return nil, err
	}
	return res, err
}

func (ps *PartyService) UpdateParty(ctx context.Context, party *pb.GetPartyResponse) (*pb.VoidPartyResponse, error) {
	_, err := ps.stg.Party().Update(party)
	if err != nil {
		return nil, err
	}

	return &pb.VoidPartyResponse{}, err
}

func (ps *PartyService) DeleteParty(ctx context.Context, party *pb.ByIdPartyRequest) (*pb.VoidPartyResponse, error) {
	_, err := ps.stg.Party().Delete(party)
	if err != nil {
		return nil, err
	}

	return &pb.VoidPartyResponse{}, err
}

package service

import (
	"context"
	"log"
	pb "vote/genproto/genproto/voting"
	"vote/storage/postgres"
)

type PublicService struct {
	stg postgres.Storage
	// stg storage.StorageI
	pb.UnimplementedPublicServiceServer
}

func NewPublicService(stg *postgres.Storage) *PublicService {
	return &PublicService{stg: *stg}
}

func (ps *PublicService) CreatePublic(ctx context.Context, public *pb.CreatePublicRequest) (*pb.VoidPublicResponse, error) {
	res, err := ps.stg.Public().Create(public)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (ps *PublicService) GetByIdPublic(ctx context.Context, public *pb.ByIdPublicRequest) (*pb.GetPublicResponse, error) {
	res, err := ps.stg.Public().GetById((*pb.ByIdPublicRequest)(public))
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return res, nil
}

func (ps *PublicService) GetAllPublic(ctx context.Context, public *pb.FilterPublicRequest) (*pb.GetAllPublicResponse, error) {
	res, err := ps.stg.Public().GetAll((*pb.FilterPublicRequest)(public))
	if err != nil {
		return nil, err
	}
	return res, err
}

func (ps *PublicService) UpdatePublic(ctx context.Context, public *pb.GetPublicResponse) (*pb.VoidPublicResponse, error) {
	_, err := ps.stg.Public().Update(public)
	if err != nil {
		return nil, err
	}

	return &pb.VoidPublicResponse{}, err
}

func (ps *PublicService) DeletePublic(ctx context.Context, public *pb.ByIdPublicRequest) (*pb.VoidPublicResponse, error) {
	_, err := ps.stg.Public().Delete(public)
	if err != nil {
		return nil, err
	}

	return &pb.VoidPublicResponse{}, err
}

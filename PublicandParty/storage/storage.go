package storage

import (
	pb "vote/genproto/genproto/voting"
)

type PublicI interface {
	Create(public *pb.CreatePublicRequest) (*pb.VoidPublicResponse, error)
	GetById(public *pb.ByIdPublicRequest) (*pb.GetPublicResponse, error)
	GetAll(public *pb.FilterPublicRequest) (*pb.GetAllPublicResponse, error)
	Update(public *pb.GetPublicResponse) (*pb.VoidPublicResponse, error)
	Delete(public *pb.ByIdPublicRequest) (*pb.VoidPublicResponse, error)
}

type PartyI interface {
	Create(party *pb.CreatePartyRequest) (error)
	GetById(party *pb.ByIdPartyRequest) (*pb.GetPartyResponse, error)
	GetAll(party *pb.FilterPartyRequest) (*pb.GetAllPartyResponse, error)
	Update(party *pb.GetPartyResponse) (*pb.VoidPartyResponse, error)
	Delete(party *pb.ByIdPartyRequest) (*pb.VoidPartyResponse, error)
}

package grpc

import (
	"context"

	"github.com/feynmaz/hex/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetAddition(
		req.GetA(),
		req.GetB(),
	)
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (a Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetSubtraction(
		req.GetA(),
		req.GetB(),
	)
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (a Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetMultiplication(
		req.GetA(),
		req.GetB(),
	)
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (a Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetDivision(
		req.GetA(),
		req.GetB(),
	)
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

package api

import (
	"strings"

	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	KV keyvalue.KeyValue `inj:""`

	logger *logrus.Entry
}

func NewEmailMappingServer(logger *logrus.Entry) protocolbuffer.EmailMappingServer {
	return &server{
		logger: logger.WithField("app", "core/api"),
	}
}

func (s *server) Register(ctx context.Context, data *protocolbuffer.RegisterRequest) (*protocolbuffer.RegisterReply, error) {
	// Getting value first and setting right after might make two threads go through. It is fine in this case, second one will overwrite and it's not a big deal.
	current, err := s.KV.GetString(data.GetWallet())
	if err != nil {
		s.logger.WithError(err).Error("Error checking current data")
		return &protocolbuffer.RegisterReply{}, status.Error(codes.Internal, "Error checking current data")
	}
	if current != "" {
		return &protocolbuffer.RegisterReply{}, status.Error(codes.FailedPrecondition, "Mapping already exists. Unregister first")
	}

	err = s.KV.PutString(strings.ToLower(data.GetWallet()), data.GetEmail())
	if err != nil {
		s.logger.WithError(err).Error("Error storing data")
		return &protocolbuffer.RegisterReply{}, status.Error(codes.Internal, "Error storing data")
	}
	return &protocolbuffer.RegisterReply{}, nil
}
func (s *server) Unregister(ctx context.Context, data *protocolbuffer.UnregisterRequest) (*protocolbuffer.UnregisterReply, error) {
	current, err := s.KV.GetString(data.GetWallet())
	if err != nil {
		s.logger.WithError(err).Error("Error checking current data")
		return &protocolbuffer.UnregisterReply{}, status.Error(codes.Internal, "Error checking current data")
	}
	if current == "" {
		return &protocolbuffer.UnregisterReply{}, status.Error(codes.FailedPrecondition, "There's no data registered to this wallet")
	}

	err = s.KV.Delete(data.GetWallet())
	if err != nil {
		s.logger.WithError(err).Error("Error deleting data")
		return &protocolbuffer.UnregisterReply{}, status.Error(codes.Internal, "Error deleting data")
	}
	return &protocolbuffer.UnregisterReply{}, nil
}

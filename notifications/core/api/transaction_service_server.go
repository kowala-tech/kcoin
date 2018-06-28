package api

import (
	"context"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/notifications/persistence"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type transactionServiceServer struct {
	Persistence persistence.TransactionRepository `inj:""`

	logger *logrus.Entry
}

func NewTransactionServiceServer(logger *logrus.Entry) protocolbuffer.TransactionServiceServer {
	return &transactionServiceServer{
		logger: logger.WithField("app", "core/api"),
	}
}

func (s *transactionServiceServer) GetTransactions(ctx context.Context, data *protocolbuffer.GetTransactionsRequest) (*protocolbuffer.GetTransactionsReply, error) {
	account := common.HexToAddress(data.GetAccount())

	txs, err := s.Persistence.GetTxsFromAccount(account)
	if err != nil {
		s.logger.WithError(err).Error(codes.Internal, "Error getting transactions")
		return &protocolbuffer.GetTransactionsReply{}, nil
	}

	return &protocolbuffer.GetTransactionsReply{
		Transactions: txs,
	}, nil
}

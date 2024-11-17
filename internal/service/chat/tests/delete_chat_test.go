package chat

import (
	"context"
	"testing"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	"github.com/Paul1k96/microservices_course_chat_service/internal/repository/mocks"
	"github.com/Paul1k96/microservices_course_chat_service/internal/service"
	"github.com/Paul1k96/microservices_course_chat_service/internal/service/chat"
	"github.com/Paul1k96/microservices_course_platform_common/pkg/client/db/transaction"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}

type DeleteSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller

	userRepo    *mocks.MockUserRepository
	chatRepo    *mocks.MockChatRepository
	messageRepo *mocks.MockMessageRepository

	service service.ChatService
}

func (t *DeleteSuite) SetupTest() {
	t.Assertions = require.New(t.T())
	t.ctrl = gomock.NewController(t.T())

	t.userRepo = mocks.NewMockUserRepository(t.ctrl)
	t.chatRepo = mocks.NewMockChatRepository(t.ctrl)
	t.messageRepo = mocks.NewMockMessageRepository(t.ctrl)

	t.service = chat.NewService(t.chatRepo, t.userRepo, t.messageRepo, transaction.NewNopTxManager())
}

func (t *DeleteSuite) TearDownTest() {
	t.ctrl.Finish()
}

type DeleteArgs struct {
	ctx context.Context
	id  model.ChatID
}

type DeleteWant struct {
	err error
}

func (t *DeleteSuite) do(args DeleteArgs, want DeleteWant) {
	err := t.service.Delete(args.ctx, args.id)

	if want.err == nil {
		t.Require().NoError(err)
	} else {
		t.Require().ErrorContains(err, want.err.Error())
	}
}

func (t *DeleteSuite) TestDelete_Ok() {
	args := DeleteArgs{
		ctx: context.Background(),
		id:  model.ChatID(gofakeit.Int64()),
	}

	want := DeleteWant{
		err: nil,
	}

	t.chatRepo.EXPECT().Delete(args.ctx, args.id).Return(nil)

	t.do(args, want)
}

func (t *DeleteSuite) TestDelete_RepoError() {
	args := DeleteArgs{
		ctx: context.Background(),
		id:  model.ChatID(gofakeit.Int64()),
	}

	want := DeleteWant{
		err: gofakeit.Error(),
	}

	t.chatRepo.EXPECT().Delete(args.ctx, args.id).Return(want.err)

	t.do(args, want)
}

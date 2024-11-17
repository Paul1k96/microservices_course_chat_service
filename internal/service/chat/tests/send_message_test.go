package chat

import (
	"context"
	"testing"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	"github.com/Paul1k96/microservices_course_chat_service/internal/repository/mocks"
	"github.com/Paul1k96/microservices_course_chat_service/internal/service"
	"github.com/Paul1k96/microservices_course_chat_service/internal/service/chat"
	tm "github.com/Paul1k96/microservices_course_chat_service/internal/testmodel"
	"github.com/Paul1k96/microservices_course_platform_common/pkg/client/db/transaction"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

func TestSendMessageSuite(t *testing.T) {
	suite.Run(t, new(SendMessageSuite))
}

type SendMessageSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller

	userRepo    *mocks.MockUserRepository
	chatRepo    *mocks.MockChatRepository
	messageRepo *mocks.MockMessageRepository

	service service.ChatService
}

func (t *SendMessageSuite) SetupTest() {
	t.Assertions = require.New(t.T())
	t.ctrl = gomock.NewController(t.T())

	t.userRepo = mocks.NewMockUserRepository(t.ctrl)
	t.chatRepo = mocks.NewMockChatRepository(t.ctrl)
	t.messageRepo = mocks.NewMockMessageRepository(t.ctrl)

	t.service = chat.NewService(t.chatRepo, t.userRepo, t.messageRepo, transaction.NewNopTxManager())
}

func (t *SendMessageSuite) TearDownTest() {
	t.ctrl.Finish()
}

type SendMessageArgs struct {
	ctx     context.Context
	message *model.Message
}

type SendMessageWant struct {
	err error
}

func (t *SendMessageSuite) do(args SendMessageArgs, want SendMessageWant) {
	err := t.service.SendMessage(args.ctx, args.message)

	if want.err == nil {
		t.Require().NoError(err)
	} else {
		t.Require().ErrorContains(err, want.err.Error())
	}
}

func (t *SendMessageSuite) TestSendMessage_Ok() {
	args := SendMessageArgs{
		ctx:     context.Background(),
		message: tm.NewMessage(),
	}

	want := SendMessageWant{
		err: nil,
	}

	t.messageRepo.EXPECT().Create(args.ctx, args.message).Return(nil)

	t.do(args, want)
}

func (t *SendMessageSuite) TestSendMessage_RepoError() {
	args := SendMessageArgs{
		ctx:     context.Background(),
		message: tm.NewMessage(),
	}

	want := SendMessageWant{
		err: gofakeit.Error(),
	}

	t.messageRepo.EXPECT().Create(args.ctx, args.message).Return(want.err)

	t.do(args, want)
}

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

func TestCreateChatSuite(t *testing.T) {
	suite.Run(t, new(CreateChatSuite))
}

type CreateChatSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller

	userRepo    *mocks.MockUserRepository
	chatRepo    *mocks.MockChatRepository
	messageRepo *mocks.MockMessageRepository

	service service.ChatService
}

func (t *CreateChatSuite) SetupTest() {
	t.Assertions = require.New(t.T())
	t.ctrl = gomock.NewController(t.T())

	t.userRepo = mocks.NewMockUserRepository(t.ctrl)
	t.chatRepo = mocks.NewMockChatRepository(t.ctrl)
	t.messageRepo = mocks.NewMockMessageRepository(t.ctrl)

	t.service = chat.NewService(t.chatRepo, t.userRepo, t.messageRepo, transaction.NewNopTxManager())
}

func (t *CreateChatSuite) TearDownTest() {
	t.ctrl.Finish()
}

type CreateChatArgs struct {
	ctx     context.Context
	userIDs model.UserIDs
}

type CreateChatWant struct {
	id  model.ChatID
	err error
}

func (t *CreateChatSuite) do(args CreateChatArgs, want CreateChatWant) {
	id, err := t.service.Create(args.ctx, args.userIDs)

	t.Require().Equal(want.id, id)

	if want.err == nil {
		t.Require().NoError(err)
	} else {
		t.Require().ErrorContains(err, want.err.Error())
	}
}

func (t *CreateChatSuite) TestCreateChat_Ok() {
	users := tm.NewUsers(2)
	newChat := tm.NewChat()

	args := CreateChatArgs{
		ctx:     context.Background(),
		userIDs: users.IDs(),
	}

	want := CreateChatWant{
		id: newChat.ID,
	}

	t.userRepo.EXPECT().List(args.ctx, args.userIDs).Return(users, nil)

	t.chatRepo.EXPECT().Create(args.ctx).Return(newChat.ID, nil)

	users.SetChatID(newChat.ID)

	t.chatRepo.EXPECT().AddUsers(args.ctx, users).Return(nil)

	t.do(args, want)
}

func (t *CreateChatSuite) TestCreateChat_GetUsersError() {
	users := tm.NewUsers(2)

	args := CreateChatArgs{
		ctx:     context.Background(),
		userIDs: users.IDs(),
	}

	want := CreateChatWant{
		err: gofakeit.Error(),
	}

	t.userRepo.EXPECT().List(args.ctx, args.userIDs).Return(nil, want.err)

	t.do(args, want)
}

func (t *CreateChatSuite) TestCreateChat_CreateChatError() {
	users := tm.NewUsers(2)

	args := CreateChatArgs{
		ctx:     context.Background(),
		userIDs: users.IDs(),
	}

	want := CreateChatWant{
		err: gofakeit.Error(),
	}

	t.userRepo.EXPECT().List(args.ctx, args.userIDs).Return(users, nil)

	t.chatRepo.EXPECT().Create(args.ctx).Return(want.id, want.err)

	t.do(args, want)
}

func (t *CreateChatSuite) TestCreateChat_AddUsersError() {
	users := tm.NewUsers(2)
	newChat := tm.NewChat()

	args := CreateChatArgs{
		ctx:     context.Background(),
		userIDs: users.IDs(),
	}

	want := CreateChatWant{
		err: gofakeit.Error(),
	}

	t.userRepo.EXPECT().List(args.ctx, args.userIDs).Return(users, nil)

	t.chatRepo.EXPECT().Create(args.ctx).Return(newChat.ID, nil)

	users.SetChatID(newChat.ID)

	t.chatRepo.EXPECT().AddUsers(args.ctx, users).Return(want.err)

	t.do(args, want)
}

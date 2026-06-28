package user_test

import (
	"context"
	"testing"

	"github.com/fajar3108/lms-backend/internal/user"
	"github.com/fajar3108/lms-backend/test/test_helper"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

const (
	dummyID           = "existing-uuid-123"
	dummyName         = "Fajar"
	dummyEmail        = "fajar@example.com"
	dummyPassword     = "password123"
	dummyRefreshToken = "dumm-token"
)

type UserTestSuite struct {
	suite.Suite
	db         *gorm.DB
	tx         *gorm.DB
	userAction *user.UserAction
}

func TestAuthActionSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetupSuite() {
	s.db = test_helper.NewTestDB(s.T())
}

func (s *UserTestSuite) SetupSubTest() {
	s.tx = s.db.Begin()
	s.userAction = user.NewUserAction(s.tx)
}

func (s *UserTestSuite) TearDownSubTest() {
	s.tx.Rollback()
}

func (s *UserTestSuite) seedDummyUser() user.User {
	user := user.User{
		ID:       dummyID,
		Name:     dummyName,
		Email:    dummyEmail,
		Password: dummyPassword,
	}
	s.Require().NoError(s.tx.Create(&user).Error)
	return user
}

func (s *UserTestSuite) TestCreateUser() {
	s.Run("Create user successfully", func() {
		createdUser, err := s.userAction.CreateUser(context.Background(), dummyName, dummyEmail, dummyPassword)

		s.Require().NoError(err)
		s.Require().NotNil(createdUser)
		s.NotEmpty(createdUser.ID)
		s.Equal(dummyName, createdUser.Name)
		s.Equal(dummyEmail, createdUser.Email)

		var dbUser user.User
		s.Require().NoError(s.tx.First(&dbUser, "id = ?", createdUser.ID).Error)
		s.Equal(createdUser.Email, dbUser.Email)
	})

	s.Run("Failed when user already registered", func() {
		s.seedDummyUser()

		createdUser, err := s.userAction.CreateUser(context.Background(), "Another Name", dummyEmail, dummyPassword)

		s.Require().Error(err)
		s.ErrorIs(err, user.ErrEmailAlreadyExist)
		s.Nil(createdUser)
	})
}

func (s *UserTestSuite) TestGetUserByEmail() {
	s.Run("Get user by email successfully", func() {
		s.seedDummyUser()

		user, err := s.userAction.GetUserByEmail(context.Background(), dummyEmail)

		s.Require().NoError(err)
		s.Require().NotNil(user)
		s.Equal(dummyEmail, user.Email)
	})

	s.Run("Failed when user not found", func() {
		foundUser, err := s.userAction.GetUserByEmail(context.Background(), "notfound@example.com")

		s.Require().Error(err)
		s.ErrorIs(err, user.ErrEmailNotFound)
		s.Nil(foundUser)
	})
}

func (s *UserTestSuite) TestGetUserByID() {
	s.Run("Get user by ID successfully", func() {
		s.seedDummyUser()

		user, err := s.userAction.GetUserByID(context.Background(), dummyID)

		s.Require().NoError(err)
		s.Require().NotNil(user)
		s.Equal(dummyID, user.ID)
	})

	s.Run("Failed when user not found", func() {
		foundUser, err := s.userAction.GetUserByID(context.Background(), "notfound-uuid-123")

		s.Require().Error(err)
		s.ErrorIs(err, user.ErrUserNotFound)
		s.Nil(foundUser)
	})
}

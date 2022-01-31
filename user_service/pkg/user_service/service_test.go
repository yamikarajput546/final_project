package userservice

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"

	"github.com/jcromanu/final_project/user_service/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreateUser(t *testing.T) {
	repoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName       string
		input          entities.User
		expectedOutput int32
		expectedError  error
	}{
		{
			testName:       "create user with all fields",
			input:          entities.User{Name: "Juan", Age: 30, Additional_information: "additional info", Parent: []string{"parent sample"}},
			expectedOutput: 1,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			repoMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			service := NewService(repoMock, logger)
			usr, err := service.CreateUser(ctx, tc.input)
			assert.Equal(t, tc.expectedOutput, usr.Id)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestServiceGetUser(t *testing.T) {
	repoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName       string
		input          int32
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:       "get user valid id",
			input:          1,
			expectedOutput: entities.User{Id: 1},
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			repoMock.On("GetUser", mock.Anything, mock.Anything).Return(tc.input, tc.expectedError)
			service := NewService(repoMock, logger)
			usr, err := service.GetUser(ctx, tc.input)
			assert.Equal(t, tc.expectedOutput, usr)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

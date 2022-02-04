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

func TestMakeCreateUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	serviceMock := new(ServiceMock)
	testCases := []struct {
		testName       string
		input          createUserRequest
		expectedOutput int32
		expectedError  error
	}{
		{
			testName:       "create endpoint user with all fields success ",
			input:          createUserRequest{User: entities.User{Name: "Juan", Age: 30, Additional_information: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: 1,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			ep := makeCreateUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, tc.input)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			re, ok := result.(createUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(t, tc.expectedOutput, re.User.Id, "Error on user response")
			assert.Equal(t, tc.expectedError, err, "Error on user response")
		})
	}
}

func TestMakeGetUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	serviceMock := new(ServiceMock)
	testCases := []struct {
		testName       string
		input          getUserRequest
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:       "get  endpoint user with all fields success ",
			input:          getUserRequest{Id: 1},
			expectedOutput: entities.User{Id: 1},
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock.On("GetUser", mock.Anything, mock.Anything).Return(tc.input.Id, tc.expectedError)
			ep := makeGetUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, tc.input)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			re, ok := result.(getUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(t, tc.expectedOutput, re.User, "Error on user response")
			assert.Equal(t, tc.expectedError, err, "Error on user response")
		})
	}
}

func TestMakeUpdateUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	serviceMock := new(ServiceMock)
	testCases := []struct {
		testName       string
		input          updateUserRequest
		expectedOutput error
		expectedError  error
	}{
		{
			testName:       "update endpoint user with all fields success ",
			input:          updateUserRequest{entities.User{Id: 1, Name: "Juan", Age: 30, Pwd_hash: "hash ", Additional_information: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: nil,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock.On("UpdateUser", mock.Anything, mock.Anything).Return(tc.expectedError)
			ep := makeUpdatesUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, tc.input)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			_, ok := result.(updateUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(t, tc.expectedError, err, "Error on user response")
		})
	}
}

func TestMakeDeleteUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	serviceMock := new(ServiceMock)
	testCases := []struct {
		testName      string
		input         deleteUserRequest
		expectedError error
	}{
		{
			testName:      "delete user on request ",
			input:         deleteUserRequest{id: 1},
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(tc.expectedError)
			ep := makeDeleteUserEndpoint(serviceMock, logger)
			res, err := ep(ctx, tc.input)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			_, ok := res.(deleteUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(t, tc.expectedError, err, "Error on user response")
		})
	}

}

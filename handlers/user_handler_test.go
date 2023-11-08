package handlers

import (
	"encoding/xml"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tmbrody/taskquire/internal/database"
)

// TestCreateUserHandler_Success tests the CreateUserHandler function when it succeeds.
func TestCreateUserHandler_Success(t *testing.T) {
	// Create a mock database transaction.
	mockDBTX := new(MockDBTX)

	// Create a Gin router instance.
	r := gin.Default()

	// Use a custom middleware for setting up a test database (mockDBTX).
	r.Use(SetupTestDatabaseMiddleware(mockDBTX))

	// Register a POST route for the CreateUserHandler.
	r.POST("/api/users", CreateUserHandler)

	// Configure the mockDBTX to return a successful result when ExecContext is called.
	mockDBTX.On("ExecContext", mock.Anything, mock.Anything, mock.Anything).Return(sqlmock.NewResult(1, 1), nil)

	// Create a sample XML payload for user creation.
	payload := `
        <CreateUserParams>
            <Name>John Smith</Name>
            <Email>test@example.com</Email>
            <Password>passwordpassword</Password>
        </CreateUserParams>
    `

	// Create an HTTP request with the payload.
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/xml")
	resp := httptest.NewRecorder()

	// Serve the HTTP request using the Gin router.
	r.ServeHTTP(resp, req)

	// Unmarshal the XML response into a user object for validation.
	var user database.User
	err := xml.Unmarshal(resp.Body.Bytes(), &user)
	if err != nil {
		t.Errorf("Error unmarshalling XML response: %v", err)
	}

	// Assert that the HTTP response code is StatusCreated (201).
	assert.Equal(t, http.StatusCreated, resp.Code)

	// Validate the user object's properties.
	assert.Equal(t, "John Smith", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "", user.Password)
	assert.Equal(t, user.CreatedAt.Truncate(time.Second), user.UpdatedAt.Truncate(time.Second))

	// Ensure that the expected methods on the mockDBTX were called as expected.
	mockDBTX.AssertExpectations(t)
}

// TestCreateUserHandler_Failure tests the CreateUserHandler function when it fails due to a database error.
func TestCreateUserHandler_Failure(t *testing.T) {
	// Create a mock database transaction.
	mockDBTX := new(MockDBTX)

	// Create a Gin router instance.
	r := gin.Default()

	// Use a custom middleware for setting up a test database (mockDBTX).
	r.Use(SetupTestDatabaseMiddleware(mockDBTX))

	// Register a POST route for the CreateUserHandler.
	r.POST("/api/users", CreateUserHandler)

	// Define an expected database error and mock result.
	expectedError := errors.New("Database error")
	mockResult := sqlmock.NewResult(1, 1)

	// Configure the mockDBTX to return the expected error when ExecContext is called.
	mockDBTX.On("ExecContext", mock.Anything, mock.Anything, mock.Anything).Return(mockResult, expectedError)

	// Create a sample XML payload for user creation.
	payload := `
        <CreateUserParams>
            <Name>John Smith</Name>
            <Email>test@example.com</Email>
            <Password>passwordpassword</Password>
        </CreateUserParams>
    `

	// Create an HTTP request with the payload.
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/xml")
	resp := httptest.NewRecorder()

	// Serve the HTTP request using the Gin router.
	r.ServeHTTP(resp, req)

	// Assert that the HTTP response code is InternalServerError (500).
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	// Unmarshal the XML response into an error response object for validation.
	var errorResponse ErrorResponse
	err := xml.Unmarshal(resp.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Errorf("Error unmarshalling XML response: %v", err)
	}

	// Validate the error message in the response.
	expectedErrorMessage := "Unable to create new user"
	assert.Equal(t, expectedErrorMessage, errorResponse.Message)

	// Ensure that the expected methods on the mockDBTX were called as expected.
	mockDBTX.AssertExpectations(t)
}

// TestCreateUserHandler_InvalidPayload tests the CreateUserHandler function with an invalid XML payload.
func TestCreateUserHandler_InvalidPayload(t *testing.T) {
	// Create a mock database transaction.
	mockDBTX := new(MockDBTX)

	// Create a Gin router instance.
	r := gin.Default()

	// Use a custom middleware for setting up a test database (mockDBTX).
	r.Use(SetupTestDatabaseMiddleware(mockDBTX))

	// Register a POST route for the CreateUserHandler.
	r.POST("/api/users", CreateUserHandler)

	// Create an invalid XML payload with missing fields.
	payload := `
        <InvalidCreateUserParams>
            <Name>John Smith</Name>
        </InvalidCreateUserParams>
    `

	// Create an HTTP request with the invalid payload.
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/xml")
	resp := httptest.NewRecorder()

	// Serve the HTTP request using the Gin router.
	r.ServeHTTP(resp, req)

	// Assert that the HTTP response code is BadRequest (400).
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Unmarshal the XML response into an error response object for validation.
	var errorResponse ErrorResponse
	err := xml.Unmarshal(resp.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Errorf("Error unmarshalling XML response: %v", err)
	}

	// Validate the error message in the response.
	expectedErrorMessage := "Invalid request payload: Missing or invalid fields."
	assert.Equal(t, expectedErrorMessage, errorResponse.Message)

	// Ensure that the ExecContext method on the mockDBTX was not called.
	mockDBTX.AssertNotCalled(t, "ExecContext")
}

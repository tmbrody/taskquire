package handlers

import (
	"context"
	"database/sql"
	"encoding/xml"

	"github.com/stretchr/testify/mock"
)

type ErrorResponse struct {
	XMLName xml.Name `xml:"Error"`
	Message string   `xml:"Message"`
}

type MockDBTX struct {
	mock.Mock
}

func (m *MockDBTX) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	argsPassed := m.Called(ctx, query, args)
	return argsPassed.Get(0).(sql.Result), argsPassed.Error(1)
}

func (m *MockDBTX) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	argsPassed := m.Called(ctx, query)
	return argsPassed.Get(0).(*sql.Stmt), argsPassed.Error(1)
}

func (m *MockDBTX) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	argsPassed := m.Called(ctx, query, args)
	return argsPassed.Get(0).(*sql.Rows), argsPassed.Error(1)
}

func (m *MockDBTX) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	argsPassed := m.Called(ctx, query, args)
	return argsPassed.Get(0).(*sql.Row)
}

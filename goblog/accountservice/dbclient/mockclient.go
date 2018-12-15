package dbclient

import (
	"github.com/cloudenterprise/goblog/accountservice/model"
	"github.com/stretchr/testify/mock"
)

// MockBoltClient is a mock implementation of our BoltDB for testing purposes.
type MockBoltClient struct {
	mock.Mock
}

// QueryAccount uses our MockBoltClient to query an account
func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

// OpenBoltDb is required to satisfy the interface but not implemented
func (m *MockBoltClient) OpenBoltDb() {
	// Does nothing
}

// Seed is required to satisfy the interface but not implemented
func (m *MockBoltClient) Seed() {
	// Does nothing
}

// Check is healthcheck to make sure our db is initialized
func (m *MockBoltClient) Check() bool {
	args := m.Mock.Called()
	return args.Get(0).(bool)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\Admin\go\src\url-shortener\internal\app\utility\cache\redis.go

// Package cache is a generated GoMock package.
package cache

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIRedisClient is a mock of IRedisClient interface.
type MockIRedisClient struct {
	ctrl     *gomock.Controller
	recorder *MockIRedisClientMockRecorder
}

// MockIRedisClientMockRecorder is the mock recorder for MockIRedisClient.
type MockIRedisClientMockRecorder struct {
	mock *MockIRedisClient
}

// NewMockIRedisClient creates a new mock instance.
func NewMockIRedisClient(ctrl *gomock.Controller) *MockIRedisClient {
	mock := &MockIRedisClient{ctrl: ctrl}
	mock.recorder = &MockIRedisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedisClient) EXPECT() *MockIRedisClientMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockIRedisClient) Exists(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockIRedisClientMockRecorder) Exists(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIRedisClient)(nil).Exists), key)
}

// Get mocks base method.
func (m *MockIRedisClient) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIRedisClientMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIRedisClient)(nil).Get), key)
}

// Set mocks base method.
func (m *MockIRedisClient) Set(key string, value interface{}, ttl time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, ttl)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockIRedisClientMockRecorder) Set(key, value, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockIRedisClient)(nil).Set), key, value, ttl)
}
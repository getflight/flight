package mocks

import "github.com/stretchr/testify/mock"

type WriterMock struct {
	mock.Mock
}

func (m *WriterMock) Write(p []byte) (n int, err error) {
	args := m.Called(p)

	return args.Int(0), args.Error(1)
}

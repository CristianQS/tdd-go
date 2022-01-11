package mocks

type MockLogger struct {
	message string
}

func (m *MockLogger) GetMessage() string {
	return m.message
}

func (m *MockLogger) Log(message string) {
	m.message = message
}

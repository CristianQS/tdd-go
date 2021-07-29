package pkg

type MockLogger struct {
	message string
}

func (m MockLogger) GetMessage() string {
	return m.message
}

func (m *MockLogger) log(message string) {
	m.message = message
}


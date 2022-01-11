package mocks

type MockSensor struct {
	pressure float32
}

func (s *MockSensor) SetPressure(pressure float32) {
	s.pressure = pressure
}

func (s MockSensor) PopNextPressurePsiValue() float32 {
	return s.pressure
}

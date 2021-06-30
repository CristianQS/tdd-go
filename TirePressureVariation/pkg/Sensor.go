package pkg

import "math/rand"

type Sensor struct {
}

var OFFSET float32 = 16

func (s Sensor) PopNextPressurePsiValue() float32 {
	var (
		pressureTelemetryValue float32
	)
	pressureTelemetryValue = s.samplePressure()

	return OFFSET + pressureTelemetryValue
}

func (s Sensor) samplePressure() float32 {
	var basicRandomNumbersGenerator = rand.Rand{}
	var pressureTelemetryValue = 6 * basicRandomNumbersGenerator.Float32() * basicRandomNumbersGenerator.Float32();
	return pressureTelemetryValue
}
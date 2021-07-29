package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//http://cs-guy.com/blog/2015/01/test-main/

func Test_should_alarm_be_activated_when_sensor_is_lower_than_threshold(t *testing.T) {
	sensor := MockSensor{}
	logger := MockLogger{}
	sensor.SetPressure(16)
	alarm := Alarm{sensor: sensor, logger: &logger}

	alarm.check()

	assert.Equal(t, "Alarm activated!",logger.GetMessage())
}

func Test_should_alarm_be_activated_when_sensor_is_high_than_threshold(t *testing.T) {
	sensor := MockSensor{}
	logger := MockLogger{}
	sensor.SetPressure(22)
	alarm := Alarm{sensor: sensor, logger: &logger}

	alarm.check()

	assert.Equal(t, "Alarm activated!",logger.GetMessage())
}

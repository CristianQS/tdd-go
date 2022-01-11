package pkg

import (
	"TirePressureVariation/pkg/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

//http://cs-guy.com/blog/2015/01/test-main/

func Test_should_alarm_be_activated_when_sensor_is_lower_than_threshold(t *testing.T) {
	sensor := &mocks.MockSensor{}
	logger := &mocks.MockLogger{}
	sensor.SetPressure(16)
	alarm := NewAlarm(sensor, logger)

	alarm.check()

	assert.Equal(t, "Alarm activated!", logger.GetMessage())
}

// https://blog.codecentric.de/en/2017/08/gomock-tutorial/
func Test_should_alarm_be_activated_when_sensor_is_lower_than_threshold_gomock(t *testing.T) {
	ctrl := gomock.NewController(t)
	sensor := mocks.NewMockISensor(ctrl)
	logger := mocks.NewMockILogger(ctrl)
	alarm := NewAlarm(sensor, logger)
	sensor.EXPECT().PopNextPressurePsiValue().Return(float32(16))
	logger.EXPECT().Log("Alarm activated!").Times(1)

	alarm.check()
}

func Test_should_alarm_be_activated_when_sensor_is_high_than_threshold(t *testing.T) {
	sensor := &mocks.MockSensor{}
	logger := &mocks.MockLogger{}
	sensor.SetPressure(22)
	alarm := NewAlarm(sensor, logger)

	alarm.check()

	assert.Equal(t, "Alarm activated!", logger.GetMessage())
}

func Test_should_alarm_be_not_activated_when_sensor_is_between_threshold(t *testing.T) {
	sensor := &mocks.MockSensor{}
	logger := &mocks.MockLogger{}
	sensor.SetPressure(20)
	alarm := NewAlarm(sensor, logger)

	alarm.check()

	assert.Equal(t, "", logger.GetMessage())
}

func Test_should_alarm_be_deactivated_when_alarm_has_been_activated(t *testing.T) {
	sensor := &mocks.MockSensor{}
	logger := &mocks.MockLogger{}
	sensor.SetPressure(16)
	alarm := NewAlarm(sensor, logger)

	alarm.check()

	sensor.SetPressure(20)
	alarm.check()

	assert.Equal(t, "Alarm deactivated!", logger.GetMessage())
}

func Test_should_alarm_be_activated_when_sensor_keeps_out_threshold(t *testing.T) {
	sensor := &mocks.MockSensor{}
	logger := &mocks.MockLogger{}
	sensor.SetPressure(16)
	alarm := NewAlarm(sensor, logger)

	alarm.check()
	alarm.check()

	assert.Equal(t, "Alarm activated!", logger.GetMessage())
}

package pkg

import "TirePressureVariation/pkg/mocks"

type Alarm struct {
	sensor  mocks.ISensor
	logger  mocks.ILogger
	alarmOn bool
}

func NewAlarm(sensor mocks.ISensor, logger mocks.ILogger) *Alarm {
	return &Alarm{sensor: sensor, logger: logger, alarmOn: false}
}

const LowPressureThreshold float32 = 17
const HighPressureThreshold float32 = 21

func (a *Alarm) check() {
	var psiPressureValue = a.sensor.PopNextPressurePsiValue()

	if psiPressureValue < LowPressureThreshold || HighPressureThreshold < psiPressureValue {
		if !a.IsAlarmOn() {
			a.alarmOn = true
			a.logger.Log("Alarm activated!")
		}
	} else {
		if a.IsAlarmOn() {
			a.alarmOn = false
			a.logger.Log("Alarm deactivated!")
		}
	}
}

func (a Alarm) IsAlarmOn() bool {
	return a.alarmOn
}

package pkg

type Alarm struct {
	sensor ISensor
	logger Logger
	alarmOn bool
}

func NewAlarm(sensor ISensor, logger Logger) *Alarm {
	return &Alarm{sensor: sensor, logger: logger, alarmOn: false}
}

const LowPressureThreshold float32 = 17
const HighPressureThreshold float32 = 21


func (a Alarm) check() {
	var psiPressureValue = a.sensor.PopNextPressurePsiValue()

	if psiPressureValue < LowPressureThreshold || HighPressureThreshold < psiPressureValue {
		if !a.IsAlarmOn() {
			a.alarmOn = true
			a.logger.log("Alarm activated!")
		}
	} else {
		if a.IsAlarmOn() {
			a.alarmOn = false
			a.logger.log("Alarm deactivated!")
		}
	}
}

func (a Alarm) IsAlarmOn() bool {
	return a.alarmOn
}

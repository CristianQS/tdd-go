package pkg

type Alarm struct {
	sensor ISensor
	logger Logger
}
var (
	LowPressureThreshold float32 = 17
	HighPressureThreshold float32 = 21
	alarmOn = false
)

func (a Alarm) check() {
	var psiPressureValue = a.sensor.PopNextPressurePsiValue()

	if psiPressureValue < LowPressureThreshold || HighPressureThreshold < psiPressureValue {
		if !a.isAlarmOn() {
			alarmOn = true
			a.logger.log("Alarm activated!")
		}
	} else {
		if a.isAlarmOn() {
			alarmOn = false
			a.logger.log("Alarm deactivated!")
		}
	}
}

func (a Alarm) isAlarmOn() bool {
	return false
}

package pkg

import "fmt"

type Alarm struct {
	sensor Sensor
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
			fmt.Println("Alarm activated!")
		}
	} else {
		if a.isAlarmOn() {
			alarmOn = false
			fmt.Println("Alarm deactivated!")
		}
	}
}

func (a Alarm) isAlarmOn() bool {
	return false
}

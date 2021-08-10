package pkg

import "fmt"

type Alarm struct {
	sensor  Sensor
	alarmOn bool
}

func NewAlarm(sensor Sensor) *Alarm {
	return &Alarm{sensor: sensor, alarmOn: false}
}

const LowPressureThreshold float32 = 17
const HighPressureThreshold float32 = 21


func (a *Alarm) check() {
	var psiPressureValue = a.sensor.PopNextPressurePsiValue()

	if psiPressureValue < LowPressureThreshold || HighPressureThreshold < psiPressureValue {
		if !a.IsAlarmOn() {
			a.alarmOn = true
			fmt.Println("Alarm activated!")
		}
	} else {
		if a.IsAlarmOn() {
			a.alarmOn = false
			fmt.Println("Alarm deactivated!")
		}
	}
}

func (a Alarm) IsAlarmOn() bool {
	return a.alarmOn
}

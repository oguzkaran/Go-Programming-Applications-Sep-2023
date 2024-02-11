package main

import (
	"github.com/google/wire"
)

func InitSensorEvent() SensorEvent {
	wire.Build(NewName, NewSensorInfo, NewSensorEvent)
	return SensorEvent{}
}

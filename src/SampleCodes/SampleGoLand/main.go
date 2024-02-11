package main

import (
	"fmt"
)

type NameType string

type SensorInfo struct {
	Name NameType
	//...
}

type SensorEvent struct {
	Sensor *SensorInfo
	//...
}

func NewName() NameType {
	return NameType("CSD Sensor") //Bu isim herhangi bir yerden elde edilebilir
}

func NewSensorInfo(name NameType) *SensorInfo {
	return &SensorInfo{Name: name}
}

func (si *SensorInfo) GetName() NameType {
	return si.Name
}

func NewSensorEvent(si *SensorInfo) *SensorEvent {
	return &SensorEvent{Sensor: si}
}

func (e *SensorEvent) PrintName() {
	fmt.Printf("Sensor Name:%s\n", e.Sensor.GetName())
}

func main() {
	e := InitSensorEvent()
	e.PrintName()
}

package controller

import . "HotelAutomation/service"

type MotionController struct {
	Publisher
}

func NewMotionController() *MotionController {
	return &MotionController{}
}

func (mc *MotionController) RaiseMotionDetectedEvent(request MovementDetectedEvent) error{
	return mc.NotifyAll(request)
}

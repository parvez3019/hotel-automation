package service

import (
	. "HotelAutomation/model"
)

type HotelService struct {
	hotel *Hotel
}

func NewHotelService() *HotelService {
	return &HotelService{}
}

func (hotelService *HotelService) CreateHotel(request CreateHotelRequest) {
	hotelService.hotel = NewHotelBuilder().
		WithFloors(request.NumberOfFloors).
		WithCorridors(request.MainCorridorPerFloor, request.SubCorridorPerFloor).
		WithOneLightBulbAndOneACInEveryCorridor().
		Build()
}

func (hotelService *HotelService) GetAppliancesInfo() []AppliancesInfo {
	return mapApplianceToApplianceInfo(hotelService.GetAppliances())
}

func (hotelService *HotelService) GetAppliances() []Appliances {
	return mapToAppliances(hotelService.hotel.GetFloors())
}

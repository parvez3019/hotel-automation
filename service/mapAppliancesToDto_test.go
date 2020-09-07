package service

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	"github.com/stretchr/testify/assert"
	"testing"
)

var lightBulb *LightBulb
var airConditioner *AirConditioner
var mainCorridor *Corridor
var subCorridor *Corridor

func init() {
	lightBulb = NewLightBulb(1, 5)
	airConditioner = NewAirConditioner(1, 10)
	mainCorridor = NewCorridor(MAIN, 1).
		AddLightBulb(lightBulb).
		AddAirConditioner(airConditioner)
	subCorridor = NewCorridor(SUB, 1).
		AddLightBulb(lightBulb).
		AddAirConditioner(airConditioner)
}

func TestShouldReturnEmptyListOfApplianceInCaseOfNoFloors(t *testing.T) {
	assert.Equal(t, []Appliances{}, mapToAppliances([]*Floor{}))
}

func TestShouldReturnListOfApplianceFromMainCorridorOnlyInCaseOfNoSubCorridors(t *testing.T) {
	floor := NewFloor(1).
		AddCorridors([]*Corridor{mainCorridor}, MAIN)

	expectedApplianceInfo := []AppliancesInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5,
			Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10,
			Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
	}

	appliances := mapToAppliances([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, mapApplianceToApplianceInfo(appliances))
}

func TestShouldReturnListOfApplianceFromSubCorridorOnlyInCaseOfNoMainCorridors(t *testing.T) {
	floor := NewFloor(1).
		AddCorridors([]*Corridor{subCorridor}, SUB)

	expectedApplianceInfo := []AppliancesInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5,
			Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10,
			Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
	}

	appliances := mapToAppliances([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, mapApplianceToApplianceInfo(appliances))
}

func TestShouldReturnListOfApplianceFromAllCorridor(t *testing.T) {
	floor := NewFloor(1).
		AddCorridors([]*Corridor{mainCorridor}, SUB).
		AddCorridors([]*Corridor{subCorridor}, MAIN)

	expectedApplianceInfo := []AppliancesInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5,
			Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10,
			Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5,
			Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10,
			Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
	}

	appliances := mapToAppliances([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, mapApplianceToApplianceInfo(appliances))
}

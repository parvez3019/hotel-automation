package appliances

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAirConditioner(t *testing.T) {
	ac := NewAirConditioner(1, 10)
	assert.Equal(t, 1, ac.GetId())
	assert.Equal(t, 10, ac.GetPowerConsumption())
	assert.Equal(t, AC, ac.GetType())
	assert.False(t, ac.IsOn())
}

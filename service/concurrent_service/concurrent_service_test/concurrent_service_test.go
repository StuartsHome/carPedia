package concurrent_service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	service "github.com/stuartshome/carpedia/service/concurrent_service"
)

func TestGetDetails(t *testing.T) {

	conService := service.NewConcurrentService()
	result, result1 := conService.GetDetails()

	// carService := service.NewCarService()
	// userService := service.NewUserService()
	assert.NotNil(t, result, result1)
	// expected := service.ConService{}
	// assert.Equal(t, expected, result)

}

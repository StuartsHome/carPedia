package concurrent_service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	service "github.com/stuartshome/carpedia/service/concurrent_service"
)

/*
	To do:
	mock endpoints 'user' & 'caruser'
*/
func TestGetDetails(t *testing.T) {
	conService := service.NewConcurrentService()
	result := conService.GetDetails()
	assert.NotNil(t, result)
	assert.Equal(t, "Trevor", result.User.Name)
	assert.Equal(t, "Dacia", result.Car.Make)

}

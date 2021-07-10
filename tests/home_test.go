package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stuartshome/carpedia/service"
)

func TestJsonResponse(t *testing.T) {
	expectedResponse :=
		`{"cars":[{"make":"Fiesta","model":"Ford","reg":1980},{"make":"Mondeo","model":"Ford","reg":1995}]}`

	hmm, err := service.JsonResponse()
	response, _ := json.Marshal(hmm)

	require.Nil(t, err)
	assert.Equal(t, expectedResponse, string(response))
}

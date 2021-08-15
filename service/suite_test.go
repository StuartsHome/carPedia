package service

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/stuartshome/carpedia/mock/mock_store"
	"github.com/vektra/mockery/v2/pkg/config"
)

type ServiceTestSuite struct {
	suite.Suite

	// System under test
	svc *Service

	// Mocked dependencies
	store *mock_store.Store

	config *config.Config

	w *httptest.ResponseRecorder

	username string
	password string
	user     string
}

func (ts *ServiceTestSuite) SetupTest() {
	// Create test user
	ts.username = "dummy"
	ts.password = "test"
	ts.user = "dummy"

	ts.store = new(mock_store.Store)

	// Init empty config
	ts.config = &config.Config{}

	// Now construct the system under test by passing it the mocks
	ts.svc = &Service{
		Store: ts.store,
	}
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

/*
	When is this run?
	What mock does it use? Right now it looks like as though it uses mockito mock
*/

## Mock packages to generate

mock: mock/mock_store mock/mock_client

mock/mock_store: mock/mock_store/store.go

mock/mock_store/store.go: store/store.go
	$(GOGEN) store/store.go

mock/mock_client: mock/mock_client/client.go

mock/mock_client/client.go: http_client/client.go
	$(GOGEN) http_client/client.go
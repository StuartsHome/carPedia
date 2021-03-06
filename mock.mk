## Mock packages to generate

mock: mock/mock_store mock/mock_client mock/mock_redis_cache

mock/mock_store: mock/mock_store/interface.go

mock/mock_store/interface.go: store/interface.go
	$(GOGEN) store/interface.go

mock/mock_client: mock/mock_client/client.go

mock/mock_client/client.go: http_client/client.go
	$(GOGEN) http_client/client.go

mock/mock_redis_cache: mock/mock_redis_cache/interface.go

mock/mock_redis_cache/interface.go: cache/interface.go
	$(GOGEN) cache/interface.go
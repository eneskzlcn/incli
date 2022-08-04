generate-mocks:
	mockgen -destination=mocks/mock_root_handler.go -package mocks github.com/eneskzlcn/incli RootHandler
	mockgen -destination=mocks/mock_handler.go -package mocks github.com/eneskzlcn/incli Handler
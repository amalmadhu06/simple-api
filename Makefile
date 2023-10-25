wire:
	cd pkg/di && wire

mock:
	mockgen -source=pkg/repository/interfaces/user.go -destination=pkg/repository/mock/user.go
	mockgen -source=pkg/usecase/interfaces/user.go -destination=pkg/usecase/mock/user.go

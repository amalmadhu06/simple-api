wire:
	cd pkg/di && wire

mock:
	mockgen -source=pkg/repository/interfaces/user.go -destination=pkg/repository/mockrepo/user.go -package mockrepo
	mockgen -source=pkg/usecase/interfaces/user.go -destination=pkg/usecase/mockusecase/user.go -package mockusecase	

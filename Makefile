DATABASE_ADDR ?= '127.0.0.1:13306'

.PHONEY: none

none: ;

migration_up:
	@bin/migrate -database 'mysql://docker:docker@tcp(${DATABASE_ADDR})/clean-architecture-go' -path ./migrations up

migration_down:
	@bin/migrate -database 'mysql://docker:docker@tcp(${DATABASE_ADDR})/clean-architecture-go' -path ./migrations down

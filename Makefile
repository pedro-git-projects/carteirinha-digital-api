include .env
export

run:
	cd ./src/ && go run *.go

migrate-up:
	 migrate -path=./migrations -database="$(DATABASE_URL)" up

migrate-down:
	 migrate -path=./migrations -database="$(DATABASE_URL)" down 
	# migrate create -seq -ext .sql -dir ./migrations

migrate-reset-1:
	migrate -path=./migrations -database="$(DATABASE_URL)" force 1 


DATABASE = "database"


docker-all:
	docker build -t back backend  && docker build -t front frontend  
	docker compose up

docker-compose:
	docker compose up

docker-build:	
	docker build -t back backend  && docker build -t front frontend

NUMBER ?= 

migrate-down:
	migrate --path backend/pkg/db/migrations --database sqlite://backend/pkg/db/database.db down $(NUMBER)

migrate-up:
	migrate --path backend/pkg/db/migrations --database sqlite://backend/pkg/db/database.db up $(NUMBER)

TABLE ?= add_new_table

migrate-new:
	 migrate create -ext sql -dir backend/pkg/db/migrations -seq $(TABLE) 


migrate-version:
	 migrate -path backend/pkg/db/migrations -database sqlite://backend/pkg/db/database.db version


launch-back: 
	cd backend/cmd/ && go run main.go 

launch-front:
	cd frontend && npm run dev

NAME ?= back

create-exe:
	go build -o $(NAME).exe backend/exe/exe.go



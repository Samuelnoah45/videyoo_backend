
# docker compose up
up:
	sudo docker compose up 

# docker compose down
down:
	sudo docker compose down


api_port ?= 7770
console_port ?= 7771
console:
	hasura console  --api-port $(api_port) --console-port $(console_port) --project ./hasura-graphql

run:
	go run ./server/main.go

# apply metadata
metadata_apply:
	hasura metadata apply --envfile .env --project ./hasura-graphql


#apply migrations
migrate_apply:
	hasura migrate apply --envfile .env --project ./hasura-graphql
metadata_prod:
	hasura metadata apply --envfile ../.env.prod --project ./hasura-graphql

migrate_prod:
	hasura migrate apply --envfile ../.env.prod --project ./hasura-graphql

meta_export:
	hasura metadata export --envfile ../.env --project ./hasura-graphql

ic_list:
	hasura metadata ic list --project ./hasura-graphql

# run go file
go_run:
	go run ./server/main.go



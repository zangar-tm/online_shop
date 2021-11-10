.SILENT:

run:
	docker-compose up online-shop

migrateup:
	migrate -path ./schema -database 'postgres://postgres:secret@0.0.0.0:5433/postgres?sslmode=disable' up	

migratedown:
	migrate -path ./schema -database 'postgres://postgres:secret@0.0.0.0:5433/postgres?sslmode=disable' down	
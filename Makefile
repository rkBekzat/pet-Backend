
postgres:
	docker run --name=pets -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

createdb:
	#docker exec -it pets createdb pet_project
	migrate -source file://schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
#migrate -source file://schema -database "postgres:://postgres:qwerty@localhost:5436/postgres?sslmode=disadsadble" -verbose up

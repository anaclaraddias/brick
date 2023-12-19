docker exec -i app_db psql -U sa -d brickdb -a -f 01_create_vehicle_table.up.sql

docker exec -i app_db migrate -path ./infra/database/migrations -database "mysql://root:admin@tcp(app_db)/brickdb" up

# docker exec -i emcash_simulador migrate -path ./infra/database/migrations/mariaDB -database "mysql://root:$1@tcp(emcash_db)/dbsimulador" up

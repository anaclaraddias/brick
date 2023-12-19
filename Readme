# Documentação

## Configuração docker

O primeiro passo é buildar a imagem:
- rode `docker build -t brick-backend .` 

Agora é preciso buildar o docker compose:
- rode `docker-compose build`

E por fim, é só rodar os containers:
- rode `docker compose up`

## Conexão com o banco de dados

Para conectar com o banco de dados, atente-se as informações:

- host: `postgres`
- username: `root`
- password: `admin`
- port: `5432`
- host ip: `127.0.0.1`

## Criação de tabelas do banco de dados

Copie as queries que estão dentro da pasta infra/migrations no arquivo `create_tables.sql` e rode dentro do banco de dados `brickdb`

## Utilize as rotas 
- POST user -> cria o usuário no banco de dados
- POST vehicle -> cria o veículo no banco de dados
- DELETE vechile -> deleta o veículo no banco de dados
- POST policy -> cria a apólice no banco de dados 
- POST policy/vehicle -> cria o relacionamento entre a apólice e o veículo
- POST policy/coverage -> cria o relacionamento entre a apólice e as coberturas
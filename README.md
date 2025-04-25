# go-expert-clean-arch

Implementação de um desafio de Clean Architecture. O sistema é composto por uma API RESTful, gRPC e GraphQL, e interage com o MySQL e RabbitMQ.
🛠️ Pré-requisitos

Antes de executar o projeto, certifique-se de ter os seguintes softwares instalados:

- Docker
- Go
- Docker Compose

## 🚀 Como rodar o projeto
1. Subindo o ambiente com Docker Compose

Primeiro, é necessário subir o ambiente que inclui o MySQL e o RabbitMQ, além de executar as migrations. Para isso, execute o comando abaixo na raiz do projeto:

```bash
docker compose --env-file ./cmd/ordersystem/.env up -d
```

Este comando irá:

- Subir o MySQL.
- Subir o RabbitMQ.
- Executar as migrations necessárias para o MySQL.

2. Rodando o sistema Go

Após subir o ambiente com o Docker Compose, entre no diretório cmd/ordersystem/ e execute o seguinte comando para rodar o sistema:

```bash
cd ./cmd/ordersystem/
go run main.go wire_gen.go
```

Este comando irá iniciar a aplicação. Os serviços estarão disponíveis nas seguintes portas:
- HTTP (RESTful): http://localhost:8000
- gRPC: localhost:50051
- GraphQL: localhost:8080


É possível carregar como teste algumas orders no sistema executando o arquivo api/create_order.http

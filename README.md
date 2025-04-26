# go-expert-clean-arch

Implementação de um desafio de Clean Architecture. O sistema é composto por uma API RESTful, gRPC e GraphQL, e interage com o MySQL e RabbitMQ.
🛠️ Pré-requisitos

Antes de executar o projeto, certifique-se de ter os seguintes softwares instalados:

- Docker
- Docker Compose

## 🚀 Como rodar o projeto
1. Subindo o ambiente com Docker Compose

Execute o seguinte comando apra subir o ambiente via docker compose.

```bash
docker compose --env-file ./cmd/ordersystem/.env up -d
```

Este comando irá:

- Subir o MySQL.
- Subir o RabbitMQ.
- Executar as migrations necessárias para o MySQL.
- Iniciar o servidor da applicação Go


Os serviços da aplicação estarão disponíveis nas seguintes portas:
- HTTP (RESTful): http://localhost:8000
- gRPC: localhost:50051
- GraphQL: localhost:8080


É possível carregar como teste algumas orders no sistema executando o arquivo api/create_order.http

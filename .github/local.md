# Executando o Projeto Localmente (apenas com container do MySQL e RabbitMQ)

Este guia fornece instruções para executar o projeto diretamente na sua máquina, utilizando apenas container do MySQL e RabbitMQ

## Pré-requisitos

Antes de começar, certifique-se de ter instalado:

- Docker
- docker compose

## Configuração

1. Primeiro, clone o repositório se ainda não o fez:
```bash
git clone https://github.com/dyhalmeida/go-clean-arch.git
cd go-clean-arch
```

2. Configure o arquivo de ambiente:
```bash
cd cmd/orders
cp .env.example .env
```

3. Abra o arquivo `docker-compose.yml` e comente o serviço da aplicação (app) para evitar conflitos:
```yaml
# Comente todo o bloco do serviço 'app':
#  app:
#    build: .
#    ports:
#      - "8080:8080"
#      - "50051:50051"
#    volumes:
#      - .:/go/src/
#    depends_on:
#      - mysql
#      - rabbitmq
```

4. Inicie os serviços de dependência (MySQL e RabbitMQ):
```bash
docker-compose up -d
```

5. Verifique se o arquivo `.env` em `cmd/orders/.env` está configurado corretamente. Os valores devem ser similares a:
```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=:8080
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8081
RABBITMQ_HOST=localhost
RABBITMQ_USER=guest
RABBITMQ_PASS=guest
RABBITMQ_PORT=5672
```

## Executando a Aplicação

1. Navegue até a pasta da aplicação:
```bash
cd cmd/orders
```

2. Execute a aplicação:
```bash
go run main.go wire_gen.go
```

A aplicação estará disponível com:
- API REST: http://localhost:8080
- GraphQL Playground: http://localhost:8080/graphql
- gRPC: localhost:50051
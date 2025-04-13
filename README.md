# Clean Architecture em Go

Este projeto é um desafio proposto no curso GoExpert da FullCycle. Ele consiste em um sistema básico de pedidos (orders), no qual cada pedido possui um valor e uma taxa, e o sistema é responsável por calcular o valor final.

O principal objetivo do projeto é demonstrar na prática como aplicar uma arquitetura limpa, escalável e de fácil manutenção utilizando Go, seguindo boas práticas de desenvolvimento de software.

Além disso, o projeto implementa múltiplas interfaces de comunicação, permitindo o consumo do sistema por diferentes meios:

- HTTP REST
- GraphQL
- gRPC

## 🚀 Tecnologias

- Go 1.21+
- Docker & Docker Compose
- MySQL
- RabbitMQ
- gRPC
- GraphQL
- REST API

## 📁 Estrutura do Projeto

```
.
├── cmd/           # Pontos de entrada da aplicação
├── configs/       # Configuração de variáveis de ambiente
├── internal/      # Código interno da aplicação
│   ├── entity/    # Entidades e regras de negócio
│   ├── event/     # Hanlder de eventos
│   ├── infra/     # Implementações de infraestrutura (database, graphql, grpc, web)
│   └── usecase/   # Casos de uso da aplicação
```

## 🔧 Configuração

1. Clone o repositório:
```bash
git clone https://github.com/dyhalmeida/go-clean-arch.git
cd go-clean-arch
```

2. Inicie os containers:
```bash
docker-compose up -d
```

3. Caso deseje executar esse projeto sem uso de container, siga essa [doc](.github/local.md)

## 💻 Como Usar

### REST API

A API REST está disponível em `http://localhost:8080`

Exemplos de requisições:

```bash
# Create order
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{
    "id":"ae717117-cdfb-427f-9f8d-fe34c183f5aa",
    "price": 100.5,
    "tax": 0.5
  }'


# List orders
curl http://localhost:8080/order
```
ou se preferir utilize o arquivo `create_order.http` que se encontra na pasta `api`

### GraphQL

O playground do GraphQL está disponível em `http://localhost:8080/graphql`

Exemplo de queries:

```graphql
# Create order
mutation CreateOrder {
  createOrder(input: { 
    id: "80be177d-e920-4626-a4e6-4aee05e7455c",
    Price: 10.20,
    Tax: 2.30
  }),
  {
    id
    Price
    Tax
    FinalPrice
  }
}

# List orders
query ListAllOrders {
  listOrder {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### gRPC

O servidor gRPC está disponível na porta 50051. Você pode usar o [Evans](https://github.com/ktr0731/evans) para interagir com o servidor gRPC.

1. Conecte ao servidor:
```bash
evans -r repl -p 50051
```

3. Exemplos de comandos no Evans:
```bash

# Mostra os serviços disponíveis
show service

# Mostra os métodos disponíveis
show message

# Chama CreateOrder para criar order
call CreateOrder

# Chama ListOrder para listar order
call CreateOrder
```

## 🎥 Demonstração

Abaixo uma demonstração de como utilizar o Evans CLI para interagir com o servidor gRPC:

https://github.com/dyhalmeida/go-clean-arch/blob/main/.github/evans-example.gif

Passo a passo para execução da aplicação:

* Para subir as dependencias do projeto execute o comando:
```
    docker-compose up -d
```

* Para executar o projeto execute os comandos:
```
    go mod tidy
    cd cmd/ordersystem/ && go run main.go wire_gen.go
```

A aplicação estará de pé respondendo nas seguintes portas:
* Web Server: 8000 
* GraphQL: 8080
* gRPC: 50051

1. Para testar o Web Server, execute os arquivos create_order.http e list_orders.http que estao na pasta api.

2. Para testar o GraphQL acesse a url http://localhost:8080/ no navegador e execute as seguintes os seguintes comandos:
```
mutation createOrder {
  createOrder(input:{id:"A", Price:30, Tax:2}) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

```
query ListOrders {
  ListOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

3. Para testar o gRPC utilize o evans, execute o comando:
```
evans -r repl
```
E faça chamada as seguintes services:

* call CreateOrder
* call ListOrders
Cashmachine
============

### Descrição
    Cashmachine é um serviço para gerenciar uma conta bancária e executar saques

### Linguagens
- golang: 1.14.3
  
### Frameworks
- fiber/fiber: 2.14.3


Instalation
-----------

### Build & Run

- Server.

```sh
$ docker-compose build
$ docker-compose up
```

### Configurando a database

Após iniciar os serviços com o docker-compose execute o comando:
```sh
make build-db
```

### Testando

```sh
$ make test
$ make cover
```

Exemplo
-------

Após iniciar o servidor é possível ter acesso ao serviço através da porta 3000

```sh
curl --location --request GET 'localhost:3000/balance/1'
```

API
---
### POST /new/
Cria uma nova conta e retorna o seu id e saldo incial.

form-data:
  - value float   // Deposito inicial da nova conta.

### GET /balance/:id
Retorna o saldo atual de uma conta.

### PUT /withdraw/:id
Saca um valor de uma detrminada conta.

form-data:
  - value int     // Valor a ser sacado. 


### PUT /deposit/:id
Deposita um valor em uma determinada conta

form-data:
  - value int     // Valor a ser depositado

Na pasta api/ está o arquivo de colection do postman para chamadas da api.

_________

Config
------
Os parâmetros de configuração do serviço estão no arquivo .env na raiz do projeto

Arquitetura
-------------

## Clean Architecture

Algumas caracteristicas do modelo de arquitetura "Clean Architecture" foram adotadas durante o desenvolvimento.

  - Regra de negocio da empresa (./pkg/entity): responsável pelas regras de negocios e validação de dados.
  - Casos de usos (./pkg/usecase): nesta camada estão as unidades de código responsáveis por orquestrar as entidades do sistema.
  - Camada de acesso a dados (./pkg/repository): responsável por isolar a aplicação do SGBD tornando assim, a api independente do método de armazenamento de dados e possibilitando a fácil substituição do mesmo.
  - Driver (./pkg/driver): Responsável por tratar dos detalhes de implementação do SGBD e outros tipos de conexões que possam vir a integrar o sistema.
  
## Pinrcípio da inversão de dependencia
Neste projeto foi utilizado o conceito do Princípio de Inversão de Dependencia que faz uso de interfaces para orientar o fluxo e direção das dependencias entre os modulos, destá forma os modulos menos estáveis (controllers) apontam na direção dos modulos mais estáveis (usecases). O objetivo com isso é fazer com que o design da aplicação gire em torno dos casos de uso do sistema e não do framework (que pode ser fácilmente substituido caso seja convenienete). 


### Fiber
A exposição da API REST é feita através da biblioteca Fiber, que utiliza a fasthttp que é a biblioteca http para golang com os melhores resultados em benchmarks.
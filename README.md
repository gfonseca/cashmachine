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

Na pasta api/ esta o arquivo de colection do postman para chamadas da api.

_________

Config
------
Os parâmetros de configuração do serviço estão no arquivo .env na raiz do projeto
### ./.env
    - port: default 3000
    - config-file: default input-file.txt


Arquitetura
-------------

Aplicação em 3 camadas:

  - Regra de negocio da empresa: responsável pelos calculos de rotas e validação de dados.
  - Regra de negocoi da aplicação: responsável por orquestrar os detalhes de implementação e as regras de negocio da empresa.
  - Camada de acesso a dados: Controla o acesso ao arquivo que possui os dados de rotas e custos.

### Estrutura de pastas do projeto:

- api: Arquivos para documentação de contratos de dados
- bin: Binários, scripts e outros executáveis
- cmd: Arquivos fontes de entry points do projeto
- pkg: Pacotes do projeto em golang
    - pkg/controller: Regra de negocio da aplicação
    - pkg/graph: Regra de negocio da empresa
    - pkg/repository: camada de acesso a dados
    - pkg/test: Utilitários para testes

A exposição da API REST é feita através da biblioteca Fiber, que utiliza a fasthttp que é a biblioteca http para golang com os melhores resultados em benchmarks.
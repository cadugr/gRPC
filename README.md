
# Projeto gRPC em Go

Este projeto demonstra a implementação de um serviço gRPC utilizando a linguagem Go, integrando-se a um banco de dados SQLite3. O objetivo é fornecer um exemplo prático de como estruturar um projeto gRPC em Go e interagir com um banco de dados local.

## Pré-requisitos

- **Go** instalado em sua máquina: [Instalar Go](https://go.dev/doc/install)
- **Evans**: Cliente gRPC para interagir com o serviço. Instale utilizando o comando abaixo ou acesse a documentação oficial para outras opções:
  
  ```bash
  go install github.com/ktr0731/evans@latest
  ```

- **SQLite3**: Este projeto utiliza SQLite3 como banco de dados local.

## Como baixar o projeto

Clone o repositório do projeto do GitHub usando o seguinte comando:

```bash
git git@github.com:cadugr/gRPC.git
cd gRPC
```

## Configuração do Banco de Dados

Este projeto utiliza SQLite3 como banco de dados, e o arquivo do banco de dados precisa ser criado, conforme instruções a seguir.

### Criando o Banco de Dados SQLite Localmente

Para criar o banco de dados e as tabelas localmente, siga os passos abaixo:

1. No terminal, abra o cliente SQLite e crie um banco de dados (ex.: `.db.sqlite`):

    ```bash
    sqlite3 .db.sqlite
    ```

2. Dentro do cliente SQLite, crie a tabela `categories` com o seguinte comando:

    ```sql
    CREATE TABLE IF NOT EXISTS categories (
        id TEXT PRIMARY KEY,
        name TEXT,
        description TEXT
    );
    ```

3. Para sair do SQLite, digite:

    ```sql
    .exit
    ```

O banco de dados será criado e salvo no diretório atual.

## Executando o Projeto

1. Instale as dependências:

    ```go
    go mod tidy
    ```

2. Inicie o servidor gRPC:

    ```go
    go run cmd/grpcServer/main.go
    ```

O servidor estará rodando no endereço `localhost:50051`, por padrão.

## Testando o Serviço com Evans

Utilize o **Evans** para interagir com o servidor gRPC e realizar chamadas aos métodos disponíveis:

1. Inicie o Evans no modo REPL:

    ```bash
    evans -r repl -p 50051
    ```

2. Dentro do Evans, você pode listar e chamar os métodos do serviço gRPC disponíveis.

Para mais detalhes sobre os comandos do Evans, consulte a [documentação oficial](https://github.com/ktr0731/evans).

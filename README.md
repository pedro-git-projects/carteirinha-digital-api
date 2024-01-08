# Carteirinha Digital - API

Este repositório contém a implementação da API em Golang para um aplicativo de controle de entrada de alunos, conhecido como "Carteirinha Digital". O aplicativo utiliza o framework Gin para a construção da API e a biblioteca go-qrcode para a geração de códigos QR. Além disso, o banco de dados utilizado é o PostgreSQL, gerenciado em um container Docker, e as migrações são realizadas através de arquivos .sql utilizando a ferramenta Golang Migrate.

## Configuração do Ambiente

### Requisitos

Certifique-se de ter os seguintes requisitos instalados em sua máquina:

- Golang
- Docker
- Docker Compose
- GNU Make

### Configuração do Banco de Dados

O banco de dados PostgreSQL pode ser iniciado em um container Docker utilizando o seguinte comando:

```bash
docker-compose up -d
```

Isso iniciará o PostgreSQL em um container configurado conforme as especificações do arquivo `docker-compose.yml`.

### Migrações

As migrações do banco de dados são gerenciadas através da ferramenta Golang Migrate. Certifique-se de ter a ferramenta instalada em sua máquina. As migrações estão localizadas no diretório `migrations/`. Para executar as migrações, utilize o seguinte comando:

```bash
migrate -path migrations/ -database "postgres://username:password@localhost:5432/database_name?sslmode=disable" up
```

Substitua `username`, `password`, e `database_name` com as credenciais apropriadas.

## Endpoints da API

### Autenticação

A autenticação é necessária para acessar os seguintes endpoints. O token de acesso deve ser enviado no cabeçalho da requisição.

- **GET /auth/qr-code**: Gera um código QR para autenticação.
- **POST /auth/record-attendance**: Registra a entrada de um aluno na escola.

#### Autenticação de Usuários

- **POST /auth/students/signin**: Autenticação de estudantes.
- **POST /auth/parents/signin**: Autenticação de responsáveis.
- **POST /auth/staff/signin**: Autenticação de funcionários.

### Cadastro de Usuários

- **POST /parent-student**: Criação de vínculo entre responsável e estudante.
- **POST /students**: Cadastro de estudantes.
- **POST /parents**: Cadastro de responsáveis.
- **POST /staff**: Cadastro de funcionários.

### Saúde da API

- **GET /**: Verifica a saúde da API.

## Como Executar

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/carteirinha-digital-api.git
cd carteirinha-digital-api
```

2. Configure o ambiente e inicie o banco de dados:

```bash
docker-compose up -d
```

3. Execute as migrações:

```bash
make migrate-up
```

4. Inicie a API:

```bash
make run
```

A API estará disponível em `http://localhost:8080`.

Esteja ciente de fornecer as credenciais corretas para o banco de dados e ajustar conforme necessário antes de iniciar a aplicação.

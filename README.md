# 🌍 API de Sistema Web de Viagens

Esta é uma API RESTful desenvolvida em **Go** para um sistema de viagens. A API gerencia **destinos turísticos** e **depoimentos de clientes**, com suporte a operações CRUD completas e endpoints otimizados. O projeto foi configurado para rodar em ambiente **Docker** e utiliza **Gin** e **Gorm** para gerenciamento de rotas e banco de dados.

---

<p align="center">
  <a href="#️-tecnologias-utilizadas">Tecnologias Utilizadas</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-como-configurar-e-executar-o-projeto">Como Configurar e Executar o Projeto</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#️-estrutura-do-projeto">Estrutura do Projeto</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-melhorias-futuras">Melhorias Futuras</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#memo-licença">Licença</a>
</p>

<p align="center">
  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=49AA26&labelColor=000000">
</p>

## 🛠️ **Tecnologias Utilizadas**

- **Golang (Go)**: Linguagem principal do projeto.
- **Gin**: Framework para criação de APIs rápidas e eficientes.
- **Gorm**: ORM para manipulação do banco de dados.
- **SQLite**: Banco de dados local para armazenar informações.
- **Docker**: Containerização do ambiente.
- **Postman**: Testes de endpoints e validação da API.

---

## 🚀 **Como Configurar e Executar o Projeto**

### Pré-requisitos

Certifique-se de ter os seguintes itens instalados no seu sistema:

- **Docker** e **Docker Compose**
- **Git**
- **Postman** (opcional, mas recomendado para testar a API)

### Clonar o Repositório

```bash
git clone https://github.com/usrJosephC/api-viagens.git
cd api-viagens
```

## 🗂️ Estrutura do Projeto

.
├── controllers # Lógica dos endpoints
│ ├── destination_controller.go
│ ├── testimony_controller.go
├── models # Modelos para o banco de dados
│ ├── destination.go
│ ├── testimony.go
├── routes # Configuração das rotas
│ ├── router.go
├── main.go # Entrada principal da aplicação
├── Dockerfile # Configuração do container
├── docker-compose.yml # Orquestração dos containers
└── README.md # Documentação do projeto

## 📌 Melhorias Futuras
1. Adicionar autenticação (ex.: JWT) para proteger endpoints.
2. Configurar suporte a bancos de dados externos (ex.: PostgreSQL, MySQL).
3. Implementar paginação para os endpoints GET.
4. Adicionar suporte a upload de imagens reais nos depoimentos e destinos.

### 🤝 Contribuição
Sinta-se à vontade para abrir issues ou enviar pull requests para melhorar o projeto. Feedbacks são bem-vindos!

## :memo: Licença

Esse projeto está sob a licença MIT.

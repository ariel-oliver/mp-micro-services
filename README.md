# Documentação dos Microservices

Este repositório contém três microservices escritos em GoLang.

## Pré-Requisitos

Antes de iniciar os serviços, certifique-se de ter o GoLang instalado em sua máquina. Você pode baixar e instalar o GoLang em [golang.org](https://golang.org/dl/).

## Execução dos Serviços

Para executar os serviços, siga as etapas abaixo:

1. Abra três terminais.

2. Em cada terminal, navegue até o diretório de um dos serviços:
   - Terminal 1: `cd jwt-auth-service`
   - Terminal 2: `cd product-service`
   - Terminal 3: `cd web`

3. Em cada terminal, execute o seguinte comando para iniciar o serviço:
   ```shell
      go run main.go
   ```
- `jwt-auth-service` será executado na porta **8083**
- `product-service` será executado na porta **8081**
- `web` será executado na porta **8080**

Certifique-se de que todas as dependências estejam instaladas corretamente antes de iniciar os serviços.

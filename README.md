# Aplicação de Linha de Comando

## Visão Geral

Esta aplicação foi desenvolvida em **Go** e fornece as seguintes funcionalidades:

- **Receber um endereço web**: A aplicação recebe um endereço web (por exemplo, `example.com`) como entrada.
- **Retornar o IP público**: Resolve e exibe o endereço IP público do endereço web informado.
- **Retornar o nome do servidor**: Identifica e retorna o nome do servidor onde o endereço web está hospedado.

## Funcionalidades

- Leve e rápida.
- Interface simples de linha de comando.
- Capacidade eficiente de resolução de DNS e rede.

## Pré-requisitos

- Go instalado (recomenda-se a versão 1.18 ou superior).

## Como Começar

1. Clone o repositório:
   ```bash
   git clone <repository-url>
   ```
2. Navegue até o diretório do projeto:
   ```bash
   cd <project-directory>
   ```
3. Compile a aplicação:
   ```bash
   go build -o web-info
   ```
4. Execute a aplicação:

   ```bash
   ./web-info <endereco-web>
   ```

   Exemplo:

   ```bash
   ./web-info example.com
   ```

## Exemplo de Saída

```
Endereço Web: example.com
IP Público: 93.184.216.34
Nome do Servidor: example-host
```

## Licença

Este projeto está licenciado sob a Licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.

## Contribuições

Contribuições são bem-vindas! Fique à vontade para abrir uma issue ou enviar um pull request.

## Autor

Walker Esteves Negrão

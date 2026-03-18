# Sistema Runner

Ferramenta de linha de comandos para execução de aplicações Java relacionadas à assinatura digital no contexto do HubSaúde (SES-GO).

## Visão geral

O projeto é composto por dois componentes principais:

- **`assinatura`** — CLI multiplataforma (Go) que o usuário executa diretamente
- **`assinador.jar`** — Aplicação Java que valida parâmetros FHIR e simula operações de assinatura digital

```
Usuário → assinatura (CLI) → assinador.jar → resposta formatada → Usuário
```

## Estrutura do repositório

```
runner/
├── cli/                   # Código-fonte da CLI (Go)
│   ├── cmd/
│   │   ├── assinar.go
│   │   ├── validar.go
│   │   └── simulador.go
│   ├── internal/
│   └── main.go
├── assinador/             # Código-fonte do assinador.jar (Java)
│   ├── src/
│   └── pom.xml
├── .github/
│   └── workflows/
│       └── build.yml
└── README.md
```

## Pré-requisitos

| Ferramenta | Versão mínima | Usado por |
|---|---|---|
| Go | 1.22+ | CLI `assinatura` |
| JDK | 17+ | `assinador.jar` |
| Maven | 3.9+ | Build do `assinador.jar` |

> O JDK pode ser provisionado automaticamente pela CLI se não estiver presente (ver US-04).

## Instalação e build

### Clonar o repositório

```bash
git clone https://github.com/org/runner.git
cd runner
```

### Build da CLI

```bash
cd cli
go build -o assinatura .
```

Para compilar para outras plataformas:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o assinatura-linux-amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o assinatura-windows-amd64.exe .

# macOS
GOOS=darwin GOARCH=amd64 go build -o assinatura-darwin-amd64 .
```

### Build do assinador.jar

```bash
cd assinador
mvn clean package
```

O artefato gerado fica em `assinador/target/assinador.jar`.

## Uso básico

```bash
# Criar assinatura (simulado)
./assinatura assinar --bundle <arquivo> --politica <uri>

# Validar assinatura (simulado)
./assinatura validar --jws <base64> --timestamp <unix> --politica <uri>

# Gerenciar o simulador do HubSaúde
./assinatura simulador start
./assinatura simulador stop
./assinatura simulador status

# Ajuda
./assinatura --help
./assinatura assinar --help
```

## Executar os testes

```bash
# Testes da CLI
cd cli
go test ./...

# Testes do assinador.jar
cd assinador
mvn test
```

## Parâmetros e especificação FHIR

Os parâmetros de entrada e os formatos de resposta seguem as especificações da SES-GO:

- [Criar assinatura](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-criar-assinatura.html)
- [Validar assinatura](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-validar-assinatura.html)

As respostas de sucesso e erro usam o formato `OperationOutcome` (FHIR R4).

## Integração entre componentes

A CLI invoca o `assinador.jar` de dois modos:

- **Modo CLI (padrão):** subprocess por chamada — `java -jar assinador.jar <args>`
- **Modo HTTP:** processo persistente na porta configurada — reduz latência em múltiplas chamadas

## Releases e distribuição

Os binários pré-compilados são distribuídos via [GitHub Releases](https://github.com/org/runner/releases) com versionamento semântico (SemVer).

Cada release inclui checksums SHA-256:

```bash
# Verificar integridade (exemplo Linux)
sha256sum -c assinatura-linux-amd64.sha256
```

## Contribuindo

1. Crie uma branch a partir de `main`: `git checkout -b feat/minha-feature`
2. Faça as alterações e adicione testes
3. Abra um Pull Request descrevendo o que foi alterado

## Referências

- [Especificação do projeto (este repositório)](./docs/especificacao.md)
- [Modelo C4](https://c4model.com/)
- [FHIR R4 — OperationOutcome](https://www.hl7.org/fhir/R4/operationoutcome.html)
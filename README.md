# Sistema Runner

CLIs para execução de operações de assinatura digital e gerenciamento do Simulador do HubSaúde (SES-GO).

## Componentes

| Componente | Linguagem | Papel |
|---|---|---|
| `assinador` | Go | CLI — comandos `assinar` e `validar` |
| `simulador` | Go | CLI — comandos `iniciar`, `parar`, `status` |
| `assinador.jar` | Java | Valida parâmetros FHIR e simula assinaturas |
| `simulador.jar` | Java | Simulador do HubSaúde — externo, fora do escopo |

## Estrutura do repositório

```
runner/
├── assinador/              # CLI de assinatura
│   ├── cmd/
│   │   ├── assinar.go
│   │   └── validar.go
│   ├── internal/           # código privado deste módulo
│   ├── main.go
│   └── go.mod              # module runner/assinador
│
├── simulador/              # CLI do simulador
│   ├── cmd/
│   │   ├── iniciar.go
│   │   ├── parar.go
│   │   └── status.go
│   ├── internal/           # código privado deste módulo
│   ├── main.go
│   └── go.mod              # module runner/simulador
│
├── shared/                 # código Go compartilhado entre os CLIs
│   ├── jdk/                # detecção e provisionamento do JDK
│   ├── output/             # formatação stdout/stderr e exit codes
│   ├── executor/           # invocação do assinador.jar (subprocess ou HTTP)
│   ├── processo/           # gerenciamento de processos do OS
│   └── go.mod              # module runner/shared
│
├── assinador-jar/          # aplicação Java
│   ├── src/
│   │   ├── main/java/runner/assinador/
│   │   │   ├── operacoes/  # CriarAssinatura, ValidarAssinatura
│   │   │   ├── validacao/  # validação de parâmetros FHIR
│   │   │   ├── simulacao/  # respostas pré-construídas
│   │   │   └── http/       # servidor HTTP (modo warm start)
│   │   └── test/java/runner/assinador/
│   └── pom.xml
│
├── docs/
│   ├── design/             # diagramas C4 em PlantUML
│   │   ├── contexto.puml
│   │   ├── conteineres.puml
│   │   ├── imagens/
│   │   ├── geraimagens.sh
│   │   └── geraimagens.bat
│   └── ...                 # documentação adicional (futura)
│
├── .github/
│   └── workflows/
│       ├── assinador.yml   # build + testes do CLI assinador
│       ├── simulador.yml   # build + testes do CLI simulador
│       └── assinador-jar.yml
│
├── go.work                 # workspace — vincula assinador, simulador e shared
├── go.work.sum
├── Makefile                # make build · make test · make release
├── .gitignore
└── README.md
```

## Pré-requisitos

| Ferramenta | Versão |
|---|---|
| Go | 1.22+ |
| JDK | 17+ |
| Maven | 3.9+ |

## Build

```bash
# tudo de uma vez
make build

# individualmente
cd assinador     && go build -o assinador .
cd simulador     && go build -o simulador .
cd assinador-jar && mvn clean package
```

Cross-compile:

```bash
cd assinador
GOOS=linux   GOARCH=amd64 go build -o assinador-linux-amd64 .
GOOS=windows GOARCH=amd64 go build -o assinador-windows-amd64.exe .
GOOS=darwin  GOARCH=amd64 go build -o assinador-darwin-amd64 .
```

## Uso

```bash
# Assinatura
./assinador assinar --pacote <arquivo> --politica <uri> [flags]
./assinador validar --assinatura <base64> --momento <unix> --politica <uri> [flags]

# Simulador
./simulador iniciar
./simulador parar
./simulador status
```

## Testes

```bash
# tudo de uma vez
make test

# individualmente
cd assinador     && go test ./...
cd simulador     && go test ./...
cd shared        && go test ./...
cd assinador-jar && mvn test
```

---

## Entregáveis

### Binários

| Artefato | Plataformas |
|---|---|
| `assinador` | linux-amd64, windows-amd64, darwin-amd64 |
| `simulador` | linux-amd64, windows-amd64, darwin-amd64 |

Distribuídos via GitHub Releases com checksums SHA-256 e versionamento SemVer.

### Código-fonte

| Componente | Linguagem |
|---|---|
| `assinador` CLI | Go |
| `simulador` CLI | Go |
| `assinador.jar` | Java |

### Testes

| Escopo | Tipo |
|---|---|
| `assinador.jar` — validação de parâmetros FHIR | Unitários |
| `assinador.jar` — respostas simuladas | Unitários |
| `assinador` + `assinador.jar` | Integração |
| `simulador` + `simulador.jar` | Integração |

### Documentação

| Documento | Conteúdo |
|---|---|
| `README.md` | Build, uso e referências — um por componente |
| Design C4 | Diagramas de contexto e contêineres em PlantUML |
| Especificação | Requisitos, escopo e decisões de design |

---

## Requisitos

### `assinador` CLI (Go)

**Funcionais**

- Aceita os subcomandos `assinar` e `validar`
- Valida a presença e o formato dos flags antes de invocar o `assinador.jar`
- Invoca o `assinador.jar` via subprocess (`java -jar`) ou HTTP, conforme configuração
- Exibe o resultado da operação em texto legível no `stdout`
- Exibe erros no `stderr` com mensagem descritiva e `exit code` não-zero

**Não funcionais**

- Binário único, sem dependências de runtime
- Executa em Windows, Linux e macOS (amd64)
- `exit code 0` — sucesso; `exit code 1` — erro de parâmetro; `exit code 2` — erro de execução
- `--ajuda` disponível em todos os subcomandos

---

### `simulador` CLI (Go)

**Funcionais**

- Aceita os subcomandos `iniciar`, `parar` e `status`
- `iniciar` — inicia o processo `simulador.jar`; falha com erro claro se já estiver rodando
- `parar` — encerra o processo; falha com erro claro se não estiver rodando
- `status` — exibe se o processo está rodando ou parado, com PID quando ativo
- Localiza o `simulador.jar` via flag `--jar <caminho>` ou variável de ambiente `SIMULADOR_JAR`

**Não funcionais**

- Mesmos requisitos de plataforma e `exit code` do `assinador` CLI

---

### `assinador.jar` (Java)

**Funcionais**

- Aceita dois modos de operação:
  - **CLI** — invocado via `java -jar assinador.jar <args>`; responde e encerra
  - **HTTP** — inicia servidor e aguarda requisições; reduz overhead de cold start
- Operação `criar`:
  - Valida todos os parâmetros de entrada conforme [spec FHIR](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-criar-assinatura.html)
  - Parâmetros inválidos → retorna `OperationOutcome` com código de erro específico
  - Parâmetros válidos → retorna JWS JSON Serialization pré-construída (simulada)
- Operação `validar`:
  - Valida todos os parâmetros de entrada conforme [spec FHIR](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-validar-assinatura.html)
  - Parâmetros inválidos → retorna `OperationOutcome` com código de erro específico
  - Parâmetros válidos → retorna `OperationOutcome` com `VALIDATION.SUCCESS`
- Todas as respostas usam o formato `OperationOutcome` (FHIR R4)

**Parâmetros obrigatórios — `validar`**

| Parâmetro | Tipo | Descrição |
|---|---|---|
| JWS | base64 | JWS JSON Serialization completa (RFC 7515 §3.2) |
| trust store | array de string | Hashes SHA-256 (hex, 64 chars) das AC-Raiz ICP-Brasil aceitas |
| timestamp de referência | inteiro Unix UTC | Usado em todas as verificações temporais. Faixa: `1751328000` a `4102444800` |
| política de assinatura | URI | URI versionada da política; deve ser suportada pela implementação |
| `minCertIssueDate` | inteiro Unix UTC | Data mínima de emissão dos certificados. Padrão: `1751328000` |
| `revocationPolicy` | enum | `strict`, `soft-fail` ou `warn` |
| `ocspUnknownHandling` | enum | `treat-as-revoked` ou `treat-as-warning` |
| timeouts OCSP/CRL/TSA | inteiro (segundos) | Faixa: `[5, 120]`. Padrão: `30` |
| TTL cache de revogação | inteiro (segundos) | Faixa: `[300, 86400]`. Padrão: `3600` |
| `nearExpiryThresholdDays` | inteiro (dias) | Faixa: `[1, 180]`. Padrão: `30` |
| `signatureAgeThresholdDays` | inteiro (dias) | Faixa: `[1, 1825]`. Padrão: `365` |

**Não funcionais**

- Executável em qualquer JVM 17+
- Respostas em JSON (`OperationOutcome` FHIR R4)
- Parâmetros fora de faixa ou ausentes interrompem o processamento antes de qualquer operação criptográfica

---

## Referências

- [Especificação FHIR — criar assinatura](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-criar-assinatura.html)
- [Especificação FHIR — validar assinatura](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-validar-assinatura.html)
- [Design C4](./docs/design/)
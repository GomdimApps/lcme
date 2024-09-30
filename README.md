# LCME (Linux Command and Microservices Executor)

## Como adicionar o Framework

```bash
go get github.com/GomdimApps/lcme
```

# Bash

Executa um comando Bash e retorna a saída padrão e um erro, se houver.

#### Parâmetros

- **`command`** (string): Comando Bash a ser executado.

#### Retornos

- **`string`**: Saída padrão do comando.
- **`error`**: Descrição do erro, se ocorrer.

#### Exemplo de Uso

```go
package main

import (
    "fmt"
    "github.com/GomdimApps/lcme"
)

func main() {
    comando := "rm -r teste.txt" // Comando a ser executado

    resultado, erro := lcme.Bash(comando)
    if erro != nil {
        fmt.Println("Erro:", erro) // Exibe erro, se houver
    } else {
        fmt.Println("Resultado:", resultado) // Exibe a saída do comando
    }
}
```

### Comportamento

- Retorna a saída do comando e `nil` se bem-sucedido.
- Retorna a saída e um erro descritivo se o comando falhar.

---

# ConfigRead

A função `ConfigRead` serve para carregar um arquivo de configuração (`config.conf`) e preencher a estrutura `Config` com os valores lidos. O arquivo de configuração deve seguir o formato `chave=valor`.

### Como Usar:

1. Crie um arquivo de configuração no formato adequado (ver abaixo).
2. Chame a função `ConfigRead` passando o caminho do arquivo e a estrutura `Config`.
3. A função preencherá os campos da estrutura com os valores do arquivo.

### Regras para o Arquivo de Configuração

- Cada linha do arquivo deve ter o formato `chave=valor`.
- O nome da chave deve ser exatamente igual ao nome do campo da estrutura `Config`, respeitando letras maiúsculas e minúsculas.
- Os valores devem ser compatíveis com o tipo de dado correspondente ao campo:
  - Para `bool`: Use `true` ou `false`.
  - Para `int`, `int64`, `uint64`: Use números inteiros.
  - Para `float32`, `float64`: Use números decimais (ponto `.` para separar a parte decimal).
  - Para `string`: Use qualquer sequência de texto sem espaços ao redor do valor.
  
- Comentários devem começar com o caractere `#` e serão ignorados.

### Exemplo da Estrutura `Config`

Abaixo está um exemplo de uma estrutura `Config` que pode ser usada com a função `ConfigRead`:

```go
type Config struct {
    AccessIp       bool
    MaxConnections int
    Port           int
    HostName       string
    Timeout        float64
    EnableLogs     bool
    ConnectionID   int64
    BufferSize     uint64
    ResponseTime   float32
}
```

#### Exemplo de arquivo `config.conf`:

```
# Configurações do servidor
AccessIp=true
MaxConnections=100
Port=8080
HostName=localhost
Timeout=30.5
EnableLogs=true
ConnectionID=1234567890
BufferSize=4096
ResponseTime=0.25
```

- **Chave**: Deve corresponder exatamente ao nome do campo na estrutura `Config`.
- **Valor**: Deve ser compatível com o tipo do campo (exemplo: `true` ou `false` para booleanos, números para inteiros e floats, etc.).

### Comando de Uso na `main`:

```go
package main

import (
    "fmt"
    "log"
    "github.com/GomdimApps/lcme" 
)

type Config struct {
    AccessIp       bool
    MaxConnections int
    Port           int
    HostName       string
    Timeout        float64
    EnableLogs     bool
    ConnectionID   int64
    BufferSize     uint64
    ResponseTime   float32
}

func main() {

    config := Config{}

    err := lcme.ConfigRead("config.conf", &config)
    if err != nil {
        log.Fatalf("Error loading configuration: %s", err)
    }

    fmt.Printf("AccessIp: %t\n", config.AccessIp)
    fmt.Printf("MaxConnections: %d\n", config.MaxConnections)
    fmt.Printf("Port: %d\n", config.Port)
    fmt.Printf("HostName: %s\n", config.HostName)
    fmt.Printf("Timeout: %.2f\n", config.Timeout)
    fmt.Printf("EnableLogs: %t\n", config.EnableLogs)
    fmt.Printf("ConnectionID: %d\n", config.ConnectionID)
    fmt.Printf("BufferSize: %d\n", config.BufferSize)
    fmt.Printf("ResponseTime: %.2f\n", config.ResponseTime)
}
```

# getInfoServer

---

```go
package main

import (
    "fmt"
    "github.com/GomdimApps/lcme" 
)

func main() {
    serverInfo := getInfoServer()
    if serverInfo.Error != nil {
        fmt.Printf("Erro ao obter informações do servidor: %v\n", serverInfo.Error)
        return
    }

    fmt.Printf("Memória total: %v MB\n", serverInfo.TotalMemory)
    fmt.Printf("Memória usada: %v MB\n", serverInfo.UsedMemory)
    fmt.Printf("Memória livre: %v MB\n", serverInfo.FreeMemory)
    fmt.Printf("Uso da CPU: %f%%\n", serverInfo.CPUUsage)
    fmt.Printf("Espaço total em disco: %v MB\n", serverInfo.TotalDiskSpace)
    fmt.Printf("Espaço utilizado em disco: %v MB\n", serverInfo.UsedDiskSpace)
    fmt.Printf("Espaço livre em disco: %v MB\n", serverInfo.FreeDiskSpace)
}
```



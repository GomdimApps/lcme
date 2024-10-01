# LCME (Linux Content Management Engine)

### **Propósito do Framework LCME**

O LCME é um framework projetado para facilitar a criação e gestão de conteúdo em aplicações web e aplicações de microserviços no Linux. Ele é leve e modular, permitindo que desenvolvedores integrem funcionalidades de gerenciamento de conteúdo de forma eficiente e escalável.

### **Objetivo do Framework**

O principal objetivo do LCME é fornecer uma solução simples e eficaz para o gerenciamento de conteúdo, com um foco especial em aplicações de microserviços desenvolvidas em Go. O projeto visa utilizar o mínimo de frameworks de terceiros, tornando a aplicação altamente nativa e reduzindo dependências externas.

### **Como o LCME Pode Ajudar em Projetos Go**

1. **Modularidade**: O framework é modular, o que significa que você pode escolher e integrar apenas os componentes que são necessários para o seu projeto, mantendo o sistema leve e eficiente.

2. **Escalabilidade**: Com uma arquitetura leve, o LCME permite que aplicações escalem de forma eficiente, suportando um aumento no volume de conteúdo sem comprometer o desempenho.

3. **Customização**: O LCME oferece uma alta capacidade de customização, permitindo que desenvolvedores adaptem o framework às necessidades específicas do projeto, seja na interface de usuário ou na lógica de backend.

4. **Natividade**: Ao minimizar o uso de frameworks de terceiros, o LCME garante que a aplicação permaneça altamente nativa, o que pode resultar em melhor desempenho e menor complexidade.

5. **Interação Nativa com Linux**: O LCME facilita o uso e a implementação de microserviços, trazendo maneiras nativas de interagir com o sistema operacional Linux (Server) dentro de uma aplicação, o que pode ser extremamente útil para operações de baixo nível e otimizações específicas do sistema.

6. **Documentação e Suporte**: O repositório inclui documentação detalhada e exemplos de uso, facilitando a curva de aprendizado e a implementação do framework em novos projetos.

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

A função `getInfoServer` é responsável por capturar diversas informações do sistema, como dados de distribuição Linux, memória, disco, CPU, e rede.


No seu código Go, importe o pacote `lcme` e chame a função `getInfoServer` para capturar as informações do servidor.

```go
package main

import (
	"fmt"
	"github.com/GomdimApps/lcme"
)

func main() {
	// Captura informações do servidor
	serverInfo := lcme.GetInfoServer()

	// Exibe os dados capturados
	fmt.Printf("Distribuição Linux: %s\n", serverInfo.Distribution.Name)
	fmt.Printf("Memória total: %d MB\n", serverInfo.RAM.Total)
	
}
```

---

### Tabela de Retornos

A função `getInfoServer` retorna uma estrutura. A tabela a seguir detalha os campos retornados:

| Campo                | Tipo           | Descrição                                                                           |
|----------------------|----------------|-------------------------------------------------------------------------------------|
| `Distribution.Name`   | `string`       | Nome da distribuição Linux instalada no servidor.                                   |
| `RAM.Total`           | `uint64`       | Memória RAM total em megabytes (MB).                                                |
| `RAM.Used`            | `uint64`       | Memória RAM usada em megabytes (MB).                                                |
| `RAM.Available`       | `uint64`       | Memória RAM disponível em megabytes (MB).                                           |
| `Disk.Total`          | `uint64`       | Espaço total em disco em megabytes (MB).                                            |
| `Disk.Used`           | `uint64`       | Espaço em disco utilizado em megabytes (MB).                                        |
| `Disk.Available`      | `uint64`       | Espaço em disco disponível em megabytes (MB).                                       |
| `CPU.NumCores`        | `int`          | Número total de núcleos do processador.                                             |
| `CPU.Usage`           | `float64`      | Percentual atual de uso do processador.                                             |
| `Network.IPv4`        | `[]string`     | Lista de endereços IP IPv4 associados ao servidor.                                  |
| `Network.IPv6`        | `[]string`     | Lista de endereços IP IPv6 associados ao servidor.                                  |

---

### Exemplos de Uso

#### Exemplo 1: Capturar e Exibir Informações de Rede

```go
package main

import (
	"fmt"
	"github.com/GomdimApps/lcme"
)

func main() {
	serverInfo := lcme.GetInfoServer()

	// Exibir IPs IPv4
	for _, ip := range serverInfo.Network.IPv4 {
		fmt.Println("IP da máquina (IPv4):", ip)
	}

	// Exibir IPs IPv6
	for _, ip := range serverInfo.Network.IPv6 {
		fmt.Println("IP da máquina (IPv6):", ip)
	}
}
```

#### Exemplo 2: Capturar Informações de Memória

```go
package main

import (
	"fmt"
	"github.com/GomdimApps/lcme"
)

func main() {
	serverInfo := lcme.GetInfoServer()

	fmt.Printf("Memória total: %d MB\n", serverInfo.RAM.Total)
	fmt.Printf("Memória usada: %d MB\n", serverInfo.RAM.Used)
	fmt.Printf("Memória disponível: %d MB\n", serverInfo.RAM.Available)
}
```

---


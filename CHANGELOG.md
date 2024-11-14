## Versão 1.1.6 - latest

#### Novas Funcionalidades
- **Monitoramento de Rede**: Adicionada funcionalidade para calcular as taxas de download e upload.
- **Campo GetFileInfo**: Adicionado o campo `FileExtension`, `FileData` e `FileDataBuffer` à estrutura `FileInfo` e atualizado `GetFileInfo` para incluir extensões de arquivos.

#### Refatorações
- **calculateCPUUsage**: Refatorada a função `calculateCPUUsage` para fazer a média de múltiplas amostras, melhorando a precisão.

#### Melhorias
- **Documentação**: Atualizado `README.md` para incluir a estrutura detalhada de `FileInfo` e descrições de campos.

---

## Versão 1.1.5

#### Novas Funcionalidades
- **Função GetFolderSize**: Adicionada a função 

GetFolderSize

 para calcular o tamanho de uma pasta em kilobytes.
- **Função GetFileInfo**: Adicionada a função 

GetFileInfo

 para obter informações detalhadas sobre arquivos específicos em um diretório. A função agora aceita nomes de arquivos variádicos.

#### Alterado
- **Documentação Atualizada**: Adicionada documentação para a função 

---

## Versão 1.1.4.1

### Refatoração de `configs.go`
- Aprimorada a manipulação de configurações para gerenciar adequadamente valores vazios no arquivo de configuração.

### Refatoração de `network.go`
- Melhorado o manuseio de rede para suportar endereços IPv4 e IPv6 de forma contínua, aprimorando a captura de portas e IPs tanto IPv4 quanto IPv6, proporcionando maior suporte a outras distribuições. O uso de ferramentas nativas do Linux e Go ajuda na compatibilidade com várias distribuições.

### Refatoração de `cpu.go`
- Atualizada a struct `CPUInfo` e refinado o cálculo de uso da CPU para métricas de desempenho mais precisas.

### Atualização de Versão
- Atualizada a versão do projeto para `1.1.4.1`.


### Checklist
- [x] Código refatorado para `configs.go`
- [x] Código refatorado para `network.go`
- [x] Código refatorado para `cpu.go`
- [x] Versão atualizada para `1.1.4.1`
- [x] Testes atualizados e aprovados

### Notas de Lançamento

#### Adicionado
- Suporte contínuo para endereços IPv4 e IPv6 em `network.go`.

#### Alterado
- Manipulação de configurações aprimorada em `configs.go` para gerenciar valores vazios.
- Cálculo de uso da CPU refinado em `cpu.go` para métricas de desempenho mais precisas.

#### Corrigido
- Estabilidade e compatibilidade geral do projeto melhoradas.

#### Removido
- Nenhuma funcionalidade removida nesta versão.

---

### Versão 1.1.4

#### Novas Funcionalidades
- **Captura de Distribuição Linux**: Implementada a captura detalhada de informações sobre a distribuição Linux.
- **Documentação Atualizada**: Adicionadas referências ao framework e ao logo do projeto na documentação.
- **Comentários no Código**: Inseridos comentários explicativos em todas as funções para facilitar o entendimento do código.
- **Captura de Portas de Rede**: Adicionado suporte para captura de portas TCP e UDP do sistema.

#### Correções
- **Correção de Lógica**: Corrigidos erros na lógica de processamento, resultando em maior estabilidade do sistema.



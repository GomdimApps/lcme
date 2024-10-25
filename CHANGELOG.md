## Versão 1.1.4.1

### Refatoração de `configs.go`
- Aprimorada a manipulação de configurações para gerenciar adequadamente valores vazios no arquivo de configuração.

### Refatoração de `network.go`
- Melhorado o manuseio de rede para suportar endereços IPv4 e IPv6 de forma contínua, aprimorando a captura de portas e IPs tanto IPv4 quanto IPv6, proporcionando maior suporte a outras distribuições. O uso de ferramentas nativas do Linux e Go ajuda na compatibilidade com várias distribuições.

### Refatoração de `cpu.go`
- Atualizada a struct `CPUInfo` e refinado o cálculo de uso da CPU para métricas de desempenho mais precisas.

### Atualização de Versão
- Atualizada a versão do projeto para `1.1.4.1`.

Essas mudanças visam melhorar a estabilidade e a compatibilidade geral do projeto.

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

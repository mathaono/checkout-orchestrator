# Checkout Orchestrator

## 📋 Sobre o Projeto

O Checkout Orchestrator é um sistema distribuído de processamento de pedidos baseado em eventos, projetado para alta escalabilidade e resiliência. O sistema gerencia o fluxo completo de um pedido, desde sua criação até a conclusão, passando por estados de pagamento e envio.

## 🎯 Objetivos

- Orquestrar o fluxo completo de checkout de pedidos
- Garantir consistência em um ambiente distribuído
- Implementar padrões de resiliência e alta disponibilidade
- Utilizar comunicação assíncrona baseada em eventos
- Manter rastreabilidade e observabilidade do sistema

## 🏗️ Arquitetura

### Serviços
- **order-svc**: Serviço principal de gerenciamento de pedidos (implementado)
- **payment-svc**: Serviço de processamento de pagamentos (mock)
- **shipping-svc**: Serviço de gestão de envios (mock)

### Tecnologias Utilizadas
- **Linguagem**: Go (Golang)
- **Web Framework**: Chi Router
- **Banco de Dados**: MySQL
- **Cache**: Redis
- **Mensageria**: Kafka
- **Observabilidade**: Datadog

### Eventos do Sistema
- OrderCreated
- PaymentApproved
- ShipmentRequested

## 📦 Estrutura do Projeto

```
.
├── cmd/
│   └── order-svc/         # Ponto de entrada do serviço de pedidos
├── internal/
│   └── order/            # Código interno do domínio de pedidos
└── pkg/
    └── common/           # Código compartilhado entre serviços
```

## 💻 Implementação Atual

### Modelo de Domínio (internal/order/model.go)
- Estrutura completa de Order e Item
- Sistema de estados com transições validadas
- Regras de negócio e validações
- Geração de IDs via UUID
- Cálculo automático de totais
- Validações de atualizações e mudanças de estado

### Estados do Pedido
- **CREATED**: Estado inicial após criação
- **PAID**: Pagamento aprovado
- **SHIPPED**: Pedido enviado
- **COMPLETED**: Pedido entregue com sucesso
- **FAILED**: Falha em qualquer etapa

## 🚀 Como Executar

### Pré-requisitos
- Go 1.21 ou superior
- MySQL
- Redis
- Kafka

### Configuração
1. Clone o repositório
```bash
git clone https://github.com/mathaono/checkout-orchestrator.git
cd checkout-orchestrator
```

2. Instale as dependências
```bash
go mod tidy
```

## 📝 Próximos Passos

- [ ] Implementar testes unitários
- [ ] Criar interface do repositório
- [ ] Implementar endpoints REST
- [ ] Configurar conexão com MySQL
- [ ] Implementar produção de eventos Kafka
- [ ] Adicionar circuit breaker e retry
- [ ] Implementar cache com Redis
- [ ] Configurar observabilidade
- [ ] Implementar serviços mock
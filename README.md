## CPF Generator

![Pipeline](https://github.com/dellabeneta/cpf-generator/actions/workflows/main.yaml/badge.svg)
![GHCR version](https://img.shields.io/badge/ContainerVersion-v7-blue)

Um gerador de CPF brasileiro válido, desenvolvido em Go. Gere CPFs com dígitos verificadores corretos através de uma interface web intuitiva e responsiva.

### Aviso Importante

**Este gerador é apenas para fins educacionais e de teste.** Não use CPFs gerados para fraudes ou atividades ilegais. O projeto foi criado para fins de aprendizado e desenvolvimento.

### Funcionalidades

- Geração de CPFs brasileiros válidos com dígitos verificadores corretos
- Algoritmo de validação oficial do CPF implementado
- Interface responsiva para desktop e mobile
- Exibição do CPF formatado (XXX.XXX.XXX-XX) e números puros
- Geração criptograficamente segura usando `math/rand`
- Alta disponibilidade com Kubernetes
- Containerização com Docker

### Começando

**Pré-requisitos**
- Go 1.24.3 ou superior
- Docker (opcional)
- Kubernetes/k3s (opcional)

**Instalação Local**
```bash
# Clone o repositório
git clone https://github.com/seu-usuario/cpf-generator.git
cd cpf-generator

# Execute a aplicação
go run main.go
```

A aplicação estará disponível em `http://localhost:8080`

**Usando Docker**
```bash
# Construa a imagem
docker build -t cpf-generator .

# Execute o container
docker run -p 8080:8080 cpf-generator
```

**Deploy no Kubernetes**
```bash
# Crie o namespace
kubectl apply -f k3s/namespace.yaml

# Deploy da aplicação
kubectl apply -f k3s/deployment.yaml
kubectl apply -f k3s/service.yaml
```

A aplicação estará disponível na porta `30082` do cluster.

### Tecnologias

- **Backend**: Go 1.24.3
- **Frontend**: HTML5, CSS3 (inline)
- **Container**: Docker Alpine
- **Orquestração**: Kubernetes/k3s
- **Algoritmo**: Validação oficial de CPF brasileiro

### Como Funciona

O gerador implementa o algoritmo oficial de validação de CPF:

1. **Geração dos 9 primeiros dígitos**: Números aleatórios de 0-9
2. **Cálculo do primeiro dígito verificador**: Soma ponderada dos 9 dígitos
3. **Cálculo do segundo dígito verificador**: Soma ponderada dos 10 dígitos
4. **Formatação**: Exibição no formato padrão brasileiro (XXX.XXX.XXX-XX)

### Configuração

O serviço pode ser configurado através das seguintes variáveis de ambiente:

| Variável | Descrição | Padrão |
|----------|-----------|---------|
| PORT | Porta do servidor | 8080 |

### Estrutura do Projeto
```
della@ubuntu:~/projetos/cpf-generator$ tree
.
├── cpf.png
├── Dockerfile
├── go.mod
├── k3s
│   ├── deployment.yaml
│   ├── namespace.yaml
│   └── service.yaml
├── main.go
├── nuke.sh
├── README.md
└── static
    └── favicon.ico

3 directories, 10 files
```

### Scripts Úteis

**nuke.sh**: Script para limpeza completa do Docker (containers, imagens, volumes e redes)

```bash
chmod +x nuke.sh
./nuke.sh
```

### Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
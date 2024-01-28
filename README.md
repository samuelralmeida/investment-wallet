# CARTEIRA DE INVESTIMENTO

**NADA DO ESCRITO AQUI É UMA RECOMENDAÇÃO OU PROPAGANDA. ESSE É UM PROJETO PARA MINHA GESTÃO FINANCEIRA QUE TALVEZ POSSA AJUDAR OUTRAS PESSOAS**

**SOFTWARE EM DESENVOLVIMENTO**

## INTRODUÇÃO

Eu sou cliente da [Indê Investimentos](https://indeinvestimentos.com.br/). Uma serviço de recomendação de fundos de investimento com relatórios, lives e ferramentas para gerir sua carteira de investimentos.

Resumidamente, a Luciana Seabre, fundadora da Indê, planejou uma estrutura de carteira de fundos de investimento e fornece recomedações e conhecimento para os clientes terem sua própria carteira.

Ela divide os fundos em caixas e sub-caixas:

- ESTABILIDADE:
    - BAUNILHA
    - PIMENTA

- DIVERSIFICAÇÃO:
    - VIÉS MACRO
    - EX-TESOUREIROS
    - VIÉS GLOBAL
    - INFRAESTRUTURA
    - PIMENTA

- VALORIZAÇÃO
    - QUALIDADE
    - FORA DO RADAR
    - PERFIL GLOBAL
    - VIÉS COMPRADO
    - PIMENTA

- ANTIFRAGILIDADE
    - OURO
    - DÓLAR

Essa aplicação serve para registrar os investimentos, ajuda a acompanhar os rendimentos e planejar novos investimentos.

## DETALHAMENTO

A aplicação tem três entidades básicas:

1. Fundo de investimento (fund): é um fundo em si com nome, CNPJ, banco e valor mínimo de aplicação.

- `/funds/new` -> renderiza tela para salvar no fundo

2. Investimento (invesment): cada aporte feito em um fundo é um novo investimento. Tem data e valor.

- `/investments/new` -> renderiza tela para salvar um aporte de investimento

3. Checkpoint: é uma fotografia dos investimentos de um fundo. Ou seja, quanto ele está valendo no momento. Ajuda a calcular novas alocações.

- `/checkpoints/new` -> renderiza tela para salvar um checkpoint

Além dessas entidades um resumo da carteira salva pode ser acessado em `/wallet/{nome-da-carteira}` e dados calculados de rendimento e proporção de cada caixa em `/calculate/{nome-da-carteira}`

## BANCO DE DADOS

Como a aplicação é pequena e em fase inicial é usado SQLite. Ver o arquivo `example.env` para gerar o arquivo `.env` com as variáveis de ambiente para executar a aplicação.

## EXECUÇÃO

A aplicação é feita em Golang e o backend renderiza o frontend para simplificar o desenvolvimento. O projeto usa CDN da Tailwind para recursos de CSS.

`go run main.go` -> executa a aplicação na porta 3000

`go run cmd/migrate/main.go` -> cria as tabelas no banco de dados

## A FAZER

- Melhorar design das telas do frontend
- Reaproveitar código entre os templates
- Automatizar a recomendação de alocação de investimento por box de acordo com um proporção definida
- Incluir telas e cálculos considerando as sub-caixas
- melhorar a usabilidade
- criar telas para listar fundos e investimentos
- padronizações gerais
- melhorar tratamento de erros e logs
- etc
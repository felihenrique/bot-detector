# Desafio para senior developers

Olá, dev! Estamos ansiosos para conhecer suas habilidades e avaliar seu conhecimento técnico em frontend e backend. Abaixo estão as instruções e requisitos para o teste que você deve realizar. Você terá 1 semana para concluir este teste.

## Objetivo do Teste

O objetivo deste teste é criar uma aplicação para identificação de bots. Esta aplicação possui uma camada de acompanhamento de dados on será possível ver por dia quantos bots foram identificados e um filtro (dropdown) de player.

## Instruções

- Crie uma conta no GitHub, caso ainda não tenha uma.
- Crie um novo repositório privado no GitHub (não faça fork do projeto).
- Adicione os seguintes colaboradores ao repositório: @filipealc e @welingtonsampaio.
- Crie uma branch (utilize seu nome).
- Desenvolva a aplicação seguindo os requisitos especificados abaixo.
- Documente suas decisões e progresso em um arquivo HISTORY.md.
- Atualize o arquivo README.md com instruções para instalação e uso da aplicação.
- Faça commits pequenos e bem especificados ao longo do desenvolvimento do projeto.

## Contexto

- Considere que, recebemos 66 mil requests / minuto
- Cada record precisa receber uma nova propriedade chamada `is_bot: Boolean` e ser salvo no banco de dados com esta nova informação
- O arquivo [sessions.csv.gz](https://vturb-labs.s3.amazonaws.com/challenges/sessions.csv.gz) contém 4 colunas sendo elas em ordem player_id, ip, user_agent e created_at

## Requisitos

- Desenvolver um sistema que seja capaz de receber os requests do arquivo [sessions.csv.gz](https://vturb-labs.s3.amazonaws.com/challenges/sessions.csv.gz) como input
- Desenvolver uma funcionalidade que seja capaz de hidradar os dados recebidos do input do arquivo [sessions.csv.gz](https://vturb-labs.s3.amazonaws.com/challenges/sessions.csv.gz) inserindo uma nova propriedade `is_bot: Boolean`
- Salvar o dado hidratado
- Desenvolver um sistema frontend capaz de mostrar a quantidade de bots detectados por dia com filtro por player
- Utilizar a base [dbip.mmdb](https://vturb-labs.s3.amazonaws.com/challenges/dbip.mmdb) para identificação de origem dos IPs das sessões do arquivo [sessions.csv.gz](https://vturb-labs.s3.amazonaws.com/challenges/sessions.csv.gz)

## Critérios de Avaliação

- Qualidade do código e organização do projeto.
- Documentação das decisões e progresso no arquivo HISTORY.md.
- Atualização do arquivo README.md com instruções para instalação e uso da - aplicação.
- Implementação de testes automatizados.
- Uso de HTML semântico e acessível.
- Utilização de conceitos de API RESTful.
- Escalabilidade das aplicações

# Informações extras

- O arquivo sessions.csv.gz é um dump do banco clickhouse, caso você utilize outro banco e não consiga realizar o import você pode utilizar o arquivo [sessions2.csv](https://vturb-labs.s3.amazonaws.com/challenges/sessions2.csv)
- Propositalmente este teste não especifica linguagem nem tecnologia. Você tem total liberdade de ser criativo com a solução e também possui liberdade para assumir coisas do contexto que não foram especificadas aqui.

Boa sorte! Estamos ansiosos para ver seu trabalho e avaliar suas habilidades como dev. Se você tiver alguma dúvida, não hesite em entrar em contato conosco.

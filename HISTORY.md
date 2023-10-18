# 17 de outubro de 2023

## Introdução

Esse repositório conterá a solução para o teste proposto para a vaga de senior full stack developer na vturb, cujo a descrição se encontra no arquivo DESCRICAO_TESTE.md

## Requisitos da aplicação

<p>O teste se trata de desenvolver uma API que suporte alto volume de tráfego (60k+ de requisições por minuto) e cujo o objetivo é receber informações de requisições feitas em um dado backend (ip e user_agent), e a partir delas, detectar se ela foi feita ou não por um bot. As informações precisam ser persistidas no banco de dados clickhouse.</p>
<p>Precisa ser desenvolvido também uma UI básica que lista a quantidade de bots por dia com opção de filtro pelo player.</p>

## Possíveis gargalos na aplicação

<p>
A aplicação precisará fazer várias verificações nos dados e processamentos de regexes/strings. Portanto, é mais adequado que ela seja escrita em uma linguagem compilada, para que seja possível extrair o máximo de eficiência no uso da CPU. A linguagem escolhida foi o go, mas outras opções como o rust também seriam adequadas.
</p>
<p>
Outro ponto que pode vir a ser um gargalo é o banco de dados. Em picos de requisições, as chamadas no banco de dados podem vir a se acumular e acabar causando um estouro de memória que irá crashar a aplicação. Além disso, esses picos podem deixar outras aplicações que dependem desse banco lentas.
</p>

## Eficiência nas chamadas no banco de dados
<p>
A maior parte das chamadas no banco de dados será para fazer escrita. A maneira mais eficiente de fazê-lo é através de inserções em batch. 
</p>
<p>
Segundo a documentação do clickhouse de [melhores práticas](https://clickhouse.com/docs/en/cloud/bestpractices/bulk-inserts), o ideal é que os batches tenham ao menos 10k registros e chamadas subsequentes tenham ao menos 1s de intervalo, para que o banco possa completar as inserções da chamada anterior.
</p>
<p>
Há várias estratégias para fazer as inserções em bulk no clickhouse:
<ul>
<li>
Acumulando os dados em um buffer na memória e fazendo a escrita no banco a cada 1 segundo ou quando houver x registros. É a abordagem mais eficiente. A desvantagem é que no caso um termino inesperado da aplicação os dados no buffer são perdidos.
</li>
<li>
Utilizando a [inserção assincrona](https://clickhouse.com/docs/en/integrations/go#async-insert). Uma desvantagem dessa abordagem é que o banco ainda assim estará sendo sobrecarregado de requisições.
</li>
</ul>
</p>

A abordagem escolhida foi a persistência em memória, pois é a que oferece maior escalabilidade no longo prazo, visto que dá para manter a quantidade de requisições no banco constante com o passar do tempo (apesar do tamanho do batch aumentar). Além disso, essa também é a estrategia adotada pelo clickhouse (https://clickhouse.com/docs/en/optimize/asynchronous-inserts)

# 18 de outubro de 2023

## Estrutura projeto backend
O projeto do backend terá a seguinte estrutura:
- /app: Conterá toda a lógica de negócio da aplicação
- /config: Onde se encontra arquivos de configuração e envs
- /data: Conterá toda a lógica de escrita/acesso a dados
- /http-server: Inicialização/setup do servidor http, controllers e configuração de rotas
- /utils: Funções uteis em geral

## Estratégia usada para detecção de bots
Duas abordagens serão usadas para detectar se a requisição veio de um bot:
- Verificando se o IP é de uma cloud. Requests de IPS de clouds costumam ser bots.
Foi utilizada uma lista pública de ips ranges das principais clouds: https://github.com/femueller/cloud-ip-ranges
- Checando se o user agent é usado comumente em crawlers. Foi usado como referência o artigo:
https://deviceatlas.com/blog/most-active-crawlers-list


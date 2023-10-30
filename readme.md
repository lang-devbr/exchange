## Exchange api

- Devemos retornar o preço de um determinado produto na moeda nacional (BRL), e também retornar o preço deste produto  nas moedas cadastradas no nosso sistema.
    - Se o preço for maior que R$ 500, devemos aplicar um desconto de 5%
    - Caso o preço for maior que R$ 800, devemos aplicar um desconto de 10%


- Premissas:
 - Consultar uma cotação de determinada moeda pela sua sigla na api: https://economia.awesomeapi.com.br/json/last/USD-BRL
 - Salvar as moedas e cotações em um banco de dados
 - Salvar primeiramente no sqlite
 - Alterar para salvar no postgres
 - Conteinerizar o app
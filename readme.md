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

 1 produto vão ter várias moedas 
 a moeda vai associar a um produto 
 não precisa do converter 
 api que consulta um produto 
 api que consulta todos os produtos 
 criar usecase para usar acessar a infra, vai buscar o produto 
 e vai fazer a conversão 
 
 --
 cada produto vai ter as moedas que ele vai poder converter 
 --
 
 Use case:
 - 1) Ir na camada de repositorio e buscar o produto e as moedas associadas 
 - 2) Ir na api externa de cotação e pegar as cotações de cada moeda associada do item 1. Pode colocar no currency_repository e vai ter uma func para buscar na url externa.
 - 3) Com as moedas associadas, fazer a logica na domain (entity) para o calculo de converter os preços de acordo com cada moeda
 - 4) montar o retorno pra api (por produto ou todos os produtos)
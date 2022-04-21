# golang
Este repositório foi criado para publicar os exemplos que foram construídos durante a aprendizagem da linguagem Go. 

* Linguagem igual a C, mesmos operadores &&, ||, !=
* Interessante porque identifica erros e não deixa programa continuar (mais seguro que C)
* Interessante porque não deixa executar se tem variáveis sem uso
* É totalmente case sensitive, tanto no código quanto nas letras, ‘a’ <> ‘A’
* Utilizada em cenários tecnológicos recentes, incluindo:
> 1. Cloud applications
> 2. Web development
> 3. Database implementations
> 4. Distributed networking services
> 5. Utilities
> 6. IoT devices

## Setup
Utilizei o Visual Studio Code com o pacote Go:

* Instalar o Visual Studio Code
* Instalar Go do site https://go.dev/
* Instalar o pacote Go no VSC

## Onde estudar
Oficialmente é aqui: https://go.dev/learn/, mas existem vários vídeos no Youtube sobre o assunto, comece com os beginners.

## Exemplos:
Nos arquivos de **exemplo** incluí todos os testes executados (em ordem sequencial), e para não virar uma bagunça, fui comentando cada teste separadamente. Se você executar não vai aparecer nada, pois será necessário descomentar o bloco desejado e em seguida executar o **go run.** no terminal do VSC.


[Exemplo1](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo1) Cobre os assuntos básicos a respeito da linguagem, operadores, variáveis, for, if, switch, arrays, slices, package fmt(Scan, Print), package strings(Split, Contains, Map). 

[Exemplo2](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo2) Cobre o assunto functions com vários exemplos de criação, passagem de parâmetro, retorno, variáveis globais.

[Exemplo3](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo3) Cobre o assunto divisão do main em packages, porém ainda dentro do mesmo pacote main.

[Exemplo4](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo4) Cobre o assunto de criação de packages separados em pastas, arquivos e comando import. Cenário ideal para projetos mais complexos.

[Exemplo5](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo5) Cobre o tema **map** para a criação de arrays baseados em chaves ao invés de indices. O exemplo cria um tipo map e uma lista de maps.

[Exemplo6](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo6) Cobre o tema de criação de tipos estruturados com struct. Muito interessante.

## O que ficou faltando para ser TOP
Como desenvolvedor Genexus, alguns recursos automatizados são importantes e ainda não foram identificados na linguagem, destacamos:

* **domains**: criação de nomes para tipos específicos
* **toJson e fromJson**: para parsear um tipo em formato JSON

## O que ficou mais chato
Go realiza suas ações baseadas em funções, assim como ocorre na linguagem C. Programadores de Java, C, C# e outras linguagens estão acostumados a isso. Com linguagem mais focada em lowcode como Genexus, a coisa pega um pouco porque tem muito código para se produzir algo bem simples.

*  **numeric.toString()** ficou mais chato com **strconv.FormatUint(imot64(userTicket), 10)**, é muita coisa pra fazer para se converter um número em texto

## Documentação adicional
* [Google Docs](https://docs.google.com/document/d/1d5CogFKYcD7gxHnzGoZ2b_WpSF0DbXdPbHtGTg9XJj0/edit?usp=sharing) Doc explicativo dos exemplos este Git.
* [golang.org](https://golang.org)
* [go.dev](https://go.dev/)
* [Learning Go Language in 7 hours](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1153s)



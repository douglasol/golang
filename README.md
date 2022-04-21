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

* Instalar o Visual Studio Code (https://visualstudio.microsoft.com/pt-br/downloads/)
* Instalar Go baixando do site https://go.dev/
* Instalar o pacote Go no VSC
* Instalar o pacote GitHub no VSC para publicar seus testes aqui nesse projeto.

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

# Boas praticas
Talvez isso se transforme em um capítulo a parte, pois muita coisa deve ser definida para que se produza um bom guia de programação.

## uint
Go possui vários tipos de unsigned int (uint)
``
uint
uint8       0-255
uint16      0-65535
uint32      0-4294967295
````

## var
Declaração de variáveis:
```
var nome string
var sobrenome string
var idade uint
```

fica melhor com:

```
var (
    nome string
    sobrenome string
    idade uint
)
```


## Documentação adicional
* [Google Docs](https://docs.google.com/document/d/1d5CogFKYcD7gxHnzGoZ2b_WpSF0DbXdPbHtGTg9XJj0/edit?usp=sharing) Doc explicativo dos exemplos este Git.
* [golang.org](https://golang.org)
* [go.dev](https://go.dev/)
* [Learning Go Language in 7 hours](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1153s)

## Considerações de um desenvolvedor lowcode chato
Desenvolver em um ambiente lowcode, como Genexus por exemplo, faz com que fiquemos um pouco mal acostumados com o fato de ter que escrever muito para obter algum resultado. 

Quando aprendemos uma linguagem nova tentamos encontrar aquilo que nos faz mais falta. Senti falta de alguns recursos que automatizam o código, e auxiliam na padronização da programação. 

* **domains**: criação de nomes para tipos específicos
* **toJson e fromJson**: para parsear um tipo em formato JSON

O que ficou faltando para ser TOP

Go realiza suas ações baseadas em funções, assim como ocorre na linguagem C. Programadores de Java, C, C# e outras linguagens estão acostumados a isso, mas novamente, quem se acostumou num cenário mais lowcode prefere escrever menos e obter diretamente o resultado.
Até entendo Go na parte dos tipos de dados, visto que é uma linguagem fortemente tipada que objetiva otimizar o uso da memória, mas isso se traduz em muito código.

* (gx) ***numeric.toString()** ficou mais chato em Go: **strconv.FormatUint(imot64(userTicket), 10)**, é muita coisa pra fazer para se converter um número em texto

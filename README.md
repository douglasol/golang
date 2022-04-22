Este repositório foi criado para publicar os testes, a partir de um primeiro contato com a linguagem GO. O objetivo era entender a programação nesta linguagem, e acabou se transformando em um passo-a-passo ordenado com os temas mais relevantes, sob o meu ponto de vista, na programação básica com a linguagem, e que permite compreender um pouquinho dos recursos. Não é o guia definitivo de Go, e nem pretende ser, mas foi interessante para que eu pudesse compreender um pouco a respeito da estrutura de programação, execução e da dinâmica com Go.

Aprender Go é outra coisa um pouco mais séria, pois vai exigir uma abordagem mais prática em um projeto real, avançando sobre os temas encontrados no dia-a-dia de trabalho. Quem sabe sai alguma coisa assim por aqui.

Go é extenso e cheio de recursos, os pacotes são imensos e a linguagem muito interessante. Para um tutorial mais sério, falta muita coisa para ser abordado, e mesmo nos pacotes que me deparei neste primeiro momento, muita coisa ficou pra tras.

Caso queira participar, fique a vontade para publicar seus testes, seja benvindo.

## O Que é Go?

* Criada na Google para substituir a programação Python, C/C++ e Java, por uma linguagem mais segura, concorrente e com compilação ágil
* Linguagem igual a C/Java, com os mesmos operadores &&, ||, !=, incluindo os operadores binários ! | & ^ << >>
* Fortemente tipada, incluindo tipos complexos complex64, complex128, para programação científica
* Muito interessante porque identifica os erros no código durante a edição e não deixa programa continuar (mais seguro que C)
* Interessante porque não deixa executar o prorama se existem variáveis sem uso, fornçando o desenvolvedor a ser organizado, rs
* É totalmente case sensitive, tanto no código quanto nas letras, 'a' é diferente de 'A'
* Utilizada em cenários tecnológicos atuais (2022) incluindo: 
> 1. Cloud applications
> 2. Web development
> 3. Database implementations
> 4. Distributed networking services
> 5. Utilities
> 6. IoT devices

Como tudo hoje em dia se refere a nuvem, convém uma linguagem que explore todos os recursos deste ambiente.

## Setup
Utilizei o Microsoft Visual Studio Code com o pacote Go para começar esta brincadeira. Praticamente downloads e cliques simples para começar a mexer com o trem.

>* Instalar o Visual Studio Code (https://visualstudio.microsoft.com/pt-br/downloads/)
>* Instalar Go baixando do site https://go.dev/
>* Instalar o pacote Go no VSC
>* Instalar o pacote GitHub no VSC para publicar seus testes aqui nesse projeto.

Devem existir outras IDEs interessantes, me perdoem a simplicidade deste primeiro cenário.

## Onde estudar
Oficialmente é aqui: https://go.dev/learn/, mas existem vários vídeos no Youtube sobre o assunto, comece com os beginners. Os vídeos ajudam a dar uma sequência no estudo, porém, sempre fica aquela ansiedade de ver em Go, aquilo que já sabemos em outra linguagem. No meu caso tive que conter a ansiedade até chegar ao tema *struct*, para abandonar o vídeo e começar a quebrar a cabeça com Database, Coleções, Json, e por ai vai.

Por isso nem vou recomendar algum para você começar, pois vai depender do seu ritmo e interesse. Entre no Youtube e digite 'GoLang Course' e vai aparecer um montão para voce escolher.

## Exemplos:
É uma pasta na qual criei várias outras, em ordem sequencial, contendo os temas específicos. Do mais básico e simples para os temas mais cabeludos.

Nos arquivos de **exemplo** incluí todos os testes executados, e para não virar uma bagunça, fui comentando, em um mesmo arquivo os testes separadamente. Isso nos primeiros exemplos. Depois fui criando arquivos .go em separado, devido a complexidade dos temas. Se ao executar não  aparecer nada, será necessário descomentar o bloco desejado e em seguida executar o **go run.** no terminal do VSC.

Algumas pastas incluem mais de um arquivo go, neste caso para executar o programa no termilan será necessário digitar **go run nome_programa.go**

Para os temas de SQL Server será necessário incluir pacotes externos ao Go, tudo está documentado aqui: [Laboratório.go](https://docs.google.com/document/d/1d5CogFKYcD7gxHnzGoZ2b_WpSF0DbXdPbHtGTg9XJj0/edit?usp=sharing)

[Exemplo01](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo01) Cobre os assuntos básicos a respeito da linguagem, operadores, variáveis, for, if, switch, arrays, slices, package fmt(Scan, Print), package strings(Split, Contains, Map). 

[Exemplo02](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo02) Cobre o assunto functions com vários exemplos de criação, passagem de parâmetro, retorno, variáveis globais.

[Exemplo03](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo03) Cobre o assunto divisão do main em packages, porém ainda dentro do mesmo pacote main.

[Exemplo04](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo04) Cobre o assunto de criação de packages separados em pastas, arquivos e comando import. Cenário ideal para projetos mais complexos.

[Exemplo05](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo05) Cobre o tema **map** para a criação de arrays baseados em chaves ao invés de indices. O exemplo cria um tipo map e uma lista de maps.

[Exemplo06](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo06) Explora o tema de criação de tipos estruturados com struct em vários exemplos:
* (main.go) Conceitos básicos de struct
* (colecao.go) Um segundo programa foi incluido para mostrar a operação de coleção de structs e sua carga por variáveis. Cada coleção deverá implementar seu método add(item), pois Go não automatiza isso para o tipo struct. 
* (colecaocarga) Mostra como se cria um struct simples, Pessoas, e em seguida uma variável que armazena uma coleção neste tipo  pessoas := []Pessoas{}, e na própria definição desta, se carregar os dados diretamente. 
* (colecaoforin.go) Percorre uma coleção (conceito for in)

[Exemplo07](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo07) Trata do assunto de manipulação de strings

[Exemplo08](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo08) A conexão com o banco de dados Microsoft SQL Server é vista neste exemplo. A operação é um simples select em uma tabela, mas já é possível verificar que a programação segue um padrão simples, mas não menos trabalhoso para o desenvolvedor. Siga os passos da Microsoft (https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go) para instalar os pacotes necessários para rodar o exemplo.

[Exemplo09](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo09) JSON é o assunto. 
* (main.go) apresenta um exemplo simples que pega a estrutura criada e a converte em JSON. 
* (colorjson.go) utiliza um formatador de JSON fornecido por um desenvolvedor para apresentar o JSON.
* (dbjson.go) Já avança criando JSON a partir do select na tabela de Pessoas.

[Exemplo10](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo10) Web application

[Exemplo11](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo11) Files

## Documentação adicional
* [Laboratório.go](https://docs.google.com/document/d/1d5CogFKYcD7gxHnzGoZ2b_WpSF0DbXdPbHtGTg9XJj0/edit?usp=sharing) Doc explicativo dos exemplos este Git.
* [Go Language Specification](https://go.dev/ref/spec)
* [golang.org](https://golang.org) or [go.dev](https://go.dev/)
* [Learning Go Language in 7 hours](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1153s)

# Boas praticas
Talvez isso se transforme em um capítulo a parte, pois muita coisa deve ser definida para que se produza um bom guia de programação.

## tipos
Definição de tipos em linguagens do tipo Go não é tarefa simples. O cenário de uso, com as exceções deve ser pensada antes de se definir o tamanho do número que deverá ser armazenado. Em caso de overflow muita coisa pode acontecer, como um prejuízo de U$ 137 milhões ([Curiosidade: space error](https://hownot2code.com/2016/09/02/a-space-error-370-million-for-an-integer-overflow/))

Definir se será um inteiro, ou inteiro sem sinal, ou real, string, ..., é parte do problema. O tamanho do valor a ser armazenado é outro.

```
* uint8 or byte     0-255
* uint16            0-65535
* uint32 or rune    0-4294967295
* string
* float32 
* float64
* complex64 
* complex128
````
* [Tipos básicos](https://go.dev/ref/spec#Types)

## const
Porque constantes? Os valores são definidos no momento da compilação e não execução, são imutáveis, representam valores fixos, melhora a codificação do sistema.
```
const CONFIG_HTTPPORT int = 80
```

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

## strings
Uma string já é decomposta em caracteres automaticamente. Aspas ("") são usadas para definir strings. Apóstrofe um único caracter ('')
```
texto := "Isso é uma string"
fmt.Printf("%v %v %T", string(texto[2], texto[2], texto[2]
```

> result: s 115 uint8
> na posição [2] da string temos uma letra 's', equivalente ao valor 115 do tipo uint8, representação numérica do caracter. Porém um caracter em Go é definido como uint32 (?), por causa do UTF32


## Considerações de um desenvolvedor lowcode chato
Desenvolver em um ambiente lowcode, como Genexus por exemplo, faz com que fiquemos um pouco mal acostumados com o fato de ter que escrever muito para obter algum resultado. 

Quando aprendemos uma linguagem nova tentamos encontrar aquilo que nos faz mais falta. Senti falta de alguns recursos que automatizam o código, e auxiliam na padronização da programação. 

* **domains**: criação de nomes para tipos específicos
* **toJson e fromJson**: para parsear um tipo em formato JSON
* **concorrência** no acesso a operações no banco LUW, sessão BD

O que ficou faltando para ser TOP

Go realiza suas ações baseadas em funções, assim como ocorre na linguagem C. Programadores de Java, C, C# e outras linguagens estão acostumados a isso, mas novamente, quem se acostumou num cenário mais lowcode prefere escrever menos e obter diretamente o resultado.
Até entendo Go na parte dos tipos de dados, visto que é uma linguagem fortemente tipada que objetiva otimizar o uso da memória, mas isso se traduz em muito código.

* (gx) ***numeric.toString()** ficou mais chato em Go: **strconv.FormatUint(imot64(userTicket), 10)**, é muita coisa pra fazer para se converter um número em texto

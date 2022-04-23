Este repositório foi criado para publicar os programas em Go, resultados de um primeiro contato com a linguagem. O objetivo era desenvolver testes para compreender a programação nesta linguagem, e acabou se transformando em um passo-a-passo ordenado com os temas mais relevantes, sob o meu ponto de vista, na programação básica com esta linguagem. Não é o guia definitivo de Go, e nem pretende ser, mas foi interessante para que eu pudesse compreender um pouco a respeito da estrutura de programação, execução e da dinâmica com Go, seus pontos fortes e também, alguns inconvenientes.

Aprender Go é algo um pouco mais sério, pois vai exigir uma abordagem mais prática em um projeto real, avançando sobre os temas encontrados no dia-a-dia de trabalho. Quem sabe sai alguma coisa assim por aqui. Coisa que se alcança fazendo.

A linguagem Go é extensa e cheia de recursos, os pacotes são completos e a sintaxe muito interessante. Para um tutorial mais sério, falta muita coisa para ser abordado, e mesmo nos pacotes que me deparei neste primeiro momento, muita coisa ficou pra tras.

Caso queira participar, seja benvindo, fique a vontade para publicar seus testes.

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
Oficialmente é aqui: https://go.dev/learn/, mas existem vários vídeos no Youtube sobre o assunto, comece com os beginners. Os vídeos ajudam a dar uma sequência no estudo, porém, sempre fica aquela ansiedade de ver em Go, aquilo que já sabemos em outra linguagem. No meu caso, consegui me conter até chegar ao tema *struct*, para abandonar o vídeo e começar a quebrar a cabeça com Database, Coleções, Json, e por ai vai.

Por isso nem vou recomendar algum para você começar, pois vai depender do seu ritmo e interesse. Entre no Youtube e digite 'GoLang Course' e vai aparecer um montão para voce escolher.

## Exemplos:
É uma pasta na qual criei várias outras, em ordem sequencial, contendo os temas específicos. Do mais básico e simples para os temas mais cabeludos.

Nos arquivos de **exemplo** incluí todos os testes executados, e para não virar uma bagunça, fui comentando, em um mesmo arquivo os testes separadamente. Isso nos primeiros exemplos. Depois fui criando arquivos .go em separado, devido a complexidade dos temas. Se ao executar não  aparecer nada, será necessário descomentar o bloco desejado e em seguida executar o **go run.** no terminal do VSC.

Algumas pastas incluem mais de um arquivo go, neste caso para executar o programa no termilan será necessário digitar **go run nome_programa.go**

Para os temas de SQL Server será necessário incluir pacotes externos ao Go, tudo está documentado aqui: [Laboratório.go](https://docs.google.com/document/d/1d5CogFKYcD7gxHnzGoZ2b_WpSF0DbXdPbHtGTg9XJj0/edit?usp=sharing)

[Exemplo01](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo01) Cobre os assuntos básicos a respeito da linguagem, operadores, variáveis, for, if, switch, arrays, slices, package fmt(Scan, Print), package strings(Split, Contains, Map). 

[Exemplo02](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo02) Functions: Cobre o assunto functions com vários exemplos de criação, passagem de parâmetro, retorno, variáveis globais.

[Exemplo03](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo03) Packages: Cobre o assunto divisão do main em packages, porém ainda dentro do mesmo pacote main.

[Exemplo04](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo04) Packages: Cobre o assunto de criação de packages separados em pastas, arquivos e comando import. Cenário ideal para projetos mais complexos.

[Exemplo05](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo05) Map: Cobre o tema **map** para a criação de arrays baseados em chaves ao invés de indices. O exemplo cria um tipo map e uma lista de maps.

[Exemplo06](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo06) Structs: Explora o tema de criação de tipos estruturados com struct em vários exemplos:
* (main.go) Conceitos básicos de struct
* (structreview.go) Apresenta todos os recursos possiveis de criação e inicialização de struct com Go, com destaque o Method que permite associar uma ou mais funções em uma estrutura. Praticamente transforma a estrutura em uma classe.
* (structadd.go e structadd2.go) Um segundo programa foi incluido para mostrar a operação de coleção de structs e sua carga por variáveis. Cada coleção deverá implementar seu método add(item), pois Go não automatiza isso para o tipo struct. 
* (structcarga) Mostra como se cria um struct simples, Pessoas, e em seguida uma variável que armazena uma coleção neste tipo  pessoas := []Pessoas{}, e na própria definição desta, se carregar os dados diretamente. 
* (structforin.go) Percorre uma coleção (conceito for in)
* (slicemethodsteste.go) Mostra as funções programadas no package **slicemethods**, para operação de coleções com slices.

[Exemplo07](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo07) Strings: Trata do assunto de manipulação de strings

[Exemplo08](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo08) Database: A conexão com o banco de dados Microsoft SQL Server é vista neste exemplo. A operação é um simples select em uma tabela, mas já é possível verificar que a programação segue um padrão simples, mas não menos trabalhoso para o desenvolvedor. Siga os passos da Microsoft (https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go) para instalar os pacotes necessários para rodar o exemplo.

[Exemplo09](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo09) JSON: é o assunto. 
* (main.go) apresenta um exemplo simples que pega a estrutura criada e a converte em JSON. 
* (jsoncolor.go) utiliza um formatador de JSON fornecido por um desenvolvedor terceiro para apresentar o JSON formatdado.
* (jsondb.go) Já avança criando JSON a partir do select na tabela de Pessoas.
* (jsonmarshal.go) Exemplo do uso do método Marshal para formatação de JSON

[Exemplo10](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo10) Web application:

[Exemplo11](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo11) Files:

[Exemplo12](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo11) Orientação a Objetos: Interfaces
* (interface.go) Apresenta o conceito de interfaces, sendo um método Show() usado em dois structs distintos (Artigo e Livro)
* (interfacedupla.go) Apresenta a comparação de Area de circulo e retângulo, mostrando chamada e retorno.


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
Não foi nada complicado entender Go, pois o meu background é de um professor que ensinou por muito tempo C, C++, Java para meus aluninhos. Todos sofriam muito em C devio a tipagem da linguagem, e recebi com surpresa e muito interesse como Go resolveu isso, ficou muito mais simples. Porém, inclui também no meu portfólio um modelo de programação low-code com Genexus (programadores em geral criticam muito, rs), mas encontrei muito mérito neste recurso. O problema é que em cenários low-code se busca programar muito pouco para se obter muito resultado, e isso faz com que fiquemos um pouco mal acostumados. E ao retornar a programação hardcode e ter que se escrever muito para obter algum resultado, fica aquela pulga nos dizendo que talvez seja melhor buscar um cenário mais automatizado.

Não identifiquei ainda uma ferramenta que faça isso com Go, me avise se houver alguma.

Aqui uma listinha do que eu acho que está faltando:

* não vou incluir normalização de tabelas, queries escritas automaticamente conforme se escolhe o que se quer na tabela, porque ai já seria algo, acredito, 'impossível' em uma programação hardcode, mas se tiver alguma coisa nesse sentido, vou ficar muito feliz
* **domains**: criação de nomes para tipos específicos para facilitar a manutenção do sistema como um todo. Por exemplo, se eu defino um nome como Varchar(128), gostaria que todos os programas pudessem ser corrigidos, caso queira alterar para Varchar(100).
* **toJson e fromJson**: para parsear um tipo em formato JSON, acho que já consegui resolver.
* **concorrência** na operação de gravação de atualizações de dados no banco, é importante manter a integridade dos registros, quando duas ou mais pessoas operam sobre um mesmo registro ao mesmo tempo.
* operaçoes de conversão de tipo são um pouco complexas em Go, **numeric.toString()** ficou mais chato: **strconv.FormatUint(imot64(userTicket), 10)**, é muita coisa pra fazer para se converter um número em texto, mas acho que isso se resolve.
* cada linguagem possui seu nível de chatice, Go já me mostrou uma, tipo em um struct a primeira letra do campo em maiúsculo ou minúsculo faz uma grande diferença (observei isso na conversão Json). Não sei se existe alguma razão miraculosa para isso, ainda não descobri.


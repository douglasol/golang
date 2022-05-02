## GO: GUIA DO CURIOSO

Seja benvindo ao repositório de exemplos e programas na linguagem Go, do livro GO: GUIA DO CURIOSO.

O livro é gratuito e poderá ser baixado aqui: 

* Versão draft: [ACESSE AQUI](https://drive.google.com/file/d/15TBj6SLhg2Rtqy9UUy2f5RZtiTvtXaQT/view?usp=sharing). Esta versão está sujeita a modificações.

O objetivo deste repositório é disponibilizar para os interessados em aprender esta linguagem os programas testes e exemplos, relacionados aos temas que considero relevantes para um programador.

A linguagem Go é extensa e cheia de recursos e merece alguma atenção, os pacotes são completos e extensos, e a sintaxe muito interessante. Para um tutorial mais sério, falta muita coisa aqui para ser abordado, e mesmo nos pacotes que me deparei neste primeiro momento, muita coisa ficou para trás.

## O Que é Go?

* A linguagem foi criada na Google para substituir a programação Python, C/C++ e Java, por uma linguagem mais segura, concorrente e com compilação ágil
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

Como tudo hoje em dia se refere a nuvem, convém conhecer uma linguagem que explora todos os recursos deste ambiente e produz programas enxutos, otimizados e de baixo custo operacional.

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

Algumas pastas incluem mais de um arquivo go, neste caso para executar o programa no termilan será necessário digitar **go run nome_programa.go**

[01_Basico](https://github.com/douglasol/golang/tree/main/Exemplos/01_Basico) Cobre os assuntos básicos a respeito da linguagem, operadores, variáveis, for, if, switch, arrays, slices, package fmt(Scan, Print), package strings(Split, Contains, Map). 
* array.go: exemplo simples de array (o tema vai ser explorado futuramente em um capítulo a parte)
* concatstring.go: concatenação de strings com +
* for.go: laços for, break, if e outros recursos de loop
* prints.go: tipos de Prints do Go
* scans.go: entrada de dados no Terminal
* regex.go: validação de entrada e valores com expressão regular (não é nada basico)
* slice.go: conceito de slice e exemplos simples 

[02_Functions](https://github.com/douglasol/golang/tree/main/Exemplos/02_Functions) Cobre o assunto functions com vários exemplos de criação, passagem de parâmetro, retorno, variáveis globais.
* alomamae.go: dispensa comentários
* global.go: variáveis globais em funções
* recebendo.go: passagem e retorno de array em função
* recebendovarios.go: mostra a passagem e retorno de muitos valores na mesma função
* funcoes.go: basico de funções

[03_Package](https://github.com/douglasol/golang/tree/main/Exemplos/03_Package) Cobre o assunto divisão do main em packages, dentro e fora do pacote main.

[04_Map](https://github.com/douglasol/golang/tree/main/Exemplos/04_Map) Tema **map** para a criação de arrays baseados em chaves ao invés de indices. O exemplo cria um tipo map e uma lista de maps.

[05_Slice](https://github.com/douglasol/golang/tree/main/Exemplos/05_Slice) Slices

[06_Struct](https://github.com/douglasol/golang/tree/main/Exemplos/06_Struct) Explora o tema de criação de tipos estruturados com struct em vários exemplos:
* (main.go) Conceitos básicos de struct
* (structreview.go) Apresenta todos os recursos possiveis de criação e inicialização de struct com Go, com destaque o Method que permite associar uma ou mais funções em uma estrutura. Praticamente transforma a estrutura em uma classe.
* (structadd.go e structadd2.go) Um segundo programa foi incluido para mostrar a operação de coleção de structs e sua carga por variáveis. Cada coleção deverá implementar seu método add(item), pois Go não automatiza isso para o tipo struct. 
* (structcarga) Mostra como se cria um struct simples, Pessoas, e em seguida uma variável que armazena uma coleção neste tipo  pessoas := []Pessoas{}, e na própria definição desta, se carregar os dados diretamente. 
* (structforin.go) Percorre uma coleção (conceito for in)
* (slicemethodsteste.go) Mostra as funções programadas no package **slicemethods**, para operação de coleções com slices.

[07_String](https://github.com/douglasol/golang/tree/main/Exemplos/07_String) Trata do assunto de manipulação de strings

[Exemplo08](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo08) Database: A conexão com o banco de dados Microsoft SQL Server é vista neste exemplo. A operação é um simples select em uma tabela, mas já é possível verificar que a programação segue um padrão simples, mas não menos trabalhoso para o desenvolvedor. Siga os passos da Microsoft (https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go) para instalar os pacotes necessários para rodar o exemplo.

[Exemplo09](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo09) JSON: é o assunto. 
* (main.go) apresenta um exemplo simples que pega a estrutura criada e a converte em JSON. 
* (jsoncolor.go) utiliza um formatador de JSON fornecido por um desenvolvedor terceiro para apresentar o JSON formatdado.
* (jsondb.go) Já avança criando JSON a partir do select na tabela de Pessoas.
* (jsonmarshal.go) Exemplo do uso do método Marshal para formatação de JSON

[Exemplo10](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo10) Web application:

[Exemplo11](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo11) Files: Criação e leitura de arquivo texto no diretório.

[Exemplo12](https://github.com/douglasol/golang/tree/main/Exemplos/Exemplo11) Orientação a Objetos: Interfaces
* (interface.go) Apresenta o conceito de interfaces, sendo um método Show() usado em dois structs distintos (Artigo e Livro)
* (interfacedupla.go) Apresenta a comparação de Area de circulo e retângulo, mostrando chamada e retorno.


## Documentação adicional
* [Go Language Specification](https://go.dev/ref/spec)
* [golang.org](https://golang.org) or [go.dev](https://go.dev/)

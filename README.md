# capital-gains
Desafio nubank sobre o tema Ganho de Capitais.

## Executando o código pela IDE

* Extraia o projeto compactado para o seu repositório local;
* Abra o projeto na IDE de sua escolha. Recomendo para GO estas: vs-code ou goland;

### Comandos úteis

#### Assumindo que você irá rodar com o GO previamente instalado na sua maquina:

``` bash
# baixando as dependencias para a vendor local e buildando o projeto
$ make build
```

``` bash
# testando
$ make tests
```

``` bash
# verificando a cobertura de testes
$ make coverage
```

``` bash
# baixando e atualizando dependencias
$ make deps
```

``` bash
# executando o projeto
$ make run
```

#### Assumindo que você irá rodar a partir do docker:

``` bash
# baixando as dependencias para a vendor local e buildando o projeto
$ make docker-build
```

``` bash
# executando o projeto
$ make docker-run
```

## Executando o código diretamente pelo executável

* O go é uma linguagem fortemente tipada e seu build é a nível de S.O;
* Gerei o executável `capital-gains.sh` após compilar o projeto;
* O executável `.sh` pode ser aberto em S.Os Linux e Mac.

## Decisões técnicas
``` bash
* Escolhi o go como linguagem de programação pelo meu skill na linguagem, pela facilidade de implementação que a linguagem permite, por sua robustes e tipagem forte;
* O projeto foi estruturado levando-se em consideração boas práticas da linguagem go;
```
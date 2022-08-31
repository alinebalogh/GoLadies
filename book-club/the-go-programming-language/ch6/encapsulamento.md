# 6.6 Encapsulamento

- [6.6 Encapsulamento](#66-encapsulamento)
  - [O que é?](#o-que-é)
  - [Quando usar?](#quando-usar)
  - [Como usar?](#como-usar)
  - [Alguns pontos para observar](#alguns-pontos-para-observar)


## O que é?
Padrão orientado a objetos: do jeito Go

Vamos ver como Go aplica um dos pilares da programação orientada a objetos, que são:encapsulamento, herança e polimorfismo. 

Vamos falar de *encapsulamento* hoje. Esse padrão é característica de linguagens baseadas em classe, as quais estão nas linguagens orientadas a objeto mais populares.

Go encapsula coisas a nível de pacote. Nomes que começam com letra minúscula são visíveis, apenas, dentro do pacote. Podemos definir nossos métodos como privados dentro de um pacote e expor apenas interfaces ou alguma função.

O mesmo sistema de ocultação vale para campos de uma struc ou aos métodos de um tipo. Já para encapsular um objeto, devemos transformá-lo em struct.

## Quando usar?
Para saber se devemos encapsular ou não é possível começar pensando se quero que cliets de outro pacote tenham o poder de alterar/ver meu meodo ou campo ou objeto, se sim, vou deixar público, se não, deixo privado.

## Como usar?

Um exemplo:
``` 
    type IntSet struct {
        words []uint64
    }
 ```
 Meu IntSet poderia ser apenas uma fatia:
 ``` 
    type IntSet []uint64    
 ```
 
 Mas eu não quero que um clients de outros pacotes possas alterar ou ler meu IntSet, por este motivo, declaramos como uma estrutura.

 ## Alguns pontos para observar
 1. Uma consequência desse sistema baseado em nomes é que a unidade de encapsulamento é o pacotee não o tipo, como em outros linguagens. E assim os campos de uma estrutura são visiveis para todos que estiverem no pacote.


2. O encapsulamento oferece três vantagens:
   
   1° lugar: Como os clients não podem modificar diretamente as variáveis do objeto, não precisamos nos preocupar com inspeções para mapear todos os possíveis valores da variáveis.
   
   2° lugar: Ocultar detalhes de implementação evita que os clients dependam de aspectos que podem mudar, e assim há mais liberdade para evoluir a implementação, sem quebrar compatibilidade.

   3° lugar: Evita alteração arbitrária pelos clients (regra mais importante), como as vars só podem ser alteradas pelo pacote é mais fácil garantir que nada saia do controle.



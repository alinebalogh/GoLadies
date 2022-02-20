**Table of contents**

- [4.2 Fatias (Slices)](#42-fatias-slices)
  - [O que são?](#o-que-são)
  - [Arrays e fatias, intimamente conectados](#arrays-e-fatias-intimamente-conectados)
  - [Como criar fatias a partir de um array já existente](#como-criar-fatias-a-partir-de-um-array-já-existente)
    - [Posso Estender uma fatia já existente?](#posso-estender-uma-fatia-já-existente)
    - [O que acontece se estender além da capacidade :thinking: ?](#o-que-acontece-se-estender-além-da-capacidade-thinking-)
  - [Como criar uma fatia sem explicitamente associar a um array](#como-criar-uma-fatia-sem-explicitamente-associar-a-um-array)
    - [Fatia literal](#fatia-literal)
    - [Make](#make)
  - [Particularidades das fatias](#particularidades-das-fatias)
  - [Append. O que seria das fatias sem a função append :relieved: ?](#append-o-que-seria-das-fatias-sem-a-função-append-relieved-)

# 4.2 Fatias (Slices)

<img src="https://golangforall.com/assets/gophslice.svg" alt="gopher slicing a beef" width="200"/>

[Créditos](https://golangforall.com)

## O que são?

- Sequências de tamanho variáveis de um tipo específico.

## Arrays e fatias, intimamente conectados

Uma fatia é uma estrutura de dados *leve* que da acesso a uma subsequência de elementos de um array (array subjacente).

**Componentes de uma fatia**:

- Ponteiro -> Aponta para o primeiro elemento do array 
- Tamanho -> Número de elementos da fatia
- Capacidade -> Número de elementos entre o inicio da fatia e o ultimo elemento do array subjacente.


## Como criar fatias a partir de um array já existente

```golang
    months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]
```

<img src="https://miro.medium.com/max/300/0*9R0AwP_71_bFrnHQ.png" alt="slice"/>

### Posso Estender uma fatia já existente?

```golang
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	summer := months[6:9]
	endlessSummer := summer[:5]

	fmt.Println(endlessSummer)
```

### O que acontece se estender além da capacidade :thinking: ?

```golang
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	summer := months[6:9]
    fmt.Println(summer[:20])
```

## Como criar uma fatia sem explicitamente associar a um array

### Fatia literal

```golang
    s := []string{"a", "a", "a", "4", "5", "6"}
	fmt.Println(s, len(s), cap(s))
```

### Make
Na função make devemos informar o tipo da fatia, o tamanho da fatia. A capacidade é opcional, caso for omitida sera igual ao tamanho.

```golang
    // omitindo capacidade
    s := make([]int, 20)
	fmt.Println(s, cap(s))

    // declarando capacidade
	sl := make([]int, 20, 30)
	fmt.Println(sl, cap(sl))
```

## Particularidades das fatias

- Valor zero é `nil`. E se a fatia é `nil` seu tamanho e capacidade são iguais a zero.

- Caso precise testar se uma fatia esta vazia, não use `nil`. Compare o resultado de `len` a zero.

```golang
    var s []int
	fmt.Println(len(s), s == nil)
	s = nil
	fmt.Println(len(s), s == nil)
	s = []int(nil)
	fmt.Println(len(s), s == nil)
	s = []int{}
	fmt.Println(len(s), s == nil)
```

- Não são comparavéis. Não podemos usar o operador `==` para comparar duas fatias. Existe uma exceção para comparação com `nil`.

```golang
    s := []string{"a", "a", "a", "4", "5", "6"}
	b := []string{"a", "a", "a", "4", "5", "6"}

	fmt.Println(s == b)
```


## Append. O que seria das fatias sem a função append :relieved: ?

```golang
    s := []string{"a", "a", "a", "4", "5", "6"}
	s = append(s, "new_value")
	fmt.Printf("%#v", s)
```

A funçao embutida append do golang, consegue concatenar um item em uma fatia.

Essa função consegue trazer flexibilidade as fatias, pois ela consegue verificar a capacidade da fatia e caso a capacidade tenha sido ultrapassada a função aloca mais espaço.

Como?
- Aloca um novo array subjacente com a capacidade grande o suficiente para armazenar o novo dado
- copia os valores da fatia existente para este novo array
- concatena o novo valor

Exemplo de código um [append](https://github.com/adonovan/gopl.io/blob/master/ch4/append/main.go)


Append permite concatenar multiplos valores, como no exemplo abaixo:

```golang
    s := []string{"a", "a", "a", "4", "5", "6"}
	s = append(s, "new_value", "novo_valor", "mais_valor")
	fmt.Printf("%#v", s)
```

Append permite o uso de reticência, permitindo concatenar uma lista em outra.

```golang
    s := []string{"a", "a", "a", "4", "5", "6"}
	s1 := []string{"c", "b"}
	s = append(s, s1...)
	fmt.Printf("%#v", s)
```



# Aprende Go: Guía para Principiantes

Este proyecto es una introducción práctica al lenguaje Go, ideal para principiantes que desean aprender los conceptos básicos de programación usando Go.

## ¿Qué es Go?

Go (o Golang) es un lenguaje de programación creado por Google. Es conocido por ser rápido, eficiente y fácil de aprender. Con un enfoque en la simplicidad, Go es ideal para construir aplicaciones modernas.

---

## Contenido

1. [Hola, Mundo](#1-hola-mundo)
2. [Variables](#2-variables)
3. [Constantes](#3-constantes)
4. [Tipos de Datos](#4-tipos-de-datos)
5. [Operadores](#5-operadores)
6. [Estructuras de Control](#6-estructuras-de-control)
7. [Bucles](#7-bucles)
8. [Funciones](#8-funciones)
9. [Arreglos](#9-arreglos)
10. [Slices](#10-slices)
11. [Mapas](#11-mapas)
12. [Structs](#12-structs)
13. [Punteros](#13-punteros)
14. [Interfaces](#14-interfaces)
15. [Paquetes](#15-paquetes)
16. [Manejo de Errores](#16-manejo-de-errores)

---

## 1. Hola, Mundo

El programa más básico que imprime "Hola, Mundo!" en la consola.

```go
package main

import "fmt"

func main() {
    greeting := "Hola, Mundo!"
    fmt.Println(greeting)
}
```

- `package main`: Indica que este archivo es parte del paquete principal.
- `import "fmt"`: Importa el paquete `fmt` para manejar la entrada/salida.
- `fmt.Println()`: Imprime texto en la consola.

---

## 2. Variables

Las variables son contenedores que almacenan datos para su uso en el programa.

```go
var a int = 10
var b int = 20
var c int = a + b
fmt.Println(c)
```

- `var a int = 10`: Declara una variable `a` de tipo entero y la inicializa con `10`.
- `var c = a + b`: Suma los valores de `a` y `b` y los almacena en `c`.

---

## 3. Constantes

Las constantes son valores que no pueden cambiar durante la ejecución.

```go
const d int = 30
const e int = 40
const f int = d + e
fmt.Println(f)
```

- `const`: Palabra clave para declarar constantes.
- Las constantes son útiles para valores que no deben modificarse, como `PI`.

---

## 4. Tipos de Datos

Go soporta varios tipos de datos básicos.

```go
var g int = 50
var h float64 = 60.5
var i string = "Hello, World!"
var j bool = true
fmt.Println(g, h, i, j)
```

- `int`: Números enteros.
- `float64`: Números con decimales.
- `string`: Texto.
- `bool`: Verdadero o falso.

---

## 5. Operadores

Los operadores realizan operaciones como suma, resta y comparación.

```go
var k int = 10
var l int = 20
var m int = k + l
fmt.Println(m)
```

---

## 6. Estructuras de Control

Controlan el flujo del programa basándose en condiciones.

```go
var r int = 10
if r > 5 {
    fmt.Println("r es mayor que 5")
} else {
    fmt.Println("r es menor o igual a 5")
}
```

---

## 7. Bucles

Los bucles repiten una acción varias veces.

```go
for s := 0; s < 5; s++ {
    fmt.Println(s)
}
```

---

## 8. Funciones

Las funciones son bloques reutilizables de código.

```go
func add(i1, i2 int) int {
    return i1 + i2
}
fmt.Println(add(10, 20))
```

---

## 9. Arreglos

Un arreglo almacena varios valores del mismo tipo.

```go
var t [5]int
t[0] = 10
fmt.Println(t)
```

---

## 10. Slices

Un slice es una versión más flexible de un arreglo.

```go
var u []int
u = append(u, 10, 20, 30)
fmt.Println(u)
```

---

## 11. Mapas

Un mapa almacena pares clave-valor.

```go
v := make(map[string]int)
v["a"] = 10
fmt.Println(v)
```

---

## 12. Structs

Un struct agrupa datos relacionados.

```go
type Person struct {
    Name string
    Age  int
}
w := Person{Name: "John", Age: 30}
fmt.Println(w)
```

---

## 13. Punteros

Un puntero almacena la dirección de memoria de una variable.

```go
x := 10
y := &x
fmt.Println(*y)
```

---

## 14. Interfaces

Una interfaz define un conjunto de métodos comunes.

```go
type Shape interface {
    Area() int
}
type Circle struct {
    Radius int
}
func (c Circle) Area() int {
    return 3 * c.Radius * c.Radius
}
```

---

## 15. Paquetes

Los paquetes organizan el código y permiten la reutilización.

```go
import "fmt"
```

---

## 16. Manejo de Errores

Go tiene un manejo de errores explícito.

```go
if err != nil {
    panic("Algo salió mal!")
}
```

---

## Cómo Ejecutar el Código

1. Instala Go en tu computadora.
2. Guarda el código en un archivo con extensión `.go`.
3. Ejecuta el archivo:
   ```bash
   go run archivo.go
   ```

---

¡Felicidades por empezar tu viaje con Go! Si tienes dudas, no dudes en pedir ayuda.

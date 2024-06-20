[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11062031)
# Guía 4
## Conjuntos o sets

En la carpeta `/set` se encuentra una implementación de conjunto sobre una lista enlazada simple (reescrita y optimizada para representar conjuntos)
La carpeta `/tests` contiene los tests unitarios de conjunto y los tests de los ejercicios. Por favor no modificar los nombres de estos archivos para que funcionen las pruebas automáticas

## Ejercicios

Completar los ejercicios en el archivo ejercicios.go

1. Escribir una función que reciba una cadena y devuelva el conjunto de todas las letras de la cadena:

```go
func Letras(s string) *set.Set[string]
```

> Recuerden que al recorrer una cadena deben castear cada valor a string ya que GO representa los caracteres como enteros. No se deben incluir los blancos en el conjunto. Por lo tanto se recomienda usar `github.com/agrison/go-commons-lang/stringUtils` para chequear si un caracter es un blanco (espacio, tabulador, salto de línea, etc.)

2. Escribir una función que reciba un arreglo de elementos comparables y elimine los repetidos.

```go
func EliminarRepetidos[T comparable](arreglo []T) []T
```

 3. Escribir una función que reciba dos conjuntos A y B y devuelva la diferencia simétrica entre ambos.

> La diferencia simétrica es el conjunto de elementos que solo pertenecen a A o a B pero no a ambos a la vez.

```go
func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T]
```

4. Escribir una función que reciba una cantidad variables de conjuntos de elementos comparables y devuelva la intersección de todos ellos.

```go
func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] 
```

5. Agregar casos de tests para probar las siguientes propiedades de la diferencia simétrica de conjuntos (Δ) (`/tests/ejercicios_tests/DiferenciaSimetrica_test.go`)

    - Nilpotencia: A Δ A = ∅
    - Asociativa: (A Δ B) Δ C = A Δ (B Δ C)
    - Conmutativa: A Δ B = B Δ A  
    - Neutro: A Δ ∅ = A

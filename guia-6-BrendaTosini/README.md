[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11065477)
# Guía 6
## Diccionario o dictionary

En la carpeta `/dictionary` se encuentra una implementación de diccionario sobre un map
La carpeta `/tests` contiene los tests unitarios de diccionario y los tests de los ejercicios. Por favor no modificar los nombres de estos archivos para que funcionen las pruebas automáticas

## Ejercicios

Completar los ejercicios en el archivo ejercicios.go

1. Escriba una función que reciba una cadena de caracteres y calcule la frecuencia de aparición de cada letra. Por ejemplo:

	frecuencias("hola como estas...") debe devolver 
	{a: 2. c: 1, e: 1, h: 1, l: 1, m: 1, o: 3, s: 2, t: 1,}

```go
func Frecuencia(s string) *dictionary.Dictionary[string, int]
```

> Recuerden que al recorrer una cadena deben castear cada valor a string ya que GO representa los caracteres como enteros. No se deben incluir los blancos en el conjunto. Por lo tanto se recomienda usar `github.com/agrison/go-commons-lang/stringUtils` para chequear si un caracter es un blanco (espacio, tabulador, salto de línea, etc.)

2. Escriba una función que toma un texto en castellano y un diccionario, y devuelve la oración traducida palabra por palabra según el diccionario. Si una palabra no se encuentra en el diccionario, deberá sustituirla por <error> en la cadena resultante.

```go
func Traducir(texto string, dic dictionary.Dictionary[string, string]) string
```
> Para parsear el string pueden usar: func Split(str, sep string) []string

3. Escriba una función que toma 2 slices y devuelve una nueva lista que es el resultado de la interseccion de los 2 slices anteriores, la complejidad del metodo debe ser O(n)

```go
func Interseccion(s1 []string, s2 []string) linkedlist.LinkedList[string]
```

4. Tenemos que ayudar a los docentes a preparar la información solicitada por el
Departamento de Alumnos.
Cuando toman asistencia anotan los presentes por
día de clase, pero ahora desde Alumnos se les pide que envíen la información
de asistencia por alumno.  Por ejemplo, si la
lista que se recibe es:
{“Mie 10”: [“Ana", "Pedro"], “Vie 12”:, [“Ana", "Luz”], “Mie 17”:, [“Luz”, "Pedro"]}
el resultado debe ser 
{"Ana":["Mie 10", "Vie 12"], "Luz": ["Vie 12", "Mie 17"], "Pedro": ["Mie 10", "Mie 17"]}

```go
func InformacionSolicitada(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string]
```

# Guía 11
## Implementaciones de las diapositivas

En las siguientes carpetas se encuentran las implementaciones del código suministrado en las clases:

- `/fibonacci`, el ejemplo de fibonacci, utilizando todas las estrategias de programación dinámica
- `/ladron`, el ejemplo del ladrón de casas, utilizando todas las estrategias de programación dinámica

## Ejercicios

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para las siguientes consignas.

### Programación Dinámica

Usando las estrategias vistas en clase resolver:

#### 1. Camino de costo mínimo
Dada una matriz con valores, cada celda contiene uno que representa el costo de pasar por esa celda.
Se desea encontrar un camino que conecte la celda de la esquina superior izquierda con la esquina inferior derecha, cuyo coste total es mínimo. Devolver el costo del camino.

**Ejemplo**
```
matriz = 3,  2, 12, 15, 10
         6, 19,  7, 11, 17
         8,  5, 12, 32, 21
         3, 20,  2,  9,  7

output = 52
```

**Resolver utilizando estrategia top-down.**

#### 2. Subsecuencia común más larga (LCS)
Dadas dos cadenas `s1` y `s2`, encontrar el largo de su subsecuencia en común más larga. Una subsecuencia de un string `s` es un subconjunto de sus caracteres, no necesariamente adyacentes, pero que tienen el mismo orden.
Ejemplo:
```go
s1 := "abdacbab"  
s2 := "acebfca"

len(lcs)   // 4
lcs        // abca
```

**Resolver utilizando estrategia bottom-up.**

#### 3. Formas de subir una escalera
Dada una escalera con n escalones, y una lista de formas posibles en que se puede subir, encontrar el número total de formas de llegar al último escalón, comenzando desde el suelo.

**Ejemplo**
```
input:  
n = 10  
formas = [2, 4, 5, 8]

output: 11
```

**Resolver utilizando estrategia top-down.**

#### 4. Problema de la mochila 0/1
Dado el valor y peso de cada uno de n items, queremos maximizar el valor de los items que colocamos en la mochila sin exceder su capacidad k. Se pide encontrar el valor total máximo que no exceda un peso total de k.
Cada elemento puede ser utilizado a lo sumo una vez.

**Ejemplo**
```
input:  
values = [20, 30, 15, 25, 10]  
weights = [6, 13, 5, 10, 3]  
k = 20

output: 55
```

**Resolver utilizando estrategia bottom-up.**

#### 5. Suma de subconjuntos
Dado un array `arr` de positivos estrictamente positivos, y un entero k, crear una función que retorne el número de subconjuntos de arr que suman k.

**Ejemplo 1**
```
input:  
arr = [1, 2, 3, 1]  
k = 4
output: 3
subsets: [1, 2, 1], [1, 3], [3, 1]
```

**Ejemplo 2**
```
input:  
arr = [1, 2, 3, 1, 4],  
k = 6
output: 4
subsets: [1, 2, 3], [1, 1, 4], [2, 3, 1], [2, 4] 
```

**Ejemplo 3**
```
input:  
arr = [2, 4, 2, 6, 8]  
k = 7
output: 0
subsets: []
```
**Resolver utilizando estrategia top-down.**

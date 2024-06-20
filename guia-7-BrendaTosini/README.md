[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11065478)
# Guía 7
## Implementaciones de las diapositivas

En las siguientes carpetas se encuentran las implementaciones del código suministrado en las clases:

- `/bases`, los ejemplos de descomposición de un número en base 10 o en base arbitraria
- `/busqueda`, los ejemplos de búsqueda binaria con estrategia recursiva de división y conquista
- `/hanoi`, el ejemplo de ejecución del algoritmo de torres de Hanoi, con estrategia recursiva
- `/palindromo`, los ejemplos de identificacion de palíndromos, con una consigna oculta (para valientes)
- `/potencia`, la resolución del ejercicio de clase, de la estrategia division y conquista para potencias

## Ejercicios

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para las siguientes consignas.

### Recursividad

1. Escriba un método recursivo que tome un entero `n` y devuelva la suma de los primeros `n` números.

2. Escriba un método recursivo que tome un entero `n` devuelva su factorial.

3. Escriba un método recursivo que devuelva la cantidad de unos en la representación binaria de un número `n`. Use el hecho de que es igual al número de unos en la representación binaria de `n/2`, mas `1` si es impar.

4. Escriba un método recursivo que tome una cadena y devuelva como resultado la cadena invertida.

5. Escribir un método que calcule recursivamente cuántos elementos hay en una `Pila` y no altere el contenido de la misma. La pila sólo tiene los métodos `Push()`, `Pop()` y `IsEmpty()`.

6. Escriba un método recursivo que calcule el máximo común divisor entre dos números enteros. Nota: Se puede usar el algoritmo de Euclides:

```
	MCD(a, 0) = a
	MCD(a, b) = MCD (b, a mod b)
```

7. Escriba un método recursivo que tome dos números enteros y calcule la multiplicación entre ellos, usando sólo sumas.

8. Escribir un método recursivo que devuelva el cociente y el resto de la división entera mediante restas sucesivas.

9. Escribir un método recursivo que tome un array de números enteros y devuelva la suma de todos sus elementos.

10. Escribir un método recursivo que dado un número entero decimal devuelva su equivalente en binario.

11. Escribir un método recursivo que reciba 2 enteros `n` y `b` y devuelva true si `n` es potencia de `b`. Por ejemplo: `EsPotencia(8, 2)` devuelve `true`.

12. Escriba un método recursivo que calcule Fibonacci de `n`.

### Búsquedas y Ordenamientos Recursivos
13. Escriba un método recursivo para determinar si un elemento está en un arreglo utilizando búsqueda binaria. Calcule su eficiencia.

14. Sean `as` y `bs` dos listas de enteros de tamaño `n`. Escribir una función que reciba como parámetros `as`, `bs` y un entero `x` y decida si `x` se puede escribir como suma de un elemento de `as` más un elemento de `bs` (por ejemplo, si as=[7,4,6,8], bs=[3,1,6,6] y x=7, la respuesta debe ser `true`, mientras que con los mismos `as` y `bs`, pero con x=2, la respuesta debe ser `false`). El tiempo de la función tiene que ser menor a O(n^2). Justificar el tiempo.

15. Se tiene un arreglo de `n` >= 3 elementos en forma de pico, esto es: estrictamente creciente hasta una determinada posición `p`, y estrictamente decreciente a partir de ella (con 0 < p < n-1). Por  ejemplo, en el arreglo [1,2,3,1,0,-2] la posición del pico es p=2. Se pide implementar un algoritmo de división y conquista de orden O(log n) que encuentre la posición `p` del pico.

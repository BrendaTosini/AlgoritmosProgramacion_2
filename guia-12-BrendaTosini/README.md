[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11164380)
# Guía 12
## Implementación de las diapositivas

En la siguiente carpeta se encuentra la implementación del código suministrado en la clase:

- `/binarytree`, código del binarytree como del binarynode

En la siguiente carpeta se encuentran los test del código suministrado en la clase:

- `/tests`

## Ejercicios

En la carpeta `/ejercicios` encontrarás una implementacion de un binarytree para enteros que deberas usar para resolver los ejercicios


### Ejercicio 1
Escribir un método recursivo que calcule (y devuelva) la suma de todos los nodos en un árbol binario de enteros. El método debe devolver 0 si el árbol está vacío

![](imagenes/tree.png)

Por ejemplo: 2 + 7 + 5 + 2 + 6 + 9 + 5 + 11 + 4 =  51

### Ejercicio 2
Escribir un método recursivo que calcule (y devuelva) la suma de los nodos hojas en un árbol binario de enteros. El método debe devolver 0 si el árbol está vacío

![](imagenes/tree.png)

Por ejemplo: 2 + 5 + 11 + 4 =  22

### Ejercicio 3
Escribir un método recursivo que calcule (y devuelva) la suma de los nodos internos en un árbol binario de enteros. El método debe devolver 0 si el árbol está vacío o solo tiene raiz o raiz y hojas

![](imagenes/tree.png)

Por ejemplo: 7 + 5 + 6 + 9 =  27

### Ejercicio 4
Escribir un método recursivo que calcule (y devuelva) la suma de todos los valores pares de los nodos en un árbol binario de enteros. El método debe devolver 0 si el árbol está vacío

![](imagenes/tree.png)

Por ejemplo: 2 + 2 + 6 + 4 =  14

### Ejercicio 5
Escribir un método recursivo que calcule (y devuelva) la cantidad de nodos cuyo valor es mayor o igual que 6 de un árbol binario de enteros. El método debe devolver 0 si el árbol está vacío

![](imagenes/tree.png)

Por ejemplo: 7 + 6 + 9 + 11 =  33

### Ejercicio 6
Escribir un método recursivo que calcule (y devuelva) la mayor altura entre el sub-árbol izquierdo y derecho de un árbol binario de enteros. El método debe devolver -1 si el árbol está vacío

![](imagenes/ejer6.jpg)

Por ejemplo: debe devolver 3

### Ejercicio 7
Escribir un método recursivo que calcule (y devuelva),la suma de todas las hojas de un árbol binario de enteros que son hijas izquierdas de un nodo. El método debe devolver 0 si el árbol está vacío

![](imagenes/ejer7.jpg)

Por ejemplo: 17 + 27 = 44

### Ejercicio 8

Escribir un método recursivo para un árbol binario de enteros, que calcule (y devuelva) la suma de todos los nodos derechos cuyo valor es impar. El método debe devolver 0 si el árbol está vacío

![](imagenes/ejer8.jpg)

Por ejemplo: 29 + 15 + 11 = 55

### Ejercicio 9

Agregar a la clase árbol binario dada en clase un método que permita recorrer el árbol por niveles. Pista usar una cola para registrar los hijos de un nodo que se está procesando y que se visitarán luego de terminar de visitar todos los nodos del nivel actual. Por ejemplo dado el árbol de la figura el recorrido por niveles debería mostrar en pantalla: A, B, C, D, E, F, G,H, I, J, K, L.

![](imagenes/ejer9.png)

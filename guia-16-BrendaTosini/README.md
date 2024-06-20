[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11305573)
# Guía 16
## Implementaciones de las diapositivas

En las siguientes carpetas se encuentran las implementaciones del código suministrado en las clases:

- `/heapsort`, el ejemplo de heapsort
- `/mergesort`, el ejemplo de mergesort
- `/quicksort`, el ejemplo de quicksort

## Ejercicios

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para las siguientes consignas.

### HeapSort
1. Compare las siguientes versiones de HeapSort bajo la luz de la complejidad temporal y espacial:
    1. Aquella que interpreta un array como Heap, y verifica la propiedad de montículo
    2. Aquella que toma cada elemento de un array y lo inserta en un montículo, para luego pasarlos nuevamente al array

2. Dada una lista de números enteros, se pide encontrar el tercer valor más grande de esa lista utilizando el algoritmo de HeapSort. En lugar de ordenar toda la lista, se puede utilizar HeapSort para encontrar rápidamente el tercer valor más grande sin necesidad de ordenar completamente los elementos.
Se puede hacer de la siguiente forma:
    1. Primero, construir un heap máximo utilizando los elementos de la lista. Puede hacerse utilizando una función auxiliar, como heapify, que ajusta los elementos de la lista para que cumplan con la propiedad del heap máximo.
    2. Después de construir el heap de máximo, realizar dos extracciones del máximo elemento del heap. Cada extracción del máximo elemento corresponderá a encontrar el valor más grande de la lista en ese momento. Descartar estos dos valores, ya que no son los que se están buscando.
    3. Finalmente, realizar una tercera extracción del máximo elemento del heap. Este valor corresponderá al tercer valor más grande de la lista original.


### QuickSort
1. Dado un arreglo de números enteros, diseña un algoritmo que encuentre el "k-ésimo" elemento más pequeño del arreglo. Es decir, encuentra el elemento que ocuparía la posición k si el arreglo estuviera ordenado de manera ascendente utilizando el algoritmo QuickSort.

Pasos para resolver el ejercicio:

1. Define una función llamada "EncontrarKesimo" que tome como parámetros el arreglo de números enteros y el valor k.
2. Implementa una lógica similar al particionamiento del algoritmo QuickSort para encontrar el elemento "k-ésimo" en el arreglo.
    1. Seleccione un pivote aleatorio del arreglo.
    2. Realice el particionamiento del arreglo en base al pivote, de manera que todos los elementos menores que el pivote estén a su izquierda y los mayores estén a su derecha.
    3. Verifique en qué posición quedó el pivote después del particionamiento.
    4. Si la posición del pivote es igual a k, significa que hemos encontrado el elemento "k-ésimo" y lo retornamos.
    5. Si la posición del pivote es mayor a k, aplicamos el mismo proceso recursivamente en la sublista izquierda del pivote.
    6. Si la posición del pivote es menor a k, aplicamos el mismo proceso recursivamente en la sublista derecha del pivote.
3. Utiliza la función "EncontrarKesimo" en tu programa principal para encontrar el "k-ésimo" elemento del arreglo.

### MergeSort

1. Implementar el merge sort, pero en vez de dividir en dos partes el arreglo, dividirlo en tres, llamando recursivamente al algoritmo para las tres partes, y luego hacer el merge de las 3.
Suponiendo que el merge de las tres partes es de orden lineal, calcular el orden de este algoritmo utilizando el teorema maestro.

2. Si en vez de separar en 3 partes, se divide en n partes (siendo n la cantidad de elementos del arreglo), ¿a qué otro algoritmo de ordenamiento se asemeja esta implementación? ¿cuál es el orden de dicho algoritmo?

3. Dado lo obtenido en los puntos anteriores, ¿tiene sentido implementar mergesort con separación en k arreglos (para k > 2)?
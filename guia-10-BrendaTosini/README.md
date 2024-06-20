[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11164374)
# Guía 10 - Backtracking
## Implementaciones de las diapositivas

En las siguientes carpetas se encuentran las implementaciones de algunos problemas utilizando backtracking:

- `/ratmaze`, problema del laberinto, resuelto con backtracking
- `/cambio`, problema del cambio de moneda, resuelto con backtracking
- `/sudoku`, resolución de un sudoku clásico, utilizando backtracking
- `/reinas`, resolución del problema de las n reinas, utilizando backtracking

- `/main`, ejecuciones adicionales con fines interactivos

## Ejercicios

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para las siguientes consignas.

### Backtracking

Usando los conceptos de backtracking resuelva los siguientes ejercicios. 

0. En los ejercicios del laberinto y sudoku investigar la estructura de datos que soporta la solución, e identificar las operaciones Factible, esSolución, Extender, Registrar y Borrar. Tomar como ejemplo el código de cambio de moneda y N-reinas que está comentado.
   
1. **Problema de los dados:** Se tienen n dados de k caras cada uno, se desea saber la cantidad de formas de obtener una suma de x puntos al lanzar los n dados.
   
    > Ejemplo: Si tenemos n = 3 dados de k = 6 caras y queremos obtener el valor x = 7 la solución es (1, 1, 5), (1, 2, 4), (1, 3, 3), (1, 4, 2), (1, 5, 1), (2, 1, 4), (2, 2, 3), (2, 3, 2), (2, 4, 1), (3, 1, 3), (3, 2, 2), (3, 3, 1), (4, 1, 2), (4, 2, 1), (5, 1, 1). En total 15 variantes distintas

2. **Problema de la mochila (Knapsack):** Este es uno de los 21 problemas de [NP completos](https://es.wikipedia.org/wiki/NP-completo) de Richard Karp, y ya [mencionado en 1896](https://doi.org/10.1112%2Fplms%2Fs1-28.1.486) por Mathews, G. B. Consiste en un problema de optimización combinatoria, donde se espera poder llenar una "mochila" con un peso limitado, por una cantidad de objetos, cada uno con un peso y valor específico, máximizando el valor total almacenado. Hay una simplificación del mismo, denominado problema de la mochila 0-1 donde un objeto puede estar o no dentro de la mochila, por completo, es decir no se puede fraccionar. Debe devolver el valor máximo que se puede transportar en la mochila.

3. Modificar el problema anterior para que devuelva ademas la lista de objetos a incluir en la mochila

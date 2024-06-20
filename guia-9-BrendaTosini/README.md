[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11164370)
# Guía 9
## Implementaciones de las diapositivas

En las siguientes carpetas se encuentran las implementaciones del código suministrado en las clases:

- `/cambio`, el ejemplo de cambio de monedas
- `/actividades`, el ejemplo selección de actividades

## Ejercicios

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para las siguientes consignas.

### Algoritmos Ávidos

#### 1. Selector de actividades recursivo
Reescribir la función que resuelve el problema de seleccionar actividades en forma recursiva. 


#### 2. Minimización del tiempo en el sistema
Dadas n tareas, cada una de las cuales tarda un tiempo t en ejecutarse y un procesador que ejecuta dichas tareas, se debe diseñar una función que nos devuelva una planificación de las tareas (un orden para ejecutarlas), tal que minimice el tiempo medio de finalización (promedio de los tiempos en que finalizan las tareas).

**Ejemplo:**

> t = [4, 10, 2, 20] Si las ejecuto en el orden a1, a2, a3, a4 entonces cada una de las tareas se ejecutarán en los siguientes intervamos de tiempo a1: [0, 4), a2: [4, 14), a3:[14, 16), a4:[16, 36) y el promedio de los tiempos de finalización es (4 + 14 + 16 + 36) / 4 = 17,5. En cambio si se ejecutan a3, a1, a2, a4 entonces el cada una de las tareas se ejecutarán en los siguientes intervalos de tiempo a1: [2, 6), a2: [6, 16), a3:[0, 2), a4:[16, 36) y el promedio de los tiempos de finalización es (6 + 16 + 2 + 36) / 4 = 15

#### 3. Problema de la mochila 0/1
Se tienen n objetos (cada objeto i tiene un peso y un valor); y una mochila con capacidad máxima de W. Se pretende encontrar la manera de cargar la mochila de forma que se maximice el valor de lo transportado y se respete su capacidad máxima. Los objetos se pueden fraccionar de tal forma de cargar solo una fracción del objeto

> El paso greedy es elegir primero el elemento que tenga mayor valor por unidad de peso

#### 4. Farolas públicas
Hay M farolas en las posiciones `y1,…,yM` de una recta y N puntos `x1,…,xN`. Cada farola tiene un radio de iluminación `ri`, tal que la i-ésima farola ilumina puntos en el intervalo `[yi−ri,yi+ri]`. Se quiere encender el mínimo número de farolas tales que cada uno de los N puntos `x1,…xN` esté iluminado por al menos una farola. Programar una función que devuelva una lista de farolas a encender

**Entrada:** Un arreglo de Farolas de tamaño M (de cada farola se registra posición, radio y si está encendida) y un arreglo de enteros de tamaño N que representan los puntos a iluminar

**Salida:** Lista de Farolas encendidas. Si el problema no tiene solución debe devolver un error

**Pre condiciones:** El arreglo de farolas y puntos está ordenado por posicion de menor a mayor.

> **Pista:** Para resolver este problema, nos imaginamos los puntos y las farolas en la recta, ordenados de izquierda a derecha. Se encienden las farolas más a la derecha que puedan iluminar el punto actual y se avanza al siguiente punto. 

> **Paso Greedy:** Dado un punto x, encender la farola más a la derecha que pueda iluminarlo.


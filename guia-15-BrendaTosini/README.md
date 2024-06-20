[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11166798)
# Guía 15

## Desafíos con Montículos

### Parte 0: Implementación de referencia

Ver la carpeta `/src/monticulo` para una implementación de referencia de Montículo de Máximo, junto a sus tests. También se provee una implementación de referencia para Montículo de Mínimo.

#### Ejecutar tests

    % go test ./...

### Parte I: Seguimientos


Dado el siguiente montículo de MÁXIMO (del cual sólo se muestran las prioridades de sus elementos):
```
         [25]
    [15]      [20]
  [6]  [10]
```

1. Ejecutar la siguiente secuencia: Insertar 45, Insertar 22, EliminarMáx, Insertar 5. Realice los dibujos del montículo luego de cada momento (tanto el árbol como el arreglo).

   Puede verificar sus pasos en visualgo.net/en/heap

2. En un montículo de MÍNIMO vacío, ejecutar la secuencia (de la cual sólo conocemos las prioridades de los elementos):

   10, 12, 1, 14, 6, 5, 8, 15, 3, 9, 7, 4, 11, 13, 2. EliminarMín, EliminarMín.


### Parte II: Modificando Montículos

Dado un Montículo de MÁXIMO:

3. Escriba un constructor alternativo `NuevoMonticuloMaxDesdeArreglo` que tome como parámetro un arreglo desordenado de elementos y devuelva un montículo que contenga esos elementos.

4. Escriba una función que, dado un Montículo de tamaño M y un número entero n, con n >= 1 y n <= M,
devuelva el enésimo máximo del montículo. (Ej: si n == 1, devolver el primer máximo.)

5. Escriba una función que, dados dos montículos, devuelva un tercero que sea la combinación de ambos.


### Ejercicios extra

6. Resuelva los puntos anteriores de la Parte II para la Montículo de MÍNIMO.
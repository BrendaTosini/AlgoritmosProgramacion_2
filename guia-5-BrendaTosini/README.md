[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=11065472)
# Guía 5
## Mapa de Bits y Tablas de Hash

En la carpeta `/mapadebits` se encuentra una implementación de un mapa de sobre un entero de 32 bits.
La carpeta `/tests` contiene los tests unitarios del mapa. Por favor no modificar los nombres de estos archivos para que funcionen las pruebas automáticas

## Ejercicios

Completar los ejercicios en el archivo un ejercicios.go igual que en las guías anteriores

1. Implementar el TAD Conjunto sobre un Map de Go. Tal que las operaciones Add, Remove y Contains sean O(1). Sabiendo que los Map de Go usan hashing para almacenar sus claves. La idea por lo tanto es aprovechar el map para almacenar los elementos del conjunto como las claves de un map.
```go
type Set [T comparable] struct{
  elementos map[T]struct{}
 }
 ```
  

2. Dada una tabla de hash de tamaño 7 y la función de hashing h(x) = x mod 7, inserte los números 1, 8, 27, 125, 216, 343, resolviendo los choques con:
  - el método de hashing cerrado
  - el método de hashing abierto
y dibuje las tablas resultantes en cada caso.

3. Dada una tabla de hash de tamaño 10 y la función de hashing h(x) = x mod 10, inserte los números 4.371, 1.323, 6.173, 4.199, 4.344, 9.679, 1.989, resolviendo los choques con:
  - el método de hashing cerrado
  - el método de hashing abierto
y dibuje las tablas resultantes en cada caso.

4. Se quiere crear un registro de asistencia para la clase, sabiendo que son 40 alumnos. Las operaciones que debe soportar el registro son dado un día listar los ausentes. Listar todos los alumnos que no tuvieron ninguna falta en todo el mes además de las operaciones propia del registro como colocar presente o ausente a cada alumno.



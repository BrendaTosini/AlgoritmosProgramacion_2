

# Trabajo Práctico
## Algoritmos y Programación II




### Participantes
 
Agustín Andrés Scazzino,
Brenda Tosini,
Ivan Garay,
Gonzalo Giuffre,
Martín Moreno



### Tutora: 
Victoria Becker

### Lenguaje: 
Golang

## Documento Online
https://docs.google.com/document/d/17zEb_6wfxPloi00DGiNE9t82GP6QGe_C_r9QSZIYs_w/edit#




# GESTOR DE TAREAS



### Descripción del programa

El programa es una aplicación de gestión de tareas que permite al usuario realizar diferentes operaciones, como agregar nuevas tareas, modificar tareas existentes, eliminar tareas y ver el listado completo de tareas.

El programa está implementado en el lenguaje de programación Go y utiliza estructuras de datos para representar las tareas. Cada tarea tiene un nombre, duración, prioridad, etiquetas, y puede tener subtareas.



## Funcionalidades principales

El programa ofrece las siguientes funcionalidades principales:

- Nueva Tarea: Permite al usuario ingresar los detalles de una nueva tarea, como nombre, duración, prioridad y etiquetas. La tarea se agrega a la lista de tareas existentes.

- Listar Tareas: Muestra al usuario un listado completo de todas las tareas existentes, incluyendo su estado (completada o no), prioridad, nombre, duración y etiquetas asociadas.

- Modificar Tarea: Permite al usuario seleccionar una tarea de la lista y modificar su nombre, duración, prioridad o etiquetas.

- Eliminar Tarea: Permite al usuario seleccionar una tarea de la lista y eliminarla de forma permanente.

- Salir: Permite al usuario finalizar el programa.



## Explicación de las funciones del programa



### Archivo: Tareas

A continuación, se presenta una descripción de cada una de las funciones presentes en el código tareas.go:

- NewTarea(nombre string, duracion float32, etiquetas *set.Set[string]) *Tarea: Esta función es un constructor que crea y devuelve un nuevo objeto Tarea. Recibe como argumentos el nombre de la tarea, su duración y un conjunto de etiquetas. Inicializa las propiedades de la tarea y devuelve un puntero al objeto Tarea creado.
 
- EditarNombre(nombre string): Este método permite editar el nombre de la tarea. Toma como argumento el nuevo nombre y actualiza la propiedad "nombre" de la tarea con el valor proporcionado.
 
- EditarDuracion(duracion float32): Este método permite editar la duración de la tarea. Toma como argumento la nueva duración y actualiza la propiedad "duración" de la tarea con el valor proporcionado.
 
- SubirPrioridad(): Este método aumenta la prioridad de la tarea. Actualiza la propiedad "prioridad" de la tarea a "Alta Prioridad".
 
- BajarPrioridad(): Este método disminuye la prioridad de la tarea. Actualiza la propiedad "prioridad" de la tarea a "Baja Prioridad".
 
- EditarEtiquetas(etiquetas *set.Set[string]): Este método permite editar las etiquetas asociadas a la tarea. Recibe como argumento un conjunto de etiquetas y actualiza la propiedad "etiquetas" de la tarea con el conjunto proporcionado.
 
- Iniciar() string: Este método cambia el estado de la tarea a "En Progreso" si su estado actual es "Pendiente" o "Interrumpida". Devuelve un mensaje indicando si la tarea se ha pasado a "En Progreso" o si el estado actual no permite el cambio.
 
- Interrumpir() string: Este método cambia el estado de la tarea a "Interrumpida" si su estado actual es "En Progreso". Devuelve un mensaje indicando si la tarea se ha interrumpido o si el estado actual no permite el cambio.
 
- Completar() string: Este método cambia el estado de la tarea a "Completada" si su estado actual es "En Progreso". Devuelve un mensaje indicando si la tarea se ha completado o si el estado actual no permite el cambio.
 
- Abandonar() string: Este método cambia el estado de la tarea a "Pendiente" si su estado actual es "En Progreso". Devuelve un mensaje indicando si la tarea se ha vuelto pendiente o si el estado actual no permite el cambio.
 
- AgregarSubTarea(subt Tarea): Este método agrega una sub tarea a la lista de sub tareas de la tarea actual. Recibe como argumento una tarea y la agrega a la propiedad "lista_de_subtareas" de la tarea actual.
 
- EliminarSubTarea(subt Tarea): Este método elimina una sub tarea de la lista de sub tareas de la tarea actual. Recibe como argumento una tarea y la elimina de la propiedad "lista_de_subtareas" de la tarea actual.
 
- ListaSubTareas() string: Este método devuelve una representación en forma de cadena de las subtareas de la tarea.
  
- ConsultarEstado() string: Este método devuelve el estado actual de la tarea como una cadena de texto.
 
- String() string: Este método devuelve una representación en forma de cadena de la tarea. Incluye el nombre, estado, prioridad, duración formateada y etiquetas.
 
- formatDuracion(duracion time.Duration) string: Esta función auxiliar formatea la duración de la tarea en horas y minutos. Recibe una duración de tipo time.Duration y devuelve una cadena formateada en el formato "Xh Xm".
 
- TiempoTotal() float32: Este método calcula y devuelve la duración total de la tarea, incluyendo la duración de las subtareas. Suma la duración de todas las subtareas de la lista de subtareas y la duración de la tarea actual.
 
- ContienePalabra(palabra string) bool: Este método verifica si el nombre de la tarea contiene una palabra específica. Divide el nombre en palabras y crea un conjunto (set) de palabras. Luego, verifica si el conjunto contiene la palabra específica y devuelve un valor booleano.



### Archivo: Lista de Tareas

A continuación, se presenta una descripción de cada una de las funciones presentes en el código lista_de_tareas.go:
 
- NewListaDeTareas() *ListaDeTareas: Esta función es un constructor que crea y devuelve un nuevo objeto ListaDeTareas. No recibe ningún argumento y devuelve un puntero al objeto ListaDeTareas creado.
 
- AgregarTarea(tarea Tarea): Este método permite agregar una tarea a la lista de tareas. Recibe como argumento una tarea y la agrega al final de la lista de tareas.
 
- EliminarTarea(tarea Tarea): Este método permite eliminar una tarea de la lista de tareas. Recibe como argumento una tarea y la elimina de la lista de tareas si está presente.
 
- Size() int: Este método devuelve el tamaño actual de la lista de tareas. Retorna un entero que representa la cantidad de tareas presentes en la lista.
 
- TareaActual() Tarea: Este método devuelve la tarea actual de la pila de tareas activas. Utiliza el método Top() de la pila de tareas activas para obtener la tarea en la cima de la pila y la devuelve. Si la pila está vacía, devuelve un valor predeterminado de tipo Tarea.

- ListarTareas() Tarea: Este método devuelve una representacion en forma de cadena de las tareas.

//stack
- EliminarTareaActual() Tarea: Este método elimina y devuelve la tarea actual de la pila de tareas activas utilizando el método Pop() de la pila.
 
- VerTareaActual() Tarea: Este método devuelve la tarea actual de la pila de tareas activas sin eliminarla utilizando el método Top() de la pila.
 
- NuevaTarea(): Este método solicita al usuario los datos de una nueva tarea y la agrega a la lista de tareas utilizando la función SolicitarDatosDeTarea().
 
- SolicitarDatosDeTarea() Tarea: Esta función solicita al usuario los datos de una nueva tarea, como el nombre, la duración y la etiqueta, y devuelve una nueva tarea creada con esos datos.
 
- ComenzarTarea(tarea Tarea): Este método marca una tarea como "En Progreso" y actualiza la lista de tareas. Si hay una tarea actual en la pila de tareas activas, la interrumpe y la reemplaza por la nueva tarea. Luego, agrega la nueva tarea a la pila de tareas activas.
 
- CompletarTareaActual(): Este método marca la tarea actual como "Completada" y actualiza la lista de tareas. Además, marca como "Completadas" todas las subtareas de la tarea actual. Si hay una tarea actual en la pila de tareas activas, la elimina de la pila y, si aún quedan tareas en la pila, selecciona la siguiente tarea y la marca como "En Progreso".
 
- AbandonarTareaActual(): Este método marca la tarea actual como "Pendiente" y actualiza la lista de tareas. Si hay una tarea actual en la pila de tareas activas, la elimina de la pila y, si aún quedan tareas en la pila, selecciona la siguiente tarea y la marca como "En Progreso".
 
- PrioridadValue(prioridad string) int: Esta función toma una cadena de prioridad y devuelve un valor entero correspondiente a esa prioridad. Se utiliza en la función PrepararColaDeTareas() para ordenar las tareas por prioridad.
 
- PrepararColaDeTareas(tiempoTotal float32): Este método prepara y muestra una cola de tareas ordenadas según la prioridad y el tiempo total disponible. Utiliza el método ToSlice() de la lista enlazada para obtener una representación de la lista de tareas como una rebanada (slice). Luego, ordena la rebanada de tareas utilizando la función sort.Slice() con una función de comparación personalizada. Finalmente, recorre el slice de tareas y muestra las tareas cuya duración total sea menor que el tiempo total disponible.
 
- BuscarTareas(clave, tipo string) int: Este método busca tareas en la lista de tareas según una clave y un tipo de búsqueda. Si el tipo es "palabra clave", busca tareas cuyos nombres contengan la clave. Si el tipo es "etiqueta", busca tareas que tengan la etiqueta especificada. Muestra las tareas encontradas y devuelve la cantidad de tareas encontradas.
 
- ReordenarTareas(tipo string): Este método reordena las tareas en la lista según un tipo de reordenación especificado. Los tipos de reordenación pueden ser "Prioridad", "Duracion" o "Subtareas". Utiliza la función sort.Slice() con una función de comparación personalizada para ordenar la lista de tareas según el tipo especificado.
AbandonarTareaActual():

Además de estas funciones, se hace referencia a dos tipos de datos importados desde los paquetes "bluelock/linkedlist" y "bluelock/stack". Estos tipos de datos son:
 
- linkedlist.LinkedList[Tarea]: Es una lista enlazada que almacena elementos de tipo Tarea. Se utiliza para almacenar la lista de tareas en el objeto 
- ListaDeTareas.stack.Stack[Tarea]: Es una pila que almacena elementos de tipo Tarea. Se utiliza para almacenar las tareas activas en el objeto ListaDeTareas.
 
En resumen, las funciones presentes en el código permiten agregar y eliminar tareas de una lista, obtener el tamaño de la lista y obtener la tarea actual de la pila de tareas activas. Estas funciones son parte de la implementación de la estructura de datos ListaDeTareas y proporcionan operaciones básicas para manipular y acceder a las tareas almacenadas. actual. 


## Archivo: Menú

- NewMenu() *Menu: Esta función es el constructor del tipo Menu. Crea una nueva instancia de Menu y retorna un puntero a ella. También crea una instancia de ListaDeTareas utilizando el constructor NewListaDeTareas() y la asigna al campo list_t de la estructura Menu.
 
- MenuPrincipal(): Esta función muestra el menú principal del programa y espera que el usuario ingrese una opción. Dependiendo de la opción ingresada, ejecuta la función correspondiente. Por ejemplo, si se ingresa el número 1, llama a la función ListarTareas() de la lista de tareas almacenada en list_t. Si se ingresa el número 2, llama a la función NuevaTarea() de la lista de tareas. El bucle for permite que el menú se muestre repetidamente hasta que el usuario elija la opción de salir (0).
 
- SubMenuLista(): Esta función muestra un submenú relacionado con las tareas. Permite al usuario seleccionar entre volver al menú principal o acceder al submenú de acciones de tarea. Dependiendo de la opción ingresada, llama a la función correspondiente.
 
- SubMenuTarea(): Esta función muestra un submenú de acciones relacionadas con una tarea específica. El usuario debe ingresar el número de tarea deseado. Luego, muestra las acciones disponibles para esa tarea, como editar el nombre, editar la duración, subir o bajar la prioridad, etc. Dependiendo de la opción ingresada, se realizan las acciones correspondientes sobre la tarea seleccionada.
 
- SubMenuTareaActual(): Este método muestra el submenú de la tarea actual. Muestra información sobre la tarea actual y permite al usuario completarla o abandonarla.
 
- SubMenuBuscarTarea(): Este método muestra el submenú de búsqueda de tareas. Permite al usuario buscar tareas por palabra clave o etiqueta y muestra los resultados encontrados.
 
- SubMenuReordenarTareas(): Este método muestra el submenú de reordenamiento de tareas. Permite al usuario reordenar las tareas por prioridad, duración o cantidad de subtareas.

En resumen, el archivo menu.go define y proporciona funciones para interactuar con un menú de administración de tareas, donde el usuario puede realizar diversas acciones, como crear, editar, eliminar y buscar tareas, así como realizar otras operaciones relacionadas.


## Procedimientos realizados

Durante el desarrollo del programa de gestión de tareas, se llevaron a cabo los siguientes procedimientos:

Realización del análisis de requisitos: Se realizó un análisis exhaustivo de los requisitos del programa para comprender las funcionalidades necesarias y los datos requeridos.

Diseño de la estructura del programa: Se diseñó la estructura del programa utilizando el lenguaje de programación Go. Se definieron las estructuras de datos necesarias, como la estructura de la tarea y la lista de tareas.

Implementación de las funcionalidades principales: Se implementaron las funcionalidades principales del programa, como agregar una nueva tarea, listar las tareas existentes, modificar una tarea y eliminar una tarea.

Realización de pruebas exhaustivas: Se llevaron a cabo pruebas exhaustivas del programa para comprobar su funcionamiento correcto. Se realizaron pruebas unitarias, pruebas de integración y pruebas funcionales en diferentes escenarios para asegurar la calidad del código.

Sincronización de cambios: Se utilizó un sistema de control de versiones, como Git, para gestionar los cambios en el código. Los miembros del equipo realizaron solicitudes de extracción, revisaron los cambios propuestos y se fusionaron en la rama principal del proyecto una vez aprobados.

Actualización del repositorio local: Cada persona del equipo actualizó su repositorio local para obtener los cambios más recientes realizados por otros miembros del equipo. Esto se logró ejecutando el comando `git pull origin main` en el directorio del repositorio local.

Iteración y mejora continua: El equipo continuó iterando en el desarrollo del programa, agregando nuevas funcionalidades, corrigiendo errores y mejorando el código existente. Se repitieron los pasos anteriores según las necesidades y los requisitos del proyecto.


## Conclusiones

Gracias a la realización de estos procedimientos, el equipo pudo avanzar de manera efectiva en el desarrollo del programa de gestión de tareas. Se siguió un enfoque iterativo que permitió realizar mejoras constantes y adaptarse a los cambios y requisitos en curso. La colaboración entre los miembros del equipo fue fundamental para el éxito del proyecto, con revisiones regulares, comunicación efectiva y toma de decisiones conjunta.

El programa resultante fue una herramienta funcional y eficiente para gestionar tareas, permitiendo a los usuarios realizar diversas operaciones de manera sencilla. Se logró cumplir con los requisitos establecidos y se sentaron las bases para futuras mejoras y expansiones del programa.

En resumen, los procedimientos realizados durante el desarrollo del programa fueron fundamentales para alcanzar los objetivos establecidos y garantizar la calidad del código y la satisfacción del usuario, agregando nuevas funcionalidades, corrigiendo errores y mejorando el código existente. Se repitieron los pasos anteriores según las necesidades y los requisitos del proyecto.



----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------













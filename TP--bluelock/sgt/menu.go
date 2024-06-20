package sgt

import (
	"bluelock/set"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/agrison/go-commons-lang/stringUtils"
)

type Menu struct {
	list_t ListaDeTareas
}

func NewMenu() *Menu {
	list_t := NewListaDeTareas()
	return &Menu{list_t: *list_t}
}

func (m *Menu) MenuPrincipal() {
	for {
		fmt.Println("\033[1;36mMenu Principal")
		fmt.Println("--------------\033[0m")
		fmt.Println("  1.Listar Tareas")
		fmt.Println("  2.Nueva Tarea")
		fmt.Println("  3.Tarea Actual")
		fmt.Println("  4.Preparar Cola de Tareas")
		fmt.Println("  5.Buscar Tarea")
		fmt.Println("  6.Reordenar Tareas")
		fmt.Println("  0.Salir")
		fmt.Println("")
		fmt.Print("Seleccione una funcion del menu: ")
		var opcion int
		fmt.Scanln(&opcion)
		fmt.Println("")

		switch opcion {
		case 1:
			m.list_t.ListarTareas()
			if !(m.list_t.Size() == 0) {
				m.SubMenuLista()
			}
		case 2:
			nueva_tarea := SolicitarDatosDeTarea("")
			m.list_t.AgregarNuevaTarea(nueva_tarea)
		case 3:
			m.SubMenuTareaActual()
		case 4:
			var tiempoTotal string
			fmt.Print("Ingrese el límite de tiempo que desea para sus tareas (ej: 4.2): ")
			fmt.Scanln(&tiempoTotal)
			fmt.Println("")
			value, err := strconv.ParseFloat(tiempoTotal, 32)
			if err != nil {
				fmt.Println("\033[1;31mLa duracion (" + tiempoTotal + ") ingresada en incorrecta tiene que ser un float valido\033[0m")
				fmt.Println("")
				return
			}
			m.list_t.PrepararColaDeTareas(float32(value))
		case 5:
			m.SubMenuBuscarTarea()
		case 6:
			m.SubMenuReordenarTareas()
		case 0:
			fmt.Println("\033[1;32mSaliendo del programa...\033[0m")
			os.Exit(0)
		default:
			fmt.Println("\033[1;31mOpcion Invalida, Intentelo de nuevo\033[0m")
		}

	}
}

func (m *Menu) SubMenuLista() {
	fmt.Println("\033[1;36m--------------------------\033[0m")
	fmt.Println("1.Acciones de tarea")
	fmt.Println("0.Volver al menu principal")
	fmt.Println("")
	fmt.Print("Seleccione una funcion del menu: ")
	var opcion int
	fmt.Scanln(&opcion)
	fmt.Println("")

	switch opcion {
	case 1:
		m.SubMenuTarea()
	case 0:
		m.MenuPrincipal()
	default:
		fmt.Println("\033[1;31mOpcion Invalida, Intentelo de nuevo\033[0m")
	}
}

func (m *Menu) SubMenuTarea() {
	num_tarea := m.ValidarSeleccionDeTarea()
	referencia_tarea_seleccionada := m.list_t.list.GetNode(num_tarea - 1)
	tarea_seleccionada := referencia_tarea_seleccionada.Value()

	for {
		fmt.Println("\033[1;36mAcciones de Tarea")
		fmt.Println("-----------------\033[0m")
		fmt.Println("  1.Editar Nombre")
		fmt.Println("  2.Editar Duracion")
		fmt.Println("  3.Subir Prioridad")
		fmt.Println("  4.Bajar Prioridad")
		fmt.Println("  5.Agregar Etiqueta")
		fmt.Println("  6.Eliminar Etiqueta")
		fmt.Println("  7.Comenzar Tarea")
		fmt.Println("  8.Nueva Subtarea")
		fmt.Println("  9.Borrar Tarea")
		fmt.Println("  0.Volver al menu principal")
		fmt.Println("")
		fmt.Print("Seleccione una funcion del menu: ")
		var opcion int
		fmt.Scanln(&opcion)
		fmt.Println("")

		switch opcion {
		case 1:
			nombre := ValidarNombre()
			tarea_seleccionada.EditarNombre(nombre)
			m.list_t.list.UpdateNodeValue(referencia_tarea_seleccionada.Value(), tarea_seleccionada)
		case 2:
			duracion := ValidarDuracion()
			tarea_seleccionada.EditarDuracion(duracion)
			m.list_t.list.UpdateNodeValue(referencia_tarea_seleccionada.Value(), tarea_seleccionada)
		case 3:
			tarea_seleccionada.SubirPrioridad()
			m.list_t.list.UpdateNodeValue(referencia_tarea_seleccionada.Value(), tarea_seleccionada)
		case 4:
			tarea_seleccionada.BajarPrioridad()
			m.list_t.list.UpdateNodeValue(referencia_tarea_seleccionada.Value(), tarea_seleccionada)
		case 5:
			fmt.Print("Ingrese una etiqueta: ")
			var etiqueta string
			fmt.Scanln(&etiqueta)
			fmt.Println("")
			tarea_seleccionada.AgregarEtiqueta(etiqueta)
			m.list_t.list.UpdateNodeValue(referencia_tarea_seleccionada.Value(), tarea_seleccionada)
		case 6:
			fmt.Print("Ingrese la etiqueta a eliminar: ")
			var etiqueta string
			fmt.Scanln(&etiqueta)
			fmt.Println("")
			tarea_seleccionada.EliminarEtiqueta(etiqueta)
			m.list_t.list.UpdateNodeValue(referencia_tarea_seleccionada.Value(), tarea_seleccionada)
		case 7:
			m.list_t.ComenzarTarea(tarea_seleccionada)
		case 8:
			nueva_sub_tarea := SolicitarDatosDeTarea(tarea_seleccionada.estado)
			tarea_seleccionada.AgregarSubTarea(nueva_sub_tarea)
			m.list_t.list.UpdateNodeValue(referencia_tarea_seleccionada.Value(), tarea_seleccionada)
		case 9:
			m.list_t.EliminarTarea(tarea_seleccionada)
			m.MenuPrincipal()
		case 0:
			return
		default:
			fmt.Println("\033[1;31mOpcion Invalida, Intentelo de nuevo\033[0m")
		}

	}
}

func (m *Menu) SubMenuTareaActual() {
	tarea_actual := m.list_t.VerTareaActual()
	fmt.Printf("\033[1;33mTarea Actual: %s \n\033[0m", tarea_actual.String())
	fmt.Println("")
	fmt.Println("\033[1;36m------------------------\033[0m")
	fmt.Println("1.Completar Tarea actual")
	fmt.Println("2.Abandonar")
	fmt.Println("0.Volver al menu principal")
	fmt.Println("")
	fmt.Print("Seleccione una funcion del menu: ")
	var opcion int
	fmt.Scanln(&opcion)
	fmt.Println("")

	switch opcion {
	case 1:
		m.list_t.CompletarTareaActual()
	case 2:
		m.list_t.AbandonarTareaActual()
	case 0:
		m.MenuPrincipal()
	default:
		fmt.Println("\033[1;31mOpcion Invalida, Intentelo de nuevo\033[0m")
	}
}

func (m *Menu) SubMenuBuscarTarea() {
	fmt.Println("\033[1;36mMenu de Busqueda")
	fmt.Println("----------------\033[0m")
	fmt.Println(" 1.Buscar por palabra clave")
	fmt.Println(" 2.Buscar por etiqueta")
	fmt.Println(" 0.Volver al menu principal")
	fmt.Println("")
	fmt.Print("Seleccione una funcion del menu: ")
	var opcion int
	fmt.Scanln(&opcion)
	fmt.Println("")

	switch opcion {
	case 1:
		fmt.Print("Ingrese la palabra clave: ")
		var palabra_clave string
		fmt.Scanln(&palabra_clave)
		fmt.Println("")
		m.list_t.BuscarTareas(palabra_clave, "palabra clave")
		m.SubMenuLista()
	case 2:
		fmt.Print("Ingrese la etiqueta: ")
		var etiqueta string
		fmt.Scanln(&etiqueta)
		fmt.Println("")
		m.list_t.BuscarTareas(etiqueta, "etiqueta")
		m.SubMenuLista()
	case 0:
		m.MenuPrincipal()
	default:
		fmt.Println("\033[1;31mOpcion Invalida, Intentelo de nuevo\033[0m")
	}
}

func (m *Menu) SubMenuReordenarTareas() {
	fmt.Println("\033[1;36mMenu de Reordanamiento")
	fmt.Println("----------------------\033[0m")
	fmt.Println(" 1.Reordenar por prioridad")
	fmt.Println(" 2.Reordenar por duración")
	fmt.Println(" 3.Reordenar por cantidad de subtareas")
	fmt.Println(" 0.Volver al menu principal")
	fmt.Println("")
	fmt.Print("Seleccione una funcion del menu: ")
	var opcion int
	fmt.Scanln(&opcion)
	fmt.Println("")

	switch opcion {
	case 1:
		m.list_t.ReordenarTareas("Prioridad")
	case 2:
		m.list_t.ReordenarTareas("Duracion")
	case 3:
		m.list_t.ReordenarTareas("Subtareas")
	case 0:
		m.MenuPrincipal()
	default:
		fmt.Println("\033[1;31mOpcion Invalida, Intentelo de nuevo\033[0m")

	}
}

func SolicitarDatosDeTarea(estado string) Tarea {
	var etiqueta string

	nombre := ValidarNombre()
	duracion := ValidarDuracion()

	fmt.Print("Ingrese una etiqueta: ")
	fmt.Scanln(&etiqueta)
	fmt.Println("")
	if estado == "" {
		estado = "Pendiente"
	}

	etiquetas := set.NewSet(etiqueta)
	sub_tareas := NewListaDeTareas()
	nueva_tarea := NewTarea(nombre, duracion, estado, etiquetas, sub_tareas)
	return *nueva_tarea
}

func ValidarNombre() string {
	scanner := bufio.NewScanner(os.Stdin)
	var nombre string
	fmt.Print("Ingrese el nombre de la tarea: ")

	for scanner.Scan() {
		nombre = scanner.Text()
		if stringUtils.IsEmpty(nombre) {
			fmt.Println("")
			fmt.Println("\033[1;31mNo se puede ingresar un nombre vacio\033[0m")
			fmt.Println("")
			fmt.Print("Ingrese el nombre de la tarea: ")
		} else {
			break
		}

	}
	fmt.Println("")
	return nombre
}

func ValidarDuracion() float32 {
	var duracion string
	for {
		fmt.Print("Ingrese la duracion (ej:1.2): ")
		fmt.Scanln(&duracion)
		fmt.Println("")
		value, err := strconv.ParseFloat(duracion, 32)
		if err != nil {
			fmt.Println("\033[1;31mLa duracion (" + duracion + ") ingresada es incorrecta tiene que ser un float valido\033[0m")
			fmt.Println("")
		} else {
			return float32(value)
		}
	}
}

func (m *Menu) ValidarSeleccionDeTarea() int {
	var num_tarea string

	for {
		fmt.Print("Elija una tarea por su numero: ")
		fmt.Scanln(&num_tarea)
		fmt.Println("")
		i, err := strconv.Atoi(num_tarea)
		if err != nil {
			fmt.Println("\033[1;31mEl numero de tarea (" + num_tarea + ") no es un entero valido\033[0m")
			fmt.Println("")
		} else {
			if i > m.list_t.Size() || i <= 0 {
				fmt.Println("\033[1;31mNumero de tarea no existente\033[0m")
				fmt.Println("")
			} else {
				return i
			}
		}
	}
}

package sgt

import (
	"bluelock/linkedlist"
	"bluelock/stack"
	"fmt"
)

type ListaDeTareas struct {
	list              linkedlist.LinkedList[Tarea]
	pilaTareasActivas stack.Stack[Tarea]
}

func NewListaDeTareas() *ListaDeTareas {
	return &ListaDeTareas{}
}

func (l *ListaDeTareas) ListarTareas() {
	if l.list.Size() == 0 {
		fmt.Println("\033[1;31mNo se puede listar si la lista esta vacia\033[0m")
		fmt.Println("")
		return
	}
	fmt.Println("\033[1;36mLista de Tareas")
	fmt.Println("---------------\033[0m")
	fmt.Print("\033[1;33m")
	for i := 0; i < l.list.Size(); i++ {
		tarea := l.list.Get(i)
		fmt.Println(i+1, ".", tarea.String())
		tarea.ListarSubTareas()
	}
	fmt.Print("\033[0m")
	fmt.Println("")
}

func (l *ListaDeTareas) AgregarNuevaTarea(tarea Tarea) {
	l.list.Append(tarea)
	fmt.Println("\033[1;32mTarea (" + tarea.nombre + ") agregada correctamente\033[0m")
	fmt.Println("")
}

func (l *ListaDeTareas) EliminarTarea(tarea Tarea) {
	if !l.pilaTareasActivas.IsEmpty() && l.VerTareaActual() == tarea {
		l.list.Remove(tarea)
		fmt.Println("\033[1;32mTarea (" + tarea.nombre + ") eliminada correctamente\033[0m")
		l.EliminarTareaActual()
		l.IniciarTareaInterumpida()
	} else if tarea.ConsultarEstado() == "Pendiente" || tarea.ConsultarEstado() == "Completada" {
		l.list.Remove(tarea)
		fmt.Println("\033[1;32mTarea (" + tarea.nombre + ") eliminada correctamente\033[0m")
	} else {
		fmt.Println("\033[1;31mNo se puede eliminar una tarea interumpida\033[0m")
	}
	fmt.Println("")
}

func (l *ListaDeTareas) ComenzarTarea(tarea Tarea) {
	tarea_mod := tarea
	tarea_mod.Iniciar()
	l.list.UpdateNodeValue(tarea, tarea_mod)
	if !l.pilaTareasActivas.IsEmpty() {
		tarea_actual := l.VerTareaActual()
		tarea_actual.Interrumpir()
		l.list.UpdateNodeValue(l.EliminarTareaActual(), tarea_actual)
		l.pilaTareasActivas.Push(tarea_actual)
	}
	l.pilaTareasActivas.Push(tarea_mod)
}

func (l *ListaDeTareas) PrepararColaDeTareas(tiempoTotal float32) {
	if l.list.Size() == 0 {
		fmt.Println("\033[1;31mNo se puede preparar una cola con la lista vacia\033[0m")
		fmt.Println("")
		return
	}
	fmt.Println("\033[1;36mTareas sugeridas")
	fmt.Println("---------------\033[0m")
	contador := 0
	slice_tarea := l.list.ToSlice()
	QuickSortCola(slice_tarea)
	//sort.Slice(slice_tarea, func(i, j int) bool {
	//	return (PrioridadValue(slice_tarea[i].prioridad) > PrioridadValue(slice_tarea[j].prioridad)) || ((PrioridadValue(slice_tarea[i].prioridad) == PrioridadValue(slice_tarea[j].prioridad)) && (slice_tarea[i].TiempoTotal() < slice_tarea[j].TiempoTotal()))
	//})
	fmt.Print("\033[1;33m")
	for _, tarea := range slice_tarea {
		if tarea.TiempoTotal() < tiempoTotal {
			fmt.Println(tarea.String())
			tiempoTotal -= tarea.TiempoTotal()
			contador++
		}
	}
	fmt.Print("\033[0m")
	if contador == 0 {
		fmt.Println("\033[1;31mNo se puede preparar una cola en el tiempo estimado\033[0m")
	}
	fmt.Println("")
}

func (l *ListaDeTareas) BuscarTareas(clave, tipo string) {
	contador := 0
	fmt.Println("\033[1;33m")
	for i := 0; i < l.list.Size(); i++ {
		tarea := l.list.Get(i)
		switch tipo {
		case "palabra clave":
			if tarea.ContienePalabra(clave) {
				fmt.Println(i+1, ".", tarea.String())
				tarea.ListarSubTareas()
				contador++
			}
		case "etiqueta":
			if tarea.etiquetas.Contains(clave) {
				fmt.Println(i+1, ".", tarea.String())
				tarea.ListarSubTareas()
				contador++
			}
		}
	}
	fmt.Println("\033[0m")
	if contador == 0 {
		fmt.Println("\033[1;31mEsta busqueda no a encontrado ninguna tarea\033[0m")
	}
	fmt.Println("")
}

func (l *ListaDeTareas) ReordenarTareas(tipo string) {
	if l.list.Size() == 0 {
		fmt.Println("\033[1;31mNo hay tareas para reordenar\033[0m")
		fmt.Println("")
		return
	}
	lista_ordenada := l.list.ToSlice()
	switch tipo {
	case "Prioridad":
		QuickSortPrioridad(lista_ordenada)
		//sort.Slice(lista_ordenada, func(i, j int) bool {
		//	return (PrioridadValue(lista_ordenada[i].prioridad) > PrioridadValue(lista_ordenada[j].prioridad))
		//})
	case "Duracion":
		QuickSortTiempoTotal(lista_ordenada)
		//sort.Slice(lista_ordenada, func(i, j int) bool {
		//	return lista_ordenada[i].TiempoTotal() < lista_ordenada[0].TiempoTotal()
		//})
	case "Subtareas":
		QuickSortSubTareas(lista_ordenada)
		//sort.Slice(lista_ordenada, func(i, j int) bool {
		//	return lista_ordenada[i].lista_de_subtareas.Size() < lista_ordenada[j].lista_de_subtareas.Size()
		//})
	}
	lista_nueva := linkedlist.NewLinkedList[Tarea]()
	for _, tarea := range lista_ordenada {
		lista_nueva.Append(tarea)
	}
	l.list = *lista_nueva
	fmt.Println("\033[1;32mLista de tareas reordanada correctamente, dirijase a la lista de tareas\033[0m")
	fmt.Println("")
}

func PrioridadValue(prioridad string) int {
	switch prioridad {
	case "Alta Prioridad":
		return 2
	case "Baja Prioridad":
		return 1
	default:
		return 0
	}
}

func (l *ListaDeTareas) Size() int {
	return l.list.Size()
}

//Stack

func (l *ListaDeTareas) EliminarTareaActual() Tarea {
	tarea, _ := l.pilaTareasActivas.Pop()
	return tarea
}

func (l *ListaDeTareas) VerTareaActual() Tarea {
	tarea, _ := l.pilaTareasActivas.Top()
	return tarea
}

func (l *ListaDeTareas) CompletarTareaActual() {
	if !l.pilaTareasActivas.IsEmpty() {
		tarea_actual := l.VerTareaActual()
		tarea_mod := tarea_actual
		tarea_mod.Completar()
		l.list.UpdateNodeValue(tarea_actual, tarea_mod)
		l.EliminarTareaActual()
		l.IniciarTareaInterumpida()
	} else {
		fmt.Println("\033[1;31mNo hay tarea actual para completar\033[0m")
		fmt.Println("")
	}
}

func (l *ListaDeTareas) AbandonarTareaActual() {
	if !l.pilaTareasActivas.IsEmpty() {
		tarea_actual := l.VerTareaActual()
		tarea_mod := tarea_actual
		tarea_mod.Abandonar()
		l.list.UpdateNodeValue(tarea_actual, tarea_mod)
		l.EliminarTareaActual()
		l.IniciarTareaInterumpida()
	} else {
		fmt.Println("\033[1;31mNo hay tarea actual para abandonar\033[0m")
		fmt.Println("")
	}
}

func (l *ListaDeTareas) IniciarTareaInterumpida() {
	if !l.pilaTareasActivas.IsEmpty() {
		tarea_actual := l.VerTareaActual()
		tarea_actual.Iniciar()
		l.list.UpdateNodeValue(l.EliminarTareaActual(), tarea_actual)
		l.pilaTareasActivas.Push(tarea_actual)
	}
}

func QuickSortPrioridad(array []Tarea) {
	if len(array) < 2 {
		return
	}
	pivot := array[0]
	i := 1
	j := len(array) - 1

	for i < j {
		for i < j && PrioridadValue(array[i].prioridad) > PrioridadValue(pivot.prioridad) {
			i++
		}

		for j > 0 && PrioridadValue(array[j].prioridad) < PrioridadValue(pivot.prioridad) {
			j--
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	if PrioridadValue(array[j].prioridad) > PrioridadValue(pivot.prioridad) {
		array[0], array[j] = array[j], array[0]
	}

	QuickSortPrioridad(array[:j])
	QuickSortPrioridad(array[j+1:])
}

func QuickSortTiempoTotal(array []Tarea) {
	if len(array) < 2 {
		return
	}
	pivot := array[0]
	i := 1
	j := len(array) - 1

	for i < j {

		for array[i].TiempoTotal() < pivot.TiempoTotal() && i < len(array)-1 {
			i++
		}

		for array[j].TiempoTotal() > pivot.TiempoTotal() && j > 0 {
			j--
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	if array[j].TiempoTotal() < pivot.TiempoTotal() {
		array[0], array[j] = array[j], array[0]
	}

	QuickSortTiempoTotal(array[:j])
	QuickSortTiempoTotal(array[j+1:])
}

func QuickSortSubTareas(array []Tarea) {
	if len(array) < 2 {
		return
	}
	pivot := array[0]
	i := 1
	j := len(array) - 1

	for i < j {

		for array[i].lista_de_subtareas.Size() < pivot.lista_de_subtareas.Size() && i < len(array)-1 {
			i++
		}

		for array[j].lista_de_subtareas.Size() > pivot.lista_de_subtareas.Size() && j > 0 {
			j--
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	if array[j].lista_de_subtareas.Size() < pivot.lista_de_subtareas.Size() {
		array[0], array[j] = array[j], array[0]
	}

	QuickSortSubTareas(array[:j])
	QuickSortSubTareas(array[j+1:])
}

func QuickSortCola(array []Tarea) {
	if len(array) < 2 {
		return
	}
	pivot := array[0]
	i := 1
	j := len(array) - 1

	for i < j {

		for (PrioridadValue(array[i].prioridad) > PrioridadValue(pivot.prioridad)) || ((PrioridadValue(array[i].prioridad) == PrioridadValue(pivot.prioridad)) && (array[i].TiempoTotal() < pivot.TiempoTotal())) && i < len(array)-1 {
			i++
		}

		for (PrioridadValue(array[j].prioridad) < PrioridadValue(pivot.prioridad)) || ((PrioridadValue(array[j].prioridad) == PrioridadValue(pivot.prioridad)) && (array[j].TiempoTotal() > pivot.TiempoTotal())) && j > 0 {
			j--
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	if (PrioridadValue(array[j].prioridad) > PrioridadValue(pivot.prioridad)) || ((PrioridadValue(array[j].prioridad) == PrioridadValue(pivot.prioridad)) && (array[j].TiempoTotal() < pivot.TiempoTotal())) {
		array[0], array[j] = array[j], array[0]
	}

	QuickSortCola(array[:j])
	QuickSortCola(array[j+1:])
}

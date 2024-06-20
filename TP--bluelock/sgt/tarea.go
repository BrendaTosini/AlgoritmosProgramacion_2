package sgt

import (
	"bluelock/set"
	"fmt"
	"strings"
	"time"
)

type Tarea struct {
	nombre             string
	duracion           float32
	prioridad          string
	lista_de_subtareas ListaDeTareas
	estado             string
	etiquetas          set.Set[string]
}

func NewTarea(nombre string, duracion float32, estado string, etiquetas *set.Set[string], lista_de_subtareas *ListaDeTareas) *Tarea {
	return &Tarea{nombre: nombre, duracion: duracion, prioridad: "Sin Prioridad", lista_de_subtareas: *lista_de_subtareas, estado: estado, etiquetas: *etiquetas}
}

// Acciones sobre la tarea

func (t *Tarea) EditarNombre(nombre string) {
	fmt.Println("\033[1;32mNombre: " + t.nombre + " cambiado a: " + nombre + "\033[0m")
	fmt.Println("")
	t.nombre = nombre
}

func (t *Tarea) EditarDuracion(duracion float32) {
	duracion_vieja := formatDuracion(time.Duration((t.duracion + 0.01) * float32(time.Hour)))
	duracion_nueva := formatDuracion(time.Duration((duracion + 0.01) * float32(time.Hour)))
	fmt.Println("\033[1;32mDuracion: " + string(duracion_vieja) + " cambiada a: " + duracion_nueva + "\033[0m")
	fmt.Println("")
	t.duracion = duracion
}

func (t *Tarea) SubirPrioridad() {
	t.prioridad = "Alta Prioridad"
	fmt.Println("\033[1;32mPrioridad subida correctamente a la tarea (" + t.nombre + ")\033[0m")
	fmt.Println("")
}

func (t *Tarea) BajarPrioridad() {
	t.prioridad = "Baja Prioridad"
	fmt.Println("\033[1;32mPrioridad bajada correctamente a la tarea (" + t.nombre + ")\033[0m")
	fmt.Println("")
}

func (t *Tarea) AgregarEtiqueta(etiqueta string) {
	t.etiquetas.Add(etiqueta)
	fmt.Println("\033[1;32mEtiqueta (" + etiqueta + ") agregada correctamente a la tarea (" + t.nombre + ")\033[0m")
	fmt.Println("")
}

func (t *Tarea) EliminarEtiqueta(etiqueta string) {
	if t.etiquetas.Contains(etiqueta) {
		t.etiquetas.Remove(etiqueta)
		fmt.Println("\033[1;32mEtiqueta (" + etiqueta + ") eliminada correctamente a la tarea (" + t.nombre + ")\033[0m")
	} else {
		fmt.Println("\033[1;31mNo se puede eliminar una etiqueta que no existe\033[0m")
	}
	fmt.Println("")
}

func (t *Tarea) Iniciar() {
	if t.estado == "Pendiente" || t.estado == "Interrumpida" {
		t.estado = "En Progreso"
		fmt.Println("\033[1;32mTarea (" + t.nombre + ") En Progreso\033[0m")
		for i := 0; i < t.lista_de_subtareas.Size(); i++ {
			sub_tarea := t.lista_de_subtareas.list.Get(i)
			sub_tarea_mod := sub_tarea
			sub_tarea_mod.Iniciar()
			t.lista_de_subtareas.list.UpdateNodeValue(sub_tarea, sub_tarea_mod)
		}
	} else {
		fmt.Println("\033[1;31mEl estado actual no permite pasar a en progreso\033[0m")
	}
	fmt.Println("")
}

func (t *Tarea) Interrumpir() {
	if t.estado == "En Progreso" {
		t.estado = "Interrumpida"
		fmt.Println("\033[1;32mTarea (" + t.nombre + ") Interrumpida\033[0m")
		for i := 0; i < t.lista_de_subtareas.Size(); i++ {
			sub_tarea := t.lista_de_subtareas.list.Get(i)
			sub_tarea_mod := sub_tarea
			sub_tarea_mod.Interrumpir()
			t.lista_de_subtareas.list.UpdateNodeValue(sub_tarea, sub_tarea_mod)
		}
	} else {
		fmt.Println("\033[1;31mEl estado actual no permite pasar a interrumpida\033[0m")
	}
	fmt.Println("")
}

func (t *Tarea) Completar() {
	if t.estado == "En Progreso" {
		t.estado = "Completada"
		fmt.Println("\033[1;32mTarea (" + t.nombre + ") completada\033[0m")
		for i := 0; i < t.lista_de_subtareas.Size(); i++ {
			sub_tarea := t.lista_de_subtareas.list.Get(i)
			sub_tarea_mod := sub_tarea
			sub_tarea_mod.Completar()
			t.lista_de_subtareas.list.UpdateNodeValue(sub_tarea, sub_tarea_mod)
		}
	} else {
		fmt.Println("\033[1;31mEl estado actual no permite pasar a completada\033[0m")
	}
	fmt.Println("")
}

func (t *Tarea) Abandonar() {
	if t.estado == "En Progreso" {
		t.estado = "Pendiente"
		fmt.Println("\033[1;32mTarea (" + t.nombre + ") a pendiente\033[0m")
		for i := 0; i < t.lista_de_subtareas.Size(); i++ {
			sub_tarea := t.lista_de_subtareas.list.Get(i)
			sub_tarea_mod := sub_tarea
			sub_tarea_mod.Abandonar()
			t.lista_de_subtareas.list.UpdateNodeValue(sub_tarea, sub_tarea_mod)
		}
	} else {
		fmt.Println("\033[1;31mEl estado actual no permite pasar a pendiente\033[0m")
	}
	fmt.Println("")
}

// Acciones sobre las subtareas de la tarea

func (t *Tarea) AgregarSubTarea(subt Tarea) {
	t.lista_de_subtareas.AgregarNuevaTarea(subt)
}

func (t *Tarea) EliminarSubTarea(subt Tarea) {
	t.lista_de_subtareas.EliminarTarea(subt)
}

func (t *Tarea) ListarSubTareas() {
	for i := 0; i < t.lista_de_subtareas.Size(); i++ {
		subtarea := t.lista_de_subtareas.list.Get(i)
		fmt.Println("  ", i+1, ".", subtarea.String())
		subtarea.ListarSubTareas()
	}
}

func (t *Tarea) ConsultarEstado() string {
	return t.estado
}

func (t *Tarea) ConsultarDuracion() float32 {
	return t.duracion
}

func (t *Tarea) String() string {
	duracion := formatDuracion(time.Duration((t.duracion + 0.01) * float32(time.Hour)))
	tiempo_total := formatDuracion(time.Duration((t.TiempoTotal() + 0.01) * float32(time.Hour)))
	return t.nombre + " " + t.estado + " " + t.prioridad + " " + "(" + duracion + ") " + t.etiquetas.String() + " Tiempo total: " + tiempo_total
}

func formatDuracion(duracion time.Duration) string {
	horas := int(duracion.Hours())
	minutos := int(duracion.Minutes()) % 60

	return fmt.Sprintf("%dh %dm", horas, minutos)
}

// Total de duracion de las tarea + subtareas
func (t *Tarea) TiempoTotal() float32 {
	var duracionTotal float32 = 0.0
	for i := 0; i < t.lista_de_subtareas.Size(); i++ {
		sub_tarea := t.lista_de_subtareas.list.Get(i)
		duracionTotal += sub_tarea.TiempoTotal()
	}
	return duracionTotal + t.duracion

}

func (t *Tarea) ContienePalabra(palabra string) bool {
	palabras := strings.Split(t.nombre, " ")
	set_palabras := set.NewSet(palabras...)
	return set_palabras.Contains(palabra)
}

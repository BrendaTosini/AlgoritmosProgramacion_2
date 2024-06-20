package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Selector de actividades recursivo
// Reescribir la función que resuelve el problema de seleccionar actividades en forma recursiva
func SelectorRecursivo(actividades []Actividad) []Actividad {

	var seleccionadas []Actividad
	var aux []Actividad

	if len(actividades) == 0 {
		return seleccionadas
	}

	seleccionadas = append(seleccionadas, actividades[0])

	for i := 1; i < len(actividades); i++ {
		if actividades[i].Inicio >= actividades[0].Fin {
			aux = append(aux, actividades[i])
		}
	}

	seleccionadas = append(seleccionadas, SelectorRecursivo(aux)...)

	return seleccionadas

}

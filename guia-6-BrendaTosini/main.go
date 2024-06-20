package main

import (
	"fmt"
)

func main() {
	/*dic := dictionary.NewDictionary[string, []string]()
	fmt.Println(AprobadasPorAlumno(dic))*/

	/*d := dictionary.NewDictionary[int, []string]()
	s := []string{"Individual F US Open", "Doble Mixto Wimbledon"}
	s2 := []string{"Dobles F US Open"}
	s3 := []string{"Dobles F Wimbledon"}
	s4 := []string{"Dobles F US Open"}
	d.Put(1990, s)
	d.Put(1988, s2)
	d.Put(1991, s3)
	d.Put(1988, s4)
	diccionario := dictionary.NewDictionary[string, dictionary.Dictionary[int, []string]]()
	diccionario.Put("GB", d)
	diccionario.Put("Steffi", d)
	DeportistasPremiadosXAño(diccionario)*/

	/*d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 60)
	d.Put("Leo", 61)
	d.Put("Fabi", 36)
	d.Put("Leo", 60)
	d.Put("Fede", 36)
	d.Put("Vale", 40)
	d.Put("Leo", 50)
	fmt.Println("Clave: valor en el diccionario (String, Int)")
	fmt.Println(d.String())
	fmt.Println("--------------------")
	d.Remove("Fede")
	fmt.Println("Borre a fede")
	fmt.Println("Al pedirle al diccionario el valor para fede me da como resultado: ", d.Get("Fede"))
	fmt.Println("--------------------")
	fmt.Println("Clave: valor en el diccionario sin fede(String, Int)")
	fmt.Println(d.String())
	fmt.Println("--------------------")
	ds := dictionary.NewDictionary[string, set.Set[int]]()
	s := set.NewSet[int]()
	s.Add(100)
	s.Add(222)
	s.Add(333)
	ss := set.NewSet[int]()
	ss.Add(1)
	ss.Add(2)
	ss.Add(3)
	ds.Put("uno", *s)
	ds.Put("dos", *ss)
	ds.Put("tres", *ss)
	fmt.Println("Clave: valor en el diccionario (String, Set[Int])")
	fmt.Println(ds.String())
	fmt.Println("--------------------")
	dl := dictionary.NewDictionary[string, linkedlist.LinkedList[string]]()
	l := linkedlist.NewLinkedList[string]()
	l.InsertAt("A", 0)
	l.InsertAt("B", 1)
	l.InsertAt("C", 2)
	dl.Put("Lista 1", *l)
	ll := linkedlist.NewLinkedList[string]()
	ll.InsertAt("D", 0)
	ll.InsertAt("E", 1)
	ll.InsertAt("F", 2)
	dl.Put("Lista 2", *ll)
	fmt.Println("Clave: valor en el diccionario (String, LinkedList[String])")
	fmt.Println(dl.String())*/
	cadena := "banana"
	a := CantidadCaracter(cadena, "a")
	fmt.Println(a)
}

/*func DeportistasPremiadosXAño(deportistas dictionary.Dictionary[string, dictionary.Dictionary[int, []string]]) dictionary.Dictionary[int, dictionary.Dictionary[string, []string]] {

	diccionario := dictionary.NewDictionary[int, dictionary.Dictionary[string, []string]]()
	nombres := deportistas.GetKeys()
	for _, nombre := range nombres {
		diccionarioCampeonatos := deportistas.Get(nombre)
		sliceAños := diccionarioCampeonatos.GetKeys()
		for _, año := range sliceAños {

			dd2 := dictionary.NewDictionary[string, []string]()
			if diccionario.Contains(año) {
				dd2 = diccionario.Get(año)
			}

			sliceCampeonatos := diccionarioCampeonatos.Get(año)
			for _, campeonato := range sliceCampeonatos {
				if dd2.Contains(nombre) {
					sliceValores := dd2.Get(nombre)
					sliceValores = append(sliceValores, campeonato)
					dd2.Put(nombre, sliceValores)
				} else {
					slice := []string{}
					slice = append(slice, campeonato)
					dd2.Put(nombre, slice)
				}
			}
			diccionario.Put(año, dd2)

		}
	}
	fmt.Println(diccionario)
	return diccionario
}*/

/*func AprobadasPorAlumno(materias dictionary.Dictionary[string, []string]) dictionary.Dictionary[string, []string] {

	diccionario := dictionary.NewDictionary[string, []string]()

	if materias.Size() == 0 {
		return diccionario
	}

	claves := materias.GetKeys()

	for _, materia := range claves {
		alumnos := materias.Get(materia)
		for _, alumno := range alumnos {
			if !diccionario.Contains(alumno) {
				materias := []string{}
				materias = append(materias, materia)
				diccionario.Put(alumno, materias)
			} else {
				slice := diccionario.Get(alumno)
				slice = append(slice, materia)
				diccionario.Put(alumno, slice)
			}

		}
	}
	return diccionario
}
*/

func CantidadCaracter(cadena, caracter string) int {

	if len(cadena) == 0 {
		return 0
	}

	if len(cadena) == 1 && string(cadena[0]) == caracter {
		return 1
	}
	if len(cadena) == 1 && string(cadena[0]) != caracter {
		return 0
	}

	m := len(cadena) / 2
	izCadena := cadena[:m]
	derCadena := cadena[m:]

	ci := CantidadCaracter(izCadena, caracter)
	cd := CantidadCaracter(derCadena, caracter)

	return ci + cd

}

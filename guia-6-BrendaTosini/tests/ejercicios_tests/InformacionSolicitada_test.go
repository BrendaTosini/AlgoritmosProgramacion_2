package tests
import (
	"guia6/ejercicios"
	"guia6/dictionary"
	"testing"
	"reflect"
	"sort"
)

func TestInformacionSolicitada(t *testing.T) {
	entrada := dictionary.NewDictionary[string, []string]()
	sl1:= []string{"Ana", "Pedro"}
	sl2 :=[]string{"Ana"}
	entrada.Put("Mie 10", sl1)
	entrada.Put("Vie 12", sl2)
	salida := ejercicios.InformacionSolicitada(entrada)
	val1 := []string{"Mie 10"}
	val2 := []string{"Mie 10", "Vie 12"}
	if salida.Size() != 2 {
		t.Errorf("Tamaño %d, debería ser %d", salida.Size(), 2)
	}
	if !salida.Contains("Ana"){
		t.Errorf("Contains(Ana) = %t, debería ser %t", salida.Contains("Ana"), true)
	}else{
		k1 := salida.Get("Ana")
		sort.Strings(k1)
		if !reflect.DeepEqual(k1, val2) {
			t.Errorf("Los valores son %v, pero deberían ser %v", k1, val2)
		}
	}
	if !salida.Contains("Pedro"){
		t.Errorf("Contains(Pedro) = %t, debería ser %t", salida.Contains("Pedro"), true)
	}else{
		k2 := salida.Get("Pedro")
		sort.Strings(k2)
		if !reflect.DeepEqual(k2, val1) {
			t.Errorf("Los valores son %v, pero deberían ser %v", k2, val1)
		}
	}
}
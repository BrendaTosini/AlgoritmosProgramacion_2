package ordenamiento

func Burbujeo(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-1; j++ {

			if arr[j] > arr[j+1] {
				aux := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = aux
			}

		}

	}
	return arr

}

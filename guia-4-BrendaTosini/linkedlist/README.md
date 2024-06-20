## linkedlist

Nueva versión de la lista enlazada simple, con la siguiente estructura: 

```go
type LinkedList[T comparable] struct {
	head *node[T] // puntero al primer nodo
	tail *node[T] // puntero al último nodo
	size int // tamaño de la lista
}
```

y las siguientes operaciones

```go

func NewLinkedList[T comparable]() *LinkedList[T]
func (l *LinkedList[T]) InsertAt(value T, position int)
func (l *LinkedList[T]) RemoveAt(position int)
func (l *LinkedList[T]) Search(value T) int 
func (l *LinkedList[T]) Get(position int) T
func (l *LinkedList[T]) Size() int
func (l *LinkedList[T]) Contains(value T) bool
func (l *LinkedList[T]) IsEmpty() bool
func (l *LinkedList[T]) String() string 
```
Ver el código para más detalles de la implementación

El archivo `linkedlist_test.go` contiene pruebas unitarias de la lista

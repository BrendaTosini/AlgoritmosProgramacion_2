## set
Implementa el TAD conjunto sobre una lista enlazada simple.

> Un conjunto no contiene elementos duplicados ni preserva el orden de los elementos

```go
type Set[T comparable] struct {
	elementos linkedlist.LinkedList[T]
}
```
set contiene los siguientes métodos:

```go
func (s *Set[T]) Contains(element T) bool
func (s *Set[T]) Add(element T)
func (s *Set[T]) Remove(element T)
func (s *Set[T]) Size() int
func (s *Set[T]) String() string
func (s *Set[T]) Values() []T 
```
y soporta las siguientes operaciones que devuelven un nuevo conjunto:
```go
func NewSet[T comparable](elementos ...T) *Set[T]
func Union[T comparable](s1, s2 *Set[T]) *Set[T]
func Intersection[T comparable](s1, s2 *Set[T]) *Set[T]
func Difference[T comparable](s1, s2 *Set[T]) *Set[T]
func Subset[T comparable](s1, s2 *Set[T]) bool
func Equal[T comparable](s1, s2 *Set[T]) bool
```
Ver el código para más detalles
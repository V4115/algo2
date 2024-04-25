package lista

const (
	PANIC_MSG_BORRAR_LISTA_VACIA      = "La lista esta vacia"
	PANIC_MSG_VER_PRIMERO_LISTA_VACIA = "La lista esta vacia"
	PANIC_MSG_VER_ULTIMO_LISTA_VACIA  = "La lista esta vacia"
)

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func _crear_nodo[T any](dato T, prox *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato, prox}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	if lista.EstaVacia() {
		lista.primero = _crear_nodo(dato, nil)
		lista.ultimo = lista.primero
	} else {
		lista.primero = _crear_nodo(dato, lista.primero)
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	if lista.EstaVacia() {
		lista.primero = _crear_nodo(dato, nil)
		lista.ultimo = lista.primero
	} else {
		aux := lista.ultimo
		lista.ultimo = _crear_nodo(dato, nil)
		aux.prox = lista.ultimo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(PANIC_MSG_BORRAR_LISTA_VACIA)
	} else if lista.largo == 1 {

		aux := lista.primero.dato
		lista.primero, lista.ultimo = nil, nil
		lista.largo--
		return aux
	} else {

		aux := lista.primero.dato
		lista.primero = lista.primero.prox
		lista.largo--
		return aux
	}
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(PANIC_MSG_VER_PRIMERO_LISTA_VACIA)
	} else {
		return lista.primero.dato
	}
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(PANIC_MSG_VER_ULTIMO_LISTA_VACIA)
	} else {
		return lista.ultimo.dato
	}
}

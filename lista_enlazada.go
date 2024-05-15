package lista

const (
	PANIC_MSG_BORRAR_LISTA_VACIA      = "La lista esta vacia"
	PANIC_MSG_VER_PRIMERO_LISTA_VACIA = "La lista esta vacia"
	PANIC_MSG_VER_ULTIMO_LISTA_VACIA  = "La lista esta vacia"
	PANIC_ITERADOR_FINALIZO           = "El iterador termino de iterar"
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
		lista.ultimo.prox = _crear_nodo(dato, nil)
		lista.ultimo = lista.ultimo.prox
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(PANIC_MSG_BORRAR_LISTA_VACIA)
	}

	aux := lista.primero.dato
	lista.primero = lista.primero.prox
	lista.largo--
	return aux
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

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {

	for actual := lista.primero; actual != nil; actual = actual.prox {
		if visitar(actual.dato) == false {
			break
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{list: lista, actual: lista.primero, anterior: nil}
}

type iteradorListaEnlazada[T any] struct {
	list     *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func (iterador *iteradorListaEnlazada[T]) sinElementos() {
	if !iterador.HaySiguiente() {
		panic(PANIC_ITERADOR_FINALIZO)
	}
}

func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	iterador.sinElementos()
	return iterador.actual.dato
}

func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorListaEnlazada[T]) Siguiente() {
	iterador.sinElementos()
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.prox
}

func (iterador *iteradorListaEnlazada[T]) Insertar(dato T) {
	if iterador.anterior == nil {
		iterador.list.InsertarPrimero(dato)
		iterador.actual = iterador.list.primero
	} else {
		iterador.anterior.prox = _crear_nodo(dato, iterador.actual)
		iterador.actual = iterador.anterior.prox
		iterador.list.largo++
	}

}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	iterador.sinElementos()
	dato := iterador.actual.dato

	if iterador.HaySiguiente() {
		iterador.actual = iterador.actual.prox
	}
	if iterador.anterior != nil {
		iterador.anterior.prox = iterador.actual
	}

	iterador.list.largo--
	return dato
}

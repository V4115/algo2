package lista

type Lista[T any] interface {

	//Devuelve verdadero si la lista esta vacio, sino falso
	EstaVacia() bool

	//Recibe un dato T y lo inserta en la primera posicion de la lista
	InsertarPrimero(T)

	//Recibe un dato y lo inserta en la ultima posicion de la lista
	InsertarUltimo(T)

	//Borra el primer dato de la lista
	BorrarPrimero() T

	//Muestra el primer dato alojado en la lista
	VerPrimero() T

	//Duestra el ultimo dato alojado en la lista
	VerUltimo() T

	//Devuelve la cantidad de datos actuales en la lista
	Largo() int

	//Iterar(visitar func(T) bool)
	//Iterador() IteradorLista[T]
}

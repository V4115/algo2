package algo2_tdas_lista

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

	//Itera la lista mientras que tenga elementos y el resultado de la funcion anonima que recibe sea true.
	Iterar(visitar func(T) bool)

	//Crea un dato de tipo Iterador que permite recorrer la lista
	Iterador() IteradorLista[T]
}
type IteradorLista[T any] interface {

	//Devuelve verdadero si hay al menos un elemento por ver
	HaySiguiente() bool

	//Devuelve el dato T del actual
	VerActual() T

	//Se mueve a la posicion siguiente de la lista
	Siguiente()

	//Inserta un dato T en la lista. Haciendo que el actual sea el siguiente de ese dato.
	Insertar(T)

	//Borra el dato actual de la lista.
	Borrar() T
}

package lista_test

import (
	"strconv"
	"strings"

	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	LONGITUD_VACIA          = 0
	PANIC_LISTA_VACIA       = "La lista esta vacia"
	PANIC_ITERADOR_FINALIZO = "El iterador termino de iterar"
	testVolLen6             = 10
	testVolLen5             = 100
	testVolLen4             = 1000
	testVolLen3             = 10000
	testVolLen2             = 100000
	testVolLen1             = 1000000

	testBordeVaciaLen = 200
)

func TestListaaVacia(t *testing.T) {

	lista1 := TDALista.CrearListaEnlazada[int]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	lista3 := TDALista.CrearListaEnlazada[rune]()
	lista4 := TDALista.CrearListaEnlazada[float64]()
	lista5 := TDALista.CrearListaEnlazada[bool]()

	_estaVacia[int](t, lista1)
	_estaVacia[string](t, lista2)
	_estaVacia[rune](t, lista3)
	_estaVacia[float64](t, lista4)
	_estaVacia[bool](t, lista5)
}
func _estaVacia[T any](t *testing.T, lista TDALista.Lista[T]) {

	require.True(t, lista.EstaVacia())
	require.EqualValues(t, lista.Largo(), 0)
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista.BorrarPrimero() }, "No hay un Panic al querer borrar el rpimer dato de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_ULTIMO_LISTA_VACIA, func() { lista.VerUltimo() }, "No hay un Panic al intentar ver el ultimo de una lista vacia")
}

func TestInsertarPrimero(t *testing.T) {
	var (
		vec1 = []int{10, 9, -15, 0, 7, -12, 1}
		vec2 = []string{"hola", " ", "1234", "sdafasdfgasgdcvfdsafs", "#$%&/(*][]:;,.,)", "", "josé Valls"}
		vec3 = []rune{'s', 'a', 'ñ', ' ', '今', '日', '✌'}
		vec4 = []float32{1.0, 0.1111, 1.11, -2.3, 4.332344, -10.0, 324.134, 2}
		vec5 = []bool{false, true, true, true, false, true, false, false, true, false}
	)

	lista1 := TDALista.CrearListaEnlazada[int]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	lista3 := TDALista.CrearListaEnlazada[rune]()
	lista4 := TDALista.CrearListaEnlazada[float32]()
	lista5 := TDALista.CrearListaEnlazada[bool]()

	_insertarPrimero(t, func(a int, b int) bool { return a == b }, lista1, vec1)
	_insertarPrimero(t, func(a string, b string) bool { return strings.Compare(a, b) == 0 }, lista2, vec2)
	_insertarPrimero(t, func(a rune, b rune) bool { return a == b }, lista3, vec3)
	_insertarPrimero(t, func(a float32, b float32) bool { return a == b }, lista4, vec4)
	_insertarPrimero(t, func(a bool, b bool) bool { return a == b }, lista5, vec5)
}

func TestInsertarUltimo(t *testing.T) {
	var (
		vec1 = []int{10, 9, -15, 0, 7, -12, 1}
		vec2 = []string{"hola", " ", "1234", "sdafasdfgasgdcvfdsafs", "#$%&/(*][]:;,.,)", "", "josé Valls"}
		vec3 = []rune{'s', 'a', 'ñ', ' ', '今', '日', '✌'}
		vec4 = []float32{1.0, 0.1111, 1.11, -2.3, 4.332344, -10.0, 324.134, 2}
		vec5 = []bool{false, true, true, true, false, true, false, false, true, false}
	)

	lista1 := TDALista.CrearListaEnlazada[int]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	lista3 := TDALista.CrearListaEnlazada[rune]()
	lista4 := TDALista.CrearListaEnlazada[float32]()
	lista5 := TDALista.CrearListaEnlazada[bool]()

	_insertarUltimo(t, func(a int, b int) bool { return a == b }, lista1, vec1)
	_insertarUltimo(t, func(a string, b string) bool { return strings.Compare(a, b) == 0 }, lista2, vec2)
	_insertarUltimo(t, func(a rune, b rune) bool { return a == b }, lista3, vec3)
	_insertarUltimo(t, func(a float32, b float32) bool { return a == b }, lista4, vec4)
	_insertarUltimo(t, func(a bool, b bool) bool { return a == b }, lista5, vec5)
}

func _insertarPrimero[T any](t *testing.T, comparador func(a T, b T) bool, lista TDALista.Lista[T], datos []T) {

	for i := 0; i < len(datos); i++ {
		lista.InsertarPrimero(datos[i])
		require.True(t, comparador(lista.VerPrimero(), datos[i]))
		require.True(t, comparador(lista.VerUltimo(), datos[0]))
		require.False(t, lista.EstaVacia())
	}
	for i := len(datos) - 1; i >= 0; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), datos[i])
	}
	_estaVacia(t, lista)
}

func _insertarUltimo[T any](t *testing.T, comparador func(a T, b T) bool, lista TDALista.Lista[T], datos []T) {

	for i := 0; i < len(datos); i++ {
		lista.InsertarUltimo(datos[i])
		require.True(t, comparador(lista.VerUltimo(), datos[i]))
		require.True(t, comparador(lista.VerPrimero(), datos[0]))
		require.False(t, lista.EstaVacia())
	}
	for i := 0; i < len(datos); i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), datos[i])
	}
	_estaVacia(t, lista)
}

func TestVolumen(t *testing.T) {
	_testVolumen(t, testVolLen1)
	_testVolumen(t, testVolLen2)
	_testVolumen(t, testVolLen3)
	_testVolumen(t, testVolLen4)
	_testVolumen(t, testVolLen5)
	_testVolumen(t, testVolLen6)
}

func _testVolumen(t *testing.T, vol int) {
	lista1 := TDALista.CrearListaEnlazada[int]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	lista3 := TDALista.CrearListaEnlazada[rune]()
	lista4 := TDALista.CrearListaEnlazada[float64]()
	lista5 := TDALista.CrearListaEnlazada[bool]()

	//int
	for i := 0; i < vol; i++ {
		lista1.InsertarPrimero(i)
	}

	for i := vol - 1; i >= 0; i-- {
		require.Equal(t, lista1.BorrarPrimero(), i)
	}
	_estaVacia(t, lista1)
	//string
	for i := 0; i < vol; i++ {
		lista2.InsertarPrimero(strconv.Itoa(i))
	}

	for i := vol - 1; i >= 0; i-- {
		require.Equal(t, lista2.BorrarPrimero(), strconv.Itoa(i))
	}
	_estaVacia(t, lista2)

	//rune
	for i := 0; i < vol; i++ {
		lista3.InsertarPrimero(rune('a' + i%('z'-'a'+1)))
	}

	for i := vol - 1; i >= 0; i-- {
		require.Equal(t, lista3.BorrarPrimero(), rune('a'+i%('z'-'a'+1)))
	}
	_estaVacia(t, lista3)

	//float
	for i := 0; i < vol; i++ {
		lista4.InsertarPrimero(float64(i) + 0.534)
	}

	for i := vol - 1; i >= 0; i-- {
		require.Equal(t, lista4.BorrarPrimero(), float64(i)+0.534)
	}
	_estaVacia(t, lista4)

	//bool
	for i := 0; i < vol; i++ {
		lista5.InsertarPrimero(i%5 == 0)
	}
	for i := vol - 1; i >= 0; i-- {
		require.Equal(t, lista5.BorrarPrimero(), i%5 == 0)
	}
	_estaVacia(t, lista5)
}

func TestBordeEstaVacia(t *testing.T) {
	lista1 := TDALista.CrearListaEnlazada[int]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	lista3 := TDALista.CrearListaEnlazada[rune]()
	lista4 := TDALista.CrearListaEnlazada[float64]()
	lista5 := TDALista.CrearListaEnlazada[bool]()

	//int
	for i := 0; i <= testBordeVaciaLen; i++ {
		lista1.InsertarPrimero(i)
		require.False(t, lista1.EstaVacia())
	}
	for i := testBordeVaciaLen; i >= 0; i-- {
		require.False(t, lista1.EstaVacia())
		require.Equal(t, lista1.BorrarPrimero(), i)
	}
	_estaVacia(t, lista1)

	//string
	for i := 0; i < testBordeVaciaLen; i++ {
		lista2.InsertarPrimero(strconv.Itoa(i))
		require.False(t, lista2.EstaVacia())
	}
	for i := testBordeVaciaLen - 1; i >= 0; i-- {
		require.False(t, lista2.EstaVacia())
		require.Equal(t, lista2.BorrarPrimero(), strconv.Itoa(i))
	}
	_estaVacia(t, lista2)

	//rune
	for i := 0; i < testBordeVaciaLen; i++ {
		lista3.InsertarPrimero(rune('a' + i%('z'-'a'+1)))
		require.False(t, lista3.EstaVacia())
	}
	for i := testBordeVaciaLen - 1; i >= 0; i-- {
		require.False(t, lista3.EstaVacia())
		require.Equal(t, lista3.BorrarPrimero(), rune('a'+i%('z'-'a'+1)))
	}
	_estaVacia(t, lista3)

	//float
	for i := 0; i < testBordeVaciaLen; i++ {
		lista4.InsertarPrimero(float64(i) + 0.456)
		require.False(t, lista4.EstaVacia())
	}
	for i := testBordeVaciaLen - 1; i >= 0; i-- {
		require.False(t, lista4.EstaVacia())
		require.Equal(t, lista4.BorrarPrimero(), float64(i)+0.456)
	}

	_estaVacia(t, lista4)

	//bool
	for i := 0; i < testBordeVaciaLen; i++ {
		lista5.InsertarPrimero(i%7 == 0)
		require.False(t, lista5.EstaVacia())
	}
	for i := testBordeVaciaLen - 1; i >= 0; i-- {
		require.False(t, lista5.EstaVacia())
		require.Equal(t, lista5.BorrarPrimero(), i%7 == 0)
	}

	_estaVacia(t, lista5)
}

func TestBordeInvalidFunc(t *testing.T) {
	lista1 := TDALista.CrearListaEnlazada[int]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	lista3 := TDALista.CrearListaEnlazada[rune]()
	lista4 := TDALista.CrearListaEnlazada[float64]()
	lista5 := TDALista.CrearListaEnlazada[bool]()

	//int
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista1.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_ULTIMO_LISTA_VACIA, func() { lista1.VerUltimo() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista1.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	lista1.InsertarPrimero(1)
	lista1.BorrarPrimero()
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista1.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista1.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	//string
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista2.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_ULTIMO_LISTA_VACIA, func() { lista2.VerUltimo() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista2.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	lista2.InsertarPrimero("hola")
	lista2.BorrarPrimero()
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista2.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista2.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	//rune
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista3.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_ULTIMO_LISTA_VACIA, func() { lista3.VerUltimo() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista3.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	lista3.InsertarPrimero('b')
	lista3.BorrarPrimero()
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista3.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista3.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	//float
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista4.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_ULTIMO_LISTA_VACIA, func() { lista4.VerUltimo() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista4.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	lista4.InsertarPrimero(0.445)
	lista4.BorrarPrimero()
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista4.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista4.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	//bool
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista5.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_ULTIMO_LISTA_VACIA, func() { lista5.VerUltimo() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista5.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

	lista5.InsertarPrimero(false)
	lista5.BorrarPrimero()
	require.PanicsWithValue(t, TDALista.PANIC_MSG_VER_PRIMERO_LISTA_VACIA, func() { lista5.VerPrimero() }, "No hay un Panic al intentar ver el primero de una lista vacia")
	require.PanicsWithValue(t, TDALista.PANIC_MSG_BORRAR_LISTA_VACIA, func() { lista5.BorrarPrimero() }, "No hay un Panic al intentar BorrarPrimero una lista vacia")

}

func TestIteradorVacio(t *testing.T) {
	lista1 := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista1.EstaVacia())
	require.PanicsWithValue(t, PANIC_LISTA_VACIA, func() { lista1.BorrarPrimero() })
	require.PanicsWithValue(t, PANIC_LISTA_VACIA, func() { lista1.VerPrimero() })
	require.PanicsWithValue(t, PANIC_LISTA_VACIA, func() { lista1.VerUltimo() })
	require.NotNil(t, lista1)
	require.EqualValues(t, LONGITUD_VACIA, lista1.Largo())
	iterador1 := lista1.Iterador()
	require.NotNil(t, iterador1)
	require.PanicsWithValue(t, PANIC_ITERADOR_FINALIZO, func() { iterador1.VerActual() })
	require.PanicsWithValue(t, PANIC_ITERADOR_FINALIZO, func() { iterador1.Borrar() })
	require.PanicsWithValue(t, PANIC_ITERADOR_FINALIZO, func() { iterador1.Siguiente() })
	require.False(t, iterador1.HaySiguiente())
}

func TestIterador_InsertarPrimero(t *testing.T) {
	//Verifica que al insertar en la posicion inicial de un iterador recien creado va a ser igual a insertarPrimero
	lista1 := TDALista.CrearListaEnlazada[int]()
	lista2 := TDALista.CrearListaEnlazada[int]()
	_estaVacia(t, lista1)
	_estaVacia(t, lista2)

	iterador1 := lista1.Iterador()
	require.NotNil(t, iterador1)
	require.False(t, iterador1.HaySiguiente())

	lista2.InsertarPrimero(9)
	require.False(t, lista2.EstaVacia())

	iterador1.Insertar(5)
	require.True(t, iterador1.HaySiguiente())
	iterador1.Insertar(9)

	require.EqualValues(t, iterador1.VerActual(), lista2.VerPrimero())
}

func TestIteradorL_InsertarFinal(t *testing.T) {
	//Verifica que al insertar al final del iterador, es el mismo resultado que insertarUltimo
	lista1 := TDALista.CrearListaEnlazada[string]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	_estaVacia(t, lista1)
	_estaVacia(t, lista2)

	lista2.InsertarPrimero("iman")
	lista2.InsertarPrimero("foca")
	lista2.InsertarUltimo("mate")

	iterador1 := lista1.Iterador()
	require.NotNil(t, iterador1)
	require.False(t, iterador1.HaySiguiente())
	iterador1.Insertar("Que")
	require.True(t, iterador1.HaySiguiente())
	iterador1.Siguiente()
	iterador1.Insertar("tal?")
	require.True(t, iterador1.HaySiguiente())
	iterador1.Siguiente()
	iterador1.Insertar("mate")
	require.True(t, iterador1.HaySiguiente())
	require.EqualValues(t, iterador1.VerActual(), lista2.VerUltimo())
}

func TestIterador_InsertarMedio(t *testing.T) {
	//se corrobora que al insertar al medio se hace en la pos correcta
	lista1 := TDALista.CrearListaEnlazada[float64]()
	_estaVacia(t, lista1)
	expectedValues := []float64{1, 2.1, 4.65, 5.5, 6.00}

	for _, dato := range expectedValues {
		lista1.InsertarUltimo(dato)
	}
	posInsertar := lista1.Largo() / 2 //para este caso deberia ser 3 en la posicion que inserta
	posActual := 0
	iterador1 := lista1.Iterador()

	for iterador1.HaySiguiente() && posInsertar > posActual {
		require.EqualValues(t, iterador1.VerActual(), expectedValues[posActual])
		iterador1.Siguiente()
		posActual++
	}
	require.EqualValues(t, 4.65, iterador1.VerActual()) //estamos en la posicion que queremos insertar
	iterador1.Insertar(3.79)
	require.EqualValues(t, 3.79, iterador1.VerActual())
	require.True(t, iterador1.HaySiguiente())
	iterador1.Siguiente()
	require.EqualValues(t, 4.65, iterador1.VerActual())
	expectedValues = []float64{1, 2.1, 3.79, 4.65, 5.5, 6.00}
	posActual = 0
	iterador2 := lista1.Iterador()

	for iterador2.HaySiguiente() { //Para verificar que se inserto correctamente en el medio como queriamos
		require.EqualValues(t, iterador2.VerActual(), expectedValues[posActual])
		iterador2.Siguiente()
		posActual++
	}
	posActual = -1 //para que empiece del 0 cuando haga ++ en cada iteracion
	lista1.Iterar(func(i float64) bool {
		posActual++
		return i == expectedValues[posActual]
	})
	require.EqualValues(t, posActual, lista1.Largo()-1) //Si se cumple es porque coincide hasta el final de la lista esperada
}

func TestIterador_BorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	_estaVacia(t, lista)
	lista.InsertarPrimero("haces?")
	lista.InsertarPrimero("que")
	lista.InsertarPrimero("hola")

	iterador := lista.Iterador()
	require.NotNil(t, iterador)
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, lista.VerPrimero(), iterador.VerActual())
	require.EqualValues(t, "hola", iterador.VerActual())
	require.EqualValues(t, lista.VerPrimero(), iterador.Borrar())
}

func TestIterador_BorrarFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	_estaVacia(t, lista)
	lista.InsertarPrimero(false)
	lista.InsertarPrimero(false)
	lista.InsertarUltimo(true)

	iterador := lista.Iterador()
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	iterador.Siguiente()
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, lista.VerUltimo(), iterador.VerActual())
	require.EqualValues(t, true, iterador.Borrar())
}

func TestIterador_BorrarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	_estaVacia(t, lista)
	expectedValues := []int{18, 12, 20, 9, 0, 3, 14}

	for _, dato := range expectedValues {
		lista.InsertarUltimo(dato)
	}
	posBorrar := 2
	iterador := lista.Iterador()

	for i := 0; iterador.HaySiguiente() && i <= posBorrar; i++ {
		if posBorrar == i {
			require.EqualValues(t, iterador.VerActual(), expectedValues[i])
			require.EqualValues(t, 20, iterador.Borrar())
		}
		iterador.Siguiente()
	}
	expectedValues = []int{18, 12, 9, 0, 3, 14}
	posActual := -1 //para que empiece del 0 cuando haga ++ en cada iteracion

	lista.Iterar(func(i int) bool {
		posActual++
		return i == expectedValues[posActual]
	})
	require.EqualValues(t, posActual, lista.Largo()-1) //Si se cumple es porque coincide hasta el final de la lista esperada
}

func TestIterador_Volumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	_estaVacia(t, lista)
	iterador0 := lista.Iterador()
	for i := 0; i < testVolLen4; i++ {
		iterador0.Insertar(i)
		iterador0.Siguiente()
	} //Llega hasta la constante -1
	iterador1 := lista.Iterador()
	for i := 0; iterador1.HaySiguiente() && i <= testVolLen4; i++ {
		require.EqualValues(t, i, iterador1.VerActual())
		require.EqualValues(t, i, iterador1.Borrar())
	}
	require.False(t, iterador1.HaySiguiente())
	_estaVacia(t, lista)
}

func TestIteradorInterno_Volumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	_estaVacia(t, lista)
	for i := 0; i < testVolLen4; i++ {
		lista.InsertarUltimo(i)
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i, lista.VerUltimo())
	}
	valorActual := 0
	lista.Iterar(func(dato int) bool {
		require.EqualValues(t, valorActual, dato)
		valorActual++
		return valorActual == dato
	})

}

func TestIteradorInterno_Bordes(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	_estaVacia(t, lista)
	//caso de que no deberia entrar ya que la lista esta vacia
	lista.Iterar(func(dato int) bool {
		require.False(t, lista.EstaVacia())
		return lista.EstaVacia()
	})
	//caso un elemento por ver
	lista.InsertarPrimero(99)
	lista.Iterar(func(dato int) bool {
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, 99, lista.VerPrimero())
		return true
	})

}

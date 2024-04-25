package lista_test

import (
	"strconv"
	"strings"

	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testVolLen6 = 10
	testVolLen5 = 100
	testVolLen4 = 1000
	testVolLen3 = 10000
	testVolLen2 = 100000
	testVolLen1 = 1000000

	testBordeVaciaLen = 200
)

func TestListaaVacia(t *testing.T) {

	require.True(t, TDALista.CrearListaEnlazada[int]().EstaVacia())
	require.True(t, TDALista.CrearListaEnlazada[string]().EstaVacia())
	require.True(t, TDALista.CrearListaEnlazada[rune]().EstaVacia())
	require.True(t, TDALista.CrearListaEnlazada[float64]().EstaVacia())
	require.True(t, TDALista.CrearListaEnlazada[bool]().EstaVacia())
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
	require.True(t, lista.EstaVacia())
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
	require.True(t, lista.EstaVacia())
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

	//string
	for i := 0; i < vol; i++ {
		lista2.InsertarPrimero(strconv.Itoa(i))
	}

	for i := vol - 1; i >= vol; i-- {
		require.Equal(t, lista2.BorrarPrimero(), strconv.Itoa(i))
	}

	//rune
	for i := 0; i < vol; i++ {
		lista3.InsertarPrimero(rune('a' + i%('z'-'a'+1)))
	}

	for i := vol - 1; i >= 0; i-- {
		require.Equal(t, lista3.BorrarPrimero(), rune('a'+i%('z'-'a'+1)))
	}

	//float
	for i := 0; i < vol; i++ {
		lista4.InsertarPrimero(float64(i) + 0.534)
	}

	for i := vol - 1; i >= 0; i-- {
		require.Equal(t, lista4.BorrarPrimero(), float64(i)+0.534)
	}

	//bool
	for i := 0; i < vol; i++ {
		lista5.InsertarPrimero(i%5 == 0)
	}
	for i := vol - 1; i >= 0; i-- {
		require.Equal(t, lista5.BorrarPrimero(), i%5 == 0)
	}
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
	require.True(t, lista1.EstaVacia())

	//string
	for i := 0; i < testBordeVaciaLen; i++ {
		lista2.InsertarPrimero(strconv.Itoa(i))
		require.False(t, lista2.EstaVacia())
	}
	for i := testBordeVaciaLen - 1; i >= 0; i-- {
		require.False(t, lista2.EstaVacia())
		require.Equal(t, lista2.BorrarPrimero(), strconv.Itoa(i))
	}
	require.True(t, lista2.EstaVacia())

	//rune
	for i := 0; i < testBordeVaciaLen; i++ {
		lista3.InsertarPrimero(rune('a' + i%('z'-'a'+1)))
		require.False(t, lista3.EstaVacia())
	}
	for i := testBordeVaciaLen - 1; i >= 0; i-- {
		require.False(t, lista3.EstaVacia())
		require.Equal(t, lista3.BorrarPrimero(), rune('a'+i%('z'-'a'+1)))
	}
	require.True(t, lista3.EstaVacia())

	//float
	for i := 0; i < testBordeVaciaLen; i++ {
		lista4.InsertarPrimero(float64(i) + 0.456)
		require.False(t, lista4.EstaVacia())
	}
	for i := testBordeVaciaLen - 1; i >= 0; i-- {
		require.False(t, lista4.EstaVacia())
		require.Equal(t, lista4.BorrarPrimero(), float64(i)+0.456)
	}

	require.True(t, lista4.EstaVacia())

	//bool
	for i := 0; i < testBordeVaciaLen; i++ {
		lista5.InsertarPrimero(i%7 == 0)
		require.False(t, lista5.EstaVacia())
	}
	for i := testBordeVaciaLen - 1; i >= 0; i-- {
		require.False(t, lista5.EstaVacia())
		require.Equal(t, lista5.BorrarPrimero(), i%7 == 0)
	}

	require.True(t, lista5.EstaVacia())
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

package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Struct= podemos ver este operador como un tipo de dato que nosotros mismos creamos y definimos, a la cual le ponemos propiedades, que tenga métodos, etc.
type Calc struct{} //asi se define un struct, el nombre de Calc es determiando por nosotros mismos, es decir podemos colocarle cualquier nombre

//cuando le colocamos a la funcion el struct entre parentesis, se conoce como Receiver Function, esta le da la propiedad a cal de tener este metodo dentro de el.
//Quiere decir que cuando instanciemos a Calc vamos a poder llamar Calc.operator y vamos a llamar a esta funcion que estamos creando
func (Calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	//Si queremos agregarle mas de una opcion y por buenas practicas no debemos usar varios if(solo para decisiones binarias usamos if), entonces usamos switch, este operador nos permite definir varios casos como se ve acontinuacion
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, "No esta soportado")
		//Lo que nos va a permitir es manejar cuando todos los demas casos fallan, si ninguna condicion se cumple, el entra al defoult
		return 0
	}
}
func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

// lo que esta dentro del parentesis es lo que entra y lo que esta por fuera es lo que devuelve, como no sabemos el valor exact0 que devuelve, pero si sabemos que tipo de dato es, en este caso es un string, entonces se define
func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

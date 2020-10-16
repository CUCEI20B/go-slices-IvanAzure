package main
import "fmt"

func main()  {
	var cantidad int
	var entero int
	var s []int

	fmt.Print("Ingrese la cantidad de numeros a sumar: ")
	fmt.Scan(&cantidad)

	for i := 0; i < cantidad; i++{
		fmt.Scan(&entero)
		s = append(s,entero)
	}

	var res = 0

	for _, num := range s{
		res += num
	}
	fmt.Println(entero)
	fmt.Println(res)
}
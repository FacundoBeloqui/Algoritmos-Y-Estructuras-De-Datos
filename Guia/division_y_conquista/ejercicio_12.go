package division_y_conquista

/*
(★★★★) ♠♠ Implementar un algoritmo que, por división y conquista,
permita obtener la parte entera de la raíz cuadrada de un número n,
en tiempo O(logn). Por ejemplo, para n=10 debe devolver 3, y para n=25 debe devolver 5.
*/

func CalcularRaizEntera(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n%10 + CalcularRaizEntera(n/10)
}

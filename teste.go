package main

import (
	"fmt"
	"math"
)

func main() {
	peso := []float64{4199, 4179, 2923, 2815}
	forc := []float64{3453, 3120, 1962, 1452}
	fsum := 0.0
	psum := 0.0
	for _, n1 := range forc {
		fsum += n1
	}
	for _, n2 := range peso {
		psum += n2
	}
	efic := int(math.Round((fsum / psum) * 100))

	fmdi := 0.0
	fmndi := 0.0
	if forc[0] > forc[1] {
		fmdi = forc[0]
		fmndi = forc[1]
	} else {
		fmdi = forc[1]
		fmndi = forc[0]
	}
	fmtr := 0.0
	fmntr := 0.0
	if forc[2] > forc[3] {
		fmtr = forc[2]
		fmntr = forc[3]
	} else {
		fmtr = forc[3]
		fmntr = forc[2]
	}
	deseqd := int(math.Round(100 * ((fmdi - fmndi) / fmndi)))
	deseqt := int(math.Round(100 * ((fmtr - fmntr) / fmntr)))

	if efic > 55 {
		fmt.Printf("Eficiência aprovada com valor de %d%%.\n", efic)
	} else {
		fmt.Printf("Eficiência reprovada com valor de %d%%.\n", efic)
	}
	if deseqd <= 20 && deseqt <= 20 {
		fmt.Printf("Desequilibrio aprovado com valores dianteiro %d%% e traseiro %d%%.\n", deseqd, deseqt)
	} else {
		fmt.Printf("Desequilibrio reprovado com valores dianteiro %d%% e traseiro %d%%.\n", deseqd, deseqt)
	}
}

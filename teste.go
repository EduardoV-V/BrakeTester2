package main

//Importando pacotes
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

//FUnção principal (Necessária pro código funcionar)
func main() {
//Essa parte define o tipo de dado do arquivo json (data = float64) *obs: float e int tem que ser declarados como "float64" e "int64"
	type Data struct {
		Data []float64 `json:"data"`
	}
	//report de erro na leitura, caso aconteça
	file, err := os.Open("report-data.json")
	if err != nil {
		fmt.Println("Erro abrindo", err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Erro lendo", err)
		return
	}
	//Essa parte serve para ler o arquivo json, com a função "unmarshal" e salva ela na variavel jsonData
	var jsonData Data
	if err := json.Unmarshal(content, &jsonData); err != nil {
		fmt.Println("Erro parsing", err)
	}
	med := jsonData.Data
	peso := med[0:4]
	forc := med[4:8]
	//pbrake := med[9]
	fsum := 0.0
	psum := 0.0
	//aqui são os cálculos matemáticos, funciona igual ao python, praticamente
	for _, n := range forc {
		fsum += n
	}
	for _, n := range peso {
		psum += n
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
	
	//Essa parte lê os resultados e define se foram bons ou não.
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

func eqt() {

}

package interfaces

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

//ex 1
type Circulo struct {
	Raio float64
}

type Quadrado struct {
	Lado float64
}

type Forma interface {
	calculaArea() float64
}

func (c Circulo) calculaArea() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func (q Quadrado) calculaArea() float64 {
	return math.Pow(q.Lado, 2)
}

func CalculaAreaGeometrica(f Forma) float64 {
	return f.calculaArea()
}

//ex 2
type Cachorro struct {
	Tamanho   string
	Som       string
	Brinquedo string
	Comida    string
}

type Gato struct {
	Som       string
	Brinquedo string
	Comida    string
}

type Animal interface {
	apresenta() string
}

func (c Cachorro) apresenta() string {
	return fmt.Sprintf("Cachorro | Tamanho: %s | Som: %s | Brinquedo: %s | Comida: %s", c.Tamanho, c.Som, c.Brinquedo, c.Comida)
}

func (g Gato) apresenta() string {
	return fmt.Sprintf("Gato | Som: %s | Brinquedo: %s | Comida: %s", g.Som, g.Brinquedo, g.Comida)
}

func ApresentaAnimal(a Animal) string {
	return a.apresenta()
}

//ex extra 1

type DataJson struct {
	Data os.File
}

type Person struct {
	Name   string
	Age    int
	Gender string
}

type Print interface {
	printContent()
}

func (d DataJson) printContent() {
	content, _ := ioutil.ReadAll(&d.Data)
	fmt.Printf("%v\n", string(content))
}

func (p Person) printContent() {
	fmt.Printf("Nome: %s | Idade: %v | Sexo: %s\n", p.Name, p.Age, p.Gender)
}

func PrintInfo(p Print) {
	p.printContent()
}

//ex extra 2
func PrintType(i interface{}) (string, error) {
	var result string
	var err error

	switch i.(type) {
	case int32:
		result = fmt.Sprintf("Tipo int %d", i)
	case float32:
		result = fmt.Sprintf("Tipo float %.2f", i)
	default:
		err = errors.New("invalid Type")
	}

	return result, err
}

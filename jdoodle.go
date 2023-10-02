package main
import "fmt"
type Animal struct {
    nombre   string
    color string
}

func (u Animal) getAnimal() string {
    return fmt.Sprintf("nombre: %s, color: %s", u.nombre, u.color);
}

type Marino struct {
    *Animal
    aletas int
}

func (a Marino) getMarino() string {
    return fmt.Sprintf("%s - Aletas: %d", a.getAnimal(), a.aletas);
}

type Aereo struct {
    *Animal
    plumas int
    colorPico string
}

func (a Aereo) getAereo() string {
    return fmt.Sprintf("%s - plumas: %d, color Pico: %s", a.getAnimal(), a.plumas, a.colorPico);
}

type Pez struct {
    *Marino
    colorSecundario string
}

func (a Pez) getPez() string {
    return fmt.Sprintf("%s - color Secundario: %s", a.getMarino(), a.colorSecundario);
}

type Tiburon struct {
    *Marino
    dientes int
}

func (a Tiburon) getTiburon() string {
    return fmt.Sprintf("%s - dientes: %d", a.getMarino(), a.dientes);
}

type Aguila struct {
    *Aereo
}

func (a Aguila) getAguila() string {
    return fmt.Sprintf("%s ", a.getAereo());
}

type Paloma struct {
    *Aereo
}

func (a Paloma) getPaloma() string {
    return fmt.Sprintf("%s ", a.getAereo());
}

func main() {
    var Marino1 = Marino{&Animal{"AnimalMarino", "Azul"}, 1};
    var Aereo1 = Aereo{&Animal{"AnimalAereo", "Gris"}, 100, "Marron"};
    var pez1 = Pez{&Marino{&Animal{ "pescadito", "Azul Oscuro"}, 1},  "verde"};
    var tiburon1 = Tiburon{&Marino{&Animal{ "tiburonsin", "gris"}, 3},  50};
     var aguila1 = Aguila{&Aereo{&Animal{"aguilucho", "dorada"}, 116, "marron"}};
     var paloma1 = Paloma{&Aereo{&Animal{"palomin", "negro"}, 1, "violeta"}};

    fmt.Println(Marino1.getAnimal());
    fmt.Println(Marino1.getMarino());
    fmt.Println(Aereo1.getAereo());
    fmt.Println(pez1.getPez());
     fmt.Println(tiburon1.getTiburon());
     fmt.Println(aguila1.getAguila());
    fmt.Println(paloma1.getPaloma());
    
}
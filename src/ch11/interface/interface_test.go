package interface_test

import (
	"errors"
	"fmt"
	"testing"
)

type Programer interface {
	WriteHelloWorld() string
}

type GoProgramer struct{

}

func (g *GoProgramer) WriteHelloWorld() string{
	return "Hello World"
}


func TestClient(t *testing.T){
	var p Programer
	p = new(GoProgramer)
	t.Log(p.WriteHelloWorld())
}

type IntCov func(op int) int

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct{
	p *Pet
}

func (d *Dog) Speak() {
	d.p.Speak()
}

func (d *Dog) SpeakTo(host string){
	d.p.SpeakTo(host)
}

func TestDog(t *testing.T){
	var dog *Dog
	dog = new(Dog)
	dog.SpeakTo("ddd")
	errors.New("xxxxx")
}



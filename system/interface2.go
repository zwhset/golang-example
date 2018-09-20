package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}


type Stra struct {

}

func (s Stra) Read()  {
	fmt.Println("s.Read")
}

func (s Stra) Write()  {
	fmt.Println("s.Write")
}

type Intb struct {

}

func (i Intb) Read()  {
	fmt.Println("i.Read")
}

// func interface call
func Redis(r ReadWriter)  {
	r.Read()
	fmt.Println("Redis Read")
}

func Redis1(s Stra)  {

}

func File(i Intb)  {

}

func main() {
	//s := Stra{}
	i := Intb{}
	Redis(i)

	methods
}
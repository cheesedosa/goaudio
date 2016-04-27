package goaudio

import "fmt"

type Destination struct{
	
	node Node
}

func (d *Destination) Connect(c Component){
	
	fmt.Println(c)
	
}

func (d *Destination) process(*[]float32){
	fmt.Println("Dest")
}

func (d *Destination) getNode() *Node{
	return &d.node
}

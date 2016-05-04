package goaudio

import "fmt"

type Destination struct{
	
	node Node
}

func (d *Destination) Connect(c Component){
	
	fmt.Println(c)
	
}

func (d *Destination) process(){
	
	//fmt.Println(d.node.buffer)
	//Clear the buffer of the previous cycle values
	copy(d.node.buffer, emptyBuffer)
	
	//Process immediate input nodes. The nodes handle their own processes independently on their inputs
	for _, comp := range d.node.input {
		comp.process()
		compnode := comp.getNode()
		
		//accumulate the values per cycle
		for i := range d.node.buffer {
			d.node.buffer[i] = d.node.buffer[i] + compnode.buffer[i]
		}
	}
}

func (d *Destination) getNode() *Node{
	return &d.node
}

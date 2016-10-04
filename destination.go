package goaudio

import "fmt"

type Destination struct {
	node Node
}

func (d *Destination) Connect(c Component) {

	fmt.Println(c)

}

func (d *Destination) Disconnect() ([]Component, []Component) {
	return nil, nil
}

func (d *Destination) process() {

	d.node.tickCount = d.node.tickCount + 1
	//fmt.Println(d.node.buffer)
	//Clear the buffer of the previous cycle values
	copy(d.node.buffer, emptyBuffer)

	//Process immediate input nodes. The nodes handle their own processes independently on their inputs
	for _, comp := range d.node.input {
		compnode := comp.getNode()
		if d.node.tickCount > compnode.tickCount {
			comp.process()
		}

		//accumulate the values per cycle
		for i := range d.node.buffer {
			d.node.buffer[i] = d.node.buffer[i] + compnode.buffer[i]
		}
	}
}

func (d *Destination) getNode() *Node {
	return &d.node
}

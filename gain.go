package goaudio


//import "fmt"

type GainNode struct {
	Gain AudioParam
	node Node
}


func (g *GainNode) Connect(c Component){
	
	cnode:= c.getNode()
	g.node.output = append(g.node.output,c)
	(*cnode).input = append((*cnode).input,g)
}

/*Disconnect() disconnects the current component. Sets the output of all its input to a null slice 
 * and the input of all its output components to null slice.
 * Returns slices of all the input components and all the output components.
 * Note that this currently does NOT support rewiring. 
 * You'll manually have to iterate through the returned slice and use Connect() to rewire the way you like.
 */
 
func (g *GainNode) Disconnect() ([]Component,[]Component){
	var input,output []Component
	input = g.node.input
	output = g.node.output
	
	for _,c := range(input) {
		c.getNode().output = []Component{}
	}
	
	for _,c := range(output) {
		c.getNode().input = []Component{}
	}
	
	return input, output
	
}

func (g *GainNode) process() {
	
	g.node.tickCount = g.node.tickCount + 1
	//Process immediate input nodes. The nodes handle their own processes independently on their inputs
	for _,comp := range g.node.input {
		compnode := comp.getNode()
		if g.node.tickCount > compnode.tickCount {
			comp.process()
		}
		for i := range g.node.buffer {
			g.node.buffer[i] = g.node.buffer[i] + compnode.buffer[i]
		}
	}
	
	gain := fixGain(g.Gain.Value)
	
	for i := range g.node.buffer{
		g.node.buffer[i] = g.node.buffer[i] * gain
	}
	
	
}

func(g *GainNode) getNode() *Node{
	return &g.node
}

//gain helpers

func fixGain(gain float32) float32{
	var g float32
	
	switch {
		case gain < 0:
			g = 0
		case gain > 1:
			g = 1
		default :
			g = gain
	}
	
	return g
}



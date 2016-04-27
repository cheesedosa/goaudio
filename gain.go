package goaudio


//import "fmt"

type GainNode struct {
	Gain AudioParam
	node Node
}


func (g *GainNode) Connect(c Component){
	
	cnode:= c.getNode()
	g.node.output = c
	(*cnode).input = g
}

func (g *GainNode) process(data *[]float32) {
	gain := fixGain(g.Gain.Value)
	for i := range *data{
		(*data)[i] = (*data)[i] * gain
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



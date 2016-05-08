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



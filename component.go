package goaudio


//Interface Component to allow generic processing of each block

type Component interface {
	
	Connect(Component)
	process(*[]float32)
	getNode() *Node
}

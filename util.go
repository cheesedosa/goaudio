package goaudio

//import "fmt"
import "github.com/gordonklaus/portaudio"

//basic types that build up components

type Node struct {
	input Component
	output Component
}

type AudioParam struct {
	Value float32
	prevValue float32
}

//audioparam helpers 

func(ap * AudioParam) valueChanged() bool {
	
	if ap.Value != ap.prevValue {
		ap.prevValue = ap.Value
		return true
	}
	
	return false
}

//audioContext helpers; the main guys that do all the work

func (a *AudioContext) playGraph() {
	
	portaudio.Initialize()
	defer portaudio.Terminate()
	data := make([]float32, 1024)
	stream, err := portaudio.OpenDefaultStream(0,1,a.sampleRate, len(data), &data)
	checkErr(err)
	defer stream.Close()
	
	checkErr(stream.Start())
	defer stream.Stop()
	
	
	
	
	for {
		getAudioData(*(a.Dest), &data)
		checkErr(stream.Write())
	}
}


func getAudioData(d Destination, data *[]float32){
	
	c := d.node.input
	processStuff(c, data)
}


func processStuff(c Component, data *[]float32){
	
	cnode := c.getNode()
	if cnode.input == nil {
		c.process(data)
		return
	}
	nc := cnode.input
	processStuff(nc, data)
	c.process(data)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

package goaudio

//import "fmt"
import "github.com/gordonklaus/portaudio"

//basic types that build up components

type Node struct {
	input []Component
	output []Component
	buffer []float32
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
	
	//data := make([]float32, 1024)
	stream, err := portaudio.OpenDefaultStream(0,1,a.sampleRate, 1024, a.Dest.node.buffer)
	checkErr(err)
	defer stream.Close()
	
	checkErr(stream.Start())
	defer stream.Stop()
	
	for {
		
		a.Dest.process()
		//fmt.Println("Writing to stream..")
		checkErr(stream.Write())
	}
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

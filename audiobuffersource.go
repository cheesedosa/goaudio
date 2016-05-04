package goaudio

import "fmt"
//import "time"

type AudioBufferSource struct {
	Buffer *[]float32
	bufferlength int
	Loop bool
	node Node
	on bool
	end bool
	currentidx int
	//timestamp time.Time
	//startStamp float64
	//stopStamp float64
}

func (abs *AudioBufferSource) Start(){
	
	//todo
	
	//abs.startStamp = x
	
	abs.on = true
}

func (abs *AudioBufferSource) Stop(x float64){
	
	//todo
	//abs.stopStamp = x
	abs.on = false
}

func (abs *AudioBufferSource) Connect(c Component){
	
	cnode := c.getNode()
	abs.node.output = append(abs.node.output,c)
	(*cnode).input = append((*cnode).input,abs)
}

func (abs *AudioBufferSource) getNode() *Node{
	
	return &abs.node
}

func (abs *AudioBufferSource) isOn() bool {
	
	//todo: multiple start stops
	
	//fmt.Println(time.Since(o.timestamp).Seconds(), len(o.startStamp), len(o.stopStamp))
	//if len(o.startStamp) > 0 && time.Since(o.timestamp).Seconds() >= o.startStamp[0] {
		//o.startStamp = o.startStamp[1:]
		//o.on = true
		//return true
	//} else if len(o.stopStamp) > 0 &&time.Since(o.timestamp).Seconds() >= o.stopStamp[0] {
		//o.stopStamp = o.startStamp[1:]
			//o.on = false
			//return false
		//}
	//return false
	
	
	//fmt.Println(time.Since(abs.timestamp).Seconds())
	//if time.Since(abs.timestamp).Seconds() >= abs.startStamp {
		//abs.on = true
		//return true
	//} else if time.Since(abs.timestamp).Seconds() >= abs.stopStamp {
			//abs.on = false
			//return false
		//}
	//return false
	
	return abs.on
}

func (abs *AudioBufferSource) process() {
	
	//Use a default empty buffer to fill, instead of the slice passed to the function; this avoids DC values
	abs.node.buffer = emptyBuffer
	
	abs.bufferlength = len(*abs.Buffer)
	frames:= len(abs.node.buffer)
	
	//fmt.Println(abs.bufferlength, abs.currentidx*frames, abs.on, abs.end)
	
	//fmt.Println((*data)[:5])
	//check start/stop
	if !abs.isOn(){
		fmt.Println("off")
		
		//fmt.Println(data)
		return
	}
	//check if playback finished
	if abs.end{
		//if loop is set; play again
		if abs.Loop {
		fmt.Println("looping")
		abs.currentidx = 0
		abs.end = false
		}
		return
	}
	//check if final frame; otherwise just play
	
	if abs.currentidx*frames+frames > abs.bufferlength {
	//For the ending bit of the buffer, concatenate the slice with 0s to make the length equal to Node.buffer
		abs.node.buffer = append((*abs.Buffer)[abs.currentidx*frames:],make([]float32, len(abs.node.buffer)-len((*abs.Buffer)[abs.currentidx*frames:]))...)
		abs.end = true
	}else {
	//fmt.Println("Playing", abs.currentidx)
	abs.node.buffer = (*abs.Buffer)[abs.currentidx*frames:abs.currentidx*frames+frames]
	}
	abs.currentidx += 1	
}

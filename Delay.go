/*Delay Node implementation
 * AudioContext.CreateDelay() returns a pointer to this type
 * The CreateDelay function requires a delay time provided in seconds.
 * This delay time is public and can be accessed or changed: Delay.DelayTime.Value
 */

package goaudio

//import "fmt"

type Delay struct {
	
	idx int
	delayBuffer []float32
	node Node
	DelayTime AudioParam
}

func (dly *Delay) Connect(c Component){
	
	cnode := c.getNode()
	dly.node.output = append(dly.node.output,c)
	(*cnode).input = append((*cnode).input,dly)
}

func (dly *Delay) getNode() *Node{
	
	return &dly.node
}

func (dly *Delay) process() {
	

	
	//Prepare an in buffer to load input values from the process calls of the immediate inputs of this Delay node.
	
	var in [1024]float32
	
	for _,comp := range dly.node.input {
		comp.process()
		compnode := comp.getNode()
		for i := range in {
			in[i] = in[i] + compnode.buffer[i]
		}
	}
	
	//Check if the delay value was changed. If so, make an appropriate buffer to hold inputs which would be delayed
	// and written in a later cycle.
	if dly.DelayTime.valueChanged(){
		
		dly.delayBuffer = make([]float32, int(44100*dly.DelayTime.Value))
	}
	//process data and fill it up
	for i:= range dly.node.buffer {
		
		//delayBuffer holds the data copied to in at intervals specified by the delaytime;
		//the delaybuffer goes from 0 to samplingrate*delaytime -1 ;
		// so, every pass of the delaybuffer accounts for 1 unit of delaytime specified
		dly.node.buffer[i] = dly.delayBuffer[dly.idx]
		dly.delayBuffer[dly.idx] = in[i]
		dly.idx = (dly.idx + 1) % len(dly.delayBuffer)
	}
}

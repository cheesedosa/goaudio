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
	dly.node.output = c
	(*cnode).input = dly
}

func (dly *Delay) getNode() *Node{
	
	return &dly.node
}

func (dly *Delay) process(data *[]float32) {
	
	if dly.DelayTime.valueChanged(){
		
		dly.delayBuffer = make([]float32, int(44100*dly.DelayTime.Value))
	}
	
	//copy the data we receive into an in slice
	in := make([]float32, len(*data))
	copy(in, *data)
	
	//process data and fill it up
	for i:= range(*data) {
		
		//delayBuffer holds the data copied to in at intervals specified by the delaytime;
		//the delaybuffer goes from 0 to samplingrate*delaytime -1 ;
		// so, every pass of the delaybuffer accounts for 1 unit of delaytime specified
		(*data)[i] = dly.delayBuffer[dly.idx]
		dly.delayBuffer[dly.idx] = in[i]
		dly.idx = (dly.idx + 1) % len(dly.delayBuffer)
	}
}

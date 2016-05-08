/* Audio Library based on the webaudio API written in Golang.
 * This work in progress and is highly experimental.
 * Yet to implement multiple chains on Destination. Currently, it supports only one chain.
 * The Chain may have any number of the currently implemented nodes but only the last chain connected to the 
 * destination is processed as of now.
 */

package goaudio

import "fmt"
//import "time"

//some globals; maybe these need refactoring
var emptyBuffer []float32 = make([]float32, 1024)

type AudioContext struct{
	sampleRate float64
	Dest *Destination
}

func (a *AudioContext) GetSampleRate() float64{
	
	return a.sampleRate
}

func NewAudioContext(sr float64) *AudioContext{

	return &AudioContext{sampleRate: sr, Dest: &Destination{node: Node{buffer: make([]float32, 1024)}}}
}

//func (a *AudioContext) Close(){
	
	//fmt.Println("Stopped.")
//}

func (a *AudioContext) Play(){
	
	// plays the "graph", a result of all the connects called on individual components as a goroutine to allow further changes to the graph
	go a.playGraph()
	
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("This starts a goroutine that handles the playing.")
	fmt.Println("To prevent the program from exiting you should have implemented a block inside \nyour main function.")
	fmt.Println("Even a simple input scan would do. Check /examples for ideas.")
	fmt.Println("Starting...")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

}

func (a *AudioContext) CurrentTime(){
	
	fmt.Println("Current time")
}

//Various methods to create different crap

func (a *AudioContext) CreateBuffer(f string) ([]float32, float32){
	
	//Takes in a file name and returns the parsed wav data as a slice of []float32. Also, returns the sampling rate of the wav file
	
	wavdata := simpleWavReader(f)
	
	buffer := normalizeWavData(wavdata.data)
	
	return buffer , float32(wavdata.SampleRate)
}

func (a *AudioContext) CreateBufferSource() *AudioBufferSource{
	
	audiobuffersource := AudioBufferSource{node: Node{buffer: make([]float32, 1024)}}
	return &audiobuffersource
	
}

func (a *AudioContext) StereoPanner(){
	
	fmt.Println("Stereo Panner")
}

func (a *AudioContext) CreateBiquadFilter(){
	
	fmt.Println("Biquad Filer")
}

func (a *AudioContext) CreateChannelMerger(){
	
	fmt.Println("Merge channels")
}

func (a *AudioContext) CreateChannelSplitter(){
	
	fmt.Println("Current time")
}

func (a *AudioContext) CreateConvolver(){
	
	fmt.Println("Convolver")
}

func (a *AudioContext) CreateDelay(dtime float64) *Delay{
	
	dly := Delay{node:Node{buffer: make([]float32, 1024)},delayBuffer: make([]float32, int(a.sampleRate*dtime)), DelayTime: AudioParam{float32(dtime), float32(0.0)}}
	
	return &dly
}

func (a *AudioContext) CreateDynamicsCompressor(){
	
	fmt.Println("Dynamics Compressor")
}

func (a *AudioContext) CreateGain() *GainNode{
	
	gain := GainNode{Gain: AudioParam{1.0, 1.0}, node: Node{buffer: make([]float32, 1024)}}
	
	return &gain
}

func (a *AudioContext) CreateIIRFilter(){
	
	fmt.Println("IIR Filter")
}

func (a *AudioContext) CreateOscillator(freq float32, osctype string, det float32) *Oscillator{
	
	osc := Oscillator{Frequency: AudioParam{freq, 220.0}, Detune: AudioParam{det, 0.0}, OscType: osctype, wave: &Wave{float64(freq/44100), 0}, on:false, node: Node{buffer: make([]float32, 1024)}}
	
	fmt.Println(osc.wave.phase)
	
	return &osc
}

func (a *AudioContext) CreatePeriodicWave(){
	
	fmt.Println("Periodic wave")
}

func (a *AudioContext) CreateWaveShapper() {
	
	fmt.Println("Wave Shapper")
}

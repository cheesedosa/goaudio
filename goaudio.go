package goaudio

import "fmt"
import "time"

type AudioContext struct{
	sampleRate float64
	Dest *Destination

}

func (a *AudioContext) GetSampleRate() float64{
	
	return a.sampleRate
}

func NewAudioContext(sr float64) *AudioContext{

	return &AudioContext{sampleRate: sr, Dest: &Destination{node: Node{}}}
}

func (a *AudioContext) Close(){
	
	fmt.Println("Stopped.")
}

func (a *AudioContext) Play(){
	
	// plays the "graph", a result of all the connects called on individual components as a goroutine to allow further changes to the graph
	go a.playGraph()
	
	fmt.Println("Starting...")
}

func (a *AudioContext) CurrentTime(){
	
	fmt.Println("Current time")
}

//Various methods to create different crap

func (a *AudioContext) CreateBuffer(){
	
	fmt.Println("Current Buffer")
}

func (a *AudioContext) CreateBufferSource() *AudioBufferSource{
	
	audiobuffersource := AudioBufferSource{node: Node{},timestamp: time.Now()}
	fmt.Println("Current Buffer to use as source")
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

func (a *AudioContext) CreateDelay(){
	
	fmt.Println("Delay")
}

func (a *AudioContext) CreateDynamicsCompressor(){
	
	fmt.Println("Dynamics Compressor")
}

func (a *AudioContext) CreateGain() *GainNode{
	
	gain := GainNode{Gain: AudioParam{1.0, 1.0}, node: Node{}}
	
	return &gain
}

func (a *AudioContext) CreateIIRFilter(){
	
	fmt.Println("IIR Filter")
}

func (a *AudioContext) CreateOscillator(freq float32, osctype string, det float32) *Oscillator{
	
	osc := Oscillator{Frequency: AudioParam{freq, 220.0}, Detune: AudioParam{det, 0.0}, OscType: osctype, wave: &Wave{float64(freq/44100), 0}, on:false, node: Node{}, timestamp: time.Now()}
	
	fmt.Println(osc.wave.phase)
	
	return &osc
}

func (a *AudioContext) CreatePeriodicWave(){
	
	fmt.Println("Periodic wave")
}

func (a *AudioContext) CreateWaveShapper() {
	
	fmt.Println("Wave Shapper")
}

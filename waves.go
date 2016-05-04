package goaudio

import "math"

//wave type for storing wave sample data
type Wave struct{
	step, phase float64
}

//oscillator wave methods

func (o *Oscillator) getSine() {
	phase := o.wave.phase
	for i := range o.node.buffer{
		o.node.buffer[i] = float32(math.Sin(2 * math.Pi * phase))
		_, phase = math.Modf(phase + o.wave.step)
	}
	o.wave.phase = phase
}



func (o *Oscillator) getTri() {
	phase := o.wave.phase
	for i := range o.node.buffer{
		o.node.buffer[i] = float32(math.Abs(2 * (phase - math.Floor(phase)) - 1))
		_, phase = math.Modf(phase + o.wave.step)
	}
	o.wave.phase = phase
}


func (o *Oscillator) getSqr() {
	phase := o.wave.phase
	for i := range o.node.buffer{
		sin := float32(math.Sin(2 * math.Pi * phase))
		o.node.buffer[i] = cmp(sin) - (1-cmp(sin))
		_, phase = math.Modf(phase + o.wave.step)
	}
	o.wave.phase = phase
}

func (o *Oscillator) getSaw() {
	phase := o.wave.phase
	for i := range o.node.buffer{
		o.node.buffer[i] = float32(2 * (phase - math.Floor(phase)) - 1)
		_, phase = math.Modf(phase + o.wave.step)
	}
	o.wave.phase = phase
}

//wave helpers 

func cmp(v float32) float32{
	
	if v > 0 {
		return 1
	} else {
		return 0
	}
}

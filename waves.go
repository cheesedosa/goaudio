package goaudio

import "math"

//wave type for storing wave sample data
type Wave struct{
	step, phase float64
}

//oscillator wave methods

func (o *Oscillator) getSine(data *[]float32) {
	phase := o.wave.phase
	for i := range *data{
		(*data)[i] = float32(math.Sin(2 * math.Pi * phase))
		_, phase = math.Modf(phase + o.wave.step)
	}
	o.wave.phase = phase
}



func (o *Oscillator) getTri(data *[]float32) {
	phase := o.wave.phase
	for i := range *data{
		(*data)[i] = float32(math.Abs(2 * (phase - math.Floor(phase)) - 1))
		_, phase = math.Modf(phase + o.wave.step)
	}
	o.wave.phase = phase
}


func (o *Oscillator) getSqr(data *[]float32) {
	phase := o.wave.phase
	for i := range *data{
		sin := float32(math.Sin(2 * math.Pi * phase))
		(*data)[i] = cmp(sin) - (1-cmp(sin))
		_, phase = math.Modf(phase + o.wave.step)
	}
	o.wave.phase = phase
}

func (o *Oscillator) getSaw(data *[]float32) {
	phase := o.wave.phase
	for i := range *data{
		(*data)[i] = float32(2 * (phase - math.Floor(phase)) - 1)
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

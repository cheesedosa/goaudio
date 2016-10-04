/*Still highly experimental
 * Press Any Key to exit, hahaha
 * Check README.md
 *
 */

package main

import "github.com/mrnikho/goaudio"
import "fmt"
import "time"

func main() {

	//create audio context
	context := goaudio.NewAudioContext(44100)

	fmt.Println(context.GetSampleRate())

	var freq, detune float32

	freq = 110
	osctype := "SINE"
	detune = 0

	//Create oscillator
	osc := context.CreateOscillator(freq, osctype, detune)

	//Gain
	gain := context.CreateGain()
	gain.Gain.Value = 0.2

	//Destination
	dest := context.Dest

	//Connections similar to webaudio api
	osc.Connect(gain)
	gain.Connect(dest)
	//This initialises the portaudio stream and calls process recursively starting from the destination
	context.Play()

	osc.Start()

	time.Sleep(time.Second * 2) // 2 sec
	gain.Gain.Value = 0.4
	time.Sleep(time.Second * 2) // 4 sec
	gain.Gain.Value = 0.6

	osc.OscType = "SAW"

	time.Sleep(time.Second * 2) // 6 sec
	gain.Gain.Value = 0.8
	time.Sleep(time.Second * 2) // 8 sec
	osc.Frequency.Value = 440
	gain.Gain.Value = 0.4
	time.Sleep(time.Second * 2) //10 sec
	gain.Gain.Value = 0         // shhh

	//blocking the main from returning and exiting the program; press any key to exit
	var input string
	fmt.Scanln(&input)
	fmt.Print("Done.")

}

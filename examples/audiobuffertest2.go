/*Still highly experimental
 * Press any key to exit, hahaha
 * Check README.md
 * 
 */

//


// same example as audiobuffertest.go only this one uses AudioContext.CreateBuffer to create a buffer from a wav file
package main

import "github.com/mrnikho/goaudio"
import "fmt"
import "math/rand"
func main(){
	
	context:= goaudio.NewAudioContext(44100)
	
	fmt.Println("The Sampling rate is: ", context.GetSampleRate())
	
	audiosource := context.CreateBufferSource()
	
	//Create buffer returns the audio data as a []float32 slice and the sampling rate
	data, sr := context.CreateBuffer("piano.wav")
	fmt.Println("The Sampling rate is: ", sr, " Hz.")
	audiosource.Buffer = &data
	
	audiosource.Connect(context.Dest)
	audiosource.Loop = true
	audiosource.Start()
	context.Play()	
    
   	//blocking the main from returning and exiting the program; press any key to exit
    var input string
    fmt.Scanln(&input)
    fmt.Print("Done.")
	
}

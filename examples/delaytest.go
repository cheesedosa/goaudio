/*Still highly experimental
 * Press any key to exit, hahaha
 * Check README.md
 * 
 */

//


// Loads up a wav file "piano.wav" and plays back with the delay specified in CreateDelay() call
package main

import "github.com/mrnikho/goaudio"
import "fmt"
func main(){
	
	context:= goaudio.NewAudioContext(44100)
	
	fmt.Println("The Sampling rate is: ", context.GetSampleRate())
	
	audiosource := context.CreateBufferSource()
	
	data, _ := context.CreateBuffer("piano.wav")
	
	audiosource.Buffer = &data
	
	delay := context.CreateDelay(1) // 1 sec delay
		
	//Connect our source to delay
	audiosource.Connect(delay)
	
	//Connect our delay line to our destination
	delay.Connect(context.Dest)
	
	//Loop the source
	audiosource.Loop = true
	
	//Connect the source to the destination too. So, we have source + delayed source
	audiosource.Connect(context.Dest)
	
	//Start this guy
	audiosource.Start()
	
	//Start our context; this starts portaudio and the processing calls to each created node
	context.Play()	
    
   	//blocking the main from returning and exiting the program; press any key to exit
    var input string
    fmt.Scanln(&input)
    fmt.Print("Done.")
	
}

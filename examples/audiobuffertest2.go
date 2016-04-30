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
	
	fmt.Println(context.GetSampleRate())
	
	audiosource := context.CreateBufferSource()
	
	data, sr := context.CreateBuffer("piano.wav")
	
	audiosource.Buffer = &data
	
	//fmt.Println(audiosource.Buffer)
	audiosource.Connect(context.Dest)
	audiosource.Loop = true
	audiosource.Start()
	fmt.Println(audiosource)
	context.Play()	
    
   	//blocking the main from returning and exiting the program; press any key to exit
    var input string
    fmt.Scanln(&input)
    fmt.Print("Done.")
	
}

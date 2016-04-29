/*Still highly experimental
 * Press any key to exit, hahaha
 * Check README.md
 * 
 */

//

package main

import "github.com/mrnikho/goaudio"
import "fmt"
import "math/rand"
func main(){
	
	context:= goaudio.NewAudioContext(44100)
	
	fmt.Println(context.GetSampleRate())
	
	audiosource := context.CreateBufferSource()
	
	//make a buffer to be held by the audio node; sampling rate x 3 = 3 seconds of noise
	data := make([]float32,44100*3)
	
	//make some noise
	for i:= range(data) {
		data[i] = rand.Float32()*2-1
	}
	
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

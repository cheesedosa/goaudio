package goaudio

/* Contains the code for AudioContext.CreateBuffer()
 * Reuses a simple wav parser I wrote from here: https://github.com/mrnikho/yingo/blob/master/wavstuff.go
 * The CreateBuffer function takes a wav file as input and returns a slice of float32s; []float32 of the parsed wav data.
 * The values are normalised to be from -1 to 1. Use a gain node if the volume is insufficient.
 */
 
import (
	"bytes"
	"encoding/binary"
	//"fmt"
	//"math"
	"io"
	"io/ioutil"
)

// WavFormat : data structure
type WavFormat struct {
	ChunkID       uint32
	ChunkSize     uint32
	Format        uint32
	Subchunk1ID   uint32
	Subchunk1Size uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	Subchunk2ID   uint32
	Subchunk2Size uint32
	data          []int16
}

// decode : decode wav data
func (w *WavFormat) decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, &w.ChunkID); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.ChunkSize); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &w.Format); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &w.Subchunk1ID); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.Subchunk1Size); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.AudioFormat); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.NumChannels); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.SampleRate); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.ByteRate); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.BlockAlign); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.BitsPerSample); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &w.Subchunk2ID); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &w.Subchunk2Size); err != nil {
		return err
	}

	d := make([]byte, w.Subchunk2Size)
	
	// get data bytes
	if _, err := io.ReadFull(r, d); err != nil {
		return err
	}
	
	//parse them into int16 
	data := make([]int16, w.Subchunk2Size/2)
	rr:= bytes.NewReader(d)
	if err := binary.Read(rr, binary.LittleEndian, &data); err != nil {
		return err
	}
	
	//fmt.Println(data)
	
	w.data = data
	return nil
}

// simpleWavFile takes the file name and spits out the WavFormat type. The data is held in WavFormat.data
func simpleWavReader(f string) WavFormat{
	
	data, err := ioutil.ReadFile(f)
	
	if err !=nil{
		panic(err)
	}
	
	w := WavFormat{}
	r := bytes.NewReader(data)
	err = w.decode(r)
	
	if err != nil {
		panic(err)
	}
	
	return w
	
}

//takes the parsed data, a slice of uint16s and converts them to a slice of float32s normalising the value to from -1 to 1.
func normalizeWavData(data []int16) []float32{
	
	var max int16
	if len(data) > 0 {
		max = data[0]
	}
	if len(data) > 1 {
		for i:= 1; i< len(data); i++ {
			if data[i] < max{
				max = data[i]
			}
		}
	}
	buf := make([]float32, len(data))
	
	for i := range(data) {
		buf[i] = float32(data[i])/float32(max)
	}
	
	return buf
}
	

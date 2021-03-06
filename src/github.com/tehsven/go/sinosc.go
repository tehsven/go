package sinosc

import (
	"log"
	"math"
	"time"
	"os"
	"strconv"
	"code.google.com/p/portaudio-go/portaudio"
)

const tao = 2 * math.Pi
const sampleRate = 44100.0
const framesPerBuffer = 256.0
const outputChannels = 2
const outBufSize = outputChannels * framesPerBuffer

var freq = 0.0
var duration = 0.0
var interval = 0.0

func sinwave() {
	freq, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal("invalid input arg: ", err)
	}
	duration, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal("invalid input arg: ", err)
	}
	interval = freq * tao / sampleRate

	portaudio.Initialize()
	stream, err := portaudio.OpenDefaultStream(
		0,
		outputChannels,
		sampleRate,
		256,
		audioCallback)
	if err != nil {
		log.Fatal("openDefault: ", err)
	}

	err = stream.Start()
	if err != nil {
		log.Fatal("Start: ", err)
	}

	time.Sleep(time.Duration(duration) * time.Second)

	err = stream.Stop()
	if err != nil {
		log.Fatal("Stop: ", err)
	}
}

var x = 0.0
func audioCallback(
	in []float32,
	out []float32,
	timeInfo portaudio.StreamCallbackTimeInfo,
	flags portaudio.StreamCallbackFlags) {
	
	for i := 0; i < outBufSize; i+=outputChannels {
		y := float32(math.Sin(x))
		x += interval
		for j := 0; j < outputChannels; j++ {
			out[i+j] = y
		}
	}
}

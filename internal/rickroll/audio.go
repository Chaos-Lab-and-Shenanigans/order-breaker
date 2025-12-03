package rickroll

import (
	"bytes"
	"io"
	"time"

	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

var (
	stopCh = make(chan string, 1)
	player *oto.Player
	otoCtx *oto.Context
)

func playAudio(errCh chan error) {
	//check if player is already initializated
	if player != nil {
		player.Play()
		go handlePause()
		errCh <- nil
		return
	}

	audioReader := bytes.NewReader(*config.Cfg.RickyAudioBytes)
	decodedMP3, err := mp3.NewDecoder(audioReader)
	if err != nil {
		errCh <- err
		return
	}

	//Initialize audio hardware configuration
	op := oto.NewContextOptions{}
	op.ChannelCount = 2
	op.SampleRate = 44100
	op.Format = oto.FormatSignedInt16LE

	//Initialize context
	var readyChan chan struct{}
	if otoCtx == nil { //Skip if context exists
		otoCtx, readyChan, err = oto.NewContext(&op)
		if err != nil {
			errCh <- err
			return
		}
	}

	errCh <- nil

	//Wait for hardware initialization to complete
	<-readyChan

	player = otoCtx.NewPlayer(decodedMP3)
	player.SetVolume(1)
	player.Play()
	go handlePause()
}

func stopAudio() {
	if player == nil || !player.IsPlaying() {
		return
	}

	//Send stop signal if channel is empty, skip otherwise
	select {
	case stopCh <- "stop":
		return
	default:
		return
	}
}

func handlePause() {
playAgain:
	for player.IsPlaying() {
		select {
		case msg := <-stopCh:
			if msg == "stop" {
				player.Pause()
				return
			}

		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
	//Repeat song if not stopped
	player.Seek(0, io.SeekStart)
	player.Play()
	goto playAgain
}

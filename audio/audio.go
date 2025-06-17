package audio

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const soundAssetFilePrefix = "assets/audio"

type AudioPlayer struct {
	deathSound   rl.Sound
	lazerSound   rl.Sound
	splitSound   rl.Sound
	bloopLoSound rl.Sound
	bloopHiSound rl.Sound
}

func (ap *AudioPlayer) Init() {
	ap.deathSound = rl.LoadSound(fmt.Sprintf("%s/death.wav", soundAssetFilePrefix))
	ap.lazerSound = rl.LoadSound(fmt.Sprintf("%s/lazer.wav", soundAssetFilePrefix))
	ap.splitSound = rl.LoadSound(fmt.Sprintf("%s/asteroidSplit.wav", soundAssetFilePrefix))
	ap.bloopLoSound = rl.LoadSound(fmt.Sprintf("%s/bloopLo.wav", soundAssetFilePrefix))
	ap.bloopHiSound = rl.LoadSound(fmt.Sprintf("%s/bloopHi.wav", soundAssetFilePrefix))
}

func (ap AudioPlayer) PlayLazer() {
	if rl.IsAudioDeviceReady() {
		rl.PlaySound(ap.lazerSound)
	}
}

func (ap AudioPlayer) PlayDeath() {
	if rl.IsAudioDeviceReady() {
		rl.PlaySound(ap.deathSound)
	}
}

func (ap AudioPlayer) PlayAsteroidSplit() {
	if rl.IsAudioDeviceReady() {
		rl.PlaySound(ap.splitSound)
	}
}

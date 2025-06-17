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
	musicIsLow   bool
	bloopCanPlay bool
}

func (ap *AudioPlayer) Init() {
	ap.deathSound = rl.LoadSound(fmt.Sprintf("%s/death.wav", soundAssetFilePrefix))
	ap.lazerSound = rl.LoadSound(fmt.Sprintf("%s/lazer.wav", soundAssetFilePrefix))
	ap.splitSound = rl.LoadSound(fmt.Sprintf("%s/asteroidSplit.wav", soundAssetFilePrefix))
	ap.bloopLoSound = rl.LoadSound(fmt.Sprintf("%s/bloopLo.wav", soundAssetFilePrefix))
	ap.bloopHiSound = rl.LoadSound(fmt.Sprintf("%s/bloopHi.wav", soundAssetFilePrefix))
	ap.musicIsLow = true
	ap.bloopCanPlay = true
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

func (ap *AudioPlayer) HandleMusic() {
	// Per original asteroids game, "music" is two notes

	var sound rl.Sound
	now := int64(rl.GetTime())

	if now%2 == 0 {
		if !ap.bloopCanPlay {
			return
		}

		if ap.musicIsLow {
			sound = ap.bloopLoSound
		} else {
			sound = ap.bloopHiSound
		}

		rl.PlaySound(sound)
		ap.bloopCanPlay = false
	} else {
		ap.bloopCanPlay = true
	}

}

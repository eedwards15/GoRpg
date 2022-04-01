package systems

import (
	"GoRpg/assets"
	"GoRpg/src/models"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

var (
	MUSICSYSTEM *MusicSystem
)

type MusicSystem struct {
	currentSong *mp3.Stream
	sampleRate  int
	player      *audio.Player
}

func InitMusicSystem(music models.BackgroundMusic) {
	MUSICSYSTEM = &MusicSystem{}
	MUSICSYSTEM.sampleRate = music.SampleRate
	f, _ := assets.Assets.Open(music.Path)
	MUSICSYSTEM.currentSong, _ = mp3.DecodeWithSampleRate(music.SampleRate, f)

}

func NewMusicSystem(music models.BackgroundMusic) *MusicSystem {
	musicSystem := MusicSystem{}
	musicSystem.sampleRate = music.SampleRate
	f, _ := assets.Assets.Open(music.Path)
	musicSystem.currentSong, _ = mp3.DecodeWithSampleRate(music.SampleRate, f)
	return &musicSystem
}

func (musicSystem *MusicSystem) LoadSong(music models.BackgroundMusic) *MusicSystem {

	musicSystem.player.Close()
	musicSystem.sampleRate = music.SampleRate
	f, _ := assets.Assets.Open(music.Path)
	musicSystem.currentSong, _ = mp3.DecodeWithSampleRate(music.SampleRate, f)
	return musicSystem
}

func (musicSystem *MusicSystem) PlaySong() {
	if audio.CurrentContext() == nil {
		audioContext := audio.NewContext(musicSystem.sampleRate)
		loop := audio.NewInfiniteLoop(musicSystem.currentSong, musicSystem.currentSong.Length())
		musicSystem.player, _ = audioContext.NewPlayer(loop)
		musicSystem.player.Play()
		return
	}

	loop := audio.NewInfiniteLoop(musicSystem.currentSong, musicSystem.currentSong.Length())
	musicSystem.player, _ = audio.CurrentContext().NewPlayer(loop)
	musicSystem.player.Play()
}

func (musicSystem *MusicSystem) Pause() {
	musicSystem.player.Pause()
}

func (musicSystem *MusicSystem) Rewind() {
	musicSystem.player.Rewind()
}

func (musicSystem *MusicSystem) SetVolume(vol float64) {
	musicSystem.player.SetVolume(vol)
}

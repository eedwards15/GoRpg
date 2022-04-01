package models

import (
	"GoRpg/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font/opentype"
	_ "image/png"
	"log"
	"path"
)

type SceneAssets struct {
	BackgroundMusic BackgroundMusic
	Images          map[string]*ebiten.Image
	SoundEffects    map[string]*mp3.Stream
	Fonts           map[string]*opentype.Font
}

func NewSceneAssets(assetConfig *AssetConfig) *SceneAssets {
	l := SceneAssets{}

	//Background Music
	if assetConfig.BackgroundMusic != (BackgroundMusic{}) {
		l.BackgroundMusic = BackgroundMusic{
			Path:       assetConfig.BackgroundMusic.Path,
			SampleRate: assetConfig.BackgroundMusic.SampleRate,
		}
	}

	//Images
	l.Images = make(map[string]*ebiten.Image)
	for i := 0; i < len(assetConfig.Images); i++ {
		record := assetConfig.Images[i]
		filePath := path.Join(record.Location, record.FileName)
		l.Images[record.Key] = openImage(filePath)
	}

	//Sound Effects
	l.SoundEffects = make(map[string]*mp3.Stream)
	for i := 0; i < len(assetConfig.SoundEffects); i++ {
		record := assetConfig.SoundEffects[i]
		filePath := path.Join(record.Location, record.FileName)
		l.SoundEffects[record.Key] = openSound(filePath, record.SampleRate)
	}

	//Fonts
	l.Fonts = make(map[string]*opentype.Font)
	for i := 0; i < len(assetConfig.Fonts); i++ {
		record := assetConfig.Fonts[i]
		filePath := path.Join(record.Location, record.FileName)
		l.Fonts[record.Key] = openFont(filePath)
	}

	return &l
}

type BackgroundMusic struct {
	Path       string
	SampleRate int
}

func openImage(location string) *ebiten.Image {
	fs, _ := assets.Assets.Open(location)
	img, _, err := ebitenutil.NewImageFromReader(fs)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
func openSound(location string, sampleRate int) *mp3.Stream {
	f, _ := assets.Assets.Open(location)
	fireSound, _ := mp3.DecodeWithSampleRate(sampleRate, f)
	return fireSound
}
func openFont(location string) *opentype.Font {
	f, _ := assets.Assets.ReadFile(location)
	tt, err := opentype.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	return tt
}

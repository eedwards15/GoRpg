package models

type AssetConfig struct {
	Scene           string `json:"Scene"`
	BackgroundMusic struct {
		Path       string
		SampleRate int
	}
	Images []struct {
		Key      string `json:"Key"`
		Location string `json:"location"`
		FileName string `json:"FileName"`
	} `json:"Images"`
	SoundEffects []struct {
		Key        string `json:"Key"`
		Location   string `json:"location"`
		FileName   string `json:"FileName"`
		SampleRate int    `json:"SampleRate"`
	} `json:"Sounds"`
	Fonts []struct {
		Key      string `json:"Key"`
		Location string `json:"location"`
		FileName string `json:"FileName"`
	} `json:"Fonts"`
}

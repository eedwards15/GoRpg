package systems

import (
	"GoRpg/assets"
	"GoRpg/src/models"
	"encoding/json"
	"fmt"
	"log"
	"path"
	"sync"
)

var ASSETSYSTEM *AssetSystem

type AssetSystem struct {
	Assets map[string]*models.SceneAssets
	Credit []models.CreditConfig
}

func InitAssetSystem() {
	ASSETSYSTEM = &AssetSystem{}
	ASSETSYSTEM.Assets = make(map[string]*models.SceneAssets)
	configValues, _ := loadAssetConfigs()
	credits, _ := loadCredits()
	ASSETSYSTEM.Credit = credits

	var wg sync.WaitGroup
	for i := 0; i < len(configValues); i++ {
		wg.Add(1)
		r := configValues[i]
		go func(record *models.AssetConfig) {
			defer wg.Done()
			ASSETSYSTEM.Assets[record.Scene] = models.NewSceneAssets(record)

		}(r)
	}
	wg.Wait()

}

func loadAssetConfigs() ([]*models.AssetConfig, error) {
	files, err := assets.Assets.ReadDir("configs")
	if err != nil {
		fmt.Println("Failed Loading Configs")
		log.Fatal(err)
	}

	assetConfigs := []*models.AssetConfig{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileValue, _ := assets.Assets.ReadFile(path.Join("configs", file.Name()))
		assetConfig := models.AssetConfig{}

		json.Unmarshal(fileValue, &assetConfig)
		assetConfigs = append(assetConfigs, &assetConfig)
	}

	return assetConfigs, nil
}

func loadCredits() ([]models.CreditConfig, error) {
	files, err := assets.Assets.ReadDir("credits")
	if err != nil {
		fmt.Println("Failed Loading Configs")
		log.Fatal(err)
	}

	configs := []models.CreditConfig{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileValue, _ := assets.Assets.ReadFile(path.Join("credits", file.Name()))
		config := models.CreditConfig{}

		json.Unmarshal(fileValue, &config)
		configs = append(configs, config)
	}

	return configs, nil
}

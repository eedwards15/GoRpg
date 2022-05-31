package scenes

import (
	"GoRpg/src/entities"
	"GoRpg/src/helpers"
	"GoRpg/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"strconv"
)

type MainScene struct {
	menuFont   font.Face
	hasLoaded  bool
	tileMap    *ebiten.Image
	gameMap    *tiled.Map
	levelWidth int
	world      *ebiten.Image
	tileSize_s int
	camera     systems.Camera

	walls       []entities.Wall
	spawnPoints []entities.SpawnPoint
	player      *entities.Player
}

func NewMainScene() *MainScene {
	m := &MainScene{}
	m.hasLoaded = false
	return m
}

func (m *MainScene) Init() {
	systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.Assets["Forest"].BackgroundMusic).PlaySong()
	if !m.hasLoaded {
		m.world = ebiten.NewImage(800, 800)
		//Camera
		//m.camera = systems.Camera{ViewPort: f64.Vec2{800, 800}, ZoomFactor: 0}

		m.player = entities.InitPlayer(10, 10, systems.ASSETSYSTEM.Assets["Global"].Images["Player"])

		//m.camera.Follow(&m.player.Transform)

		m.tileSize_s = 32

		mapPath := "assets/maps/Forest.tmx" // Path to your Tiled Map.
		m.tileMap = systems.ASSETSYSTEM.Assets["Forest"].Images["DarkForest"]
		m.gameMap, _ = tiled.LoadFile(mapPath)
		m.levelWidth, _ = strconv.Atoi(m.gameMap.Properties.Get("screenWidth")[0])

		fontFromSystem := systems.ASSETSYSTEM.Assets["Global"].Fonts["KennySquare"]
		m.menuFont, _ = opentype.NewFace(fontFromSystem, &opentype.FaceOptions{
			Size:    60,
			DPI:     72,
			Hinting: font.HintingFull,
		})

		for i := 0; i < len(m.gameMap.ObjectGroups); i++ {
			layer := m.gameMap.ObjectGroups[i]
			if layer.Name == "wall" {
				for i := 0; i < len(layer.Objects); i++ {
					record := layer.Objects[i]
					newPlatform := entities.NewWall(record.X, record.Y, record.Width, record.Height)
					m.walls = append(m.walls, *newPlatform)
				}
			}

			if layer.Name == "runes" {
				for i := 0; i < len(layer.Objects); i++ {
					record := layer.Objects[i]
					spawnPoint := entities.NewSpawnPoint(record.X, record.Y, systems.ASSETSYSTEM.Assets["Global"].Images["Rune"])
					m.spawnPoints = append(m.spawnPoints, *spawnPoint)
				}
			}

			if layer.Name == "start" {
				record := layer.Objects[0]
				m.player.Xpos = record.X
				m.player.Ypos = record.Y
			}

		}

	}

}

func (m *MainScene) Update() error {

	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) || inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		systems.SCENEMANAGER.Pop()
	}

	m.player.IsColliding = false

	for i := 0; i < len(m.walls); i++ {
		record := m.walls[i]
		if helpers.Colision(record.Xpos, record.Ypos, record.Width, record.Height, m.player.Xpos, m.player.Ypos, float64(m.player.ImageWidth+(m.player.ImageWidth/2)), float64(m.player.ImageHeight)) {
			m.player.IsColliding = true

		}
	}

	m.player.Update()
	m.camera.Update()
	return nil
}

func (m MainScene) Draw(screen *ebiten.Image) {
	//text.Draw(screen, "MainScene", m.menuFont, ui.CenterTextXPos("MainScene", m.menuFont, systems.WINDOWMANAGER.ScreenWidth), systems.WINDOWMANAGER.YCenter, color.White)
	var xNum = m.levelWidth

	//for i, _ := range m.gameMap.Layers[0].Tiles {
	//	op := &ebiten.DrawImageOptions{}
	//	op.GeoM.Translate(float64((i%xNum)*m.tileSize_s), float64((i/xNum)*m.tileSize_s))
	//	m.world.DrawImage(systems.ASSETSYSTEM.Assets["Global"].Images["BG"], op)
	//}

	for i, tile := range m.gameMap.Layers[0].Tiles {
		if tile.Nil == false {
			spriteRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := m.tileMap.SubImage(spriteRect).(*ebiten.Image)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*m.tileSize_s), float64((i/xNum)*m.tileSize_s))
			m.world.DrawImage(tileImage, op)
		}
	}

	for i, tile := range m.gameMap.Layers[1].Tiles {
		if tile.Nil == false {
			spriteRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := m.tileMap.SubImage(spriteRect).(*ebiten.Image)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*m.tileSize_s), float64((i/xNum)*m.tileSize_s))
			m.world.DrawImage(tileImage, op)
		}
	}

	for i, tile := range m.gameMap.Layers[2].Tiles {
		if tile.Nil == false {
			spriteRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := m.tileMap.SubImage(spriteRect).(*ebiten.Image)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*m.tileSize_s), float64((i/xNum)*m.tileSize_s))
			m.world.DrawImage(tileImage, op)
		}
	}

	for _, sp := range m.spawnPoints {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(sp.Xpos, sp.Ypos)
		m.world.DrawImage(sp.CurrentImage, op)
	}

	m.player.Draw(m.world)

	m.camera.Render(m.world, screen)

}

func (m MainScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return systems.WINDOWMANAGER.ScreenWidth, systems.WINDOWMANAGER.ScreenHeight
}

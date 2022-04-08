package scenes

import (
	"GoRpg/src/entities"
	"GoRpg/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/f64"
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

	player *entities.Player
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
		m.camera = systems.Camera{ViewPort: f64.Vec2{200, 200}, ZoomFactor: 100}

		m.player = entities.InitPlayer(10, 10, systems.ASSETSYSTEM.Assets["Global"].Images["Player"])

		m.camera.Follow(&m.player.Transform)

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
	}

}

func (m *MainScene) Update() error {

	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) || inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		systems.SCENEMANAGER.Pop()
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

	m.player.Draw(m.world)
	m.camera.Render(m.world, screen)

}

func (m MainScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return systems.WINDOWMANAGER.ScreenWidth, systems.WINDOWMANAGER.ScreenHeight
}

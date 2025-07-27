package main

import (
	"fotr-rpg/world"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	TileSize     int = 16
	WindowWidth      = 1024
	WindowHeight     = 768

	CameraWidth = 800
	CameraEight = 600
)

var (
	cam *world.Camera
)

type Game struct{}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		cam.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		cam.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		cam.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		cam.X += 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	scene := GetSceneTileMapFile()

	tilemap, err := LoadScene(scene[Shire].tilemapJson)

	tileset, _, err := ebitenutil.NewImageFromFile("assets/tilesets/" + scene[Shire].tilesetName)

	if err != nil {
		log.Fatal(err)
	}

	opts := ebiten.DrawImageOptions{}

	for _, layer := range tilemap.Layers {
		for index, tileId := range layer.Data {
			x := index % layer.Width
			y := index / layer.Width

			x *= TileSize
			y *= TileSize

			srcX := (tileId - 1) % 32
			srcY := (tileId - 1) / 32

			srcX *= TileSize
			srcY *= TileSize

			opts.GeoM.Translate(float64(cam.X+x), float64(cam.Y+y))

			screen.DrawImage(
				tileset.SubImage(image.Rect(srcX, srcY, srcX+TileSize, srcY+TileSize)).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")

	cam = world.NewCamera(0, 0, 480, 360)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

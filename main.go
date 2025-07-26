package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	TileSize int = 16
)

type Game struct{}

func (g *Game) Update() error {
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

			opts.GeoM.Translate(float64(x), float64(y))

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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

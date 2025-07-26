package main

import (
	"encoding/json"
	"os"
)

type SceneId int

const (
	Shire SceneId = 1
)

type Scene struct {
	tilesetName string
	tilemapJson string
}

func GetSceneTileMapFile() map[SceneId]Scene {
	return map[SceneId]Scene{
		Shire: Scene{tilesetName: "sample.png", tilemapJson: "shire.json"},
	}
}

type TilemapLayer struct {
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type Tilemap struct {
	Layers []TilemapLayer `json:"layers"`
}

func LoadScene(tilemapJson string) (*Tilemap, error) {
	data, err := os.ReadFile("assets/tilesets/" + tilemapJson)

	if err != nil {
		return nil, err
	}

	var tilemap Tilemap
	json.Unmarshal(data, &tilemap)

	return &tilemap, nil
}

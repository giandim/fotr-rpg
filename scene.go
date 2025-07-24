package main 

type Scene struct {
  Name string
}

type TilemapLayer struct {
  Data []int `json:"data"`
  Width int `json:"width"`
  Height int `json:"height"`
} 

type Tilemap struct {
  Layers []TilemapLayer `json:"layers"`
}

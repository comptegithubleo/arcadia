package engine

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileMap struct {
	Dest       rl.Rectangle
	Src        rl.Rectangle
	Map        []int
	srcMap     []string
	mapW, mapH int
}

func (e *Engine) InitMap(mapFile string) {
	file, err := os.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var remainingNewLines = strings.Replace(string(file), "\n", " ", -1)
	var sliced = strings.Split(remainingNewLines, " ")
	fmt.Println(sliced)
	e.TileMap.mapW = -1
	e.TileMap.mapH = -1
	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		if e.TileMap.mapW == -1 {
			e.TileMap.mapW = m
		} else if e.TileMap.mapH == -1 {
			e.TileMap.mapH = m
		} else if i < e.TileMap.mapW*e.TileMap.mapH+2 {
			e.TileMap.Map = append(e.TileMap.Map, m+1)
		} else {
			e.TileMap.srcMap = append(e.TileMap.srcMap, sliced[i])
		}
	}
	if len(e.TileMap.Map) > e.TileMap.mapW*e.TileMap.mapH {
		e.TileMap.Map = e.TileMap.Map[:len(e.TileMap.Map)-1]
	}
}

func (e *Engine) RenderMap() {

	//Plants Shadow
	rl.DrawTexturePro(
		e.Sprites.PlantsShadow,
		rl.NewRectangle(0, 0, 128, 128), //source rectangle
		rl.NewRectangle(0, 0, 128, 128),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
	//Plants
	rl.DrawTexturePro(
		e.Sprites.Plants,
		rl.NewRectangle(0, 0, 128, 128), //source rectangle
		rl.NewRectangle(0, 0, 128, 128),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

	for i := 0; i < len(e.TileMap.Map); i++ {
		if e.TileMap.Map[i] != 0 {
			e.TileMap.Dest.X = e.TileMap.Dest.Width * float32(i%e.TileMap.mapW)
			e.TileMap.Dest.Y = e.TileMap.Dest.Height * float32(i/e.TileMap.mapW)
		}
	}
}

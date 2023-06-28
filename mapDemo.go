package main

import (
	"fmt"
	"os"

	"github.com/lafriks/go-tiled"
)

const mapPath = "demoMap.tmx" // Path to your Tiled Map.

type mapGame struct {
}

func main() {
	// Parse .tmx file.
	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	fmt.Println("tilesets:", gameMap.Tilesets[0])
	//fmt.Println("layers:", gameMap.Layers[0].Tiles)
	fmt.Println("type:")
}

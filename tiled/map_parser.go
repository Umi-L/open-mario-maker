package tiled

import (
	"github.com/umi-l/waloader"
	"log"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
)

func ParseCsv(file string, sheet waloader.Sheet) Map {

	file = strings.ReplaceAll(file, "\n", ",")

	strData := strings.Split(file, ",")

	var data []int

	for _, num := range strData {

		if num == "" {
			continue
		}

		value, err := strconv.Atoi(num)

		if err != nil {
			log.Print("Invalid Value In File!")
			log.Fatal(err)
		}

		data = append(data, value)
	}

	container := make([]TileLayer, 1)
	container[0] = TileLayer{
		TileData: data,
		width:    40,
		height:   30,
		id:       0,
		name:     "csvmap (depricated)",
		visible:  true,
		x:        0,
		y:        0,
	}

	return Map{
		Sheet:      sheet,
		TileLayers: container,
	}
}

func ParseJson(file []byte, sheet waloader.Sheet) Map {

	var TileLayers []TileLayer

	jsonparser.ArrayEach(file, func(layerContainer []byte, dataType jsonparser.ValueType, offset int, err error) {
		// log.Print(string(layerContainer))

		if val, _ := jsonparser.GetString(layerContainer, "type"); val == "tilelayer" {

			// log.Print(jsonparser.Array)

			var data []int

			jsonparser.ArrayEach(layerContainer, func(tileid []byte, dataType jsonparser.ValueType, offset int, err error) {
				// data = append(data, )
				value, err := strconv.Atoi(string(tileid))

				if err != nil {
					log.Print("Invalid Value In File!")
					log.Fatal(err)
				}

				data = append(data, value)
			}, "data")

			layer := TileLayer{
				TileData: data,

				width:  safeGetValueInt("width", layerContainer),
				height: safeGetValueInt("height", layerContainer),

				x: safeGetValueInt("x", layerContainer),
				y: safeGetValueInt("y", layerContainer),

				id: safeGetValueInt("id", layerContainer),

				name:    safeGetValueString("name", layerContainer),
				visible: safeGetValueBool("visible", layerContainer),
			}

			TileLayers = append(TileLayers, layer)
		}

	}, "layers")

	return Map{
		Sheet:      sheet,
		TileLayers: TileLayers,
	}
}

func safeGetValueInt(name string, json []byte) int {
	val, err := jsonparser.GetInt(json, name)

	if err != nil {
		log.Print("error parsing map file")
		log.Fatal(err)
	}

	return int(val)
}

func safeGetValueString(name string, json []byte) string {
	val, err := jsonparser.GetString(json, name)

	if err != nil {
		log.Print("error parsing map file")
		log.Fatal(err)
	}

	return val
}

func safeGetValueBool(name string, json []byte) bool {
	val, err := jsonparser.GetBoolean(json, name)

	if err != nil {
		log.Print("error parsing map file")
		log.Fatal(err)
	}

	return val
}

package tiled

import (
	"github.com/umi-l/open-mario-maker/loader"
	"log"
	"strconv"
	"strings"
)

func ParseCsv(file string, sheet loader.Sheet) Map {

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

	return Map{
		Sheet: sheet,
		Data:  data,
	}
}

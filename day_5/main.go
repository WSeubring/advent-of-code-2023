package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

type AlmanacMap struct {
	dest_range_start   int
	source_range_start int
	range_length       int
}

type Almanac struct {
	seeds                  []int
	seed_2_soil            []AlmanacMap
	soil_2_fertilizer      []AlmanacMap
	fertilizer_2_water     []AlmanacMap
	water_2_light          []AlmanacMap
	ligth_2_temperature    []AlmanacMap
	temperature_2_humidity []AlmanacMap
	humidity_2_location    []AlmanacMap
}

func parseAlamnacMap(lines []string) []AlmanacMap {
	maps := make([]AlmanacMap, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		source, _ := strconv.Atoi(parts[0])
		source_range_start, _ := strconv.Atoi(parts[1])
		range_length, _ := strconv.Atoi(parts[2])
		maps[i] = AlmanacMap{source, source_range_start, range_length}
	}
	return maps
}

func parseAlamnac(lines []string) Almanac {
	sections := make([][]string, 0)
	section := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			sections = append(sections, section)
			section = make([]string, 0)
		} else {
			section = append(section, line)
		}
	}

	sections = append(sections, section)
	seeds := make([]int, 0)
	// Select all after the first work in the first line
	seedsInput := strings.Fields(sections[0][0])[1:]
	for _, seed := range seedsInput {
		seedInt, _ := strconv.Atoi(seed)
		seeds = append(seeds, seedInt)
	}

	return Almanac{
		seeds,
		parseAlamnacMap(sections[1][1:]),
		parseAlamnacMap(sections[2][1:]),
		parseAlamnacMap(sections[3][1:]),
		parseAlamnacMap(sections[4][1:]),
		parseAlamnacMap(sections[5][1:]),
		parseAlamnacMap(sections[6][1:]),
		parseAlamnacMap(sections[7][1:]),
	}
}

func (almanacMap *AlmanacMap) IsWithinRange(source int) bool {
	return source >= almanacMap.source_range_start && source < almanacMap.source_range_start+almanacMap.range_length
}

func (almanacMap *AlmanacMap) Destination(source int) int {
	if !almanacMap.IsWithinRange(source) {
		panic("Source out of range, this should not happen")
	}
	// return the output range with the offset of the input range
	return almanacMap.dest_range_start + (source - almanacMap.source_range_start)
}

func (almanac *Almanac) ApplyMap(source int, maps []AlmanacMap) int {
	for _, almanacMap := range maps {

		if almanacMap.IsWithinRange(source) {
			return almanacMap.Destination(source)
		}
	}
	fmt.Printf("No map found for %v \n", source)
	return source
}

func (almanac *Almanac) MinimumLocation() int {
	locations := make([]int, len(almanac.seeds))
	for i, seed := range almanac.seeds {
		locations[i] = almanac.GetLocation(seed)
	}
	return slices.Min(locations)
}

func (almanac *Almanac) GetLocation(seed int) int {
	soil := almanac.ApplyMap(seed, almanac.seed_2_soil)
	fertilizer := almanac.ApplyMap(soil, almanac.soil_2_fertilizer)
	water := almanac.ApplyMap(fertilizer, almanac.fertilizer_2_water)
	light := almanac.ApplyMap(water, almanac.water_2_light)
	temperature := almanac.ApplyMap(light, almanac.ligth_2_temperature)
	humidity := almanac.ApplyMap(temperature, almanac.temperature_2_humidity)
	location := almanac.ApplyMap(humidity, almanac.humidity_2_location)
	return location
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)

	panicIfError(err)
	// Close the file when we leave the scope of the current function,
	defer file.Close()

	// Make a buffer to keep chunks that are read.
	fileScanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	almanac := parseAlamnac(lines)

	fmt.Println(almanac.MinimumLocation())

	panicIfError(fileScanner.Err())
}

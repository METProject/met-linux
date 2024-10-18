package util

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	var cfg Config
	if _, err = toml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

type Config struct {
	Gen  GenConfig  `toml:"gen"`
	Path PathConfig `toml:"path"`
	Osm  OsmConfig  `toml:"osm"`
}

type GenConfig struct {
	WorldName              string `toml:"world_name"`
	BlocksPerTile          int    `toml:"blocks_per_tile"`
	DegreePerTile          int    `toml:"degree_per_tile"`
	HeightRatio            int    `toml:"height_ratio"`
	Threads                int    `toml:"threads"`
	UseHeighQualityTerrain bool   `toml:"use_heigh_quality_terrain"`
}

type PathConfig struct {
	Pbf                   string `toml:"pbf_path"`
	OsmFolder             string `toml:"osm_folder_path"`
	QgisProject           string `toml:"qgis_project_path"`
	QgisBathymetryProject string `toml:"qgis_bathymetry_project_path"`
	QgisTerrainProject    string `toml:"qgis_terrain_project_path"`
	QgisHeightmapProject  string `toml:"qgis_heightmap_project_path"`
	ScriptsFolder         string `toml:"scripts_folder_path"`
	OutputFolder          string `toml:"output_folder_path"`
}

type OsmConfig struct {
	Switch SwitchConfig `toml:"switch"`
	Rivers string       `toml:"rivers"`
}

type SwitchConfig struct {
	Aerodrome   bool `toml:"aerodrome"`
	BareRock    bool `toml:"bare_rock"`
	Beach       bool `toml:"beach"`
	BigRoad     bool `toml:"big_road"`
	Border      bool `toml:"border"`
	Farmland    bool `toml:"farmland"`
	Forest      bool `toml:"forest"`
	Glacier     bool `toml:"glacier"`
	Grass       bool `toml:"grass"`
	Highway     bool `toml:"highway"`
	Meadow      bool `toml:"meadow"`
	MiddleRoad  bool `toml:"middle_road"`
	Quarry      bool `toml:"quarry"`
	River       bool `toml:"river"`
	SmallRoad   bool `toml:"small_road"`
	Stateborder bool `toml:"stateborder"`
	Stream      bool `toml:"stream"`
	Swamp       bool `toml:"swamp"`
	Urban       bool `toml:"urban"`
	Volcano     bool `toml:"volcano"`
	Water       bool `toml:"water"`
	Wetland     bool `toml:"wetland"`
	Vineyard    bool `toml:"vineyard"`
}

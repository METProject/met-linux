package qgis

// Commands for QGIS preprocessing
var PreprocessCommands = map[string]string{
	"test":        "ls -la",
	"highway":     "osmium tags-filter {P_FILE} w/highway=motorway,trunk -o {OUT}highway.osm",
	"big_road":    "osmium tags-filter {P_FILE} w/highway=primary,secondary -o {OUT}big_road.osm",
	"middle_road": "osmium tags-filter {P_FILE} w/highway=tertiary -o {OUT}middle_road.osm",
	"small_road":  "osmium tags-filter {P_FILE} w/highway=residential -o {OUT}small_road.osm",
	"stream":      "osmium tags-filter {P_FILE} w/waterway=river,stream w/water=river -o {OUT}stream.osm",
	"aerodrome":   "osmium tags-filter {P_FILE} w/aeroway=launchpad -o {OUT}aerodrome.osm",
	"urban":       "osmium tags-filter {P_FILE} w/landuse=commercial,construction,industrial,residential,retail -o {OUT}urban.osm",
	"wetland":     "osmium tags-filter {P_FILE} w/natural=wetland -o {OUT}wetland.osm",
	// logical and is needed here
	"volcano":   "osmium tags-filter -f pbf {P_FILE} w/natural=volcano | osmium tags-filter -o {OUT}volcano.osm - w/volcano:status=active -R -F pbf",
	"bare_rock": "osmium tags-filter {P_FILE} w/natural=scree,shingle w/landuse=bare_rock -o {OUT}bare_rock.osm",
	"quarry":    "osmium tags-filter {P_FILE} w/landuse=quarry -o {OUT}quarry.osm",
	"grass":     "osmium tags-filter {P_FILE} w/landuse=grass,fell,heath,scrub w/natural=grassland -o {OUT}grass.osm",
	"meadow":    "osmium tags-filter {P_FILE} w/landuse=meadow -o {OUT}meadow.osm",
	"vineyard":  "osmium tags-filter {P_FILE} w/landuse=vineyard -o {OUT}vineyard.osm",
	"farmland":  "osmium tags-filter {P_FILE} w/landuse=farmland -o {OUT}farmland.osm",
	"forest":    "osmium tags-filter {P_FILE} w/landuse=forest -o {OUT}forest.osm",
	"beach":     "osmium tags-filter {P_FILE} w/natural=beach -o {OUT}beach.osm",
	"glacier":   "osmium tags-filter {P_FILE} w/natural=glacier -o {OUT}glacier.osm",
	"river":     "osmium tags-filter {P_FILE} w/waterway=river,riverbank,canal w/water=river -o {OUT}river.osm",
	// logical and is needed here
	"swamp": "osmium tags-filter -f pbf {P_FILE} w/natural=wetland | osmium tags-filter -o {OUT}swamp.osm - w/wetland=swamp -R -F pbf",
	"water": "osmium tags-filter {P_FILE} w/natural=water w/water=lake,reservoir w/landuse=reservoir -o {OUT}water.osm",
	// fucked up things needed here
	// first we pipe 2 commands into each other to create a logical "AND" then write a temporary file
	// ten, we parse this file again with a logical "AND"
	"stateborder": "osmium tags-filter -f pbf {P_FILE} w/boundary=administrative | osmium tags-filter -o {OUT}state_raw.osm - w/admin_level=3,4 -R -F pbf && osmium tags-filter -f pbf {OUT}state_raw.osm w/natural!=coastline | osmium tags-filter -o {OUT}stateborder.osm - w/admin_level!=2,5,6,7,8,9,10,11 -R -F pbf && rm -rf {OUT}state_raw.osm",
	"border":      "osmium tags-filter -f pbf {P_FILE} w/boundary=administrative | osmium tags-filter -o {OUT}border_raw.osm - w/admin_level=2 -R -F pbf && osmium tags-filter -f pbf {OUT}border_raw.osm w/natural!=coastline | osmium tags-filter -o {OUT}border.osm - w/admin_level!=3,4,5,6,7,8,9,10,11 -R -F pbf && rm -rf {OUT}border_raw.osm",
}

// OSM Postfix things
var OsmPostfix = map[string]struct {
	inputFile string
	suffix    string
}{
	"urban":        {"urban", "|layername=multipolygons"},
	"broadleaved":  {"forest", "|layername=multipolygons|\nsubset=\"other_tags\" = '\"leaf_type\"'=>\"broadleaved\""},
	"needleleaved": {"forest", "|layername=multipolygons|\nsubset=\"other_tags\" = '\"leaf_type\"'=>\"needleleaved\""},
	"mixedforest":  {"forest", "|layername=multipolygons"},
	"beach":        {"beach", "|layername=multipolygons"},
	"grass":        {"grass", "|layername=multipolygons"},
	"farmland":     {"farmland", "|layername=multipolygons"},
	"meadow":       {"meadow", "|layername=multipolygons"},
	"quarry":       {"quarry", "|layername=multipolygons"},
	"water":        {"water", "|layername=multipolygons|subset=\"natural\" = 'water'"},
	"glacier":      {"glacier", "|layername=multipolygons"},
	"wetland":      {"wetland", "|layername=multipolygons"},
}

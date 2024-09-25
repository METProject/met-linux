# IMPORTANT

This is an outdated, potentially broken version that was NOT maintained by me. A full rewrite is occuring in the **next** branch (yet to be made).
Using this may be error prone!

# world-generator

The World Generator inspired by [Minecraft Earth Map](https://earth.motfe.net/) but runs in parallel.

## Usage

0. Docker is required to run the world generator. Please install Docker first. See [Docker Installation](https://docs.docker.com/get-docker/).
1. Download all the required data. See [Tiles Installation](https://earth.motfe.net/tiles-installation/).Put them all in the `Data` folder and unzip.
2. Download the project Copy the `config.example.yaml` to `config.yaml` and modify the `config.yaml` to your needs.
3. Run the following command in the project root directory.

```bash
docker pull alicespaceli/trumancrafts_builder:v0.0.3
docker run -idt --rm -v $(pwd):/workspace alicespaceli/trumancrafts_builder:v0.0.3
```

You can check the progress by just looking into the `generator.log`.

The folders in the project, all the folder names should be the same as your `config.yaml`

```
.
├── Data                 /* `scripts_folder_path` in config, unzip the worldpainter-script.zip here */
│   ├── voidscript.js
│   ├── worldpainter-script.zip
│   ├── wpscript
│   ├── wpscript.js
│   ├── osm              /* `osm_folder_path` in config, osm files will be generated here */
│   │   └── all
│   ├── qgis-bathymetry  /* unzip the qgis-bathymetry.zip here */
│   │   └── QGIS
│   ├── qgis-heightmap   /* unzip the qgis-heightmap.zip.001... here */
│   │   └── QGIS
│   ├── qgis-terrain     /* unzip the qgis-terrain.zip here */
│   │   └── QGIS
│   ├── qgis-project     /* unzip the qgis-project.zip here */
│   │   └── QGIS
│   └── wpscript
│       ├── backups
│       ├── exports
│       ├── farm
│       ├── layer
│       ├── ocean
│       ├── ores
│       ├── roads
│       ├── schematics
│       ├── terrain
│       └── worldpainter_files
├── Docker


```

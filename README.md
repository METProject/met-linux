# MET-Linux

Dockerized and parallelized version of the [Minecraft Earth Map](https://earth.motfe.net/) generator.
Recently open sourced!

## Installation - INCOMPLETE DOCUMENTATION, PLEASE WAIT

We recommend using [Docker](https://docs.docker.com/get-docker/) with [Docker Compose](https://docs.docker.com/compose/install/) to run the generator. However, you can still run it from source.and

### Docker Compose (recommended)

Ensure you have Docker and Docker Compose installed.
```bash
docker -v
docker compose
```

Then, continue to the Wiki page for [Installation](https://github.com/imide/met-linux/wiki/Installation).

### Docker (legacy)

Documentation is incomplete. Follow (legacy instructions)[https://github.com/imide/met-linux/tree/main] at your own risk.

## Contributing

Contributions are welcome!

## TODO/Checklist

- [x] Finish this checklist
- [ ] Rewrite the legacy generator with optimised code
- [ ] Add better documentation (wiki pages or docs website on installation, usage, etc.)
- [ ] Merge the bash processor for speed improvements (possibly up to 4x)?
- [ ] Release beta version 0.1.0

### Stable Release

- [ ] Automate downloading of the required data with either configuration (most likely) or a script
- [ ] Stability improvements
- [ ] Rewrite in lower level language? (C++ as all the neccessary libraries are already available)

## License

This project is licensed under the GNU AGPLv3 License as per (upstream)[https://github.com/DerMattinger/MinecraftEarthTiles] - see the [LICENSE](LICENSE) file for details.

## Credits

- [DerMattinger](https://github.com/DerMattinger) - for the original [Minecraft Earth Tiles](https://earth.motfe.net/) project.
- [imide](https://github.com/imide) - for the [MET Linux](https://github.com/imide/met-linux) project.
- [AliceSpaceLi](https://github.com/AliceSpaceLi) - for the [world-generator](https://github.com/truman-crafts/world-generator) project (legacy generator).
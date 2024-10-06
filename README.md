# MET-Linux

Dockerized and parallelized version of the [Minecraft Earth Map](https://earth.motfe.net/) generator.
Recently open sourced!

## Installation - INCOMPLETE DOCUMENTATION, PLEASE WAIT

We recommend using [Docker](https://docs.docker.com/get-docker/) with [Docker Compose](https://docs.docker.com/compose/install/) to run the generator. However, you can still run it from source (Nix is required to avoid issues!)

### Docker Compose (recommended)

Ensure you have Docker and Docker Compose installed.
```bash
docker -v
docker compose
```

Then, continue to the Wiki page for [Installation](https://github.com/imide/met-linux/wiki/Installation).

### Docker (legacy)

Documentation is incomplete. Follow (legacy instructions)[https://github.com/imide/met-linux/tree/main] at your own risk.

## Development

### Nix

To set up a development environment, the following is needed:

- A UNIX-like environment (Linux, macOS, WSL, etc.)
- Nix (see [Determinate Systems](https://github.com/DeterminateSystems/nix-installer))
- [Direnv](https://direnv.net/) hooked into [shell](https://direnv.net/docs/hook.html)
- Git (duh)

Then, run the following commands:

```bash
git clone https://github.com/imide/met-linux.git
cd met-linux
direnv allow

# Open your IDE of choice from the shell (to avoid nix-related issues) and start developing! For example with VSCodium or Cursor:
cursor .
codium .
```

## Contributing

Contributions are welcome!

## TODO/Checklist

- [x] Finish this checklist
- [ ] Rewrite in Go and bash
- [ ] Add better documentation (wiki pages or docs website on installation, usage, etc.)
- [ ] Merge the bash processor for speed improvements (possibly up to 4x)?
- [ ] Release beta version 0.1.0

### Stable Release

- [ ] Automate downloading of the required data with either configuration (most likely) or a script
- [ ] Stability improvements

## License

This project is licensed under the GNU AGPLv3 License as per (upstream)[https://github.com/DerMattinger/MinecraftEarthTiles] - see the [LICENSE](LICENSE) file for details.

## Credits

- [DerMattinger](https://github.com/DerMattinger) - for the original [Minecraft Earth Tiles](https://earth.motfe.net/) project.
- [imide](https://github.com/imide) - for the [MET Linux](https://github.com/imide/met-linux) project.
- [MeerBiene](https://github.com/MeerBiene) - for the shell implementation (not in this project yet) of the osmium generator. (literally made this possible to make without python)
- [AliceSpaceLi](https://github.com/AliceSpaceLi) - for the [world-generator](https://github.com/truman-crafts/world-generator) project (legacy generator).

<h1 align="center">
  WIP: fallout-termlink
</h1>

<p align="center">
  <strong>
    Recreation of Fallout's hacking mini-game implemented in Golang.
  </strong>
</p>

<p align="center">
  <a href="https://github.com/rmpalgo/fallout-termlink/blob/main/LICENSE"><img src=https://img.shields.io/github/license/rmpalgo/fallout-termlink
 alt="License Status"></a>
</p>

### 
<em>Disclaimer: This project is not affiliated with the Fallout game series, Zenimax, or Bethesda. - Please contact me [here](mailto:ronnelpalencia@proton.com) if there are any issues.</em>

<img src="./docs/demo_wip.gif" width="250" height="250"/>

## Installation

### Go

```bash
go install github.com/rmpalgo/fallout-termlink@latest
```

## Usage

### Controls

The default game controls:

- Move Left: `A`
- Move Right: `D`
- Move Down: `S`
- Move Up: `W`
- Choose Selection: `Enter`
- Exit Game: `q`
- Force Quit: `ctrl+c`

### Help

```
usage: termlink [<options>]
  -v	Print version and exit
```

## Configuration

### CLI

Starting the termlink without flags or subcommands will automatically start the mini-game in normal mode:

```bash
./fallout-termlink
````
## License

[MIT](https://github.com/rmpalgo/fallout-termlink/blob/main/LICENSE)

### Acknowledgments
Built with [bubbletea](https://github.com/charmbracelet/bubbletea)

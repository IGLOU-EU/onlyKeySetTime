# onlyKeySetTime [![Go Report Card](https://goreportcard.com/badge/git.iglou.eu/adrien/onlyKeySetTime)](https://goreportcard.com/report/git.iglou.eu/adrien/onlyKeySetTime)

An OnlyKey set time, implemented on GO.
This project only need a system with udev to run (or binary usage only), no more Py script.
Because it is a cool first Go project and i don't like Py (especially 2.7)

This udev rules are base on [official rules](https://raw.githubusercontent.com/trustcrypto/trustcrypto.github.io/master/49-onlykey.rules), but add (0x1d50, 0x60fc) and (0x16C0, 0x0486) devices support into.

## Installation

```
git clone https://git.iglou.eu/adrien/onlyKeySetTime.git
cd onlyKeySetTime
make install
```

## Binary only

You can use makefile for Cross-Compiling, ex : 'make build GOOS='linux' GOARCH='arm''"
Binary are on `./bin`

```
git clone https://git.iglou.eu/adrien/onlyKeySetTime.git
cd onlyKeySetTime
make build
```

## License

MIT License, see [LICENSE.md](LICENSE.md).
# onlyKeySetTime [![Go Report Card](https://goreportcard.com/badge/git.iglou.eu/adrien/onlyKeySetTime)](https://goreportcard.com/report/git.iglou.eu/adrien/onlyKeySetTime)

An OnlyKey set time, implemented on GO.
This project only need a system with udev to run (or binary usage only), no more Py script.
Because it is a cool first Go project and i don't like Py (especially 2.7)

This udev rules are base on [official rules](https://raw.githubusercontent.com/trustcrypto/trustcrypto.github.io/master/49-onlykey.rules), but add (0x1d50, 0x60fc) and (0x16C0, 0x0486) devices support into.

## Installation

**From releases** *Replace ARCH - regular udev and local/bin - need wget and to be root*
```
ARCH='amd64' # 386 amd64 arm arm64
OS='linux'

bin="onlykey-settime_${OS}_${ARCH}"
rules="49-onlykey-settime.rules"
binlink="https://git.iglou.eu/adrien/onlyKeySetTime/releases/download/v1.0.0/"
tmp="$(mktemp -d)" && cd "$tmp"
wget "${binlink}${bin}" && chmod +x "${bin}"
wget "${binlink}${bin}.sha256"
wget "${binlink}${rules}"
sha256sum -c "${bin}.sha256" && { cp "${bin}" "/usr/local/bin/${bin%%_*}"; cp ${rules} /etc/udev/rules.d/; udevadm control --reload; }
cd "/tmp" && rm -rf "${tmp}"
echo "DONE !"
```

**From source code**
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

## Clean and Uninstall

**For clean repos**

```
make clean
```

**For uninstall**

```
make remove
```

## License

MIT License, see [LICENSE.md](LICENSE.md).

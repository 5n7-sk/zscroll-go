# zscroll-go

`zscroll` is a text scroller for panels or terminals.

## Example

```sh
zscroll -a " <<<" -b ">>> " -d 0.3 -l 20 `spotify-cli status --kind title`
```

![demo](./misc/demo.gif)

## Installation

Get the binary from [GitHub Releases](https://github.com/skmatz/zscroll-go/releases).

Or, if you have Go, you can install `zscroll` with the following command.

```console
go get github.com/skmatz/zscroll-go/...
```

## Usage

```console
> zscroll --help

A text scroller for panels or terminals.

Usage:
  zscroll [flags]
  zscroll [command]

Available Commands:
  completion  Output shell completion (bash/fish/powershell/zsh)
  help        Help about any command
  version     Show version

Flags:
  -a, --after-text string       stationary paddin text to display to the right side of the scroll-text
  -b, --before-text string      stationary paddin text to display to the left side of the scroll-text
  -d, --delay float             delay in seconds for scrolling update (default 0.4)
  -h, --help                    help for zscroll
  -l, --length int              length of the scroll-text (default -1)
  -n, --new-line                whether to print a new line after each update (default true)
  -p, --padding-text string     padding text to diplay between the end and the head of the scroll-text (default " - ")
  -r, --reverse                 scroll from left to right
  -s, --scroll                  whether to scroll (default true)
  -S, --scroll-rate int         number of characters to scroll (default 1)
  -t, --timeout int             time in seconds to exit (default -1)
  -u, --update-command string   update command to change the scroll-text
  -U, --update-interval int     time in seconds to execute the update command (default 1)
  -V, --version                 show version

Use "zscroll [command] --help" for more information about a command.
```

### Dynamic Updating Text

You can change the scroll-text dynamically by giving the command to update the text as an argument.

```sh
zscroll \
  -a " <<<" \
  -b ">>> " \
  -d 0.3 \
  -l 20 \
  -u "spotify-cli status --kind title" \
  `spotify-cli status --kind title`
```

## References

- [noctuid/zscroll](https://github.com/noctuid/zscroll)

# Search me

Open your browser from terminal to search things.

## Usage

### Syntax

```
search-me [engine] and things to search
```

### Examples

```
search-me ddg vim screencast !yt
search-me google vim screencast
search-me vim screencast
```

#### Engines

- main: `http://google.com/`
- google: `http://google.com/`
- ddg: `https://duckduckgo.com/`

### Define your engines

See [config_example][1]. Feel free to place your favorite search engines
to `~/.search.json`:

```
engine-name http://example.com/?q=%s
```

### Aliases

You can alias your favorite engines. For example

```bash
alias ddg=search-me ddg
alias google=searh-me google
alias jquery=search-me jquery
```

### Dotfiles

You can save your favorite search engines into [dotfiles][2] to never miss it!

### Development

```
go get
go run search.go ddg adventure time !v
```

[1]: https://github.com/shuvalov-anton/go-search-me/blob/master/config_example
[2]: http://dotfiles.github.io/

# Search me

Open your browser from terminal to search things.

## Usage

### Development

```
go get
go run search.go ddg adventure time !v
```


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


#### Aliases

You can just alias for your favorite engines. For example

```bash
alias ddg=search-me ddg
alias google=searh-me google
alias jquery=search-me jquery
```

#### Predefined Engines

- main: `http://google.com/`
- google: `http://google.com/`
- ddg: `https://duckduckgo.com/`


#### Define your engines

Just place your search engines to `~/.search.json`:

```
{
  "engine-name": "http://example.com/?q=%s"
}
```

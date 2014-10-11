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
- jquery: `http://api.jquery.com/`
- mdn: `http://developer.mozilla.org/?q=%s`
- compass: `http://compass-style.org/search?q=%s`
- html5please: `http://html5please.com#%s`
- caniuse: `http://caniuse.com/#search=%s`
- codepen: `http://codepen.io/search?q=%s`
- bem: `http://google.com//search?as_q=%s&as_sitesearch=bem.info`
- angularjs: `http://google.com//search?as_q=%s&as_sitesearch=angularjs.org`
- reactjs: `http://google.com//search?as_q=%s&as_sitesearch=facebook.github.io/react`
- emberjs: `http://emberjs.com/api/#stq=%s&stp=1`
- hello: `http://example.com/?q=%s`


#### Define your engines

Just place your search engines to `~/.search.json`:

```
{
  "engine-name": "http://example.com/?q=%s"
}
```

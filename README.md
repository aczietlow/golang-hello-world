# Playground

What started out as a hello world example has turned into a sandbox of chaos of the last year. At this point it's become a bit of a scratch pad as I'm working through the core language specification or digging more into the standard library.

## Resources
* [personal miro board notes](https://miro.com/app/board/uXjVPYZIqT0=/?share_link_id=198914050024)
* https://golangbot.com/variables/

### More resources for future musings
* https://www.ardanlabs.com/blog/2015/09/composition-with-go.html
* https://github.com/golang/go/wiki/LearnConcurrency
* https://jdanger.com/build-a-web-crawler-in-go.html

## TIL

### Error formatters exist

fmt has a formatter that returns a string that satisfies an error type.

```go
// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.

fmt.Errorf("parsing %s as HTML: %v", url, err)
```

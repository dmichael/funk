funk
====

A port of Underscore.js function functions to Go.

If you would like to contribute, please fork the project and issue a pull request.

### Throttle

```Go
import "github.com/dmichael/funk"

// The func to be throttled
func notify(message string) {
  // ping remote service or something with a message
}

func main(){
  // funk needs a var to populate with the 'throttled' func,
  // it must be the same signature of the func to throttle
  var throttled func(string)

  funk.Throttle(notify, &throttled, 100)

  // Only fired once
  for i := 0; i < 1000; i++ {
    throttled("Oops! Something went wrong!")
  }
  
}
```

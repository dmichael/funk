funk
====

A port of Underscore.js function functions to Go.

If you would like to contribute, please fork the project and issue a pull request.

### Throttle

Throttle will only execute the passed in func at most once per interval specified. It does not distinguish which values you call it with, wrapping only the first one - you should manage this manually.

```Go
import "github.com/dmichael/funk"
import "fmt"

// The func to be throttled
func RickJames(message string) {
  fmt.Println(message)
}

func main(){
  throttled := Throttle{wait: 500 * time.Millisecond}
  
  // Do only accepts void funcs, so wrap it
  for i := 0; i < 10; i++ {
    // This is only triggered once after 500 ms elaspsed since first called
    throttled.Do(func() { 
      RickJames("Whoah-u-o-u-o-u sixty-nine")
    })
  }
  
  
}
```

TODO: 

### Delay
TODO

### Debounce
TODO
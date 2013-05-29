funk
====

A port of Underscore.js function functions to Go

### Throttle

```Go
func notifyError(message string) {
  // ping remote service or something with a message
}

func main(){
  var throttledNotifyError func(string, int)

  Throttle(notifyError, &throttledNotifyError, 100)

  // Only fired once
  for i := 0; i < 1000; i++ {
    throttledNotifyError("Oops! Something went wrong!")
  }
  
}


```

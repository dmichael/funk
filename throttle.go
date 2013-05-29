package funk

import (
  "reflect"
  "time"
)

func Throttle(fn interface{}, fnptr interface{}, milliseconds int) {
  ff := reflect.ValueOf(fn)
  ft := ff.Type()
  throttling := false

  if ft.NumOut() > 0 {
    panic("Throttled funcs may not have return values!")
  }

  var out []reflect.Value

  // This is the returned function
  // MakeFunc requires this signature, however, the return values will be empty
  prototype := func(in []reflect.Value) []reflect.Value {
    // Already called
    if throttling { return out }

    countdown := func() {
      time.AfterFunc(time.Duration(milliseconds) * time.Millisecond, func() {
        throttling = false
        // Call the original func that was passed in.
        // We dont care about the output since it cannot be returned
        ff.Call(in)
      })
    }

    throttling = true
    go countdown()

    return out
  }

  // Create a function that wraps the original
  v := reflect.MakeFunc(ft, prototype)

  // Set pointer to function that was passed in to the created function
  reflect.ValueOf(fnptr).Elem().Set(v)
}

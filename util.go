package funk


func createOutForFunc(ft reflect.Type){
  // Here is something we can put our original result into
  // For the moment, these are meaningless
  var out []reflect.Value
  // Initialize the output values since it has to be of the same size and type
  for i := 0; i < ft.NumOut(); i++ {
    val := reflect.New(ft.Out(i))

    if t := ft.Out(i); t.Kind().String() != "Ptr" {
      val = reflect.Indirect(val)
    }

    out = append(out, val)
  }
}
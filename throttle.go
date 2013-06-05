package funk

import (
  "sync/atomic"
  "time"
)

// Once is an object that will perform exactly one action.
type Throttle struct {
  waiting int32
  wait    time.Duration
}

func (self *Throttle3) Do(fn func()) {
  // If the swap was not successful, bounce out without execution
  if !atomic.CompareAndSwapInt32(&self.waiting, 0, 1) {
    return
  }

  // After the elapsed wait time, execute the input function and open the lock
  go func() {
    time.AfterFunc(self.wait, func() {
      fn()
      atomic.StoreInt32(&self.waiting, 0)
    })
  }()
}

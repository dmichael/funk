package underscore

import (
  . "launchpad.net/gocheck"
  "testing"
  "time"
)

// Hook up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

// Some ceremony
type MySuite struct {
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpSuite(c *C) {
}

func (s *MySuite) TestThrottle(c *C) {
  // Create the throttled func
  throttled := Throttle{wait: 100 * time.Millisecond}
  count := 0

  for i := 0; i < 1000; i++ {
    throttled.Do(func() { count +=1 })
    throttled.Do(func() { count +=1 })
  }

  time.Sleep(120 * time.Millisecond)
  c.Assert(count, Equals, 1)
}


func (s *MySuite) TestThrottleConcurrent(c *C) {
  // Create the throttled func
  throttled := Throttle{wait: 500 * time.Millisecond}
  count := 0

  for i := 0; i < 100; i++ {
    go throttled.Do(func() { count += 1 })
  }

  // Give the test time to complete
  time.Sleep(510 * time.Millisecond)
  c.Assert(count, Equals, 1)
}
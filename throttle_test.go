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

var count = 0

func MyThrottledFunc(message string){
  count += 1
}

func (s *MySuite) TestThrottle(c *C) {
  f := Throttle(MyThrottledFunc, 1000)
  f("TESTES 0")
  time.Sleep(500 * time.Millisecond)
  f("TESTES 1")
  f("TESTES 1")
  time.Sleep(500 * time.Millisecond)
  f("TESTES 2")
  f("TESTES 2")
  f("TESTES 2")
  f("TESTES 2")
  f("TESTES 2")
  f("TESTES 2")
  time.Sleep(500 * time.Millisecond)
  f("TESTES 3")
  time.Sleep(500 * time.Millisecond)
  f("TESTES 4")

  // Give the test time to complete
  time.Sleep(2 * time.Second)

  // Count should only be incremented by 5
  c.Assert(count, Equals, 5)
}

func ThrottleThis(message string, y int) {
  count2 += 1
}

var count2 = 0

func (s *MySuite) TestThrottle2(c *C) {
  // Create the throttled func
  var throttled func(string, int)
  Throttle2(ThrottleThis, &throttled, 100)

  for i := 0; i < 10; i++ {
    throttled("x", 6)
  }

  time.Sleep(120 * time.Millisecond)

  for i := 0; i < 10; i++ {
    throttled("x", 6)
  }



  // Give the test time to complete
  time.Sleep(120 * time.Millisecond)
  c.Assert(count2, Equals, 2)
}



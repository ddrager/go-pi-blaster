package piblaster

import (
  "os"
  "fmt"
  //"bytes"
  "strconv"
)

type Blaster struct {
  active []int64
  Pins []float64
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func Start(start []int64) {

  var b Blaster 
  const fifo = "/dev/pi-blaster"

  copy(start, b.active)
  b.active = start
  // initialize array used to keep track of pin value on updates
  b.Pins = make([]float64, 26)

  f, err := os.Create(fifo)
  check(err)
  defer f.Close()

  // set all active pin values to 0
  for i := range b.active {
    fmt.Printf("Set pin %d to 0\n", b.active[i])
    f.WriteString(strconv.FormatInt(b.active[i], 10) + "=0\n") 
    f.Sync()
    //check(err)
  }

  f.Close()
}

func Apply(pin int64, value int64) {
  f, err := os.Create("/dev/pi-blaster")
  check(err)
  defer f.Close()

  n1, err := f.WriteString(strconv.FormatInt(pin, 10) + "=" + strconv.FormatInt(value, 10) + "\n")
  fmt.Printf("wrote %d bytes (%d = %d)\n", n1, pin, value)
  f.Sync()

}

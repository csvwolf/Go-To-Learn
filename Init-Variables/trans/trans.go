package trans

import "math"

var Pi float64

const (
  Zero = iota
  First
  Second
)

func init() {
  Pi = 4 * math.Atan(1) // init() funciton computes Pi
}

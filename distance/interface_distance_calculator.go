// Distance package for various algorithms that allow distance estimation
// based on two pairs of coordinate points.

package distance

import "github.com/sebastian-sz/GotwockAppServer/model"

// Interface defining method signatures for various algorithms that allow distance estimation.
// The implemented algorithm should have at least a CalculateDistance method, that accepts a pair of
// coordinates and returns approximated distance as a float.
type Calculator interface {
	CalculateDistance(firstCoordinate, secondCoordinate model.Coordinates) float32
}

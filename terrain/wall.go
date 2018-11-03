package terrain

import (
	tl "github.com/JoelOtter/termloop"
)

// Wall is a termloop Rectangle wrapper for identification of impassible terrain.
type Wall struct {
	*tl.Rectangle
}

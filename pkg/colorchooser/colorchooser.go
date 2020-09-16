// Package colorchooser will return a string colored with a random 16-bit color,
// given the same string again, it will be returned in the same color.
package colorchooser

import (
	"math/rand"
	"sync"
	"time"

	"github.com/logrusorgru/aurora/v3"
)

var defaultInstance = New(
	aurora.RedFg,
	aurora.GreenFg,
	aurora.YellowFg,
	aurora.MagentaFg,
	aurora.CyanFg,
	aurora.BrightFg|aurora.RedFg,
	aurora.BrightFg|aurora.GreenFg,
	aurora.BrightFg|aurora.YellowFg,
	aurora.BrightFg|aurora.BlueFg,
	aurora.BrightFg|aurora.MagentaFg,
	aurora.BrightFg|aurora.CyanFg,
)

func Choose(prefix string) aurora.Color {
	return defaultInstance.Choose(prefix)
}

func Sprint(prefix string) string {
	return defaultInstance.Sprint(prefix)
}

type ColorChooser struct {
	allColors        []aurora.Color
	chosenColors     *sync.Map
	chosenCount      int
	prefixToColorMap *sync.Map
	randGen          *rand.Rand
}

func New(colors ...aurora.Color) *ColorChooser {
	return &ColorChooser{
		randGen:          rand.New(rand.NewSource(time.Now().Unix())),
		chosenColors:     &sync.Map{},
		chosenCount:      0,
		prefixToColorMap: &sync.Map{},
		allColors:        colors,
	}
}

func (c *ColorChooser) Choose(prefix string) aurora.Color {
	// Return a cached value if it exists
	v, exists := c.prefixToColorMap.Load(prefix)
	if exists {
		return v.(aurora.Color)
	}

	// Construct a list of available colors
	availableColors := []aurora.Color{}
	if c.chosenCount >= len(c.allColors) {
		// We reached the maximum number of available colors so
		// we will just have to reuse a color.
		availableColors = c.allColors
	} else {
		// Restrict available color to ones we have not used yet
		for _, v := range c.allColors {
			if _, chosen := c.chosenColors.Load(v); !chosen {
				availableColors = append(availableColors, v)
			}
		}
	}

	// Choose a new color
	var chosen aurora.Color
	if len(availableColors) == 1 {
		chosen = availableColors[0]
	} else {
		chosen = availableColors[c.randGen.Intn(len(availableColors))]
	}

	// Cache the result for next time
	c.chosenCount = c.chosenCount + 1
	c.chosenColors.Store(chosen, true)
	c.prefixToColorMap.Store(prefix, chosen)

	return chosen
}

func (c *ColorChooser) Sprint(prefix string) string {
	return aurora.Colorize(prefix, c.Choose(prefix)).String()
}

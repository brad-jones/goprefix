package colorchooser

import (
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

var defaultInstance = New(
	color.FgRed,
	color.FgGreen,
	color.FgYellow,
	//color.FgBlue, too hard to read on many terminals
	color.FgMagenta,
	color.FgCyan,
	color.FgHiRed,
	color.FgHiGreen,
	color.FgHiYellow,
	color.FgHiBlue,
	color.FgHiMagenta,
	color.FgHiCyan,
)

func Choose(prefix string) color.Attribute {
	return defaultInstance.Choose(prefix)
}

func Sprint(prefix string) string {
	return defaultInstance.Sprint(prefix)
}

type ColorChooser struct {
	allColors        []color.Attribute
	chosenColors     *sync.Map
	chosenCount      int
	prefixToColorMap *sync.Map
	randGen          *rand.Rand
}

func New(colors ...color.Attribute) *ColorChooser {
	return &ColorChooser{
		randGen:          rand.New(rand.NewSource(time.Now().Unix())),
		chosenColors:     &sync.Map{},
		chosenCount:      0,
		prefixToColorMap: &sync.Map{},
		allColors:        colors,
	}
}

func (c *ColorChooser) Choose(prefix string) color.Attribute {
	// Return a cached value if it exists
	v, exists := c.prefixToColorMap.Load(prefix)
	if exists {
		return v.(color.Attribute)
	}

	// Construct a list of avaliable colors
	availableColors := []color.Attribute{}
	if c.chosenCount >= len(c.allColors) {
		// We reached the maximum number of available colors so
		// we will just have to reuse a color.
		availableColors = c.allColors
	} else {
		// Restrict avaliable color to ones we have not used yet
		for _, v := range c.allColors {
			if _, chosen := c.chosenColors.Load(v); !chosen {
				availableColors = append(availableColors, v)
			}
		}
	}

	// Choose a new color
	var choosen color.Attribute
	if len(availableColors) == 1 {
		choosen = availableColors[0]
	} else {
		choosen = availableColors[c.randGen.Intn(len(availableColors))]
	}

	// Cache the result for next time
	c.chosenCount = c.chosenCount + 1
	c.chosenColors.Store(choosen, true)
	c.prefixToColorMap.Store(prefix, choosen)

	return choosen
}

func (c *ColorChooser) Sprint(prefix string) string {
	return color.New(c.Choose(prefix)).Sprint(prefix)
}

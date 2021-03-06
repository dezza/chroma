package chroma

import (
	"testing"

	"github.com/dezza/assert"
)

func TestStyleInherit(t *testing.T) {
	s, err := NewStyle("test", StyleEntries{
		Name:         "bold #f00",
		NameVariable: "#fff",
	})
	assert.NoError(t, err)
	assert.Equal(t, StyleEntry{Colour: 0x1000000, Bold: Yes}, s.Get(NameVariable))
}

func TestStyleColours(t *testing.T) {
	s, err := NewStyle("test", StyleEntries{
		Name: "#f00 bg:#001 border:#ansiblue",
	})
	assert.NoError(t, err)
	assert.Equal(t, StyleEntry{Colour: 0xff0001, Background: 0x000012, Border: 0x000100}, s.Get(Name))
}

func TestStyleClone(t *testing.T) {
	parent, err := NewStyle("test", StyleEntries{
		Background: "bg:#ffffff",
	})
	assert.NoError(t, err)
	clone, err := parent.Builder().Add(Comment, "#0f0").Build()
	assert.NoError(t, err)

	assert.Equal(t, "bg:#ffffff", clone.Get(Background).String())
	assert.Equal(t, "#00ff00 bg:#ffffff", clone.Get(Comment).String())
	assert.Equal(t, "bg:#ffffff", parent.Get(Comment).String())
}

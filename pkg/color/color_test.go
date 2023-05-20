package color_test

import (
	"testing"

	"github.com/PunGrumpy/dokcercolorize/pkg/color"
	"github.com/PunGrumpy/dokcercolorize/pkg/util/assert"
)

//nolint:funlen
func TestColors(t *testing.T) {
	t.Parallel()

	msg := "Message"

	t.Run("Black", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[0;30m"+msg+"\u001B[0m", color.Black(msg))
	})

	t.Run("DarkGray", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[1;30m"+msg+"\u001B[0m", color.DarkGray(msg))
	})

	t.Run("Red", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[0;31m"+msg+"\u001B[0m", color.Red(msg))
	})

	t.Run("LightRed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[1;31m"+msg+"\u001B[0m", color.LightRed(msg))
	})

	t.Run("Green", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[0;32m"+msg+"\u001B[0m", color.Green(msg))
	})

	t.Run("LightGreen", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[1;32m"+msg+"\u001B[0m", color.LightGreen(msg))
	})

	t.Run("Brown", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[0;33m"+msg+"\u001B[0m", color.Brown(msg))
	})

	t.Run("Yellow", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[1;33m"+msg+"\u001B[0m", color.Yellow(msg))
	})

	t.Run("Blue", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[0;34m"+msg+"\u001B[0m", color.Blue(msg))
	})

	t.Run("LightBlue", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[1;34m"+msg+"\u001B[0m", color.LightBlue(msg))
	})

	t.Run("Purple", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[0;35m"+msg+"\u001B[0m", color.Purple(msg))
	})

	t.Run("LightPurple", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[1;35m"+msg+"\u001B[0m", color.LightPurple(msg))
	})

	t.Run("Cyan", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[0;36m"+msg+"\u001B[0m", color.Cyan(msg))
	})

	t.Run("LightCyan", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[1;36m"+msg+"\u001B[0m", color.LightCyan(msg))
	})

	t.Run("LightGray", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[0;37m"+msg+"\u001B[0m", color.LightGray(msg))
	})

	t.Run("White", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\u001B[1;37m"+msg+"\u001B[0m", color.White(msg))
	})
}

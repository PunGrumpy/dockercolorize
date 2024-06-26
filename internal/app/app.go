package app

import (
	"errors"
	"strings"

	"github.com/PunGrumpy/dockercolorize/internal/cmd"
	"github.com/PunGrumpy/dockercolorize/internal/layout"
	"github.com/PunGrumpy/dockercolorize/internal/util"
	"github.com/PunGrumpy/dockercolorize/pkg/color"
)

var (
	ErrNoFirstLine     = errors.New("unable to parse first line")
	ErrNullableColumns = errors.New("unable to parse nullable columns")
)

func Run(in []string) ([]string, error) {
	if len(in) == 0 {
		return nil, ErrNoFirstLine
	}

	header := layout.ParseHeader(in)
	rows := layout.ParseRows(in, &header)

	if header.NullableCols() > 1 {
		return nil, ErrNullableColumns
	}

	command, err := cmd.Parse(header)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	res := make([]string, len(in))

	// First line
	var sb strings.Builder
	for _, col := range header {
		sb.WriteString(util.Pad(color.LightBlue(string(col.Name)), col.MaxLength))
	}

	res[0] = sb.String()

	// Rows
	for i, row := range rows {
		sb.Reset()

		for _, col := range header {
			sb.WriteString(util.Pad(command.Format(row, col.Name), col.MaxLength))
		}

		res[i+1] = sb.String()
	}

	return res, nil
}

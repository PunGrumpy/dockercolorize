package layout

import (
	"strings"

	"github.com/PunGrumpy/dockercolorize/internal/util"
)

func ParseHeader(rows []string) Header {
	parts := util.Split(rows[0])

	res := make(Header, len(parts))
	for i, part := range parts {
		res[i] = &Cell{
			Name:      Column(part),
			MaxLength: len(part),
		}
	}

	return res
}

func ParseRows(rows []string, header *Header) []Row {
	res := make([]Row, len(rows)-1)

	for i, row := range rows[1:] {
		offset := 0
		parts := util.Split(row)
		res[i] = make(Row, len(*header))
		mismatch := len(parts) < len(*header)

		for j, col := range *header {
			if mismatch && col.IsNullable() {
				offset++

				continue
			}

			v := parts[j-offset]

			length := len(v)
			if strings.Contains(v, "…") {
				length -= 2
			}

			if length > col.MaxLength {
				col.MaxLength = length
			}

			res[i][col.Name] = Value(v)
		}
	}

	return res
}

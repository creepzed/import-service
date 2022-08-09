package vo

type Cell struct {
	value string
}

func NewCell(value string) (cell Cell, err error) {
	cell = Cell{value: value}
	if err = cell.hasError(); err != nil {
		return Cell{}, err
	}
	return
}

func (c Cell) hasError() error {
	return nil
}

func (c Cell) Value() string {
	return c.value
}

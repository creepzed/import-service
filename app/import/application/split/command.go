package split

import "bitbucket.org/ripleyx/import-service/app/shared/application/command"

const ImportSplitType command.Type = "import.split"

type ImportSplitCommand struct {
	filename string
}

func NewImportSplitCommand(filename string) ImportSplitCommand {
	return ImportSplitCommand{
		filename: filename,
	}
}

func (c *ImportSplitCommand) Filename() string {
	return c.filename
}

func (c ImportSplitCommand) Type() command.Type {
	return ImportSplitType
}

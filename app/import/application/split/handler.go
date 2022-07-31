package split

import (
	"bitbucket.org/ripleyx/import-service/app/shared/application/command"
	"context"
	"errors"
	"fmt"
)

type SplitCommandHandler struct {
	applicationService SplitService
}

var (
	ErrUnexpectedCommand = errors.New("unexpected command")
)

func NewSplitCommandHandler(applicationService SplitService) SplitCommandHandler {
	return SplitCommandHandler{
		applicationService: applicationService,
	}
}

func (h SplitCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	command, ok := cmd.(ImportSplitCommand)
	if !ok {
		return fmt.Errorf("%w: %s", ErrUnexpectedCommand, cmd.Type())
	}
	return h.applicationService.Do(ctx, command)
}

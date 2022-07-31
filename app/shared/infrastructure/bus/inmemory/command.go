package inmemory

import (
	"bitbucket.org/ripleyx/import-service/app/shared/application/command"
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
	"context"
	"errors"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBusMemory() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (cb *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := cb.handlers[cmd.Type()]
	if !ok {
		return errors.New("error: command not found")
	}
	err := handler.Handle(ctx, cmd)
	if err != nil {
		log.Error("error: while command handling %s - %s", cmd.Type(), err)
		return err
	}
	return nil
}

func (cb CommandBus) Register(cmdType command.Type, handler command.Handler) {
	cb.handlers[cmdType] = handler
}

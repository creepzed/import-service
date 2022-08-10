package subscriber

import (
	"bitbucket.org/ripleyx/import-service/app/import/application/split"
	"bitbucket.org/ripleyx/import-service/app/shared/application/command"
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"time"
)

type message struct {
	EventId     string    `json:"event_id"`
	EventType   string    `json:"event_type"`
	AggregateId string    `json:"aggregate_id"`
	OccurredOn  time.Time `json:"occurred_on"`
	Filename    string    `json:"filename"`
}

type subscriberSplitter struct {
	commandBus command.CommandBus
	dialer     *kafka.Dialer
	groupID    string
	brokers    []string
	topic      string
}

func NewSubscriberSplitter(commandBus command.CommandBus, groupID string, topic string, dialer *kafka.Dialer, brokers ...string) *subscriberSplitter {
	return &subscriberSplitter{
		commandBus: commandBus,
		dialer:     dialer,
		groupID:    groupID,
		brokers:    brokers,
		topic:      topic,
	}
}

func (s *subscriberSplitter) getKafkaReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: s.brokers,
		GroupID: s.groupID,

		Topic:          topic,
		MinBytes:       1e6,  // 1MB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		Dialer:         s.dialer,
	})
}

func (s *subscriberSplitter) ReadMessage(ctx context.Context) {
	reader := s.getKafkaReader(s.topic)
	defer reader.Close()
	for {
		select {
		case <-ctx.Done():
			log.Debug("shutting down subscriber")
			return
		default:
			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				log.WithError(err).Fatal("error reading kafka messages")
				continue
			}
			payload := new(message)
			err = json.Unmarshal(msg.Value, &payload)
			if err != nil {
				log.WithError(err).Fatal("error unmarshalling kafka messages")
				//TODO: agregar Time sleep
				continue
			}

			cmd := split.NewImportSplitCommand(payload.Filename)

			err = s.commandBus.Dispatch(ctx, cmd)
			if err != nil {
				log.WithError(err).Fatal("error saving on import")
				continue
			}
			//TODO: Validar si se guardo?
			err = reader.CommitMessages(ctx, msg)
			if err != nil {
				log.WithError(err).
					Fatal("%s error committing kafka message: %s", s.topic, err.Error())
			}
		}
	}
}

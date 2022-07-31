package subscriber

import (
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"time"
)

type message struct {
	EventId     string    `json:"eventId"`
	EventType   string    `json:"event_type"`
	AggregateId string    `json:"aggregate_id"`
	OccurredOn  time.Time `json:"occurred_on"`
}

type subscriberSplitter struct {
	dialer  *kafka.Dialer
	groupID string
	brokers []string
	topic   string
}

func NewSubscriberSplitter(groupID string, topic string, dialer *kafka.Dialer, brokers ...string) *subscriberSplitter {
	return &subscriberSplitter{
		dialer:  dialer,
		groupID: groupID,
		brokers: brokers,
		topic:   topic,
	}
}

func (s *subscriberSplitter) getKafkaReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        s.brokers,
		GroupID:        s.groupID,
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
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.WithError(err).Error("error reading kafka messages")
				continue
			}
			payload := new(message)
			err = json.Unmarshal(msg.Value, &payload)
			if err != nil {
				log.WithError(err).Error("error unmarshalling kafka messages")
				continue
			}

			//doc := redis.UrlShortenerRedis{
			//	UrlId:       payload.UrlId,
			//	UrlEnable:   payload.UrlStatus,
			//	OriginalUrl: payload.OriginUrl,
			//	UserId:      payload.UserId,
			//}

			//err = s.cache.Set(ctx, doc.UrlId, doc)
			if err != nil {
				log.WithError(err).Error("error saving on cache")
				continue
			}
		}
	}
}

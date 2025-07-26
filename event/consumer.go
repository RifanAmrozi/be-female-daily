package event

import (
	"be-female-daily/storage"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(broker, topic string, store storage.Storage) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		GroupID:   "brand-event-group",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	log.Println("Kafka consumer started...")
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		var evt map[string]interface{}
		if err := json.Unmarshal(m.Value, &evt); err != nil {
			log.Println("Invalid event:", err)
			continue
		}

		if evt["type"] == "brand_view" {
			brandID := int(evt["brand_id"].(float64))

			data, _ := store.GetAllData()
			for i, brand := range data.Brands {
				if brand.ID == brandID {
					data.Brands[i].CrowdCount++
					data.Brands[i].IncreaseCount++
					data.Users[0].Point += 10
					break
				}
			}
			_ = store.SaveAllData(data)
			log.Printf("Updated brand %d from event\n", brandID)
		}
	}
}

package go_amqp_2phase

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	AMQPStatusPending  = 1
	AMQPStatusComplete = 2
)

var (
	AMQPStatuses = map[int]string{
		AMQPStatusPending:  "В процессе отправки",
		AMQPStatusComplete: "Отправлен",
	}
)

type AMQPEvent struct {
	ID         uint `gorm:"primary_key"`
	Status     int
	CreateDate time.Time
	SendDate   *time.Time
	Exchange   string
	Type       string
	Data       string
}

func (AMQPEvent) TableName() string {
	return "amqp_events"
}

func CreateAMQPEvent(db *gorm.DB, exchange string, typ string, data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ev := &AMQPEvent{}
	ev.Exchange = exchange
	ev.Type = typ
	ev.Data = string(jsonData)

	ev.CreateDate = time.Now().In(time.UTC)
	ev.Status = AMQPStatusPending

	return db.Create(ev).Error
}

package delivery

import (
	"SimpleServer/internal/app/providers/Provider"
	v1 "SimpleServer/internal/delivery/v1"
	"time"
)

// В данном файле должна быть структура данного сервиса
// КОНСТРУКТОР
// фукнция запуска сервиса
//

func CreateAndRunUserCenter(address string, defaultExpiration time.Duration, cleanUpInterval time.Duration, endlessLifeTimeAvailability bool, db *Provider.DataBase, promotionInterval time.Duration) *v1.Server {
	return nil

}

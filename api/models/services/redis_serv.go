package services

import "be/models/domains"

type RedisService interface {
	SetData(key string, value interface{}, durationMinute int) error
	GetData(key string) (interface{}, error)
	DeleteData(key string) error
	DeleteAllWithPrefix(prefix string) error
	ClearAllData() error
	GeoAdd(
		key string,
		longitude float64,
		latitude float64,
		member string,
	) error
	GeoRadius(
		key string,
		longitude float64,
		latitude float64,
		radiusKm float64,
		limit int,
	) ([]string, error)
	GeoRadiusWithDist(
		key string,
		longitude float64,
		latitude float64,
		radiusKm float64,
		limit int,
	) ([]domains.GeoResult, error)
}

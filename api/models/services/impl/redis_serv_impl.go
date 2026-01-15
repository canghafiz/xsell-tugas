package impl

import (
	"be/models/domains"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisServiceImpl struct {
	rdb    *redis.Client
	config domains.RedisConfig
}

func NewRedisServiceImpl(config domains.RedisConfig) *RedisServiceImpl {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	return &RedisServiceImpl{
		rdb:    rdb,
		config: config,
	}
}

/* ===================== BASIC CACHE ===================== */

func (serv *RedisServiceImpl) SetData(
	key string,
	value interface{},
	durationMinute int,
) error {

	ctx := context.Background()
	expiration := time.Duration(durationMinute) * time.Minute

	return serv.rdb.Set(
		ctx,
		serv.config.Prefix+key,
		value,
		expiration,
	).Err()
}

func (serv *RedisServiceImpl) GetData(key string) (interface{}, error) {
	ctx := context.Background()

	val, err := serv.rdb.Get(ctx, serv.config.Prefix+key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("redis key not found: %s", key)
	}
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (serv *RedisServiceImpl) DeleteData(key string) error {
	ctx := context.Background()
	return serv.rdb.Del(ctx, serv.config.Prefix+key).Err()
}

func (serv *RedisServiceImpl) DeleteAllWithPrefix(prefix string) error {
	ctx := context.Background()

	var cursor uint64
	for {
		keys, nextCursor, err := serv.rdb.Scan(
			ctx,
			cursor,
			serv.config.Prefix+prefix+"*",
			100,
		).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			if err := serv.rdb.Del(ctx, keys...).Err(); err != nil {
				log.Printf("warning delete prefix %s: %v", prefix, err)
			}
		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}

func (serv *RedisServiceImpl) ClearAllData() error {
	ctx := context.Background()
	return serv.rdb.FlushDB(ctx).Err()
}

/* ===================== GEO LOCATION ===================== */

func (serv *RedisServiceImpl) GeoAdd(
	key string,
	longitude float64,
	latitude float64,
	member string,
) error {

	ctx := context.Background()

	return serv.rdb.GeoAdd(
		ctx,
		serv.config.Prefix+key,
		&redis.GeoLocation{
			Name:      member,
			Longitude: longitude,
			Latitude:  latitude,
		},
	).Err()
}

func (serv *RedisServiceImpl) GeoRadius(
	key string,
	longitude float64,
	latitude float64,
	radiusKm float64,
	limit int,
) ([]string, error) {

	ctx := context.Background()

	results, err := serv.rdb.GeoRadius(
		ctx,
		serv.config.Prefix+key,
		longitude,
		latitude,
		&redis.GeoRadiusQuery{
			Radius: radiusKm,
			Unit:   "km",
			Sort:   "ASC", // ✅ nearest first
			Count:  limit,
		},
	).Result()

	if err != nil {
		return nil, err
	}

	members := make([]string, 0, len(results))
	for _, r := range results {
		members = append(members, r.Name)
	}

	return members, nil
}

func (serv *RedisServiceImpl) GeoRadiusWithDist(
	key string,
	longitude float64,
	latitude float64,
	radiusKm float64,
	limit int,
) ([]domains.GeoResult, error) {

	ctx := context.Background()

	results, err := serv.rdb.GeoRadius(
		ctx,
		serv.config.Prefix+key,
		longitude,
		latitude,
		&redis.GeoRadiusQuery{
			Radius:   radiusKm,
			Unit:     "km",
			Sort:     "ASC", // nearest first
			Count:    limit,
			WithDist: true,
		},
	).Result()

	if err != nil {
		return nil, err
	}

	out := make([]domains.GeoResult, 0, len(results))
	for _, r := range results {
		out = append(out, domains.GeoResult{
			Member:   r.Name,
			Distance: r.Dist,
		})
	}

	return out, nil
}

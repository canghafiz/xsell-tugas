package helpers

import (
	"encoding/json"
	"fmt"
)

type GetData func(key string) (interface{}, error)
type SetData func(key string, value interface{}, durationMinute int) error

func GetFromCache(dataGetter GetData, cacheKey string, dest interface{}) (interface{}, error) {
	// 1. Call Redis
	cachedData, err := dataGetter(cacheKey)
	if err != nil {
		LogError("Error retrieving data from cache for key %s: %v", cacheKey, err)
		return nil, err
	}

	// 2. Doing type assertion to make data string, and check is empty.
	data, ok := cachedData.(string)
	if !ok || data == "" {
		return nil, fmt.Errorf("cached data is empty or not a string")
	}

	// 3. Unmarshal string JSON to destination struct.
	if err := json.Unmarshal([]byte(data), dest); err != nil {
		LogError("Error unmarshalling cached data for key %s: %v", cacheKey, err)
		return nil, fmt.Errorf("failed to unmarshal cached data: %w", err)
	}

	LogSuccess("Successfully retrieved and unmarshalled data for key %s.", cacheKey)

	// 4. Return Data
	return dest, nil
}

func SetToCache(dataSetter SetData, cacheKey string, data interface{}, durationMinute int) error {
	// Marshal data to JSON.
	dataToCache, err := json.Marshal(data)
	if err != nil {
		LogError("Failed to marshal data for cache: %v", err)
		return fmt.Errorf("failed to marshal data for cache: %w", err)
	}

	// Set the marshaled data in Redis.
	err = dataSetter(cacheKey, string(dataToCache), durationMinute)
	if err != nil {
		LogError("Failed to set cache for key %s: %v", cacheKey, err)
		return fmt.Errorf("failed to set cache for key %s: %w", cacheKey, err)
	}

	LogSuccess("Successfully set data to cache for key %s.", cacheKey)
	return nil
}

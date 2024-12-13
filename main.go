package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Define the Redis Sentinel addresses and master name
	sentinelAddrs := []string{"redis-sentinel-sentinel:26379"}
	masterName := "myMaster"

	// Create the failover cluster client
	client := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:      masterName,
		SentinelAddrs:   sentinelAddrs,
		RouteByLatency:  false,                  // Optional: Routes by latency to replicas
		RouteRandomly:   true,                   // Optional: Routes randomly to replicas for readonly commands
		MaxRetries:      10,                     // Retry up to 10 times
		MinRetryBackoff: 500 * time.Millisecond, // Minimum retry delay
		MaxRetryBackoff: 2 * time.Second,        // Maximum retry delay
	})

	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("Failed to close client: %v\n", err)
		}
	}()

	ctx := context.Background()

	// Ensure the connection is established
	if err := client.Ping(ctx).Err(); err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return
	}

	fmt.Println("Connected to Redis.")
	firstKey := "key_first_timestamp"

	timestampKey := fmt.Sprintf("key_%d", time.Now().UnixNano())
	value := "example_value"
	if err := client.Set(ctx, firstKey, value, 0).Err(); err != nil {
		fmt.Printf("Failed to set key: %v\n", err)
		return
	}
	fmt.Printf("Initial write of key '%s' set successfully.\n", timestampKey)

	var wg sync.WaitGroup

	// Start the write operation in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// Write a unique key based on timestamp
			timestampKey := fmt.Sprintf("key_%d", time.Now().UnixNano())
			value := "example_value"
			if err := client.Set(ctx, timestampKey, value, 0).Err(); err != nil {
				fmt.Printf("Failed to set key: %v\n", err)
				return
			}
			fmt.Printf("THIS IS WRITE OPERATION: Key '%s' set successfully.\n", timestampKey)
			time.Sleep(1 * time.Second) // Wait before the next write
		}
	}()

	// Start the read operation in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			readonlyValue, err := client.Get(ctx, firstKey).Result()
			if err != nil {
				fmt.Printf("Failed to get value for first timestamp key '%s': %v\n", firstKey, err)
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Printf("THIS IS READ ONLY OPERATION: Key '%s' has value '%s'.\n", firstKey, readonlyValue)
			time.Sleep(1 * time.Second) // Wait before the next read
		}
	}()

	wg.Wait()
}

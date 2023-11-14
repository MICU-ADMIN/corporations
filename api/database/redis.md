```go

package database

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func connect() *redis.Client {
	opt, err := redis.ParseURL("redis://default:2bd4464bf59f4e00bb3106965089250a@eu1-touching-garfish-38487.upstash.io:38487")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	return client
}

func SetRedis(key, value string, expiration time.Duration) error {
	client := connect()
	defer client.Close()

	return client.Set(ctx, key, value, expiration).Err()
}

func GetRedisHTTP(key string) (string, error) {
    url := "https://eu1-touching-garfish-38487.upstash.io/get/" + key
    bearerToken := "ApZXACQgMDdjZWM4NWEtMGVhZi00NGE1LWFkOWYtNWZjOThiMzEyY2Nh7WdzLBzCXMr0IbOwHeCcupCNK6NroXD_dkA5d_usRRA="

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Authorization", "Bearer "+bearerToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("error making request: %v", err)
    }
    defer resp.Body.Close()

    fmt.Println("Status Code:", resp.Status)

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response body: %v", err)
    }

    // Convert the response body to a string and return
    return string(body), nil
}

func SetRedisHTTP(key, value string) error {
    url := "https://eu1-touching-garfish-38487.upstash.io/set/" + key + "/" + value
    bearerToken := "AZZXACQgMDdjZWM4NWEtMGVhZi00NGE1LWFkOWYtNWZjOThiMzEyY2NhMmJkNDQ2NGJmNTlmNGUwMGJiMzEwNjk2NTA4OTI1MGE="

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Authorization", "Bearer "+bearerToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("error making request: %v", err)
    }
    defer resp.Body.Close()

    fmt.Println("Status Code:", resp.Status)

    // Check the response status code to ensure the set operation was successful
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("set operation failed with status code: %s", resp.Status)
    }

    return nil
}


func SetAndGetRedis(ctx context.Context, key, value string, expiration time.Duration) (string, error) {
    client := connect()
    defer client.Close()

    err := client.Set(ctx, key, value, expiration).Err()
    if err != nil {
        return "", err
    }

    // Retrieve the value using the Get method
    storedValue, err := client.Get(ctx, key).Result()
    if err != nil {
        return "", err
    }

    return storedValue, nil
}


func GetRedis(key string) (string, error) {
	client := connect()
	defer client.Close()

	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key %s does not exist", key)
	} else if err != nil {
		return "", err
	}

	return val, nil
}

// func main() {
// 	if err := set("foo", "bar", 0); err != nil {
// 		fmt.Println("Error setting value:", err)
// 		return
// 	}

// 	val, err := get("foo")
// 	if err != nil {
// 		fmt.Println("Error getting value:", err)
// 		return
// 	}

// 	fmt.Println(val)
// }


```

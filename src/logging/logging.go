package logging

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis"
)

type I_Logging interface {
    LogError(string, ...interface{})
    LogData(string, ...interface{})
    LogTrade(string, ...interface{})
    LogDataPrice(string, ...interface{})
}

type Logging struct {
	mu   sync.Mutex
}

func GetTraded() string {
    rdb := redis.NewClient(&redis.Options{
            Addr:     "localhost:6379",
            Password: "", // no password set
            DB:       0,  // use default DB
        })

    val, err := rdb.Get("traded").Result()

    if err != nil {
        fmt.Printf("[ERROR] %s",err)
    }
    
    return val 
}

func SetTraded(value string) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set("traded", value, 0).Err()
    if err != nil {
        fmt.Printf("[ERROR] %s", err)
    }

}
func (lg *Logging) LogError(text string, a ...interface{}) {
	lg.mu.Lock()
	defer lg.mu.Unlock()
    output := fmt.Sprintf(text, a...)
	// Ouvrir le fichier de log en mode ajout
	file, err := os.OpenFile("log/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime)

	log.Println(output)
	defer file.Close()
}

func (lg *Logging) LogData(text string, a ...interface{}) {
    output := fmt.Sprintf(text, a...)
	// Ouvrir le fichier de log en mode ajout
	file, err := os.OpenFile("log/data.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime)

	log.Println(output)
	defer file.Close()
}

func (lg *Logging) LogTrade(text string, a ...interface{}) {
    output := fmt.Sprintf(text, a...)
	// Ouvrir le fichier de log en mode ajout
	file, err := os.OpenFile("log/trades.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime)

	log.Println(output)
	defer file.Close()
}

func (lg *Logging) LogDataPrice(text string, a ...interface{}) {
    output := fmt.Sprintf(text, a...)
	// Ouvrir le fichier de log en mode ajout
	file, err := os.OpenFile("log/price.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime)

	log.Println(output)
	defer file.Close()
}

func LogError(text string, a ...interface{}) {
    output := fmt.Sprintf(text, a...)
	// Ouvrir le fichier de log en mode ajout
	file, err := os.OpenFile("log/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime)

	log.Println(output)
	defer file.Close()
}

func LogData(text string, a ...interface{}) {
    output := fmt.Sprintf(text, a...)
	// Ouvrir le fichier de log en mode ajout
	file, err := os.OpenFile("log/data.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime)

	log.Println(output)
	defer file.Close()
}

func LogTrade(text string, a ...interface{}) {
    output := fmt.Sprintf(text, a...)
	// Ouvrir le fichier de log en mode ajout
	file, err := os.OpenFile("log/trades.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime)

	log.Println(output)
	defer file.Close()
}

func LogDataPrice(text string, a ...interface{}) {
    output := fmt.Sprintf(text, a...)
	// Ouvrir le fichier de log en mode ajout
	file, err := os.OpenFile("log/price.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime)

	log.Println(output)
	defer file.Close()
}

package account

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"gopkg.in/yaml.v3"
)

type ApiKey struct {
    ApiKey string `yaml:"apiKey"`
    SecretKey string `yaml:"secretKey"`
} 

type Symbols struct {
    Symbols []string `yaml:"symbols"`
}

func GetClient() *binance.Client {
    yamlFile, err := os.Open("config/apikey.yaml")

    if err != nil {
		log.Fatalf("Impossible d'ouvrir le fichier YAML : %v", err)
    }

    defer yamlFile.Close()

	// Lecture du fichier YAML
	api, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		log.Fatalf("Impossible de lire le fichier YAML : %v", err)
	}


    ap := ApiKey{}
    err_yaml := yaml.Unmarshal(api, &ap)

    if err_yaml != nil {
        fmt.Println(err)
    }

    return binance.NewClient(ap.ApiKey,ap.SecretKey)
}


func GetSymbols() []string {
    // Chemin du fichier YAML
	filePath := "config/symbols.yaml"

	// Ouverture du fichier YAML
	yamlFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir le fichier YAML : %v", err)
	}
	defer yamlFile.Close()

	// Lecture du fichier YAML
	yamlData, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		log.Fatalf("Impossible de lire le fichier YAML : %v", err)
	}

	// Structure pour stocker les données YAML
    symbols := Symbols{}

	// Décodage du fichier YAML
	err = yaml.Unmarshal(yamlData, &symbols)
	if err != nil {
		log.Fatalf("Impossible de décoder le fichier YAML : %v", err)
	}

    return symbols.Symbols
}

func GetBalance(client *binance.Client, asset string) (float64, error) {
    res, err := client.NewGetAccountService().Do(context.Background())

    if err != nil {
        fmt.Println(err)
        return 0, err
    }

    for _, i := range res.Balances {
        if i.Asset == asset  {
            value, _ := strconv.ParseFloat(i.Free, 64)
            return value, nil
        }
    } 
    return 0, errors.New("Error: not asset found")
}

func GetBalanceString(client *binance.Client, asset string) (string, error) {
    res, err := client.NewGetAccountService().Do(context.Background())

    if err != nil {
        fmt.Println(err)
        return "", err
    }

    for _, i := range res.Balances {
        if i.Asset == asset  {
            return i.Free, nil
        }
    } 
    return "", errors.New("Error: not asset found")
}


func GetStepSymbol(client *binance.Client, asset string) int {
    res, err := client.NewExchangeInfoService().Do(context.Background())
    
    if err != nil {
        fmt.Println(err)
    }

    var stepsize string
    for _, i := range res.Symbols {
        if i.Symbol == asset  {
            stepsize = i.LotSizeFilter().StepSize
        }
    }

    rBool, _ := regexp.Compile("1\\.[0]+")
    r, _ := regexp.Compile("(0*)1")


    if !rBool.MatchString(stepsize) {
        lenStep := r.FindString(stepsize)
        return len(lenStep)
    } else {
        return 0 
    }
}

func GetAllSymbol(client *binance.Client, symbol string) []string {
    res, err := client.NewExchangeInfoService().Do(context.Background())

    if err != nil {
        fmt.Println(err)
    }

    var allSymbol []string

    r, _ := regexp.Compile(symbol)
    rup, _ := regexp.Compile("UP")
    rbull, _ := regexp.Compile("BULL")
    rdown, _ := regexp.Compile("DOWN")
    rbear, _ := regexp.Compile("BEAR")
    for _, i := range res.Symbols {
        if r.MatchString(i.Symbol) {
            if !rup.MatchString(i.Symbol) && !rdown.MatchString(i.Symbol) && !rbull.MatchString(i.Symbol) && !rbear.MatchString(i.Symbol) {
                allSymbol = append(allSymbol, i.Symbol)
            }
        }
    } 
    return allSymbol
}

func GetStepUSDTSymbol(client *binance.Client, asset string) int {
    res, err := client.NewExchangeInfoService().Do(context.Background())
    
    if err != nil {
        fmt.Println(err)
    }

    var stepsize string
    for _, i := range res.Symbols {
        if i.Symbol == asset  {
            stepsize = i.PriceFilter().MinPrice
        }
    }

    rBool, _ := regexp.Compile("1\\.[0]+")
    r, _ := regexp.Compile("(0*)1")


    if !rBool.MatchString(stepsize) {
        lenStep := r.FindString(stepsize)
        return len(lenStep)
    } else {
        return 0 
    }
}

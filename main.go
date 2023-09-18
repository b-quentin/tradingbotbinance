package main

import (
	"context"
	"fmt"
	"tradingBot/src/account"

)

func main() {
    client := account.GetClient()

    res, err := client.NewGetAccountService().Do(context.Background())

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(res)

    //strategy.Launch(client)
}

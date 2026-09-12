package main

import (
	"context"
	"fmt"

	"github.com/abhinaaaavvv/ilovego/application"
)

func main() {
	app := application.New()

	if err := app.Start(context.TODO()); err != nil {
		fmt.Println("Failed to start app:", err)
	}
}

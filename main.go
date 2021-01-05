/*
Copyright © 2020 John Arroyo
@author John Arroyo, john.arroyo@gmail.com

Search service server
*/
package main

// import "github.com/arroyo/serverless-search/cmd"
import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	fmt.Printf("User entered %v\n", name.Name)
	return fmt.Sprintf(`{"message": "Hello %s!"}`, name.Name ), nil
}

func main() {
	lambda.Start(HandleRequest)
}

/*
func main() {
	cmd.Execute()
}
*/

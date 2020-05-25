package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	drv "github.com/uber/athenadriver/go"
)

//State - Struct to store our query results
type State struct {
	State string `json:"state"`
	Count int64  `json:"count"`
}

//Handler - Is called after main
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Get Parameter from Query String - Return if empty
	state := request.QueryStringParameters["state"]
	if state == "" {
		return events.APIGatewayProxyResponse{Body: "No State Provided", StatusCode: 400}, nil
	}

	//Set Region
	S3Region := os.Getenv("REGION")
	S3Bucket := os.Getenv("BUCKET")
	DB := os.Getenv("DB")
	TABLE := os.Getenv("TABLE")

	//Start Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(S3Region),
	})
	if err != nil {
		fmt.Println(err)
	}

	//Store Session Credentials
	c, err := sess.Config.Credentials.Get()
	if err != nil {
		fmt.Println(err)
	}

	//Set AWS Credential in Driver Config.
	conf, err := drv.NewDefaultConfig(S3Bucket,
		S3Region,
		c.AccessKeyID,
		c.SecretAccessKey,
	)
	if err != nil {
		fmt.Println(err)
	}

	// Open Connection.
	db, err := sql.Open(drv.DriverName, conf.Stringify())
	if err != nil {
		fmt.Println(err)
	}

	//Store Results
	var results State

	// Build Query
	query := fmt.Sprintf(`
	SELECT State,
	Count(*) as "Count"
	
    FROM "%v"."%v"

	WHERE State = '%v'

	GROUP BY State`, DB, TABLE, state)

	//Run Query
	err = db.QueryRow(query).Scan(&results.State, &results.Count)
	if err != nil {
		fmt.Println(err)
	}

	//Marshal JSON result
	r, _ := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
	}

	//Return
	return events.APIGatewayProxyResponse{Body: string(r), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}

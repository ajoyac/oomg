package oomg

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	client     *mongo.Client
	ctxTimeout time.Duration
	currentDB  *mongo.Database
)

func init() {
	ctxTimeout = 10 * time.Second
}

// Most likely to change to work
// as mongo have many many other options to connect
type ConnectOptions struct {
	URL        string
	CtxTimeout time.Duration
	DbName string
}

// newCtx function create and return new context with your specified timeout.
func newCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ctxTimeout)
}

func setCtxTimeOut(d time.Duration) {
	if d > 0 {
		ctxTimeout = d
	}
}

func Connect(co ConnectOptions) (err error) {

	log.Info("Connect to Mongo")

	clientOptions := options.Client().ApplyURI(co.URL)
	
	setCtxTimeOut(co.CtxTimeout)

	ctx,cancel := newCtx()
	
	defer cancel()

	// Connect to MongoDB
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error("Mongo Connection Failed")
        return
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error("Failed to check Connection")
		return 
	}
	// wonder if should I check if exist... not even the driver check it
	currentDB = client.Database(co.DbName)
	
	return
	
}

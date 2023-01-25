package config

import (
	"os"
	"time"
)

var (
	StanUrl           = "nats://" + os.Getenv("STAN_HOST_URL") + ":4222"
	StanCluster       = os.Getenv("STAN_CLUSTER")
	StanClient        = os.Getenv("STAN_CLIENT")
	StanSubj          = os.Getenv("STAN_SUBJECT")
	StanDurableName   = os.Getenv("STAN_DURABLE_NAME")
	ReconnectInterval = 12 * time.Second
	ServPort          = ":" + os.Getenv("SERVER_PORT")
)

var (
	Host   = os.Getenv("POSTGRES_HOST")
	Port   = os.Getenv("POSTGRES_PORT")
	User   = os.Getenv("POSTGRES_USER")
	Pass   = os.Getenv("POSTGRES_PASSWORD")
	Dbname = os.Getenv("POSTGRES_DB")
)

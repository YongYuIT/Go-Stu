package main

import (
	"./viperutil"
	"fmt"
	"github.com/Shopify/sarama"
	bccsp "github.com/hyperledger/fabric/bccsp/factory"
	"github.com/spf13/viper"
	"strings"
	"time"
)

// TopLevel directly corresponds to the orderer config YAML.
// Note, for non 1-1 mappings, you may append
// something like `mapstructure:"weirdFoRMat"` to
// modify the default mapping, see the "Unmarshal"
// section of https://github.com/spf13/viper for more info.
type TopLevel struct {
	General    General
	FileLedger FileLedger
	RAMLedger  RAMLedger
	Kafka      Kafka
	Debug      Debug
	Consensus  interface{}
	Operations Operations
	Metrics    Metrics
}

// General contains config which should be common among all orderer types.
type General struct {
	LedgerType     string
	ListenAddress  string
	ListenPort     uint16
	TLS            TLS
	Cluster        Cluster
	Keepalive      Keepalive
	GenesisMethod  string
	GenesisProfile string
	SystemChannel  string
	GenesisFile    string
	Profile        Profile
	LocalMSPDir    string
	LocalMSPID     string
	BCCSP          *bccsp.FactoryOpts
	Authentication Authentication
}

type Cluster struct {
	ListenAddress                        string
	ListenPort                           uint16
	ServerCertificate                    string
	ServerPrivateKey                     string
	ClientCertificate                    string
	ClientPrivateKey                     string
	RootCAs                              []string
	DialTimeout                          time.Duration
	RPCTimeout                           time.Duration
	ReplicationBufferSize                int
	ReplicationPullTimeout               time.Duration
	ReplicationRetryTimeout              time.Duration
	ReplicationBackgroundRefreshInterval time.Duration
	ReplicationMaxRetries                int
	SendBufferSize                       int
	CertExpirationWarningThreshold       time.Duration
}

// Keepalive contains configuration for gRPC servers.
type Keepalive struct {
	ServerMinInterval time.Duration
	ServerInterval    time.Duration
	ServerTimeout     time.Duration
}

// TLS contains configuration for TLS connections.
type TLS struct {
	Enabled            bool
	PrivateKey         string
	Certificate        string
	RootCAs            []string
	ClientAuthRequired bool
	ClientRootCAs      []string
}

// SASLPlain contains configuration for SASL/PLAIN authentication
type SASLPlain struct {
	Enabled  bool
	User     string
	Password string
}

// Authentication contains configuration parameters related to authenticating
// client messages.
type Authentication struct {
	TimeWindow time.Duration
}

// Profile contains configuration for Go pprof profiling.
type Profile struct {
	Enabled bool
	Address string
}

// FileLedger contains configuration for the file-based ledger.
type FileLedger struct {
	Location string
	Prefix   string
}

// RAMLedger contains configuration for the RAM ledger.
type RAMLedger struct {
	HistorySize uint
}

// Kafka contains configuration for the Kafka-based orderer.
type Kafka struct {
	Retry     Retry
	Verbose   bool
	Version   sarama.KafkaVersion // TODO Move this to global config
	TLS       TLS
	SASLPlain SASLPlain
	Topic     Topic
}

// Retry contains configuration related to retries and timeouts when the
// connection to the Kafka cluster cannot be established, or when Metadata
// requests needs to be repeated (because the cluster is in the middle of a
// leader election).
type Retry struct {
	ShortInterval   time.Duration
	ShortTotal      time.Duration
	LongInterval    time.Duration
	LongTotal       time.Duration
	NetworkTimeouts NetworkTimeouts
	Metadata        Metadata
	Producer        Producer
	Consumer        Consumer
}

// NetworkTimeouts contains the socket timeouts for network requests to the
// Kafka cluster.
type NetworkTimeouts struct {
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Metadata contains configuration for the metadata requests to the Kafka
// cluster.
type Metadata struct {
	RetryMax     int
	RetryBackoff time.Duration
}

// Producer contains configuration for the producer's retries when failing to
// post a message to a Kafka partition.
type Producer struct {
	RetryMax     int
	RetryBackoff time.Duration
}

// Consumer contains configuration for the consumer's retries when failing to
// read from a Kafa partition.
type Consumer struct {
	RetryBackoff time.Duration
}

// Topic contains the settings to use when creating Kafka topics
type Topic struct {
	ReplicationFactor int16
}

// Debug contains configuration for the orderer's debug parameters.
type Debug struct {
	BroadcastTraceDir string
	DeliverTraceDir   string
}

// Operations configures the operations endpont for the orderer.
type Operations struct {
	ListenAddress string
	TLS           TLS
}

// Operations confiures the metrics provider for the orderer.
type Metrics struct {
	Provider string
	Statsd   Statsd
}

// Statsd provides the configuration required to emit statsd metrics from the orderer.
type Statsd struct {
	Network       string
	Address       string
	WriteInterval time.Duration
	Prefix        string
}

func main() {
	config := viper.New()
	config.AddConfigPath("/home/yong/Go-Stu20191008001/ReadFabricSourceCode/stu_orderer_load_InitViper/tmp_fabric/")
	config.SetConfigName("orderer")
	config.SetEnvPrefix("ORDERER")
	config.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)

	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configuration: %s", err)
	}
	fmt.Println(config.ConfigFileUsed())
	///////////////////////////////////////////////////////////////////////////////////////
	var uconf TopLevel
	if err := viperutil.EnhancedExactUnmarshal(config, &uconf); err != nil {
		fmt.Printf("Error unmarshaling config into struct: %s", err)
	}
	fmt.Println("EnhancedExactUnmarshal success")
}

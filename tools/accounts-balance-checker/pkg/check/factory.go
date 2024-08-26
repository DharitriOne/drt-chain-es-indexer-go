package check

import (
	"math"
	"time"

	"github.com/DharitriOne/drt-chain-core-go/core/pubkeyConverter"
	"github.com/DharitriOne/drt-chain-es-indexer-go/client/logging"
	"github.com/DharitriOne/drt-chain-es-indexer-go/process/elasticproc/converters"
	"github.com/DharitriOne/drt-chain-es-indexer-go/tools/accounts-balance-checker/pkg/config"
	"github.com/DharitriOne/drt-chain-es-indexer-go/tools/accounts-balance-checker/pkg/esclient"
	"github.com/DharitriOne/drt-chain-es-indexer-go/tools/accounts-balance-checker/pkg/rest"
	"github.com/elastic/go-elasticsearch/v7"
)

// CreateBalanceChecker will create a new instance of balanceChecker
func CreateBalanceChecker(cfg *config.Config, repair bool) (*balanceChecker, error) {
	esClient, err := esclient.NewElasticClient(elasticsearch.Config{
		Addresses: []string{cfg.Elasticsearch.URL},
		Username:  cfg.Elasticsearch.Username,
		Password:  cfg.Elasticsearch.Password,
		Logger:    &logging.CustomLogger{},
		RetryBackoff: func(i int) time.Duration {
			// A simple exponential delay
			d := time.Duration(math.Exp2(float64(i))) * time.Second
			log.Info("elastic: retry backoff", "attempt", i, "sleep duration", d)
			return d
		},
		MaxRetries:    5,
		RetryOnStatus: []int{429, 502, 503, 504},
	})
	if err != nil {
		return nil, err
	}

	restClient, err := rest.NewRestClient(cfg.Proxy.URL)
	if err != nil {
		return nil, err
	}

	pubKeyConverter, err := pubkeyConverter.NewBech32PubkeyConverter(32, log)
	if err != nil {
		return nil, err
	}

	balanceToFloat, err := converters.NewBalanceConverter(18)
	if err != nil {
		return nil, err
	}

	return NewBalanceChecker(esClient, restClient, pubKeyConverter, balanceToFloat, repair, cfg.Proxy.MaxNumberOfParallelRequests)
}

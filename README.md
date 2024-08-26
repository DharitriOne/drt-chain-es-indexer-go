## drt-chain-es-indexer-go


Is a versatile component designed to enhance the data management capabilities of the `drt-chain-go` repository. 

### Overview

This module's flexibility enables it to be used in two distinct modes: as a Go module or as a separate microservice.

### Features

- **Data Transformation**: `drt-chain-es-indexer-go` excels at transforming raw data generated by the `drt-chain-go` repository into a structured format suitable for indexing and storage.

- **Indexing**: The module indexes the transformed data within Elasticsearch, enhancing query performance and data retrieval.

- **Microservice Mode**: In microservice mode, drt-chain-es-indexer-go operates as a standalone microservice that communicates with drt-chain-go via WebSocket connections. This architecture promotes modularity and scalability.

- **Go Module Mode**: The `drt-chain-go` repository can include the module directly. This makes it easy to manage data without requiring a separate microservice.

### Running as a Separate Microservice

Running the `drt-chain-es-indexer-go` module as a **microservice** allows you to efficiently manage data indexing and storage while maintaining 
modularity and scalability. This guide outlines the steps to deploy the module as a separate microservice that communicates with 
the `drt-chain-go` repository over `WebSocket` connections.

### Monitoring Endpoints

The `drt-chain-es-indexer-go` microservice provides monitoring endpoints that allow you to keep track of its health, performance, and other vital statistics.
These endpoints are essential for maintaining the stability and efficiency of the microservice.

#### Monitoring Endpoints

`/status/metrics`

This endpoint exposes various metrics about the microservice's internal operations. These metrics are formatted in a way that is suitable for consumption
by monitoring and alerting systems.

HTTP Method: **GET**

Response: Metrics are presented in JSON format for easy integration with monitoring and alerting systems.

`/status/prometheus-metrics`

This endpoint provides Prometheus-compatible metrics in a specific format for easy integration with 
Prometheus monitoring systems.

HTTP Method: **GET**

Response: Metrics are formatted in a way that Prometheus can scrape and ingest for monitoring and alerting purposes.



### Prerequisites
Before proceeding, ensure you have the following prerequisites:
- Go programming environment set up.
- Access to an Elasticsearch database instance.
- One has to setup one or multiple observers. For running an observing squad, these [docs](https://docs.DharitriOne.com/integrators/observing-squad/) cover the whole process.
The required configs for launching an observer/s with a driver attached, can be found [here](https://github.com/DharitriOne/drt-chain-go/blob/master/cmd/node/config/external.toml).

The corresponding config section for enabling the driver:

```toml
[[HostDriversConfig]]
    # This flag shall only be used for observer nodes
    Enabled = true
    # This flag will start the WebSocket connector as server or client (can be "client" or "server")
    Mode = "client"
    # URL for the WebSocket client/server connection
    # This value represents the IP address and port number that the WebSocket client or server will use to establish a connection.
    URL = "127.0.0.1:22111"
    # After a message will be sent it will wait for an ack message if this flag is enabled
    WithAcknowledge = true
    # The duration in seconds to wait for an acknowledgment message, after this time passes an error will be returned
    AcknowledgeTimeoutInSec = 60
    # Possible values: json, gogo protobuf. Should be compatible with drt-chain-es-indexer-go config
    MarshallerType = "json"
    # The number of seconds when the client will try again to send the data
    RetryDurationInSec = 5
    # Sets if, in case of data payload processing error, we should block or not the advancement to the next processing event. Set this to true if you wish the node to stop processing blocks if the client/server encounters errors while processing requests.
    BlockingAckOnError = true
    # Set to true to drop messages if there is no active WebSocket connection to send to.
    DropMessagesIfNoConnection = false
```


#### Install
Using the `cmd/elasticindexer` package as root, execute the following commands:
- install go dependencies: `go install`
- build executable: `go build -o elasticindexer`

#### Launching the elasticindexer

CLI: run `--help` to get the command line parameters
```
./elasticindexer --help
```

Before launching the `elasticindexer` service, it has to be configured so that it runs with the correct configuration.

The **_[prefs.toml](./cmd/elasticindexer/config/prefs.toml)_** file:

```toml
[config]
    disabled-indices = []
    [config.web-socket]
        # URL for the WebSocket client/server connection
        # This value represents the IP address and port number that the WebSocket client or server will use to establish a connection.
        url = "localhost:22111"
        # This flag describes the mode to start the WebSocket connector. Can be "client" or "server"
        mode = "server"
        # Possible values: json, gogo protobuf. Should be compatible with drt-chain-node outport driver config
        data-marshaller-type = "json"
        # Retry duration (receive/send ack signal) in seconds
        retry-duration-in-seconds = 5
        # Signals if in case of data payload processing error, we should send the ack signal or not
        blocking-ack-on-error = true
        # After a message will be sent it will wait for an ack message if this flag is enabled
        with-acknowledge = true
        # The duration in seconds to wait for an acknowledgment message, after this time passes an error will be returned
        acknowledge-timeout-in-seconds = 50
    
    [config.elastic-cluster]
        use-kibana = false
        url = "http://localhost:9200"
        username = ""
        password = ""
        bulk-request-max-size-in-bytes = 4194304 # 4MB
```

The _**[api.toml](./cmd/elasticindexer/config/api.toml)**_ file:
```toml
rest-api-interface = ":8080"

[api-packages]

[api-packages.status]
    routes = [
        { name = "/metrics", open = true },
        { name = "/prometheus-metrics", open = true }
    ]
```

After the configuration file is set up, the `elasticindexer` instance can be launched.

### Contribution

Contributions to the `drt-chain-es-indexer-go` module are welcomed. Whether you're interested in improving its features, 
extending its capabilities, or addressing issues, your contributions can help the community make the module even more robust.

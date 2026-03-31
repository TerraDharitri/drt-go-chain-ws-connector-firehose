package factory

import (
<<<<<<< HEAD
	"github.com/TerraDharitri/drt-go-chain-communication/websocket/data"
	factoryHost "github.com/TerraDharitri/drt-go-chain-communication/websocket/factory"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	"github.com/TerraDharitri/drt-go-chain-core/marshal/factory"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-chain-ws-connector-template/config"
	"github.com/TerraDharitri/drt-go-chain-ws-connector-template/process"
=======
	"os"

	"github.com/TerraDharitri/drt-go-chain-communication/websocket/data"
	factoryHost "github.com/TerraDharitri/drt-go-chain-communication/websocket/factory"
	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	"github.com/TerraDharitri/drt-go-chain-core/marshal/factory"
	logger "github.com/TerraDharitri/drt-go-chain-logger"

	"github.com/TerraDharitri/drt-go-chain-ws-connector-firehose/config"
	"github.com/TerraDharitri/drt-go-chain-ws-connector-firehose/process"
>>>>>>> 346a896 (test)
)

var log = logger.GetOrCreate("ws-connector")

// CreateWSConnector will create a ws connector able to receive and process incoming data
// from a dharitri node
func CreateWSConnector(cfg config.WebSocketConfig) (process.WSConnector, error) {
	marshaller, err := factory.NewMarshalizer(cfg.MarshallerType)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	dataProcessor, err := process.NewLogDataProcessor(marshaller, log)
=======
	blockContainer, err := createBlockContainer()
	if err != nil {
		return nil, err
	}

	dataProcessor, err := process.NewFirehoseDataProcessor(
		os.Stdout, // DO NOT CHANGE
		blockContainer,
		&marshal.GogoProtoMarshalizer{}, // DO NOT CHANGE
	)
>>>>>>> 346a896 (test)
	if err != nil {
		return nil, err
	}

	wsHost, err := createWsHost(marshaller, cfg)
	if err != nil {
		return nil, err
	}

	err = wsHost.SetPayloadHandler(dataProcessor)
	if err != nil {
		return nil, err
	}

	return wsHost, nil
}

<<<<<<< HEAD
=======
func createBlockContainer() (process.BlockContainerHandler, error) {
	container := block.NewEmptyBlockCreatorsContainer()

	err := container.Add(core.ShardHeaderV1, block.NewEmptyHeaderCreator())
	if err != nil {
		return nil, err
	}
	err = container.Add(core.ShardHeaderV2, block.NewEmptyHeaderV2Creator())
	if err != nil {
		return nil, err
	}
	err = container.Add(core.MetaHeader, block.NewEmptyMetaBlockCreator())
	if err != nil {
		return nil, err
	}

	return container, nil
}

>>>>>>> 346a896 (test)
func createWsHost(wsMarshaller marshal.Marshalizer, cfg config.WebSocketConfig) (factoryHost.FullDuplexHost, error) {
	return factoryHost.CreateWebSocketHost(factoryHost.ArgsWebSocketHost{
		WebSocketConfig: data.WebSocketConfig{
			URL:                        cfg.Url,
			WithAcknowledge:            cfg.WithAcknowledge,
			Mode:                       cfg.Mode,
			RetryDurationInSec:         int(cfg.RetryDuration),
			BlockingAckOnError:         cfg.BlockingAckOnError,
<<<<<<< HEAD
			DropMessagesIfNoConnection: false,
=======
			DropMessagesIfNoConnection: cfg.DropMessagesIfNoConnection,
			AcknowledgeTimeoutInSec:    cfg.AcknowledgeTimeoutInSec,
			Version:                    cfg.Version,
>>>>>>> 346a896 (test)
		},
		Marshaller: wsMarshaller,
		Log:        log,
	})
}

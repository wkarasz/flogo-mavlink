package gomavlib

import (
//	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"

//	"context"
//	"strconv"
	"time"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/carlescere/scheduler"
	"fmt"

	"github.com/wkarasz/gomavlib"
	"github.com/wkarasz/gomavlib/dialects/ardupilotmega"
)

// Create a new logger
var log = logger.GetLogger("trigger-mavlink-heartbeat")

const (
	// constant for ip-based endpoints
	_NET_BUFFER_SIZE      = 512 // frames cannot go beyond len(header) + 255 + len(check) + len(sig)
	_NET_CONNECT_TIMEOUT  = 10 * time.Second
	_NET_RECONNECT_PERIOD = 2 * time.Second
	_NET_READ_TIMEOUT     = 60 * time.Second
	_NET_WRITE_TIMEOUT    = 10 * time.Second
)

// Version allows to set the frame version used to wrap outgoing messages.
type Version int

const (
	// V2 wrap outgoing messages in v2 frames.
	V2 Version = iota
	// V1 wrap outgoing messages in v1 frames.
	V1
)


// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct{
	metadata *trigger.Metadata
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata:md}
}

// New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config:config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	timers   []*scheduler.Job
//	handlers []*trigger.Handler
	conf     *gomavlib.NodeConf
	n        *gomavlib.Node
}

// Initialize implements trigger.Init.Initialize
func (t *MyTrigger) Initialize(ctx trigger.InitContext) error {
	// Retrieve settings from user configuration
	if t.config.Settings == nil {
		return fmt.Errorf("No settings found for trigger '%s'", t.config.Id)
	}

	if _, ok := t.config.Settings["port"]; !ok {
		return fmt.Errorf("No port found for trigger '%s' in settings", t.config.Id)
	}

	port := t.config.GetSetting("port")

	t.conf = &gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointUdpServer{"0.0.0.0:"+port},
		},
		Dialect:	ardupilotmega.Dialect,
		OutSystemId:	10,
	}

	return nil
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	// start the trigger
	log.Debug("Start")
	var err error	
	t.n, err = gomavlib.NewNode(*t.conf)
	if err != nil {
		panic(err)
	}
	
	return nil
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger

	t.n.Close()

	return nil
}

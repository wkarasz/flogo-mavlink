package mavlinktrigger

import (
//	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
//	"github.com/TIBCOSoftware/flogo-lib/logger"
	
	"github.com/project-flogo/core/trigger"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/coerce"
	
	"context"
//	"github.com/carlescere/scheduler"
//	"fmt"

	"github.com/wkarasz/gomavlib"
//	"github.com/wkarasz/gomavlib/dialects/ardupilotmega"
)

// Create a new logger
//var log = logger.GetLogger("trigger-mavlink-datastream")

// Create a new metadata
var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&Trigger{}, &TriggerFactory{})
}

// TriggerFactory My Trigger factory
type TriggerFactory struct{
	metadata *trigger.Metadata
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &TriggerFactory{metadata: md}
}

// New Creates a new trigger instance for a given id
func (t *TriggerFactory) New(config *trigger.Config) (trigger.Trigger, error) {
	settings := &Settings{}
	err := metadata.MapToStruct(config.Settings, settings, true)
	if err != nil {
		return nil, err
	}
	if settings.Connection !=  "" {
		mConn, toConnerr := coerce.ToConnection(settings.Connection)
		if toConnerr != nil {
			return nil, toConnerr
		}
		node := mConn.GetConnection().(*gomavlib.Node)
		return &Trigger{metadata: t.metadata, settings: settings, id: config.Id, n: node}, nil
	}
	return nil, nil
}

// Metadata implements trigger.Factory.Metadata
func (*TriggerFactory) Metadata() *trigger.Metadata {
	return triggerMd
}

// Trigger is a stub for your Trigger implementation
type Trigger struct {
	metadata  *trigger.Metadata
	settings  *Settings
	evntLsnrs []*EventListener
	//conf     *gomavlib.NodeConf
	n         *gomavlib.Node
	logger    log.Logger
	id        string
}

// EventListener is structure of a single EventListener
type EventListener struct {
	handler   trigger.Handler
	settings  *HandlerSettings
	done      chan bool
	logger    log.Logger
	n         *gomavlib.Node
}

// Initialize implements trigger.Init.Initialize
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	t.logger = log.ChildLogger(ctx.Logger(), "mavlink-event-listener")
	t.logger.Infof("===========initializing event listener==")
	for _, handler := range ctx.GetHandlers() {
		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}
		evntLsnr := &EventListener{}
		evntLsnr.settings = s
		evntLsnr.handler = handler
		evntLsnr.logger = t.logger
		evntLsnr.done = make(chan bool)
		evntLsnr.n = t.n
		t.evntLsnrs = append(t.evntLsnrs, evntLsnr)
	}

	return nil
}

// Metadata implements trigger.Trigger.Metadata
func (t *Trigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *Trigger) Start() error {
	// start the trigger
	t.logger.Infof("Starting Trigger = %s", t.id)
	for _, evntLsnr := range t.evntLsnrs {
		go evntLsnr.listen()

	}
	
	return nil
}

func (evntLsnr *EventListener) listen() {
	evntLsnr.logger.Infof("=========listening==")
	for evt := range evntLsnr.n.Events() {
		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			//fmt.Printf("received: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
		

			trgData := make(map[string]interface{})
			//trgData["SystemId"] = frm.SystemId()
			//trgData["ComponentId"] = frm.ComponentId()
			trgData["MessageId"] = frm.Message().GetId()
			trgData["Message"] = frm.Message()

			results, err := evntLsnr.handler.Handle(context.Background(), trgData)
			if err != nil {
				evntLsnr.logger.Errorf("Failed to process, due to error - %s",err.Error())
			} else {
				// successfully processed
				evntLsnr.logger.Infof("Successfully processed: %s", results)
			}

		}
	}	
}

// Stop implements trigger.Trigger.Stop
func (t *Trigger) Stop() error {
	// stop the trigger
	t.n.Close()
	return nil
}

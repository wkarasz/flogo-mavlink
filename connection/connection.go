package mavlinkconnection

import (
	"context"
	"time"

	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/connection"
	"github.com/project-flogo/core/support/log"
	
	"github.com/wkarasz/gomavlib"
	"github.com/wkarasz/gomavlib/dialects/ardupilotmega"
)

var logmavlinkconn = log.ChildLogger(log.RootLogger(), "mavlink-connection")
var factory = &mavlinkFactory{}

// Settings struct
type Settings struct {
	Name          string `md:"name,required"`
	Description   string `md:"description"`
	Port          string `md:"port"`
}

func init() {
	err := connection.RegisterManagerFactory(factory)
	if err != nil {
		panic(err)
	}
}

type mavlinkFactory struct {
}

func (*mavlinkFactory) Type() string {
	return "mavlink"
}

func (*mavlinkFactory) NewManager(settings map[string]interface{}) (connection.Manager, error) {
	sharedConn := &MavlinkSharedConfigManager{}
	var err error

	// Retrieve config settings, if defined
	sharedConn.config, err = getmavlinkClientConfig(settings)
	if err != nil {
		return nil, err
	}
	if sharedConn.node != nil {
		// Already have a connection established,
		// Return the shared connection
		return sharedConn, nil
	}

	// Create new shared connection
	port := sharedConn.config.Port
	node, err1 := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointUdpServer{"0.0.0.0:"+port},
		},
		Dialect:	ardupilotmega.Dialect,
		OutSystemId: 	10,
	})
	if err1 != nil {
		return nil, err
	}
	sharedConn.node = node
	
	/*
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	// do a test here to validate shared connection is valid?
	
	err = client.Connect(ctx)
	if err != nil {
		logmavlinkconn.Errorf("===connect error==", err)
		return nil, err
	}
	*/
	return sharedConn, nil
}

// MavlinkSharedConfigManager Structure
type MavlinkSharedConfigManager struct {
	config  *Settings
	name    string
	node	*gomavlib.Node
}

// Type of SharedConfigManager
func (k *MavlinkSharedConfigManager) Type() string {
	return "mavlink"
}

// GetConnection ss
func (k *MavlinkSharedConfigManager) GetConnection() interface{} {
	return k.node
}

// ReleaseConnection ss
func (k *MavlinkSharedConfigManager) ReleaseConnection(connection interface{}) {

}

// Start connection manager
func (k *MavlinkSharedConfigManager) Start() error {
	return nil
}

// Stop connection manager
func (k *MavlinkSharedConfigManager) Stop() error {
	logmavlinkconn.Debug("Cleaning up client connection cache")
	return nil
}

// GetSharedConfiguration function to return Mavlink connection manager
func GetSharedConfiguration(conn interface{}) (connection.Manager, error) {
	var cManager connection.Manager
	var err error
	cManager, err = coerce.ToConnection(conn)
	if err != nil {
		return nil, err
	}
	return cManager, nil
}

func getmavlinkClientConfig(settings map[string]interface{}) (*Settings, error) {
	connectionConfig := &Settings{}
	err := metadata.MapToStruct(settings, connectionConfig, false)
	if err != nil {
		return nil, err
	}
	return connectionConfig, nil
}

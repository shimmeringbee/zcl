package zcl

import (
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_CommandRegistryGlobal(t *testing.T) {
	t.Run("getting a global command that does not exist results in an error", func(t *testing.T) {
		type ThisCommand struct{}
		expectedIdentifier := CommandIdentifier(1)

		cr := NewCommandRegistry()

		_, err := cr.GetGlobalCommandIdentifier(&ThisCommand{})
		assert.Error(t, err)

		_, err = cr.GetGlobalCommand(expectedIdentifier)
		assert.Error(t, err)
	})

	t.Run("registering a global command can be retrieved", func(t *testing.T) {
		type ThisCommand struct{}
		expectedIdentifier := CommandIdentifier(1)
		expectedType := reflect.TypeOf(&ThisCommand{})

		cr := NewCommandRegistry()

		cr.RegisterGlobal(expectedIdentifier, &ThisCommand{})

		actualIdentifier, err := cr.GetGlobalCommandIdentifier(&ThisCommand{})
		assert.NoError(t, err)
		assert.Equal(t, expectedIdentifier, actualIdentifier)

		cmd, err := cr.GetGlobalCommand(expectedIdentifier)
		actualType := reflect.TypeOf(cmd)

		assert.NoError(t, err)
		assert.Equal(t, expectedType, actualType)
	})

	t.Run("getting a local command that does not exist results in an error", func(t *testing.T) {
		type ThisCommand struct{}
		expectedIdentifier := CommandIdentifier(1)
		clusterId := zigbee.ClusterID(0x1020)
		manufacturer := zigbee.ManufacturerCode(0x3040)

		cr := NewCommandRegistry()

		_, err := cr.GetLocalCommandIdentifier(clusterId, manufacturer, ClientToServer, &ThisCommand{})
		assert.Error(t, err)

		_, err = cr.GetGlobalCommand(expectedIdentifier)
		assert.Error(t, err)
	})

	t.Run("registering a local command can be retrieved", func(t *testing.T) {
		type ThisCommand struct{}
		expectedIdentifier := CommandIdentifier(1)
		expectedType := reflect.TypeOf(&ThisCommand{})
		clusterId := zigbee.ClusterID(0x1020)
		manufacturer := zigbee.ManufacturerCode(0x3040)

		cr := NewCommandRegistry()

		cr.RegisterLocal(clusterId, manufacturer, ClientToServer, expectedIdentifier, &ThisCommand{})

		actualIdentifier, err := cr.GetLocalCommandIdentifier(clusterId, manufacturer, ClientToServer, &ThisCommand{})
		assert.NoError(t, err)
		assert.Equal(t, expectedIdentifier, actualIdentifier)

		cmd, err := cr.GetLocalCommand(clusterId, manufacturer, ClientToServer, expectedIdentifier)
		actualType := reflect.TypeOf(cmd)

		assert.NoError(t, err)
		assert.Equal(t, expectedType, actualType)
	})
}

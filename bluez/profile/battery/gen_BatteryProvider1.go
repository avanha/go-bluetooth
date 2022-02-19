// Code generated by go-bluetooth generator DO NOT EDIT.

package battery

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/muka/go-bluetooth/bluez"
	"github.com/muka/go-bluetooth/props"
	"github.com/muka/go-bluetooth/util"
)

var BatteryProvider1Interface = "org.bluez.BatteryProvider1"

// NewBatteryProvider1 create a new instance of BatteryProvider1
//
// Args:
// - servicePath: <client D-Bus address>
// - objectPath: {provider_root}/{unique battery object path}
func NewBatteryProvider1(servicePath string, objectPath dbus.ObjectPath) (*BatteryProvider1, error) {
	a := new(BatteryProvider1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: BatteryProvider1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(BatteryProvider1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	return a, nil
}

/*
BatteryProvider1 Battery Provider hierarchy

*/
type BatteryProvider1 struct {
	client                 *bluez.Client
	propertiesSignal       chan *dbus.Signal
	objectManagerSignal    chan *dbus.Signal
	objectManager          *bluez.ObjectManager
	Properties             *BatteryProvider1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// BatteryProvider1Properties contains the exposed properties of an interface
type BatteryProvider1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
		Device The object path of the device that has this battery.
	*/
	Device dbus.ObjectPath
}

//Lock access to properties
func (p *BatteryProvider1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *BatteryProvider1Properties) Unlock() {
	p.lock.Unlock()
}

// SetDevice set Device value
func (a *BatteryProvider1) SetDevice(v dbus.ObjectPath) error {
	return a.SetProperty("Device", v)
}

// GetDevice get Device value
func (a *BatteryProvider1) GetDevice() (dbus.ObjectPath, error) {
	v, err := a.GetProperty("Device")
	if err != nil {
		return dbus.ObjectPath(""), err
	}
	return v.Value().(dbus.ObjectPath), nil
}

// Close the connection
func (a *BatteryProvider1) Close() {
	a.unregisterPropertiesSignal()
	a.client.Disconnect()
}

// Path return BatteryProvider1 object path
func (a *BatteryProvider1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return BatteryProvider1 dbus client
func (a *BatteryProvider1) Client() *bluez.Client {
	return a.client
}

// Interface return BatteryProvider1 interface
func (a *BatteryProvider1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *BatteryProvider1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}

// ToMap convert a BatteryProvider1Properties to map
func (a *BatteryProvider1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an BatteryProvider1Properties
func (a *BatteryProvider1Properties) FromMap(props map[string]interface{}) (*BatteryProvider1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an BatteryProvider1Properties
func (a *BatteryProvider1Properties) FromDBusMap(props map[string]dbus.Variant) (*BatteryProvider1Properties, error) {
	s := new(BatteryProvider1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *BatteryProvider1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *BatteryProvider1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *BatteryProvider1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *BatteryProvider1) GetProperties() (*BatteryProvider1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *BatteryProvider1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *BatteryProvider1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *BatteryProvider1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *BatteryProvider1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *BatteryProvider1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *BatteryProvider1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}
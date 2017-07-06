package container

import (
	"sync"

	"github.com/mitchellh/mapstructure"
	"errors"
)

var (
	DefaultContainer = NewContainer()
)

type Creator func(cfg interface{}) (interface{}, error)

type Container struct {
	sync.RWMutex
	instances map[string]interface{}
	creators  map[string]Creator
}

func NewContainer() *Container {
	c := &Container{
		instances: make(map[string]interface{}),
		creators:  make(map[string]Creator),
	}
	c.Register("container", c.Creator)
	return c
}

func (this *Container) Set(name string, instance interface{}) {
	this.Lock()
	defer this.Unlock()
	this.instances[name] = instance
}

func (this *Container) Get(names ...string) (interface{}, error) {
	l := len(names)
	if l == 0 {
		return this, nil
	}
	this.RLock()
	defer this.RUnlock()

	if l == 1 {
		return this.get(names[0])
	}
	c := this
	var err error
	for _, name := range names[0 : l-1] {
		c, err = c.GetContainer(name)
		if err != nil {
			return nil, err
		}
	}

	return c.Get(names[l-1])
}
func (this *Container) get(name string) (interface{}, error) {
	if instance, ok := this.instances[name]; ok {
		return instance, nil
	}
	return nil, errors.New("instance not found")
}

func (this *Container) Register(name string, creator Creator) {
	this.Lock()
	defer this.Unlock()
	this.creators[name] = creator
}

func (this *Container) Create(cfg interface{}) (interface{}, error) {
	this.RLock()
	defer this.RUnlock()
	return this.create(cfg)
}

func (this *Container) create(cfg interface{}) (interface{}, error) {
	ccfg := make(map[string]interface{})
	err := mapstructure.WeakDecode(cfg, &ccfg)
	if err != nil {
		return nil, errors.New("can not trans interface: "+cfg.(string))
	}
	typ, ok := ccfg["type"].(string)
	if ok {
		delete(ccfg, "type")
	}
	if creator, ok := this.creators[typ]; ok {
		return creator(ccfg)
	}
	return nil, errors.New(typ+" creator not found")
}

func (this *Container) Configure(name string, cfg interface{}) error {
	this.Lock()
	defer this.Unlock()

	instance, err := this.create(cfg)
	if err == nil {
		this.instances[name] = instance
	}
	return err
}

func (this *Container) ConfigureAll(cfg map[string]interface{}) error {
	this.Lock()
	defer this.Unlock()
	instances := make(map[string]interface{})
	for k, v := range cfg {
		instance, err := this.create(v)
		if err != nil {
			return err
		}
		if instance != nil {
			instances[k] = instance
		}
	}
	this.instances = instances
	return nil
}

func (this *Container) GetContainer(name string) (*Container, error) {
	instance, err := this.Get(name)
	if err != nil {
		return nil, err
	}
	if c, ok := instance.(*Container); ok {
		return c, nil
	}
	return nil, errors.New(name +" not a Container instance")
}


//当传进来的type为container时，进入到这个Creator
func (this *Container) Creator(cfg interface{}) (interface{}, error) {
	cfgMap := make(map[string]interface{})
	err := mapstructure.WeakDecode(cfg, &cfgMap)
	if err == nil {
		c := NewContainer()
		c.creators = this.creators
		err := c.ConfigureAll(cfgMap)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	return nil, errors.New("container config not valid:"+cfg.(string))
}

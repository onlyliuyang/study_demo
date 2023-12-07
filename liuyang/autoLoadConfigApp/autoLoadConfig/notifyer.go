package autoLoadConfig

type Notifyer interface {
	Callback(config *Config)
}

//添加观察者
func (c *Config) AddObserver(n Notifyer) {
	c.notifyList = append(c.notifyList, n)
}

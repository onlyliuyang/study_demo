package config

import "strings"

type DefaultConfig struct {
	configData map[string]interface{}
}

func (c *DefaultConfig) get(name string) (result interface{}, found bool) {
	data := c.configData
	for _, key := range strings.Split(name, ":") {
		result, found := data[key]
		if newSection, ok := result.(map[string]interface{}); ok && found {
			data = newSection
		} else {
			return
		}
	}
	return
}

func (c *DefaultConfig) GetSection(name string) (section Configuration, found bool) {
	value, found := c.get(name)
	if found {
		if sectionData, ok := value.(map[string]interface{}); ok {
			section = &DefaultConfig{configData:sectionData}
		}
	}
	return
}

func (c *DefaultConfig) GetString(name string) (result string, found bool) {
	value, found := c.get(name)
	if found {
		result = value.(string)
	}
	return
}

func (c *DefaultConfig) GetInt(name string) (result int, found bool) {
	value, found := c.get(name)
	if found {
		result = int(value.(float64))
	}
	return
}

func (c *DefaultConfig) GetStringDefault(name string, val string) (result string) {
	result, found := c.GetString(name)
	if !found {
		result = val
	}
	return
}

func (c *DefaultConfig) GetIntDefault(name string, val int) (result int) {
	result, found := c.GetInt(name)
	if !found {
		result = val
	}
	return
}
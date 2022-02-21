package configs

import (
	"encoding/json"
	"os"
	"strings"
)

type DefaultCofig struct{
	configData map[string]interface{}
}

func Load(fileName string) (configz Configuration, err error){
	var data []byte
    data, err = os.ReadFile(fileName)
	if (err==nil){
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		m := map[string]interface{}{}
		err = decoder.Decode(&m)
		if (err == nil) {
			configz = &DefaultCofig{configData: m}
		}
	}

	return 
}

func (c *DefaultCofig) get(name string) (result interface{}, found bool) {
    data := c.configData
    for _, key := range strings.Split(name, ":") {
		result, found = data[key]
		if newSection, ok := result.(map[string]interface{}); ok && found {
		    data = newSection
		}else {
			return
		}
	}
	return
}

func (c *DefaultCofig) GetString(name string) (result string, found bool){
	value, found := c.get(name)
    if (found) { result = value.(string) }
	return 
}

func (c *DefaultCofig) GetStringDefault(name, val string)(result string){
	result, ok := c.GetString(name)
    if !ok{
		result = val
	}

	return 
}

func (c *DefaultCofig) GetInt(name string) (result int, found bool){
	value, found := c.get(name)
    if (found) { result = value.(int) }
	return 
}

func (c *DefaultCofig) GetIntDefault(name string, val int)(result int){
	result, ok := c.GetInt(name)
    if !ok{
		result = val
	}

	return 
}

func (c *DefaultCofig) GetBool(name string) (result bool, found bool){
	value, found := c.get(name)
    if (found) { result = value.(bool) }
	return 
}

func (c *DefaultCofig) GetBoolDefault(name string, val bool)(result bool){
	result, ok := c.GetBool(name)
    if !ok{
		result = val
	}

	return 
}

func (c *DefaultCofig) GetFloat(name string)  (result float64, found bool) {
	value, found := c.get(name)
    if (found) { result = value.(float64) }
	return 
}

func (c *DefaultCofig) GetFloatDefault(name string, val float64)(result float64){
	result, ok := c.GetFloat(name)
    if !ok{
		result = val
	}

	return 
}


func (c *DefaultCofig) GetSection(name string)(section Configuration, found bool){
	value,found := c.get(name)
	if (found){
		sectionData,ok:=value.(map[string]interface{})
		if ok{
			section = &DefaultCofig{configData: sectionData}
		}
	}
	return 
}
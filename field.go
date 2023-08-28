package jsonfield

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// ReserveFieldWithObj will reserve some fields from object
func ReserveFieldWithObj(obj interface{}, path []string) ([]byte, error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("marshal object failed: %s", err.Error())
	}
	return ReserveFieldWithObj(bs, path)
}

// ReserveField will reserve some files from bytes
func ReserveField(bs []byte, path []string) ([]byte, error) {
	c, err := parsed(bs)
	if err != nil {
		return nil, fmt.Errorf("parsed bytes failed: %s", err.Error())
	}
	bs, err = c.ReservePath(path...)
	if err != nil {
		return nil, fmt.Errorf("reserve path failed: %s", err.Error())
	}
	return bs, nil
}

// container defines the jsonq traverse instance
type container struct {
	path   string
	object interface{}

	reservePath []string
}

// parsed will unmarshal the result byte to object
func parsed(bs []byte) (container, error) {
	c := container{}
	if err := json.Unmarshal(bs, &c.object); err != nil {
		return c, err
	}
	return c, nil
}

// ReservePath will reserve path from byte
func (c *container) ReservePath(path ...string) ([]byte, error) {
	c.reservePath = path
	c.traverse()
	return json.Marshal(c.object)
}

func (c *container) reservedPathMatched() bool {
	for i := range c.reservePath {
		p := c.reservePath[i]
		if strings.HasPrefix(p, c.path) || strings.HasPrefix(c.path, p) {
			return true
		}
	}
	return false
}

func (c *container) reservedPathMatchedForMetaNode() bool {
	for i := range c.reservePath {
		p := c.reservePath[i]
		if strings.HasPrefix(c.path+".", p+".") {
			return true
		}
	}
	return false
}

func (c *container) traverse() bool {
	if !c.reservedPathMatched() {
		return true
	}
	if c.object == nil {
		return false
	}
	switch reflect.TypeOf(c.object).Kind() {
	case reflect.Map:
		for k, v := range c.object.(map[string]interface{}) {
			newC := container{
				path:        c.path + k + ".",
				object:      v,
				reservePath: c.reservePath,
			}
			if newC.traverse() {
				delete(c.object.(map[string]interface{}), k)
			}
		}
	case reflect.Slice:
		for i := range c.object.([]interface{}) {
			v := c.object.([]interface{})[i]
			newC := container{
				path:        c.path,
				object:      v,
				reservePath: c.reservePath,
			}
			newC.traverse()
		}
	default:
		c.path = strings.TrimSuffix(c.path, ".")
		if !c.reservedPathMatchedForMetaNode() {
			return true
		}
		return false
	}
	return false
}

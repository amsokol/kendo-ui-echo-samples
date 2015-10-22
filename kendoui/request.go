package kendoui

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// filter[logic]=and
// filter[filters][0][field]=Title
// filter[filters][0][ignoreCase]=true
// filter[filters][0][operator]=startswith
// filter[filters][0][value]=the

const (
	LogicAnd = iota
)

const (
	OperatorStartsWith = iota
)

type RequestInput struct {
	Filter Filter
}

type Filter struct {
	Logic   int
	Filters []FilterItem
}

type FilterItem struct {
	Field      string
	IgnoreCase bool
	Operator   int
	Value      string
}

func Input(r *http.Request) *RequestInput {
	i := &RequestInput{
		Filter: Filter{
			Filters: make([]FilterItem, 0),
		},
	}
	i.Parse(r.URL.Query())
	//	i.Parse(r.Form)
	return i
}

func (i *RequestInput) Parse(vs url.Values) {
	for key, value := range vs {
		keys := strings.Split(key, "[")
		var err error
		switch keys[0] {
		case "filter":
			err = i.extractFilter(keys[1:], value)
		default:
			// ignoring parameter
			// log.Println(fmt.Sprint("unsupported request parameter key='", key, "', value='", value, "'"))
		}
		if err != nil {
			log.Println(err, "key='", key, "', value='", value, "'")
		}
	}
}

func (i *RequestInput) extractFilter(keys []string, value []string) (err error) {
	if len(keys) == 0 {
		return errors.New("filter must have sub-parameters")
	}
	switch strings.TrimRight(keys[0], "]") {
	case "logic":
		i.Filter.Logic, err = getLogicValue(value)
	case "filters":
		err = i.extractFilterItem(keys[1:], value)
	default:
		err = errors.New("filter has unsupported sub-parameter")
	}
	return err
}

func (i *RequestInput) extractFilterItem(keys []string, value []string) (err error) {
	var ok bool

	if len(keys) != 2 {
		return errors.New("filter[filters] must have two sub-parameters")
	}
	var index int
	if index, err = strconv.Atoi(strings.TrimRight(keys[0], "]")); err != nil {
		return errors.New("filter[filters] has unsupported sub-parameter - must be [<index>]")
	}
	l := len(i.Filter.Filters)
	if l <= index {
		i.Filter.Filters = append(i.Filter.Filters, make([]FilterItem, index-l+1)...)
	}
	switch strings.TrimRight(keys[1], "]") {
	case "field":
		i.Filter.Filters[index].Field, ok = getStringValue(value)
		if !ok {
			err = errors.New("filter[filters][<index>][field] must have single value")
		}
	case "ignoreCase":
		i.Filter.Filters[index].IgnoreCase, ok = getBoolValue(value)
		if !ok {
			err = errors.New("filter[filters][<index>][ignoreCase] must have single bool value")
		}
	case "operator":
		i.Filter.Filters[index].Operator, err = getOperatorValue(value)
	case "value":
		i.Filter.Filters[index].Value, ok = getStringValue(value)
		if !ok {
			err = errors.New("filter[filters][<index>][value] must have single value")
		}
	default:
		err = errors.New("filter[filters][<index>] has unsupported sub-parameter")
	}
	return
}

func getLogicValue(value []string) (v int, err error) {
	s, ok := getStringValue(value)
	if ok {
		switch s {
		case "and":
			v = LogicAnd
		default:
			err = errors.New("filter[logic] has unsupported value")
		}
	} else {
		err = errors.New("filter[logic] must have single value")
	}
	return
}

func getOperatorValue(value []string) (v int, err error) {
	s, ok := getStringValue(value)
	if ok {
		switch s {
		case "startswith":
			v = OperatorStartsWith
		default:
			err = errors.New("filter[filters][<index>][operator] has unsupported value")
		}
	} else {
		err = errors.New("filter[filters][<index>][operator] must have single value")
	}
	return
}

func getStringValue(value []string) (v string, ok bool) {
	ok = len(value) == 1
	if ok {
		v = value[0]
	}
	return
}

func getBoolValue(value []string) (v bool, ok bool) {
	var s string
	var err error
	s, ok = getStringValue(value)
	if ok {
		v, err = strconv.ParseBool(s)
		ok = err == nil
	}
	return
}

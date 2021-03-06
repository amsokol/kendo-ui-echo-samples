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

type RequestInput struct {
	Filter Filter
}

type Filter struct {
	Logic   string
	Filters []FilterItem
}

type FilterItem struct {
	Field      string
	IgnoreCase bool
	Operator   string
	Value      string
}

func Input(r *http.Request) *RequestInput {
	i := &RequestInput{
		Filter: Filter{
			Filters: make([]FilterItem, 0),
		},
	}
	i.Extract(r.URL.Query())
	i.Extract(r.Form)
	return i
}

func (i *RequestInput) Extract(vs url.Values) {
	for key, value := range vs {
		keys := strings.Split(key, "[")
		var err error
		switch keys[0] {
		case "filter":
			err = i.extractFilter(keys[1:], value)
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
		i.Filter.Logic, err = getString(value)
	case "filters":
		err = i.extractFilterItem(keys[1:], value)
	default:
		err = errors.New("filter has unsupported sub-parameter")
	}
	return err
}

func (i *RequestInput) extractFilterItem(keys []string, value []string) (err error) {
	if len(keys) != 2 {
		return errors.New("filter[filters] must have two sub-parameters")
	}
	var idx int
	if idx, err = strconv.Atoi(strings.TrimRight(keys[0], "]")); err != nil {
		return
	}
	if l := len(i.Filter.Filters); l <= idx {
		i.Filter.Filters = append(i.Filter.Filters, make([]FilterItem, idx-l+1)...)
	}
	switch strings.TrimRight(keys[1], "]") {
	case "field":
		i.Filter.Filters[idx].Field, err = getString(value)
	case "ignoreCase":
		i.Filter.Filters[idx].IgnoreCase, err = getBool(value)
	case "operator":
		i.Filter.Filters[idx].Operator, err = getString(value)
	case "value":
		i.Filter.Filters[idx].Value, err = getString(value)
	default:
		err = errors.New("filter[filters][<index>] has unsupported sub-parameter")
	}
	return
}

func getString(value []string) (v string, err error) {
	if len(value) == 1 {
		v = value[0]
	} else {
		err = errors.New("must have single value")
	}
	return
}

func getBool(value []string) (v bool, err error) {
	var s string
	if s, err = getString(value); err == nil {
		v, err = strconv.ParseBool(s)
	}
	return
}

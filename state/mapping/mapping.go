package mapping

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/s7techlab/cckit/state"
)

type (
	// StateMappers interface for mappers collection
	StateMappers interface {
		Exists(schema interface{}) (exists bool)
		Map(schema interface{}) (keyValue state.KeyValue, err error)
		Get(schema interface{}) (stateMapper StateMapper, err error)
	}

	EventMappers interface {
		Exists(schema interface{}) (exists bool)
		Map(schema interface{}) (keyValue state.KeyValue, err error)
		Get(schema interface{}) (eventMapper EventMapper, err error)
	}

	StateMapper interface {
		Schema() interface{}
		Namespace() []string
		PrimaryKey(instance interface{}) (key []string, err error)
	}
	EventMapper interface {
		Schema() interface{}
		Name() (string, error)
	}

	StateMapping struct {
		schema       interface{}
		namespace    []string
		primaryKeyer state.KeyTransformer
		//PKStringer KeyerFunc
		//PKToString KeyerFunc
		//niqKey []KeyTransformer
		//Key     []KeyTransformer
	}

	StateMappings map[string]*StateMapping

	Namer func(entity interface{}) string

	Event struct {
		schema interface{}
		name   string
	}

	Events map[string]*Event
)

func mapKey(entry interface{}) string {
	return reflect.TypeOf(entry).String()
}

func (smm StateMappings) Add(schema interface{}, namespace []string, primaryKeyer state.KeyTransformer) StateMappings {
	smm[mapKey(schema)] = &StateMapping{
		schema:       schema,
		namespace:    namespace,
		primaryKeyer: primaryKeyer,
	}
	return smm
}

func (smm StateMappings) Get(entry interface{}) (StateMapper, error) {
	m, ok := smm[mapKey(entry)]
	if !ok {
		return nil, fmt.Errorf(`%s: %s`, ErrEntryTypeNotDefined, mapKey(entry))
	}
	return m, nil
}

func (smm StateMappings) Exists(entry interface{}) bool {
	_, err := smm.Get(entry)
	return err == nil
}

func (smm StateMappings) Map(entry interface{}) (mapped state.KeyValue, err error) {
	mapping, err := smm.Get(entry)
	if err != nil {
		return nil, errors.Wrap(err, `mapping`)
	}

	switch entry.(type) {
	case proto.Message:
		return NewProtoStateMapper(entry, mapping)
	default:
		return nil, ErrEntryTypeNotSupported
	}
}

func (sm *StateMapping) Namespace() []string {
	return sm.namespace
}
func (sm *StateMapping) Schema() interface{} {
	return sm.schema
}

func (s *StateMapping) PrimaryKey(entity interface{}) ([]string, error) {
	key, err := s.primaryKeyer(entity)
	if err != nil {
		return nil, err
	}
	return append(s.namespace, key...), nil
}

func (ee Events) Add(name string, schema interface{}) Events {
	ee[name] = &Event{
		schema: schema,
		name:   name,
	}
	return ee
}

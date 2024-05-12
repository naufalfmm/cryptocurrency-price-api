package ws

import jsoniter "github.com/json-iterator/go"

type Context interface {
	Bind(o any) error
}

type context struct {
	msg []byte
	err error
}

func (c context) Bind(o any) error {
	if c.err != nil {
		return c.err
	}

	if len(c.msg) == 0 {
		return nil
	}

	return jsoniter.Unmarshal(c.msg, o)
}

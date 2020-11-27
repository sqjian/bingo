package proto

import (
	"testing"
)

func TestDataValidation(t *testing.T) {
	msg := Msg{
		DataList: []*Data{
			{
				Id:   "uuid",
				Desc: map[string][]byte{"key": []byte("val")},
				Data: []byte("data"),
			},
		},
	}
	t.Log(msg.Validate(true))
}

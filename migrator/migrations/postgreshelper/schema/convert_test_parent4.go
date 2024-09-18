// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertTestParent4FromProto converts a `*storage.TestParent4` to Gorm model
func ConvertTestParent4FromProto(obj *storage.TestParent4) (*TestParent4, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &TestParent4{
		ID:         obj.GetId(),
		ParentID:   obj.GetParentId(),
		Val:        obj.GetVal(),
		Serialized: serialized,
	}
	return model, nil
}

// ConvertTestParent4ToProto converts Gorm model `TestParent4` to its protobuf type object
func ConvertTestParent4ToProto(m *TestParent4) (*storage.TestParent4, error) {
	var msg storage.TestParent4
	if err := msg.UnmarshalVTUnsafe(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}

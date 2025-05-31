package models

import (
	"github.com/gobuffalo/pop/v6"
)

type ConnectionType struct {
	BaseEnumModel
}

type ConnectionTypes []ConnectionType

func AllConnectionTypes(tx *pop.Connection) (ConnectionTypes, error) {
	enums := ConnectionTypes{}
	if err := tx.All(&enums); err != nil {
		return nil, err
	}

	return enums, nil
}

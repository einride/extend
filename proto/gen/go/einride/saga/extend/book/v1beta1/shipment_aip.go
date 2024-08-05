// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc (unknown)
// source: einride/saga/extend/book/v1beta1/shipment.proto

package bookv1beta1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type ShipmentResourceName struct {
	Space    string
	Shipment string
}

func (n SpaceResourceName) ShipmentResourceName(
	shipment string,
) ShipmentResourceName {
	return ShipmentResourceName{
		Space:    n.Space,
		Shipment: shipment,
	}
}

func (n ShipmentResourceName) Validate() error {
	if n.Space == "" {
		return fmt.Errorf("space: empty")
	}
	if strings.IndexByte(n.Space, '/') != -1 {
		return fmt.Errorf("space: contains illegal character '/'")
	}
	if n.Shipment == "" {
		return fmt.Errorf("shipment: empty")
	}
	if strings.IndexByte(n.Shipment, '/') != -1 {
		return fmt.Errorf("shipment: contains illegal character '/'")
	}
	return nil
}

func (n ShipmentResourceName) ContainsWildcard() bool {
	return false || n.Space == "-" || n.Shipment == "-"
}

func (n ShipmentResourceName) String() string {
	return resourcename.Sprint(
		"spaces/{space}/shipments/{shipment}",
		n.Space,
		n.Shipment,
	)
}

func (n ShipmentResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *ShipmentResourceName) UnmarshalString(name string) error {
	err := resourcename.Sscan(
		name,
		"spaces/{space}/shipments/{shipment}",
		&n.Space,
		&n.Shipment,
	)
	if err != nil {
		return err
	}
	return n.Validate()
}

func (n ShipmentResourceName) SpaceResourceName() SpaceResourceName {
	return SpaceResourceName{
		Space: n.Space,
	}
}

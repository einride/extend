// Code generated by protoc-gen-go-aip-cli. DO NOT EDIT.
package bookv1beta1

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func NewShipmentServiceCommand(config aipcli.Config) *cobra.Command {
	return aipcli.NewServiceCommand(
		config,
		File_einride_saga_extend_book_v1beta1_shipment_service_proto.
			Services().ByName("ShipmentService"),
		map[protoreflect.FullName]string{
			"einride.saga.extend.book.v1beta1.ShipmentService": " Shipment service.\n",
		},
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_shipment_service_proto.
				Services().ByName("ShipmentService").Methods().ByName("CreateShipment"),
			&CreateShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.Address.city":                           " City\n",
				"einride.saga.extend.book.v1beta1.Address.co":                             " Care of (C/O)\n",
				"einride.saga.extend.book.v1beta1.Address.contact_info":                   " Contact information\n",
				"einride.saga.extend.book.v1beta1.Address.display_name":                   " The displayed name of the address\n",
				"einride.saga.extend.book.v1beta1.Address.external_reference_id":          " An external reference for this address.\n",
				"einride.saga.extend.book.v1beta1.Address.lat_lng":                        " The geographic location of the address\n",
				"einride.saga.extend.book.v1beta1.Address.line1":                          " Address line 1\n",
				"einride.saga.extend.book.v1beta1.Address.line2":                          " Address line 2\n",
				"einride.saga.extend.book.v1beta1.Address.postal_code":                    " Postal code\n",
				"einride.saga.extend.book.v1beta1.Address.recipient":                      " Recipient\n",
				"einride.saga.extend.book.v1beta1.Address.region_code":                    " Region code (Unicode CLDR region code)\n https://cldr.unicode.org/\n",
				"einride.saga.extend.book.v1beta1.Address.state_code":                     " State code\n",
				"einride.saga.extend.book.v1beta1.CreateShipmentRequest.parent":           " The parent space in which to create the shipment.\n",
				"einride.saga.extend.book.v1beta1.CreateShipmentRequest.shipment":         " The shipment to create.\n",
				"einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry.key":          "",
				"einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry.value":        "",
				"einride.saga.extend.book.v1beta1.Shipment.annotations":                   " Annotations for the shipment\n",
				"einride.saga.extend.book.v1beta1.Shipment.create_time":                   " The creation timestamp of the shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.delete_time":                   " The deletion timestamp of the shipment. Set if the shipment is deleted.\n",
				"einride.saga.extend.book.v1beta1.Shipment.delivery_address":              " The address where the shipment is delivered to.\n",
				"einride.saga.extend.book.v1beta1.Shipment.delivery_instructions":         " Instructions for the delivery.\n",
				"einride.saga.extend.book.v1beta1.Shipment.delivery_state":                " Shipment delivery state\n",
				"einride.saga.extend.book.v1beta1.Shipment.external_reference_id":         " An external reference for this shipment. If supplied then it must be unique within the space.\n",
				"einride.saga.extend.book.v1beta1.Shipment.name":                          " The resource name of the shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.pickup_address":                " The address where the shipment is picked up.\n Supply when creating shipment if different from the sender default pickup address.\n",
				"einride.saga.extend.book.v1beta1.Shipment.pickup_instructions":           " Instructions for the pickup.\n",
				"einride.saga.extend.book.v1beta1.Shipment.requested_delivery_end_time":   " The shipment should be delivered before this time\n",
				"einride.saga.extend.book.v1beta1.Shipment.requested_delivery_start_time": " The shipment should be delivered after this time\n",
				"einride.saga.extend.book.v1beta1.Shipment.requested_pickup_end_time":     " The shipment should be picked up before this time\n",
				"einride.saga.extend.book.v1beta1.Shipment.requested_pickup_start_time":   " The shipment should be picked up after this time\n",
				"einride.saga.extend.book.v1beta1.Shipment.sender":                        " Resource name of the sender of the shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.service":                       " Shipment service, defaults to PALLET\n",
				"einride.saga.extend.book.v1beta1.Shipment.space":                         " Resource name of the space that owns the shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.state":                         " Shipment state\n Can be set to ACTIVE (default) or RELEASED when creating a shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.tour":                          " Resource name of the truck tour booking this shipment relates to.\n",
				"einride.saga.extend.book.v1beta1.Shipment.tracking_id":                   " A generated tracking ID for this shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.units":                         " Shipment units\n",
				"einride.saga.extend.book.v1beta1.Shipment.update_time":                   " The last update timestamp of the shipment.\n Updated when create/update/delete operation is performed.\n",
				"einride.saga.extend.book.v1beta1.Shipment.vehicle":                       " Vehicle information for the shipment\n",
				"einride.saga.extend.book.v1beta1.ShipmentService.CreateShipment":         " Create a shipment in a space.\n\n This is an AIP standard [Create](https://google.aip.dev/133) method.\n",
				"einride.saga.extend.book.v1beta1.Unit.Length.unit":                       " Length unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Length.value":                      " Length value\n",
				"einride.saga.extend.book.v1beta1.Unit.Volume.unit":                       " Volume unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Volume.value":                      " Volume value\n",
				"einride.saga.extend.book.v1beta1.Unit.Weight.unit":                       " Weight unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Weight.value":                      " Weight value\n",
				"einride.saga.extend.book.v1beta1.Unit.description":                       " Description of the unit\n",
				"einride.saga.extend.book.v1beta1.Unit.external_reference_id":             " External reference ID of the unit\n",
				"einride.saga.extend.book.v1beta1.Unit.height":                            " Unit height\n",
				"einride.saga.extend.book.v1beta1.Unit.length":                            " Unit length\n",
				"einride.saga.extend.book.v1beta1.Unit.tracking_id":                       " Unique unit tracking ID\n",
				"einride.saga.extend.book.v1beta1.Unit.type":                              " Type of unit\n",
				"einride.saga.extend.book.v1beta1.Unit.volume":                            " Unit volume\n",
				"einride.saga.extend.book.v1beta1.Unit.weight":                            " Unit weight\n",
				"einride.saga.extend.book.v1beta1.Unit.width":                             " Unit width\n",
				"einride.saga.extend.book.v1beta1.Vehicle.carrier_reference_id":           " Carrier reference id\n",
				"einride.saga.extend.book.v1beta1.Vehicle.driver_reference_id":            " Driver reference id\n",
				"einride.saga.extend.book.v1beta1.Vehicle.reference_id":                   " Reference id\n",
				"google.protobuf.Timestamp.nanos":                                         " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                                       " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
				"google.type.LatLng.latitude":                                             " The latitude in degrees. It must be in the range [-90.0, +90.0].\n",
				"google.type.LatLng.longitude":                                            " The longitude in degrees. It must be in the range [-180.0, +180.0].\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_shipment_service_proto.
				Services().ByName("ShipmentService").Methods().ByName("GetShipment"),
			&GetShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.GetShipmentRequest.name":     " The resource name of the shipment to retrieve.\n",
				"einride.saga.extend.book.v1beta1.ShipmentService.GetShipment": " Get a shipment.\n\n This is an AIP standard [Get](https://google.aip.dev/131) method.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_shipment_service_proto.
				Services().ByName("ShipmentService").Methods().ByName("ReleaseShipment"),
			&ReleaseShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.ReleaseShipmentRequest.name":     " The resource name of the shipment to release.\n",
				"einride.saga.extend.book.v1beta1.ShipmentService.ReleaseShipment": " Release a shipment.\n\n The state of the shipment after releasing it is RELEASED.\n\n This is an AIP [state](https://google.aip.dev/216) transition method.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_shipment_service_proto.
				Services().ByName("ShipmentService").Methods().ByName("CancelShipment"),
			&CancelShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.CancelShipmentRequest.name":     " The resource name of the shipment to cancel.\n",
				"einride.saga.extend.book.v1beta1.ShipmentService.CancelShipment": " Cancel a shipment.\n\n The state of the shipment after cancelling it is CANCELLED.\n\n This is an AIP [state](https://google.aip.dev/216) transition method.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_shipment_service_proto.
				Services().ByName("ShipmentService").Methods().ByName("UpdateShipment"),
			&UpdateShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.Address.city":                           " City\n",
				"einride.saga.extend.book.v1beta1.Address.co":                             " Care of (C/O)\n",
				"einride.saga.extend.book.v1beta1.Address.contact_info":                   " Contact information\n",
				"einride.saga.extend.book.v1beta1.Address.display_name":                   " The displayed name of the address\n",
				"einride.saga.extend.book.v1beta1.Address.external_reference_id":          " An external reference for this address.\n",
				"einride.saga.extend.book.v1beta1.Address.lat_lng":                        " The geographic location of the address\n",
				"einride.saga.extend.book.v1beta1.Address.line1":                          " Address line 1\n",
				"einride.saga.extend.book.v1beta1.Address.line2":                          " Address line 2\n",
				"einride.saga.extend.book.v1beta1.Address.postal_code":                    " Postal code\n",
				"einride.saga.extend.book.v1beta1.Address.recipient":                      " Recipient\n",
				"einride.saga.extend.book.v1beta1.Address.region_code":                    " Region code (Unicode CLDR region code)\n https://cldr.unicode.org/\n",
				"einride.saga.extend.book.v1beta1.Address.state_code":                     " State code\n",
				"einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry.key":          "",
				"einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry.value":        "",
				"einride.saga.extend.book.v1beta1.Shipment.annotations":                   " Annotations for the shipment\n",
				"einride.saga.extend.book.v1beta1.Shipment.create_time":                   " The creation timestamp of the shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.delete_time":                   " The deletion timestamp of the shipment. Set if the shipment is deleted.\n",
				"einride.saga.extend.book.v1beta1.Shipment.delivery_address":              " The address where the shipment is delivered to.\n",
				"einride.saga.extend.book.v1beta1.Shipment.delivery_instructions":         " Instructions for the delivery.\n",
				"einride.saga.extend.book.v1beta1.Shipment.delivery_state":                " Shipment delivery state\n",
				"einride.saga.extend.book.v1beta1.Shipment.external_reference_id":         " An external reference for this shipment. If supplied then it must be unique within the space.\n",
				"einride.saga.extend.book.v1beta1.Shipment.name":                          " The resource name of the shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.pickup_address":                " The address where the shipment is picked up.\n Supply when creating shipment if different from the sender default pickup address.\n",
				"einride.saga.extend.book.v1beta1.Shipment.pickup_instructions":           " Instructions for the pickup.\n",
				"einride.saga.extend.book.v1beta1.Shipment.requested_delivery_end_time":   " The shipment should be delivered before this time\n",
				"einride.saga.extend.book.v1beta1.Shipment.requested_delivery_start_time": " The shipment should be delivered after this time\n",
				"einride.saga.extend.book.v1beta1.Shipment.requested_pickup_end_time":     " The shipment should be picked up before this time\n",
				"einride.saga.extend.book.v1beta1.Shipment.requested_pickup_start_time":   " The shipment should be picked up after this time\n",
				"einride.saga.extend.book.v1beta1.Shipment.sender":                        " Resource name of the sender of the shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.service":                       " Shipment service, defaults to PALLET\n",
				"einride.saga.extend.book.v1beta1.Shipment.space":                         " Resource name of the space that owns the shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.state":                         " Shipment state\n Can be set to ACTIVE (default) or RELEASED when creating a shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.tour":                          " Resource name of the truck tour booking this shipment relates to.\n",
				"einride.saga.extend.book.v1beta1.Shipment.tracking_id":                   " A generated tracking ID for this shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.units":                         " Shipment units\n",
				"einride.saga.extend.book.v1beta1.Shipment.update_time":                   " The last update timestamp of the shipment.\n Updated when create/update/delete operation is performed.\n",
				"einride.saga.extend.book.v1beta1.Shipment.vehicle":                       " Vehicle information for the shipment\n",
				"einride.saga.extend.book.v1beta1.ShipmentService.UpdateShipment":         " Update a shipment in a space.\n\n This is an AIP standard [Update](https://google.aip.dev/134) method.\n",
				"einride.saga.extend.book.v1beta1.Unit.Length.unit":                       " Length unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Length.value":                      " Length value\n",
				"einride.saga.extend.book.v1beta1.Unit.Volume.unit":                       " Volume unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Volume.value":                      " Volume value\n",
				"einride.saga.extend.book.v1beta1.Unit.Weight.unit":                       " Weight unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Weight.value":                      " Weight value\n",
				"einride.saga.extend.book.v1beta1.Unit.description":                       " Description of the unit\n",
				"einride.saga.extend.book.v1beta1.Unit.external_reference_id":             " External reference ID of the unit\n",
				"einride.saga.extend.book.v1beta1.Unit.height":                            " Unit height\n",
				"einride.saga.extend.book.v1beta1.Unit.length":                            " Unit length\n",
				"einride.saga.extend.book.v1beta1.Unit.tracking_id":                       " Unique unit tracking ID\n",
				"einride.saga.extend.book.v1beta1.Unit.type":                              " Type of unit\n",
				"einride.saga.extend.book.v1beta1.Unit.volume":                            " Unit volume\n",
				"einride.saga.extend.book.v1beta1.Unit.weight":                            " Unit weight\n",
				"einride.saga.extend.book.v1beta1.Unit.width":                             " Unit width\n",
				"einride.saga.extend.book.v1beta1.UpdateShipmentRequest.shipment":         " The resource which replaces the current resource.\n",
				"einride.saga.extend.book.v1beta1.UpdateShipmentRequest.update_mask":      " The update mask applies to the shipment.\n\n Currently, it allows updating with the following:\n * external_reference_id\n * units\n For units you cannot update individual units, so you MUST provide all the other units\n * vehicle\n * vehicle.reference_id\n * vehicle.driver_reference_id\n * vehicle.carrier_reference_id\n * pickup_address\n * pickup_instructions\n * delivery_address\n * delivery_instructions\n * annotations\n * service\n\n The star (*) update mask will update the above listed masks\n This applies also for empty update mask partial update\n For the `FieldMask` definition, see:\n https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask\n",
				"einride.saga.extend.book.v1beta1.Vehicle.carrier_reference_id":           " Carrier reference id\n",
				"einride.saga.extend.book.v1beta1.Vehicle.driver_reference_id":            " Driver reference id\n",
				"einride.saga.extend.book.v1beta1.Vehicle.reference_id":                   " Reference id\n",
				"google.protobuf.FieldMask.paths":                                         " The set of field mask paths.\n",
				"google.protobuf.Timestamp.nanos":                                         " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                                       " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
				"google.type.LatLng.latitude":                                             " The latitude in degrees. It must be in the range [-90.0, +90.0].\n",
				"google.type.LatLng.longitude":                                            " The longitude in degrees. It must be in the range [-180.0, +180.0].\n",
			},
		),
	)
}

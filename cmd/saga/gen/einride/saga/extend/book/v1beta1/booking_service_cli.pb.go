// Code generated by protoc-gen-go-aip-cli. DO NOT EDIT.
package bookv1beta1

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func NewBookingServiceCommand(config aipcli.Config) *cobra.Command {
	return aipcli.NewServiceCommand(
		config,
		File_einride_saga_extend_book_v1beta1_booking_service_proto.
			Services().ByName("BookingService"),
		map[protoreflect.FullName]string{
			"einride.saga.extend.book.v1beta1.BookingService": " Booking service.\n",
		},
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_booking_service_proto.
				Services().ByName("BookingService").Methods().ByName("CreateBooking"),
			&CreateBookingRequest{},
			&Booking{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.Address.city":                             " City\n",
				"einride.saga.extend.book.v1beta1.Address.co":                               " Care of (C/O)\n",
				"einride.saga.extend.book.v1beta1.Address.contact_info":                     " Contact information\n",
				"einride.saga.extend.book.v1beta1.Address.display_name":                     " The displayed name of the address\n",
				"einride.saga.extend.book.v1beta1.Address.lat_lng":                          " The geographic location of the address\n",
				"einride.saga.extend.book.v1beta1.Address.line1":                            " Address line 1\n",
				"einride.saga.extend.book.v1beta1.Address.line2":                            " Address line 2\n",
				"einride.saga.extend.book.v1beta1.Address.postal_code":                      " Postal code\n",
				"einride.saga.extend.book.v1beta1.Address.recipient":                        " Recipient\n",
				"einride.saga.extend.book.v1beta1.Address.region_code":                      " Region code (Unicode CLDR region code)\n https://cldr.unicode.org/\n",
				"einride.saga.extend.book.v1beta1.Address.state_code":                       " State code\n",
				"einride.saga.extend.book.v1beta1.Booking.Stop.address":                     " The address of this stop.\n",
				"einride.saga.extend.book.v1beta1.Booking.Stop.instructions":                " Instructions for this stop.\n",
				"einride.saga.extend.book.v1beta1.Booking.Stop.requested_end_time":          " The latest time to depart this stop.\n",
				"einride.saga.extend.book.v1beta1.Booking.Stop.requested_start_time":        " The earliest time to arrive this stop.\n",
				"einride.saga.extend.book.v1beta1.Booking.Stop.stop_type":                   " The type of stop.\n",
				"einride.saga.extend.book.v1beta1.Booking.Stop.unit_external_reference_ids": " The unit external reference id associated with this stop. E.g. for a pickup stop this will hold the units to pickup.\n",
				"einride.saga.extend.book.v1beta1.Booking.automation_rule":                  " The rule that is applied when a confirmed booking is accepted. Defaults to NONE.\n",
				"einride.saga.extend.book.v1beta1.Booking.create_time":                      " The creation timestamp of the booking.\n",
				"einride.saga.extend.book.v1beta1.Booking.delete_time":                      " The deletion timestamp of the booking. Set if the booking is deleted.\n",
				"einride.saga.extend.book.v1beta1.Booking.external_reference_id":            " External reference ID is a unique identifier that can be set by the client.\n",
				"einride.saga.extend.book.v1beta1.Booking.name":                             " The resource name of the booking.\n",
				"einride.saga.extend.book.v1beta1.Booking.sender":                           " Resource name of the sender of the booking.\n",
				"einride.saga.extend.book.v1beta1.Booking.service_type":                     " The type of service to create a booking for.\n",
				"einride.saga.extend.book.v1beta1.Booking.space":                            " Resource name of the space that owns the booking.\n",
				"einride.saga.extend.book.v1beta1.Booking.state":                            " The state of the booking. Set to PENDING upon creation.\n A booking is then either accepted or rejected.\n",
				"einride.saga.extend.book.v1beta1.Booking.stops":                            " The stops included in this booking.\n",
				"einride.saga.extend.book.v1beta1.Booking.type":                             " Booking type. Defaults to PROVISIONAL.\n",
				"einride.saga.extend.book.v1beta1.Booking.units":                            " The units included in this booking.\n",
				"einride.saga.extend.book.v1beta1.Booking.update_time":                      " The last update timestamp of the booking.\n Updated when create/update/delete operation is performed.\n",
				"einride.saga.extend.book.v1beta1.BookingService.CreateBooking":             " Create a booking in a space.\n\n This is an AIP standard [Create](https://google.aip.dev/133) method.\n Deprecated: Use CreateTour instead.\n",
				"einride.saga.extend.book.v1beta1.CreateBookingRequest.booking":             " The booking to create.\n",
				"einride.saga.extend.book.v1beta1.CreateBookingRequest.parent":              " The parent space in which to create the booking.\n",
				"einride.saga.extend.book.v1beta1.Unit.Length.unit":                         " Length unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Length.value":                        " Length value\n",
				"einride.saga.extend.book.v1beta1.Unit.Volume.unit":                         " Volume unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Volume.value":                        " Volume value\n",
				"einride.saga.extend.book.v1beta1.Unit.Weight.unit":                         " Weight unit\n",
				"einride.saga.extend.book.v1beta1.Unit.Weight.value":                        " Weight value\n",
				"einride.saga.extend.book.v1beta1.Unit.description":                         " Description of the unit\n",
				"einride.saga.extend.book.v1beta1.Unit.external_reference_id":               " External reference ID of the unit\n",
				"einride.saga.extend.book.v1beta1.Unit.height":                              " Unit height\n",
				"einride.saga.extend.book.v1beta1.Unit.length":                              " Unit length\n",
				"einride.saga.extend.book.v1beta1.Unit.tracking_id":                         " Unique unit tracking ID\n",
				"einride.saga.extend.book.v1beta1.Unit.type":                                " Type of unit\n",
				"einride.saga.extend.book.v1beta1.Unit.volume":                              " Unit volume\n",
				"einride.saga.extend.book.v1beta1.Unit.weight":                              " Unit weight\n",
				"einride.saga.extend.book.v1beta1.Unit.width":                               " Unit width\n",
				"google.protobuf.Timestamp.nanos":                                           " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                                         " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
				"google.type.LatLng.latitude":                                               " The latitude in degrees. It must be in the range [-90.0, +90.0].\n",
				"google.type.LatLng.longitude":                                              " The longitude in degrees. It must be in the range [-180.0, +180.0].\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_booking_service_proto.
				Services().ByName("BookingService").Methods().ByName("GetBooking"),
			&GetBookingRequest{},
			&Booking{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.BookingService.GetBooking": " Get a booking.\n\n This is an AIP standard [Get](https://google.aip.dev/131) method.\n Deprecated: Use GetTour instead.\n",
				"einride.saga.extend.book.v1beta1.GetBookingRequest.name":    " The resource name of the booking to retrieve.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_booking_service_proto.
				Services().ByName("BookingService").Methods().ByName("CreateTour"),
			&CreateTourRequest{},
			&Tour{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.Address.city":                           " City\n",
				"einride.saga.extend.book.v1beta1.Address.co":                             " Care of (C/O)\n",
				"einride.saga.extend.book.v1beta1.Address.contact_info":                   " Contact information\n",
				"einride.saga.extend.book.v1beta1.Address.display_name":                   " The displayed name of the address\n",
				"einride.saga.extend.book.v1beta1.Address.lat_lng":                        " The geographic location of the address\n",
				"einride.saga.extend.book.v1beta1.Address.line1":                          " Address line 1\n",
				"einride.saga.extend.book.v1beta1.Address.line2":                          " Address line 2\n",
				"einride.saga.extend.book.v1beta1.Address.postal_code":                    " Postal code\n",
				"einride.saga.extend.book.v1beta1.Address.recipient":                      " Recipient\n",
				"einride.saga.extend.book.v1beta1.Address.region_code":                    " Region code (Unicode CLDR region code)\n https://cldr.unicode.org/\n",
				"einride.saga.extend.book.v1beta1.Address.state_code":                     " State code\n",
				"einride.saga.extend.book.v1beta1.BookingService.CreateTour":              " Create a truck tour booking in a space.\n\n This is an AIP standard [Create](https://google.aip.dev/133) method.\n",
				"einride.saga.extend.book.v1beta1.CreateTourRequest.parent":               " The parent space in which to create the tour.\n",
				"einride.saga.extend.book.v1beta1.CreateTourRequest.tour":                 " The tour to create.\n",
				"einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry.key":          "",
				"einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry.value":        "",
				"einride.saga.extend.book.v1beta1.Shipment.annotations":                   " Annotations for the shipment\n",
				"einride.saga.extend.book.v1beta1.Shipment.booking":                       " Resource name of the booking this shipment relates to.\n",
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
				"einride.saga.extend.book.v1beta1.Shipment.tracking_id":                   " A generated tracking ID for this shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.units":                         " Shipment units\n",
				"einride.saga.extend.book.v1beta1.Shipment.update_time":                   " The last update timestamp of the shipment.\n Updated when create/update/delete operation is performed.\n",
				"einride.saga.extend.book.v1beta1.Shipment.vehicle":                       " Vehicle information for the shipment\n",
				"einride.saga.extend.book.v1beta1.Tour.AnnotationsEntry.key":              "",
				"einride.saga.extend.book.v1beta1.Tour.AnnotationsEntry.value":            "",
				"einride.saga.extend.book.v1beta1.Tour.Stop.address":                      " The address of this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.instructions":                 " Instructions for this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.requested_end_time":           " The latest time to depart this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.requested_start_time":         " The earliest time to arrive this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.shipments":                    " The shipments that are associated with this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.stop_type":                    " The type of stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.unit_external_reference_ids":  " The unit external reference id associated with this stop. E.g. for a pickup stop this will hold the units to pickup.\n",
				"einride.saga.extend.book.v1beta1.Tour.annotations":                       " Supplementary annotation metadata for the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.automation_rule":                   " The rule that is applied when a confirmed tour is accepted. Defaults to CREATE_SHIPMENTS.\n",
				"einride.saga.extend.book.v1beta1.Tour.create_time":                       " The creation timestamp of the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.delete_time":                       " The deletion timestamp of the tour. This is set if the tour is deleted.\n",
				"einride.saga.extend.book.v1beta1.Tour.external_reference_id":             " External reference ID is a unique identifier that can be set by the client.\n",
				"einride.saga.extend.book.v1beta1.Tour.name":                              " The resource name of the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.sender":                            " Resource name of the sender of the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.service_type":                      " The type of service to create a tour for.\n",
				"einride.saga.extend.book.v1beta1.Tour.shipments":                         " The shipments that are associated with the tour. These are created when a tour is CONFIRMED and ACCEPTED.\n",
				"einride.saga.extend.book.v1beta1.Tour.space":                             " Resource name of the space that owns the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.state":                             " The state of the tour. Set to PENDING upon creation.\n A tour is then either accepted or rejected.\n",
				"einride.saga.extend.book.v1beta1.Tour.stops":                             " The ordered list of stops included in this tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.type":                              " The type of tour, either PROVISIONAL or CONFIRMED. Defaults to PROVISIONAL.\n",
				"einride.saga.extend.book.v1beta1.Tour.units":                             " The units included in this tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.update_time":                       " The last update timestamp of the tour.\n Updated when create/update/delete operation is performed.\n",
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
			File_einride_saga_extend_book_v1beta1_booking_service_proto.
				Services().ByName("BookingService").Methods().ByName("GetTour"),
			&GetTourRequest{},
			&Tour{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.BookingService.GetTour": " Get an existing truck tour booking.\n\n This is an AIP standard [Get](https://google.aip.dev/131) method.\n",
				"einride.saga.extend.book.v1beta1.GetTourRequest.name":    " The resource name of the tour to retrieve.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_booking_service_proto.
				Services().ByName("BookingService").Methods().ByName("ConfirmTour"),
			&ConfirmTourRequest{},
			&Tour{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.BookingService.ConfirmTour": " Confirm a Provisional tour.\n\n Reconfirming a tour that is already confirmed will return an InvalidArgument Error.\n When a tour has been accepted by Saga and confirmed by the user, Shipments will be created.\n",
				"einride.saga.extend.book.v1beta1.ConfirmTourRequest.name":    " The resource name of the tour to confirm.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_booking_service_proto.
				Services().ByName("BookingService").Methods().ByName("UpdateTour"),
			&UpdateTourRequest{},
			&Tour{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.Address.city":                           " City\n",
				"einride.saga.extend.book.v1beta1.Address.co":                             " Care of (C/O)\n",
				"einride.saga.extend.book.v1beta1.Address.contact_info":                   " Contact information\n",
				"einride.saga.extend.book.v1beta1.Address.display_name":                   " The displayed name of the address\n",
				"einride.saga.extend.book.v1beta1.Address.lat_lng":                        " The geographic location of the address\n",
				"einride.saga.extend.book.v1beta1.Address.line1":                          " Address line 1\n",
				"einride.saga.extend.book.v1beta1.Address.line2":                          " Address line 2\n",
				"einride.saga.extend.book.v1beta1.Address.postal_code":                    " Postal code\n",
				"einride.saga.extend.book.v1beta1.Address.recipient":                      " Recipient\n",
				"einride.saga.extend.book.v1beta1.Address.region_code":                    " Region code (Unicode CLDR region code)\n https://cldr.unicode.org/\n",
				"einride.saga.extend.book.v1beta1.Address.state_code":                     " State code\n",
				"einride.saga.extend.book.v1beta1.BookingService.UpdateTour":              " Update a tour.\n\n See: https://google.aip.dev/134 (Standard methods: Update).\n",
				"einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry.key":          "",
				"einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry.value":        "",
				"einride.saga.extend.book.v1beta1.Shipment.annotations":                   " Annotations for the shipment\n",
				"einride.saga.extend.book.v1beta1.Shipment.booking":                       " Resource name of the booking this shipment relates to.\n",
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
				"einride.saga.extend.book.v1beta1.Shipment.tracking_id":                   " A generated tracking ID for this shipment.\n",
				"einride.saga.extend.book.v1beta1.Shipment.units":                         " Shipment units\n",
				"einride.saga.extend.book.v1beta1.Shipment.update_time":                   " The last update timestamp of the shipment.\n Updated when create/update/delete operation is performed.\n",
				"einride.saga.extend.book.v1beta1.Shipment.vehicle":                       " Vehicle information for the shipment\n",
				"einride.saga.extend.book.v1beta1.Tour.AnnotationsEntry.key":              "",
				"einride.saga.extend.book.v1beta1.Tour.AnnotationsEntry.value":            "",
				"einride.saga.extend.book.v1beta1.Tour.Stop.address":                      " The address of this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.instructions":                 " Instructions for this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.requested_end_time":           " The latest time to depart this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.requested_start_time":         " The earliest time to arrive this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.shipments":                    " The shipments that are associated with this stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.stop_type":                    " The type of stop.\n",
				"einride.saga.extend.book.v1beta1.Tour.Stop.unit_external_reference_ids":  " The unit external reference id associated with this stop. E.g. for a pickup stop this will hold the units to pickup.\n",
				"einride.saga.extend.book.v1beta1.Tour.annotations":                       " Supplementary annotation metadata for the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.automation_rule":                   " The rule that is applied when a confirmed tour is accepted. Defaults to CREATE_SHIPMENTS.\n",
				"einride.saga.extend.book.v1beta1.Tour.create_time":                       " The creation timestamp of the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.delete_time":                       " The deletion timestamp of the tour. This is set if the tour is deleted.\n",
				"einride.saga.extend.book.v1beta1.Tour.external_reference_id":             " External reference ID is a unique identifier that can be set by the client.\n",
				"einride.saga.extend.book.v1beta1.Tour.name":                              " The resource name of the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.sender":                            " Resource name of the sender of the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.service_type":                      " The type of service to create a tour for.\n",
				"einride.saga.extend.book.v1beta1.Tour.shipments":                         " The shipments that are associated with the tour. These are created when a tour is CONFIRMED and ACCEPTED.\n",
				"einride.saga.extend.book.v1beta1.Tour.space":                             " Resource name of the space that owns the tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.state":                             " The state of the tour. Set to PENDING upon creation.\n A tour is then either accepted or rejected.\n",
				"einride.saga.extend.book.v1beta1.Tour.stops":                             " The ordered list of stops included in this tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.type":                              " The type of tour, either PROVISIONAL or CONFIRMED. Defaults to PROVISIONAL.\n",
				"einride.saga.extend.book.v1beta1.Tour.units":                             " The units included in this tour.\n",
				"einride.saga.extend.book.v1beta1.Tour.update_time":                       " The last update timestamp of the tour.\n Updated when create/update/delete operation is performed.\n",
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
				"einride.saga.extend.book.v1beta1.UpdateTourRequest.tour":                 " The resource which replaces the current resource.\n",
				"einride.saga.extend.book.v1beta1.UpdateTourRequest.update_mask":          " The update mask applies to the tour.\n For the `FieldMask` definition, see:\n https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask\n",
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

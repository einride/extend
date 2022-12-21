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
				"einride.saga.extend.book.v1beta1.Booking.tracking_id":                      " Tracking ID for the booking.\n",
				"einride.saga.extend.book.v1beta1.Booking.type":                             " Booking type. Defaults to PROVISIONAL.\n",
				"einride.saga.extend.book.v1beta1.Booking.units":                            " The units included in this booking.\n",
				"einride.saga.extend.book.v1beta1.Booking.update_time":                      " The last update timestamp of the booking.\n Updated when create/update/delete operation is performed.\n",
				"einride.saga.extend.book.v1beta1.BookingService.CreateBooking":             " Create a booking in a space.\n\n This is an AIP standard [Create](https://google.aip.dev/133) method.\n",
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
				"einride.saga.extend.book.v1beta1.BookingService.GetBooking": " Get a booking.\n\n This is an AIP standard [Get](https://google.aip.dev/131) method.\n",
				"einride.saga.extend.book.v1beta1.GetBookingRequest.name":    " The resource name of the booking to retrieve.\n",
			},
		),
	)
}

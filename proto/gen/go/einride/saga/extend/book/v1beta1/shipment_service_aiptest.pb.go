// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package bookv1beta1

import (
	context "context"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
	time "time"
)

type ShipmentServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server ShipmentServiceServer
}

func (fx ShipmentServiceTestSuite) TestShipment(ctx context.Context, options ShipmentTestSuiteConfig) {
	fx.T.Run("Shipment", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type ShipmentTestSuiteConfig struct {
	ctx        context.Context
	service    ShipmentServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Shipment
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *ShipmentTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
}

func (fx *ShipmentTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:   "",
			Shipment: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:   "invalid resource name",
			Shipment: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:   parent,
			Shipment: fx.Create(parent),
		})
		assert.NilError(t, err)
		assert.Check(t, time.Since(msg.CreateTime.AsTime()) < time.Second)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:   parent,
			Shipment: fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetShipment(fx.ctx, &GetShipmentRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".sender", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("sender")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.recipient", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("recipient")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.line1", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("line1")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.postal_code", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("postal_code")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.city", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("city")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.region_code", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("region_code")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".delivery_address", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("delivery_address")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".units", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("units")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	// The method should fail with InvalidArgument if the resource has any
	// resource references and they are invalid.
	t.Run("resource references", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".sender", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			container.Sender = "invalid resource name"
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".booking", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			container.Booking = "invalid resource name"
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:   parent,
				Shipment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *ShipmentTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetShipment(fx.ctx, &GetShipmentRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetShipment(fx.ctx, &GetShipmentRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetShipment(fx.ctx, &GetShipmentRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.GetShipment(fx.ctx, &GetShipmentRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetShipment(fx.ctx, &GetShipmentRequest{
			Name: "spaces/-/shipments/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ShipmentTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *ShipmentTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *ShipmentTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *ShipmentTestSuiteConfig) create(t *testing.T, parent string) *Shipment {
	t.Helper()
	created, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
		Parent:   parent,
		Shipment: fx.Create(parent),
	})
	assert.NilError(t, err)
	return created
}

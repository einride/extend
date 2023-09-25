// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package bookv1beta1

import (
	context "context"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
	// IDGenerator should return a valid and unique ID to use in the Create call.
	// If non-nil, this function will be called to set the ID on all Create calls.
	// If the ID field is required, tests will fail if this is nil.
	IDGenerator func() string
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Shipment
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *ShipmentTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
}

func (fx *ShipmentTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		userSetID := ""
		if fx.IDGenerator != nil {
			userSetID = fx.IDGenerator()
		}
		_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:     "",
			Shipment:   fx.Create(fx.nextParent(t, false)),
			ShipmentId: userSetID,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		userSetID := ""
		if fx.IDGenerator != nil {
			userSetID = fx.IDGenerator()
		}
		_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:     "invalid resource name",
			Shipment:   fx.Create(fx.nextParent(t, false)),
			ShipmentId: userSetID,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		userSetID := ""
		if fx.IDGenerator != nil {
			userSetID = fx.IDGenerator()
		}
		msg, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:     parent,
			Shipment:   fx.Create(parent),
			ShipmentId: userSetID,
		})
		assert.NilError(t, err)
		assert.Check(t, time.Since(msg.CreateTime.AsTime()) < time.Second)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		userSetID := ""
		if fx.IDGenerator != nil {
			userSetID = fx.IDGenerator()
		}
		msg, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:     parent,
			Shipment:   fx.Create(parent),
			ShipmentId: userSetID,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetShipment(fx.ctx, &GetShipmentRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// If method support user settable IDs, when set the resource should
	// be returned with the provided ID.
	t.Run("user settable id", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:     parent,
			Shipment:   fx.Create(parent),
			ShipmentId: "usersetid",
		})
		assert.NilError(t, err)
		assert.Check(t, strings.HasSuffix(msg.GetName(), "usersetid"))
	})

	// If method support user settable IDs and the same ID is reused
	// the method should return AlreadyExists.
	t.Run("already exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:     parent,
			Shipment:   fx.Create(parent),
			ShipmentId: "alreadyexists",
		})
		assert.NilError(t, err)
		_, err = fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
			Parent:     parent,
			Shipment:   fx.Create(parent),
			ShipmentId: "alreadyexists",
		})
		assert.Equal(t, codes.AlreadyExists, status.Code(err), err)
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".vehicle.reference_id", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetVehicle()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("reference_id")
			container.ProtoReflect().Clear(fd)
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".tour", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			container.Tour = "invalid resource name"
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
				Parent:     parent,
				Shipment:   msg,
				ShipmentId: userSetID,
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

func (fx *ShipmentTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
			Shipment: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
			Shipment: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field update_time should be updated when the resource is updated.
	t.Run("update time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
			Shipment: created,
		})
		assert.NilError(t, err)
		assert.Check(t, updated.UpdateTime.AsTime().After(created.UpdateTime.AsTime()))
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
			Shipment: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetShipment(fx.ctx, &GetShipmentRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	// The field create_time should be preserved when a '*'-update mask is used.
	t.Run("preserve create_time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		originalCreateTime := created.CreateTime
		updated, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
			Shipment: created,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"*",
				},
			},
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, originalCreateTime, updated.CreateTime, protocmp.Transform())
	})

	parent := fx.nextParent(t, false)
	created := fx.create(t, parent)
	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = created.Name + "notfound"
		_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
			Shipment: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
			Shipment: created,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"invalid_field_xyz",
				},
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if any required field is missing
	// when called with '*' update_mask.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".sender", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("sender")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.recipient", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("recipient")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.line1", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("line1")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.postal_code", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("postal_code")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.city", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("city")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".pickup_address.region_code", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg.GetPickupAddress()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("region_code")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".delivery_address", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("delivery_address")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".units", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("units")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".vehicle.reference_id", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipment)
			container := msg.GetVehicle()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("reference_id")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipment(fx.ctx, &UpdateShipmentRequest{
				Shipment: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
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
	userSetID := ""
	if fx.IDGenerator != nil {
		userSetID = fx.IDGenerator()
	}
	created, err := fx.service.CreateShipment(fx.ctx, &CreateShipmentRequest{
		Parent:     parent,
		Shipment:   fx.Create(parent),
		ShipmentId: userSetID,
	})
	assert.NilError(t, err)
	return created
}

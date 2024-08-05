// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package bookv1beta1

import (
	context "context"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
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

type BookingServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server BookingServiceServer
}

func (fx BookingServiceTestSuite) TestTour(ctx context.Context, options BookingServiceTourTestSuiteConfig) {
	fx.T.Run("Tour", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type BookingServiceTourTestSuiteConfig struct {
	ctx        context.Context
	service    BookingServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Tour
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Tour
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *BookingServiceTourTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
	t.Run("Search", fx.testSearch)
}

func (fx *BookingServiceTourTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
			Parent: "",
			Tour:   fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
			Parent: "invalid resource name",
			Tour:   fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		beforeCreate := time.Now()
		msg, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
			Parent: parent,
			Tour:   fx.Create(parent),
		})
		assert.NilError(t, err)
		assert.Check(t, msg.CreateTime != nil)
		assert.Check(t, !msg.CreateTime.AsTime().IsZero())
		assert.Check(t, msg.CreateTime.AsTime().After(beforeCreate), "msg.CreateTime (%v) is not after beforeCreate (%v)", msg.CreateTime.AsTime(), beforeCreate)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
			Parent: parent,
			Tour:   fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetTour(fx.ctx, &GetTourRequest{
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
			_, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
				Parent: parent,
				Tour:   msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".service_type", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("service_type")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
				Parent: parent,
				Tour:   msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".preliminary_shipments", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("preliminary_shipments")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
				Parent: parent,
				Tour:   msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".stops", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("stops")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
				Parent: parent,
				Tour:   msg,
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
			_, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
				Parent: parent,
				Tour:   msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *BookingServiceTourTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetTour(fx.ctx, &GetTourRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetTour(fx.ctx, &GetTourRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetTour(fx.ctx, &GetTourRequest{
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
		_, err := fx.service.GetTour(fx.ctx, &GetTourRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetTour(fx.ctx, &GetTourRequest{
			Name: "spaces/-/tours/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *BookingServiceTourTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
			Tour: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
			Tour: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field update_time should be updated when the resource is updated.
	t.Run("update time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
			Tour: created,
		})
		assert.NilError(t, err)
		assert.Check(t, updated.UpdateTime.AsTime().After(created.UpdateTime.AsTime()))
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
			Tour: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetTour(fx.ctx, &GetTourRequest{
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
		updated, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
			Tour: created,
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
		_, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
			Tour: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
			Tour: created,
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
			msg := proto.Clone(created).(*Tour)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("sender")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
				Tour: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".service_type", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Tour)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("service_type")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
				Tour: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".preliminary_shipments", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Tour)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("preliminary_shipments")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
				Tour: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".stops", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Tour)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("stops")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateTour(fx.ctx, &UpdateTourRequest{
				Tour: msg,
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

func (fx *BookingServiceTourTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListTours(fx.ctx, &ListToursRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListTours(fx.ctx, &ListToursRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListTours(fx.ctx, &ListToursRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Tour, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListTours(fx.ctx, &ListToursRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Tours,
			cmpopts.SortSlices(func(a, b *Tour) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListTours(fx.ctx, &ListToursRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListTours(fx.ctx, &ListToursRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Tour, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListTours(fx.ctx, &ListToursRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Tours))
			msgs = append(msgs, response.Tours...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Tour) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *BookingServiceTourTestSuiteConfig) testSearch(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.SearchTours(fx.ctx, &SearchToursRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.SearchTours(fx.ctx, &SearchToursRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.SearchTours(fx.ctx, &SearchToursRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Tour, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.SearchTours(fx.ctx, &SearchToursRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Tours,
			cmpopts.SortSlices(func(a, b *Tour) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.SearchTours(fx.ctx, &SearchToursRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.SearchTours(fx.ctx, &SearchToursRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Searching resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Tour, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.SearchTours(fx.ctx, &SearchToursRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Tours))
			msgs = append(msgs, response.Tours...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Tour) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *BookingServiceTourTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *BookingServiceTourTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *BookingServiceTourTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *BookingServiceTourTestSuiteConfig) create(t *testing.T, parent string) *Tour {
	t.Helper()
	created, err := fx.service.CreateTour(fx.ctx, &CreateTourRequest{
		Parent: parent,
		Tour:   fx.Create(parent),
	})
	assert.NilError(t, err)
	return created
}

# Einride Extend APIs

Based on REST principles, with support for gRPC, the Einride Extend APIs enable integrations and extensions to Einride's Saga platform.

## API design

The Extend APIs are designed using the [AIP](https://aip.dev) design framework, which means that they are resource-oriented and work like standard [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) APIs.

## API specification

The APIs are specified and documented using the [Protocol Buffer](https://developers.google.com/protocol-buffers/docs/proto3) interface definition language.

## Protocols

Saga's APIs support both gRPC and HTTP protocols.

### gRPC

[gRPC](https://grpc.io) is a high performance, open source universal RPC framework that works across programming languages and platforms.

To call Saga's APIs via gRPC, you can use the protobuf source files to generate a client for [your preferred language](https://grpc.io/docs/languages).

The Saga APIs are also available via the [Buf Schema Registry](buf.build/einride/saga), which is a registry for openly sharing gRPC APIs between organizations. You can use one of the [remote generation plugins](https://docs.buf.build/tour/use-remote-generation) provided by Buf.

### HTTP

Saga's APIs are also available as RESTful HTTP endpoints, via automatic [gRPC to HTTP transcoding](https://google.aip.dev/127).

See the [openapiv2](https://github.com/einride/extend/openapiv2) directory for the latest OpenAPI specifications, or use the API documentation at [extend.saga.einride.tech](https://extend.saga.einride.tech) where you can also try out the REST APIs and make example requests.

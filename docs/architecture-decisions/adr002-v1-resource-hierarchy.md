______________________________________________________________________

| id          | title                         | description                           |
| ----------- | ----------------------------- | ------------------------------------- |
| adrs-adr002 | ADR002: V1 Resource Hierarchy | Resource hierarchy used in the V1 API |

______________________________________________________________________

## Context

The Extend APIs are designed using the [AIP](https://aip.dev) design framework,
which means that they are resource-oriented and work like standard
[REST](https://en.wikipedia.org/wiki/Representational_state_transfer) APIs, as
is documented in our [API Design docs](../apis.md).

The current resources under `book/v1beta1` all use spaces as their parent, such
as the
[Shipment resource](../../proto/einride/saga/extend/book/v1beta1/shipment.proto)
which is referenced with the `spaces/{space}/shipments/{shipment}` name format.
This results in the REST based URLs to include the space name, such as the
`GET /v1beta1/spaces/{space}/shipment/{shipment}` endpoint to get a shipment.

We have received feedback from API users that they have found this confusing,
and that it complicate integrations by requiring URLs with dynamic parts, even
for creating resources.

## Decision

Resources in the V1 API will be owned by an organization, which an API user is
always connected to.

The V1 APIs will by-default not require spaces or organization as part of the
resource names, and thus not part of the URLs, but will instead be inferred from
the user transparently.

## Consequences

No need to specify space or organization in the V1 APIs, leading to slightly
different but simpler REST based URLs.

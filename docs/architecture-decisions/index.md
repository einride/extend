______________________________________________________________________

| id            | title                               | sidebar_label | description                                     |
| ------------- | ----------------------------------- | ------------- | ----------------------------------------------- |
| adrs-overview | Architecture Decision Records (ADR) | Overview      | Overview of Architecture Decision Records (ADR) |

______________________________________________________________________

The architecture decisions made for the Einride Extend public API are kept here.
For more information about ADRs, when to write them, and why, please see
[this blog post](https://engineering.atspotify.com/2020/04/14/when-should-i-write-an-architecture-decision-record/).

Records are never deleted but can be marked as superseded by new decisions or
deprecated.

Records should be stored under the `architecture-decisions` directory.

## Contributing

### Creating an ADR

- Copy `docs/architecture-decisions/adr000-template.md`
  to`docs/architecture-decisions/adr000-my-decision.md` (my-decision should be
  descriptive. Do not assign an ADR number.)
- Fill in the ADR following the guidelines in the template
- Submit a pull request
- Address and integrate feedback from the community
- Eventually, assign a number
- Merge the pull request

## Superseding an ADR

If an ADR supersedes an older ADR then the status of the older ADR is changed to
"superseded by ADR-XXXX", and links to the new ADR.

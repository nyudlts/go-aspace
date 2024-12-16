## CHANGELOG

#### v0.7.1
  - remove debug print from `NewClientFromCreds()`
  - tidy go mod files (removed unused libxml2 package)
  - tweak range statement per linter recommendation
  - add missing error check
  - comment out unused code
  - add error check to (a *ADClient)GetAgent()
  - update CHANGELOG, bump to v0.7.1

#### v0.7.0
  - add `func (a *ASClient) FindArchivalObjectsByID(...)([]string, error)`
    that allows you to look up archival objects by the `ref_id` or `component_id` fields

#### v0.6.1
  - bug fix: rewrite `func (wor WorkOrderRow) String()` to use
    `encoding/csv` to handle string escapes

#### v0.6.0
  - add `...FromURI()` functions:
    - add `GetArchivalObjectFromURI()`
    - add `GetDigitalObjectFromURI()`
    - add `DeleteDigitalObjectFromURI()`
    - add `GetDigitalObjectIDsForArchivalObjectFromURI()`
  - refactor `...Object()` functions to use `...FromURI()` functions
    - refactor `GetArchivalObject()` to use `GetArchivalObjectFromURI()`
    - refactor `GetDigitalObject()` to use `GetDigitalObjectFromURI()`
    - refactor `DeleteDigitalObject()` to use `DeleteDigitalObjectFromURI()`
  - add helper type and function
    - add `CreateOrUpdateResponse` type
    - add `ParseCreateOrUpdateResponse()`
  - force `bool` key/value pairs to always be sent in marshaled JSON for selected types
    - remove `omitempty` option from `DigitalObject` type `bool` JSON tags
    - remove `omitempty` option from `FileVersion` type `bool` JSON tags

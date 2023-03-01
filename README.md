To find how you can declare resources see [examples folder](./examples/)

### Limitations:
- Sources get created first, but an aggregation has to run before the rest of the plan can complete successfully because the entitlement data lookups will fail until the aggregation has pulled the entitlements from the source into IdentityNow.

- After creating the source, you also need to go into the UI and press the "Test Connection" button to verify the source. This unlocks the ability to apply `identitynow_account_schema_attribute` and `identitynow_account_aggregation_schedule`.

- Password policies can be created, but there is a bug in Idn that makes the association to the source not work. For now, you have to go into the UI and make the association. Sailpoint ticket: https://support.sailpoint.com/hc/en-us/requests/82917

- Create/enable/disable profiles cannot be managed with Terraform yet. The API for them is highly unusual and not amenable to automation.

- Due to a bug in IdentityNow, Encrypted field in ConnectorAttributes block cannot be left null.

# Development
Edit the Go files that make up the provider, and rebuild the provider.

```bash
./scripts/build.sh
```

This script places the provider binary in an implied local mirror directory ($HOME/.terraform.d/plugins/). See build.sh
for more comments about ensuring that Terraform uses the local mirror rather than searching the remote registry. 

# Testing the Provider

In order to test the provider, you can simply run `make test`.
```sh
$ make test
```
In order to run the full suite of Acceptance test identitynow url, client id and secret, owner name and id, cluster name and id are needed to make the API call to create IdentityNow source for test.

To run acceptance tests, first you need to update the `script/gotestacc_vars.sh` with above variables values and then simply run `make testacc`.

```sh
$ make testacc
```

package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func flattenIdentity(d *schema.ResourceData, in *Identity) error {
	if in == nil {
		return nil
	}
	d.SetId(in.ID)
	d.Set("alias", in.Alias)
	d.Set("name", in.Name)
	d.Set("description", in.Description)
	d.Set("external_id", in.ExternalID)
	d.Set("date_created", in.DateCreated)
	d.Set("last_updated", in.LastUpdated)
	d.Set("email", in.Email)
	d.Set("status", in.Status)
	d.Set("enabled", in.Enabled)
	d.Set("uid", in.UID)
	d.Set("uuid", in.UUID)
	d.Set("pending", in.Pending)
	d.Set("encryption_key", in.EncryptionKey)
	d.Set("encryption_check", in.EncryptionCheck)
	d.Set("password_reset_since_last_login", in.PasswordResetSinceLastLogin)
	d.Set("usage_cert_attested", in.UsageCertAttested)
	d.Set("alt_auth_via_integration_data", in.AltAuthViaIntegrationData)
	d.Set("kba_answers", in.KbaAnswers)
	d.Set("disable_password_reset", in.DisablePasswordReset)
	d.Set("pta_source_id", in.PtaSourceID)
	d.Set("supports_password_push", in.SupportsPasswordPush)
	d.Set("role", in.Role)
	d.Set("alt_phone", in.AltPhone)
	d.Set("alt_email", in.AltEmail)
	d.Set("identity_flags", in.IdentityFlags)
	d.Set("alt_auth_via", in.AltAuthVia)
	d.Set("phone", in.Phone)
	d.Set("employee_number", in.EmployeeNumber)
	d.Set("attributes", in.Attributes)

	return nil
}

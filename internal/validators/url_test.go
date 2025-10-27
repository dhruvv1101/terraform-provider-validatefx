package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func runURLValidation(v validator.String, input types.String) diag.Diagnostics {
	resp := validator.StringResponse{}
	req := validator.StringRequest{
		ConfigValue: input,
	}

	v.ValidateString(context.Background(), req, &resp)
	return resp.Diagnostics
}

func TestURLValidatorValid(t *testing.T) {
	validURLs := []string{
		"https://example.com",
		"http://google.com",
		"https://sub.domain.com/path?query=1",
	}

	v := URL()

	for _, url := range validURLs {
		diags := runURLValidation(v, types.StringValue(url))
		if diags.HasError() {
			t.Errorf("Expected valid URL, got error for: %s", url)
		}
	}
}

func TestURLValidatorInvalid(t *testing.T) {
	invalidURLs := []string{
		"example.com",
		"http:/invalid.com",
		"https://",
		"invalid",
		"ftp://",
	}
	v := URL()
	for _, url := range invalidURLs {
		diags := runURLValidation(v, types.StringValue(url))
		if !diags.HasError() {
			t.Errorf("Expected error for invalid URL, got none for: %s", url)
		}
	}
}

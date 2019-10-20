package openpay

import "testing"

func TestAPIError_Error(t *testing.T) {
	type fields struct {
		Category    string
		Code        uint
		HTTPCode    uint
		Description string
		RequestID   string
		FraudRules  []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "OK",
			fields: fields{
				Code:        400,
				Category:    "test_category",
				Description: "test_description",
			},
			want: "400: test_category - test_description",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &APIError{
				Category:    tt.fields.Category,
				Code:        tt.fields.Code,
				Description: tt.fields.Description,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("APIError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

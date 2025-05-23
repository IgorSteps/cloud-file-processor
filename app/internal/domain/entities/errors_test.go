package entities_test

import (
	"app/internal/domain/entities"
	"testing"
)

func TestInvalidError_Error(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "simple check",
			fields: fields{
				Message: "hi",
			},
			want: "hi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &entities.InvalidError{
				Message: tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("InvalidError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

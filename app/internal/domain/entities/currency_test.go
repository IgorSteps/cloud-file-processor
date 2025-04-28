package entities_test

import (
	"app/internal/domain/entities"
	"testing"
)

func TestCurrency_String(t *testing.T) {
	tests := []struct {
		name    string
		c       entities.Currency
		want    string
		wantErr bool
	}{
		{
			name: "get usd",
			c:    0,
			want: "USD",
		},
		{
			name: "get eur",
			c:    1,
			want: "EUR",
		},
		{
			name: "get gbr",
			c:    2,
			want: "GBR",
		},
		{
			name:    "unknown",
			c:       100,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("Currency.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Currency.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

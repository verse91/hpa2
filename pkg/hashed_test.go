package pkg

import (
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		name       string
		password   string
		verifyPass string
		wantErr    bool
		matched    bool
	}{
		{
			name:       "Correct password",
			password:   "golang.Abc@123",
			verifyPass: "golang.Abc@123",
			wantErr:    false,
			matched:    true,
		},
		{
			name:       "Incorrect password",
			password:   "1234",
			verifyPass: "1235",
			wantErr:    false,
			matched:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := Hash(tt.password, CoffeeSalt())
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			matched, err := verifyHashed(tt.verifyPass, hashed)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyHashed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if matched != tt.matched {
				t.Errorf("Expected verify result: %t, got %t", tt.matched, matched)
			}
		})
	}
}

func TestVerifyHash(t *testing.T) {
	// First generate a valid hash to use in tests
	validHash, err := Hash("golang.Abc@123", CoffeeSalt())
	if err != nil {
		t.Fatalf("Failed to generate valid hash for testing: %v", err)
	}

	tests := []struct {
		name           string
		password       string
		hashedPassword string
		want           bool
		wantErr        bool
	}{
		{
			name:           "Correct password",
			password:       "golang.Abc@123",
			hashedPassword: validHash,
			want:           true,
			wantErr:        false,
		},
		{
			name:           "Incorrect password",
			password:       "wrongpassword",
			hashedPassword: validHash,
			want:           false,
			wantErr:        false,
		},
		{
			name:           "Invalid hash format",
			password:       "golang.Abc@123",
			hashedPassword: "invalidhash",
			want:           false,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched, err := verifyHashed(tt.password, tt.hashedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyHashed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if matched != tt.want {
				t.Errorf("verifyHashed() got = %v, want %v", matched, tt.want)
			}
		})
	}
}

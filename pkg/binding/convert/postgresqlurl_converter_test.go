package convert

import (
	"github.com/RHEcosystemAppEng/sbo-go-library/internal/fileconfig"
	"testing"
)

func TestPostgreSQLUrlConverter_Convert(t *testing.T) {
	type args struct {
		binding fileconfig.ServiceBinding
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Correct conversion w/o options",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password",
						"srv":      "true",
						"options":  "",
						"database": "local-db",
						"sslmode":  "disable",
					},
					BindingName: "local",
				},
			},
			want: "postgresql://a-db-user:password@example.com:10011/local-db?sslmode=disable",
		},
		{
			name: "Correct conversion with options - postgre",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password",
						"srv":      "true",
						"options":  "search=somecloudhost&geqo=on",
						"database": "local-db",
						"sslmode":  "disable",
					},
					BindingName: "local",
				},
			},
			want: "postgresql://a-db-user:password@example.com:10011/local-db?sslmode=disable&options=-c%20search=somecloudhost%20-c%20geqo=on",
		},
		{
			name: "Correct conversion with options",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "CockroachDB",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password",
						"srv":      "true",
						"options":  "--cluster=somecloudhost&statement=9999",
						"database": "local-db",
						"sslmode":  "disable",
					},
					BindingName: "local",
				},
			},
			want: "postgresql://a-db-user:password@example.com:10011/local-db?sslmode=disable&options=--cluster%3Dsomecloudhost%20-c%20statement=9999",
		},
		{
			name: "Correct conversion with sslmode enabled",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "CockroachDB",
					Properties: map[string]string{
						"host":        "example.com:10011",
						"username":    "a-db-user",
						"password":    "password",
						"srv":         "true",
						"options":     "--cluster=somecloudhost",
						"database":    "local-db",
						"sslmode":     "verify-full",
						"sslrootcert": "root.ca",
					},
					BindingName: "local",
				},
			},
			want: "postgresql://a-db-user:password@example.com:10011/local-db?sslmode=verify-full&sslrootcert=/bindings/local/root.ca&options=--cluster%3Dsomecloudhost",
		},
		{
			name: "Correct connection string returned without password",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"srv":      "true",
						"options":  "",
						"sslmode":  "disable",
						"database": "local-db",
					},
				},
			},
			want: "postgresql://a-db-user@example.com:10011/local-db?sslmode=disable",
		},
		{
			name: "Correct connection string returned without host",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "",
						"username": "a-db-user",
						"srv":      "true",
						"sslmode":  "disable",
						"database": "local-db",
					},
				},
			},
			want: "postgresql://",
		},
		{
			name: "Correct connection string returned escaping password",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "pwd//'end",
						"srv":      "true",
						"sslmode":  "disable",
						"database": "local-db",
					},
				},
			},
			want: "postgresql://a-db-user:pwd%2F%2F%27end@example.com:10011/local-db?sslmode=disable",
		},
		{
			name: "Invalid_or_incomplete options are ignored",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password'",
						"srv":      "true",
						"options":  "option1=value1&option2",
						"database": "local-db",
					},
				},
			},
			want: "postgresql://a-db-user:password%27@example.com:10011/local-db?options=-c%20option1=value1",
		},
		{
			name: "Options contain an invalid option - with no key or value",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password'",
						"srv":      "true",
						"options":  "=value1&option2=",
						"database": "local-db",
					},
				},
			},
			want: "postgresql://a-db-user:password%27@example.com:10011/local-db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &PostgreSQLUrlConverter{}
			if got := f.Convert(tt.args.binding); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

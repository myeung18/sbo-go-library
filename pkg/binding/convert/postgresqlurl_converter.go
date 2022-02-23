package convert

import (
	"github.com/RHEcosystemAppEng/sbo-go-library/internal/fileconfig"
	"net/url"
	"strings"
)

const (
	encodedSpace            = "%20"
	optionFlag              = "-c" + encodedSpace
	optionsSeparator        = "&" // separator for each option
	optionKeyValueSeparator = "=" // separator for key/value in each option
)

var escaper = strings.NewReplacer(` `, `\ `, `'`, `\'`, `\`, `\\`)

type PostgreSQLUrlConverter struct{}

/*
	https://www.postgresql.org/docs/13/libpq-connect.html (connection string)
	urlExample := "postgres://username:password@localhost:5432/database_name"
	e.g.:
	postgresql://
	postgresql://localhost
	postgresql://localhost:5433
	postgresql://localhost/mydb
	postgresql://user@localhost
	postgresql://user:secret@localhost
	postgresql://other@localhost/otherdb?connect_timeout=10&application_name=myapp
	postgresql://host1:123,host2:456/somedb?target_session_attrs=any&application_name=myapp
*/
func (f *PostgreSQLUrlConverter) Convert(binding fileconfig.ServiceBinding) string {
	prefix := "postgresql://"
	if binding.Properties[keyHost] == "" {
		return prefix
	}
	un, pwd := url.QueryEscape(binding.Properties[keyUsername]), url.QueryEscape(binding.Properties[keyPassword])
	if len(un) > 0 {
		if len(pwd) > 0 {
			un += ":" + pwd
		}
		un += "@"
	}
	//username & password are optional
	prefix += un + binding.Properties[keyHost]

	if binding.Properties[keyPort] != "" {
		prefix += ":" + binding.Properties[keyPort]
	}
	if binding.Properties[keyDatabase] != "" {
		prefix += "/" + binding.Properties[keyDatabase]
	}

	//additional parameters
	var parts []string
	addToParts := func(k, v string) {
		parts = append(parts, k+"="+escaper.Replace(v))
	}
	if binding.Properties[keySslMode] != "" {
		addToParts("sslmode", binding.Properties[keySslMode])
	}
	if binding.Properties[keySslRootCert] != "" {
		addToParts("sslrootcert", fileconfig.GetBindingRootDirectory()+"/"+
			binding.BindingName+"/"+binding.Properties[keySslRootCert])
	}

	//func returns options parameter
	buildOptions := func() string {
		//option parameters
		//options=...
		var options []string
		var crdbClusterOpt string
		addToOpt := func(k, v string) {
			//each option with format: -c opt=val
			options = append(options, (optionFlag+k)+"="+url.QueryEscape(v))
		}
		if binding.Properties[keyOptions] != "" {
			// in input string multiple options are expected to be separated by & (ampersand) sign,
			// e.g. option1=value1&option2=value2
			options := strings.Split(binding.Properties[keyOptions], optionsSeparator)
			for _, option := range options {
				optionKeyValue := strings.Split(option, optionKeyValueSeparator)
				if len(optionKeyValue) != 2 || len(optionKeyValue[0]) == 0 || len(optionKeyValue[1]) == 0 {
					continue
				}
				if optionKeyValue[0] == "--cluster" { //cockroachDB specific
					crdbClusterOpt = url.QueryEscape("--cluster=" + optionKeyValue[1])
				} else {
					addToOpt(optionKeyValue[0], optionKeyValue[1])
				}
			}
		}
		opt := crdbClusterOpt
		if len(options) > 0 {
			otherOpt := strings.Join(options, encodedSpace)
			if len(opt) > 0 {
				opt = opt + encodedSpace + otherOpt
			} else {
				opt = otherOpt
			}
		}
		return opt
	}

	optionString := buildOptions()
	if len(optionString) > 0 {
		addToParts("options", optionString)
	}

	if len(parts) > 0 {
		slash := ""
		if binding.Properties[keyDatabase] == "" {
			slash = "/"
		}
		return prefix + slash + "?" + strings.Join(parts, "&")
	}
	return prefix
}

package rule
//
// import (
// 	"testing"
// )
//
// func TestEmail(t *testing.T) {
// 	testEmails := map[string]bool{
// 		`email@example.com`:                        true,
// 		`firstname.lastname@example.com`:           true,
// 		`email@subdomain.example.com`:              true,
// 		`firstname+lastname@example.com`:           true,
// 		`email@123.123.123.123`:                    true,
// 		`"email"@example.com`:                      true,
// 		`1234567890@example.com`:                   true,
// 		`email@example-one.com`:                    true,
// 		`_______@example.com`:                      true,
// 		`email@example.name`:                       true,
// 		`email@example.museum`:                     true,
// 		`email@example.co.jp`:                      true,
// 		`firstname-lastname@example.com`:           true,
// 		`email@111.222.333.44444`:                  true,
// 		`email@example.com (Joe Smith)`:            true,
// 		`email@-example.com`:                       true,
// 		`email@example`:                            true,
// 		`email@example@example.com`:                false,
// 		`email@[123.123.123.123]`:                  false,
// 		`email.@example.com`:                       false,
// 		`email..email@example.com`:                 false,
// 		`#@%^%#$@#$@#.com`:                         false,
// 		`Abc..123@example.com`:                     false,
// 		`much.”more\ unusual”@example.com`:         false,
// 		`very.unusual.”@”.unusual.com@example.com`: false,
// 		`.email@example.com`:                       false,
// 		`email@example..com`:                       false,
// 		`”(),:;<>[\]@example.com`:                  false,
// 		`this\ is"really"not\allowed@example.com`:  false,
// 		`plainaddress`:                             false,
// 		`@example.com`:                             false,
// 		`email.example.com`:                        false,
// 	}
// 	for email, valid := range testEmails {
// 		t.Run(email, func(t *testing.T) {
// 			if err := Email(email, false)(); (err == nil) != valid {
// 				t.Errorf("Email() error = %v, wantErr %v", err, !valid)
// 			}
// 		})
// 	}
// }
//
// func TestEmailRFC5322(t *testing.T) {
// 	testEmails := map[string]bool{
// 		`email@example.com`:                        true,
// 		`firstname.lastname@example.com`:           true,
// 		`email@subdomain.example.com`:              true,
// 		`firstname+lastname@example.com`:           true,
// 		`email@123.123.123.123`:                    true,
// 		`email@[123.123.123.123]`:                  true,
// 		`"email"@example.com`:                      true,
// 		`1234567890@example.com`:                   true,
// 		`email@example-one.com`:                    true,
// 		`_______@example.com`:                      true,
// 		`email@example.name`:                       true,
// 		`email@example.museum`:                     true,
// 		`email@example.co.jp`:                      true,
// 		`firstname-lastname@example.com`:           true,
// 		`email@111.222.333.44444`:                  true,
// 		`email@example@example.com`:                false,
// 		`email.@example.com`:                       false,
// 		`email..email@example.com`:                 false,
// 		`#@%^%#$@#$@#.com`:                         false,
// 		`email@-example.com`:                       false,
// 		`Abc..123@example.com`:                     false,
// 		`much.”more\ unusual”@example.com`:         false,
// 		`very.unusual.”@”.unusual.com@example.com`: false,
// 		`email@example.com (Joe Smith)`:            false,
// 		`.email@example.com`:                       false,
// 		`email@example..com`:                       false,
// 		`”(),:;<>[\]@example.com`:                  false,
// 		`this\ is"really"not\allowed@example.com`:  false,
// 		`plainaddress`:                             false,
// 		`@example.com`:                             false,
// 		`email.example.com`:                        false,
// 		`email@example`:                            false,
// 	}
// 	for email, valid := range testEmails {
// 		t.Run(email, func(t *testing.T) {
// 			if err := EmailRFC5322(email, false)(); (err == nil) != valid {
// 				t.Errorf("Email() error = %v, wantErr %v", err, !valid)
// 			}
// 		})
// 	}
// }
//
// func TestEmailWeak(t *testing.T) {
// 	testEmails := map[string]bool{
// 		`email@example.com`:                        true,
// 		`firstname.lastname@example.com`:           true,
// 		`email@subdomain.example.com`:              true,
// 		`firstname+lastname@example.com`:           true,
// 		`email@123.123.123.123`:                    true,
// 		`email@[123.123.123.123]`:                  true,
// 		`"email"@example.com`:                      true,
// 		`1234567890@example.com`:                   true,
// 		`email@example-one.com`:                    true,
// 		`_______@example.com`:                      true,
// 		`email@example.name`:                       true,
// 		`email@example.museum`:                     true,
// 		`email@example.co.jp`:                      true,
// 		`firstname-lastname@example.com`:           true,
// 		`much.”more\ unusual”@example.com`:         true,
// 		`very.unusual.”@”.unusual.com@example.com`: true,
// 		`#@%^%#$@#$@#.com`:                         true,
// 		`email@example@example.com`:                true,
// 		`.email@example.com`:                       true,
// 		`email.@example.com`:                       true,
// 		`email..email@example.com`:                 true,
// 		`email@example.com (Joe Smith)`:            true,
// 		`email@-example.com`:                       true,
// 		`email@111.222.333.44444`:                  true,
// 		`email@example..com`:                       true,
// 		`Abc..123@example.com`:                     true,
// 		`”(),:;<>[\]@example.com`:                  true,
// 		`this\ is"really"not\allowed@example.com`:  true,
// 		`plainaddress`:                             false,
// 		`@example.com`:                             false,
// 		`email.example.com`:                        false,
// 		`email@example`:                            false,
// 	}
// 	for email, valid := range testEmails {
// 		t.Run(email, func(t *testing.T) {
// 			if err := EmailWeak(email)(); (err == nil) != valid {
// 				t.Errorf("Email() error = %v, wantErr %v", err, !valid)
// 			}
// 		})
// 	}
// }
//
// func Test_emailHostLookup(t *testing.T) {
// 	type args struct {
// 		v string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "foomo@foomo.org",
// 			args: args{v: "foomo@foomo.org"},
// 		},
// 		{
// 			name:    "foomo@foomo.not",
// 			args:    args{v: "foomo@foomo.not"},
// 			wantErr: true,
// 		},
// 		{
// 			name:    "foomo@localhost",
// 			args:    args{v: "foomo@localhost"},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := emailHostLookup(tt.args.v); (err != nil) != tt.wantErr {
// 				t.Errorf("emailHostLookup() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

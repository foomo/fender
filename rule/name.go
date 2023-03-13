package rule

type Name string

const (
	NameAlpha        Name = "alpha"
	NameAlphaNumeric Name = "alphaNumeric"
	NameBool         Name = "bool"
	NameContains     Name = "contains"
	NameEmail        Name = "email"
	NameEqual        Name = "equal"
	NameHostname     Name = "hostname"
	NameMD5          Name = "md5"
	NameMax          Name = "max"
	NameMin          Name = "min"
	NameNumeric      Name = "numeric"
	NamePrefix       Name = "prefix"
	NameRegex        Name = "regex"
	NameRequired     Name = "required"
	NameSize         Name = "size"
	NameSuffix       Name = "suffix"
	NameUUID         Name = "uuid"
	NameValid        Name = "valid"
	NameURI          Name = "uri"
	NameURL          Name = "url"
)

func (n Name) String() string {
	return string(n)
}

package core

type (
	MailSender interface {
		Send(to []string, subject string, body []byte) error
	}

	Validator interface {
		CheckEmail(string) error
		CheckRequired(string, string) error
		CheckStringLen(s string, min int, max int, field string) error
	}

	JWTSignParser interface {
		Sign(claims map[string]interface{}, secret string) (string, error)
		Parse(tokenStr string, secret string) (map[string]interface{}, error)
	}

	Factory interface {
		NewCatalog() Catalog
		NewMail() Mailer
		NewUser() User
	}

	factory struct {
		StorageFactory
		ms      MailSender
		v       Validator
		jwt     JWTSignParser
		emitter Emitter
	}
)

// New instances new engine factory
func New(sf StorageFactory, ms MailSender, v Validator, jwt JWTSignParser, emitter Emitter) Factory {
	return &factory{
		StorageFactory: sf,
		ms:             ms,
		v:              v,
		jwt:            jwt,
		emitter:        emitter,
	}
}

package auth

type Authority struct {
	Descriptor string `json:"descriptor"`
}

type authorities struct {
	COMMON    Authority
	EDITOR    Authority
	MODERATOR Authority
	ADMIN     Authority
	ROOT      Authority
}

var roles *authorities

func GetAuthorities() *authorities {
	return roles
}

func init() {
	roles = &authorities{
		COMMON:    Authority{Descriptor: "COMMON"},
		EDITOR:    Authority{Descriptor: "EDITOR"},
		MODERATOR: Authority{Descriptor: "MODERATOR"},
		ADMIN:     Authority{Descriptor: "ADMIN"},
		ROOT:      Authority{Descriptor: "ROOT"},
	}
}

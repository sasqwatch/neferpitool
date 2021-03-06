package whois

type Whois struct {
	Raw string
	Parsed
}

type Parsed struct {
	Registrant
	Registrar
}

// Registrar storing registrar info
type Registrar struct {
	RegistrarID    string `json:"registrar_id"`
	RegistrarName  string `json:"registrar_name"`
	WhoisServer    string `json:"whois_server"`
	ReferralURL    string `json:"referral_url"`
	DomainId       string `json:"domain_id"`
	DomainName     string `json:"domain_naame"`
	DomainStatus   string `json:"domain_status"`
	NameServers    string `json:"name_servers"`
	DomainDNSSEC   string `json:"domain_dnssec"`
	CreatedDate    string `json:"created_date"`
	UpdatedDate    string `json:"updated_date"`
	ExpirationDate string `json:"expiration_date"`
}

// Registrant storing registrant info
type Registrant struct {
	RegistrantID   string `json:"registrant_id"`
	RegistrantName string `json:"registrant_name"`
	Organization   string `json:"organization"`
	Street         string `json:"street"`
	StreetExt      string `json:"street_ext"`
	City           string `json:"city"`
	Province       string `json:"province"`
	PostalCode     string `json:"postal_code"`
	Country        string `json:"country"`
	Phone          string `json:"phone"`
	PhoneExt       string `json:"phone_ext"`
	Fax            string `json:"fax"`
	FaxExt         string `json:"fax_ext"`
	Email          string `json:"email"`
}

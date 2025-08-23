package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Voicev1dialingpermissionsdialingpermissionscountrybulkupdate represents the Voicev1dialingpermissionsdialingpermissionscountrybulkupdate schema from the OpenAPI specification
type Voicev1dialingpermissionsdialingpermissionscountrybulkupdate struct {
	Update_request string `json:"update_request,omitempty"` // A bulk update request to change voice dialing country permissions stored as a URL-encoded, JSON array of update objects. For example : `[ { "iso_code": "GB", "low_risk_numbers_enabled": "true", "high_risk_special_numbers_enabled":"true", "high_risk_tollfraud_numbers_enabled": "false" } ]`
	Update_count int `json:"update_count,omitempty"` // The number of countries updated
}

// Voicev1iprecord represents the Voicev1iprecord schema from the OpenAPI specification
type Voicev1iprecord struct {
	Url string `json:"url,omitempty"` // The absolute URL of the resource.
	Account_sid string `json:"account_sid,omitempty"` // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IP Record resource.
	Cidr_prefix_length int `json:"cidr_prefix_length,omitempty"` // An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32.
	Date_created string `json:"date_created,omitempty"` // The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Date_updated string `json:"date_updated,omitempty"` // The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Friendly_name string `json:"friendly_name,omitempty"` // The string that you assigned to describe the resource.
	Ip_address string `json:"ip_address,omitempty"` // An IP address in dotted decimal notation, IPv4 only.
	Sid string `json:"sid,omitempty"` // The unique string that we created to identify the IP Record resource.
}

// Voicev1archivedcall represents the Voicev1archivedcall schema from the OpenAPI specification
type Voicev1archivedcall struct {
	Date string `json:"date,omitempty"` // The date
	Sid string `json:"sid,omitempty"` // The call sid
	Url string `json:"url,omitempty"` // The absolute URL of the resource.
}

// Voicev1connectionpolicyconnectionpolicytarget represents the Voicev1connectionpolicyconnectionpolicytarget schema from the OpenAPI specification
type Voicev1connectionpolicyconnectionpolicytarget struct {
	Weight int `json:"weight,omitempty"` // The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority.
	Date_created string `json:"date_created,omitempty"` // The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Date_updated string `json:"date_updated,omitempty"` // The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Friendly_name string `json:"friendly_name,omitempty"` // The string that you assigned to describe the resource.
	Sid string `json:"sid,omitempty"` // The unique string that we created to identify the Target resource.
	Connection_policy_sid string `json:"connection_policy_sid,omitempty"` // The SID of the Connection Policy that owns the Target.
	Target string `json:"target,omitempty"` // The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported.
	Account_sid string `json:"account_sid,omitempty"` // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Target resource.
	Enabled bool `json:"enabled,omitempty"` // Whether the target is enabled. The default is `true`.
	Priority int `json:"priority,omitempty"` // The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target.
	Url string `json:"url,omitempty"` // The absolute URL of the resource.
}

// Voicev1dialingpermissions represents the Voicev1dialingpermissions schema from the OpenAPI specification
type Voicev1dialingpermissions struct {
}

// Voicev1dialingpermissionsdialingpermissionscountry represents the Voicev1dialingpermissionsdialingpermissionscountry schema from the OpenAPI specification
type Voicev1dialingpermissionsdialingpermissionscountry struct {
	Country_codes []string `json:"country_codes,omitempty"` // The E.164 assigned [country codes(s)](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html)
	High_risk_special_numbers_enabled bool `json:"high_risk_special_numbers_enabled,omitempty"` // Whether dialing to high-risk special services numbers is enabled. These prefixes include number ranges allocated by the country and include premium numbers, special services, shared cost, and others
	Links map[string]interface{} `json:"links,omitempty"` // A list of URLs related to this resource.
	Name string `json:"name,omitempty"` // The name of the country.
	Url string `json:"url,omitempty"` // The absolute URL of this resource.
	Continent string `json:"continent,omitempty"` // The name of the continent in which the country is located.
	Low_risk_numbers_enabled bool `json:"low_risk_numbers_enabled,omitempty"` // Whether dialing to low-risk numbers is enabled.
	High_risk_tollfraud_numbers_enabled bool `json:"high_risk_tollfraud_numbers_enabled,omitempty"` // Whether dialing to high-risk [toll fraud](https://www.twilio.com/blog/how-to-protect-your-account-from-toll-fraud-with-voice-dialing-geo-permissions-html) numbers is enabled. These prefixes include narrow number ranges that have a high-risk of international revenue sharing fraud (IRSF) attacks, also known as [toll fraud](https://www.twilio.com/blog/how-to-protect-your-account-from-toll-fraud-with-voice-dialing-geo-permissions-html). These prefixes are collected from anti-fraud databases and verified by analyzing calls on our network. These prefixes are not available for download and are updated frequently
	Iso_code string `json:"iso_code,omitempty"` // The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
}

// Voicev1connectionpolicy represents the Voicev1connectionpolicy schema from the OpenAPI specification
type Voicev1connectionpolicy struct {
	Friendly_name string `json:"friendly_name,omitempty"` // The string that you assigned to describe the resource.
	Links map[string]interface{} `json:"links,omitempty"` // The URLs of related resources.
	Sid string `json:"sid,omitempty"` // The unique string that we created to identify the Connection Policy resource.
	Url string `json:"url,omitempty"` // The absolute URL of the resource.
	Account_sid string `json:"account_sid,omitempty"` // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Connection Policy resource.
	Date_created string `json:"date_created,omitempty"` // The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Date_updated string `json:"date_updated,omitempty"` // The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
}

// Voicev1dialingpermissionsdialingpermissionscountrydialingpermissionshrsprefixes represents the Voicev1dialingpermissionsdialingpermissionscountrydialingpermissionshrsprefixes schema from the OpenAPI specification
type Voicev1dialingpermissionsdialingpermissionscountrydialingpermissionshrsprefixes struct {
	Prefix string `json:"prefix,omitempty"` // A prefix is a contiguous number range for a block of E.164 numbers that includes the E.164 assigned country code. For example, a North American Numbering Plan prefix like `+1510720` written like `+1(510) 720` matches all numbers inclusive from `+1(510) 720-0000` to `+1(510) 720-9999`.
}

// Voicev1sourceipmapping represents the Voicev1sourceipmapping schema from the OpenAPI specification
type Voicev1sourceipmapping struct {
	Url string `json:"url,omitempty"` // The absolute URL of the resource.
	Date_created string `json:"date_created,omitempty"` // The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Date_updated string `json:"date_updated,omitempty"` // The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Ip_record_sid string `json:"ip_record_sid,omitempty"` // The Twilio-provided string that uniquely identifies the IP Record resource to map from.
	Sid string `json:"sid,omitempty"` // The unique string that we created to identify the IP Record resource.
	Sip_domain_sid string `json:"sip_domain_sid,omitempty"` // The SID of the SIP Domain that the IP Record is mapped to.
}

// Voicev1byoctrunk represents the Voicev1byoctrunk schema from the OpenAPI specification
type Voicev1byoctrunk struct {
	Date_updated string `json:"date_updated,omitempty"` // The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Status_callback_method string `json:"status_callback_method,omitempty"` // The HTTP method we use to call `status_callback_url`. Either `GET` or `POST`.
	Voice_fallback_url string `json:"voice_fallback_url,omitempty"` // The URL that we call when an error occurs while retrieving or executing the TwiML requested from `voice_url`.
	Voice_method string `json:"voice_method,omitempty"` // The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`.
	Voice_url string `json:"voice_url,omitempty"` // The URL we call using the `voice_method` when the BYOC Trunk receives a call.
	Cnam_lookup_enabled bool `json:"cnam_lookup_enabled,omitempty"` // Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
	Date_created string `json:"date_created,omitempty"` // The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	Connection_policy_sid string `json:"connection_policy_sid,omitempty"` // The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
	Account_sid string `json:"account_sid,omitempty"` // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the BYOC Trunk resource.
	Friendly_name string `json:"friendly_name,omitempty"` // The string that you assigned to describe the resource.
	From_domain_sid string `json:"from_domain_sid,omitempty"` // The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to "call back" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to "sip.twilio.com".
	Sid string `json:"sid,omitempty"` // The unique string that that we created to identify the BYOC Trunk resource.
	Url string `json:"url,omitempty"` // The absolute URL of the resource.
	Voice_fallback_method string `json:"voice_fallback_method,omitempty"` // The HTTP method we use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	Status_callback_url string `json:"status_callback_url,omitempty"` // The URL that we call to pass status parameters (such as call ended) to your application.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"` // The absolute URL of this resource.
	High_risk_special_numbers_enabled bool `json:"high_risk_special_numbers_enabled,omitempty"` // Whether dialing to high-risk special services numbers is enabled. These prefixes include number ranges allocated by the country and include premium numbers, special services, shared cost, and others
	Iso_code string `json:"iso_code,omitempty"` // The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Continent string `json:"continent,omitempty"` // The name of the continent in which the country is located.
	Country_codes []string `json:"country_codes,omitempty"` // The E.164 assigned [country codes(s)](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html)
	High_risk_tollfraud_numbers_enabled bool `json:"high_risk_tollfraud_numbers_enabled,omitempty"` // Whether dialing to high-risk [toll fraud](https://www.twilio.com/blog/how-to-protect-your-account-from-toll-fraud-with-voice-dialing-geo-permissions-html) numbers is enabled. These prefixes include narrow number ranges that have a high-risk of international revenue sharing fraud (IRSF) attacks, also known as [toll fraud](https://www.twilio.com/blog/how-to-protect-your-account-from-toll-fraud-with-voice-dialing-geo-permissions-html). These prefixes are collected from anti-fraud databases and verified by analyzing calls on our network. These prefixes are not available for download and are updated frequently
	Links map[string]interface{} `json:"links,omitempty"` // A list of URLs related to this resource.
	Low_risk_numbers_enabled bool `json:"low_risk_numbers_enabled,omitempty"` // Whether dialing to low-risk numbers is enabled.
	Name string `json:"name,omitempty"` // The name of the country.
}

// Voicev1dialingpermissionsdialingpermissionssettings represents the Voicev1dialingpermissionsdialingpermissionssettings schema from the OpenAPI specification
type Voicev1dialingpermissionsdialingpermissionssettings struct {
	Dialing_permissions_inheritance bool `json:"dialing_permissions_inheritance,omitempty"` // `true` if the sub-account will inherit voice dialing permissions from the Master Project; otherwise `false`.
	Url string `json:"url,omitempty"` // The absolute URL of this resource.
}

package sublime

import "time"

type SublimeEvent struct {
	ID         string    `json:"id"`
	APIVersion string    `json:"api_version"`
	CreatedAt  time.Time `json:"created_at"`
	Type       string    `json:"type"`
	Data       struct {
		Message struct {
			ID              string `json:"id"`
			CanonicalID     string `json:"canonical_id"`
			ExternalID      string `json:"external_id"`
			MessageSourceID string `json:"message_source_id"`
			Mailbox         struct {
				ID         string `json:"id"`
				ExternalID string `json:"external_id"`
			} `json:"mailbox"`
		} `json:"message"`
		FlaggedRules []struct {
			ID       string   `json:"id"`
			Name     string   `json:"name"`
			Severity string   `json:"severity"`
			Tags     []string `json:"tags"`
		} `json:"flagged_rules"`
		TriggeredActions []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"triggered_actions"`
	} `json:"data"`
}

type MessageGroup struct {
	ID              string      `json:"id"`
	State           string      `json:"state"`
	Spam            string      `json:"spam"`
	ReviewStatus    string      `json:"review_status"`
	UserReports     interface{} `json:"user_reports"`
	MessageOpens    interface{} `json:"message_opens"`
	MessageReplies  interface{} `json:"message_replies"`
	MessageForwards interface{} `json:"message_forwards"`
	Previews        []struct {
		CreatedAt          time.Time `json:"created_at"`
		ID                 string    `json:"id"`
		MailboxEmail       string    `json:"mailbox_email"`
		MailboxExternalID  string    `json:"mailbox_external_id"`
		MailboxID          string    `json:"mailbox_id"`
		Recipients         []string  `json:"recipients"`
		SenderDisplayName  string    `json:"sender_display_name"`
		SenderEmailAddress string    `json:"sender_email_address"`
		Subject            string    `json:"subject"`
	} `json:"previews"`
	MailboxEmailAddresses []string `json:"mailbox_email_addresses"`
	FlaggedRules          []struct {
		RuleID   string `json:"rule_id"`
		RuleMeta struct {
			Authors           interface{} `json:"authors"`
			CreatedAt         time.Time   `json:"created_at"`
			CreatedByOrgID    interface{} `json:"created_by_org_id"`
			CreatedByOrgName  interface{} `json:"created_by_org_name"`
			CreatedByUserID   interface{} `json:"created_by_user_id"`
			CreatedByUserName interface{} `json:"created_by_user_name"`
			Description       string      `json:"description"`
			FalsePositives    interface{} `json:"false_positives"`
			ID                string      `json:"id"`
			Label             interface{} `json:"label"`
			Maturity          interface{} `json:"maturity"`
			Name              string      `json:"name"`
			OrgID             string      `json:"org_id"`
			RuleType          string      `json:"rule_type"`
			Severity          string      `json:"severity"`
			Tags              []string    `json:"tags"`
			URLReferences     interface{} `json:"url_references"`
		} `json:"rule_meta"`
	} `json:"flagged_rules"`
	ActionNames              []string    `json:"action_names"`
	AttachmentNames          []string    `json:"attachment_names"`
	FirstCreatedAt           time.Time   `json:"first_created_at"`
	LastCreatedAt            time.Time   `json:"last_created_at"`
	FirstReportedAsPhishAt   interface{} `json:"first_reported_as_phish_at"`
	HistoricallyMatchedRules []struct {
		RuleID   string `json:"rule_id"`
		RuleMeta struct {
			Authors           interface{} `json:"authors"`
			CreatedAt         time.Time   `json:"created_at"`
			CreatedByOrgID    interface{} `json:"created_by_org_id"`
			CreatedByOrgName  interface{} `json:"created_by_org_name"`
			CreatedByUserID   interface{} `json:"created_by_user_id"`
			CreatedByUserName interface{} `json:"created_by_user_name"`
			Description       string      `json:"description"`
			FalsePositives    interface{} `json:"false_positives"`
			ID                string      `json:"id"`
			Label             interface{} `json:"label"`
			Maturity          interface{} `json:"maturity"`
			Name              string      `json:"name"`
			OrgID             string      `json:"org_id"`
			RuleType          string      `json:"rule_type"`
			Severity          string      `json:"severity"`
			Tags              []string    `json:"tags"`
			URLReferences     interface{} `json:"url_references"`
		} `json:"rule_meta"`
	} `json:"historically_matched_rules"`
	MessageType struct {
		Internal bool `json:"internal"`
		Inbound  bool `json:"inbound"`
		Outbound bool `json:"outbound"`
	} `json:"message_type"`
	ReviewComment      interface{} `json:"review_comment"`
	ReviewedByUserID   interface{} `json:"reviewed_by_user_id"`
	ReviewedByUserName interface{} `json:"reviewed_by_user_name"`
	ReviewedAt         interface{} `json:"reviewed_at"`
	UpdatedByUserID    interface{} `json:"updated_by_user_id"`
	DataModel          struct {
		Attachments []struct {
			ContentTransferEncoding string `json:"content_transfer_encoding"`
			ContentType             string `json:"content_type"`
			FileType                string `json:"file_type"`
			Size                    int    `json:"size"`
			FileExtension           string `json:"file_extension"`
			FileName                string `json:"file_name"`
			Md5                     string `json:"md5"`
			Sha1                    string `json:"sha1"`
			Sha256                  string `json:"sha256"`
		} `json:"attachments"`
		Body struct {
			HTML struct {
				Charset                 string `json:"charset"`
				ContentTransferEncoding string `json:"content_transfer_encoding"`
			} `json:"html"`
			Links []struct {
				DisplayText string `json:"display_text"`
				HrefURL     struct {
					URL    string `json:"url"`
					Domain struct {
						Domain     string `json:"domain"`
						RootDomain string `json:"root_domain"`
						Sld        string `json:"sld"`
						Tld        string `json:"tld"`
						Valid      bool   `json:"valid"`
					} `json:"domain"`
					Path   string `json:"path"`
					Scheme string `json:"scheme"`
				} `json:"href_url"`
			} `json:"links"`
		} `json:"body"`
		External struct {
			CreatedAt time.Time `json:"created_at"`
			MessageID string    `json:"message_id"`
			RouteType string    `json:"route_type"`
			Spam      bool      `json:"spam"`
		} `json:"external"`
		Headers struct {
			Date               time.Time `json:"date"`
			DateOriginalOffset string    `json:"date_original_offset"`
			Domains            []struct {
				Domain     string `json:"domain"`
				RootDomain string `json:"root_domain"`
				Sld        string `json:"sld"`
				Subdomain  string `json:"subdomain"`
				Tld        string `json:"tld"`
				Valid      bool   `json:"valid"`
			} `json:"domains"`
			Ips []struct {
				IP string `json:"ip"`
			} `json:"ips"`
			MessageID  string `json:"message_id"`
			ReturnPath struct {
				Email     string `json:"email"`
				LocalPart string `json:"local_part"`
				Domain    struct {
					Domain     string `json:"domain"`
					RootDomain string `json:"root_domain"`
					Sld        string `json:"sld"`
					Subdomain  string `json:"subdomain"`
					Tld        string `json:"tld"`
					Valid      bool   `json:"valid"`
				} `json:"domain"`
			} `json:"return_path"`
			Hops []struct {
				Index  int `json:"index"`
				Fields []struct {
					Name     string `json:"name"`
					Value    string `json:"value"`
					Position int    `json:"position"`
				} `json:"fields"`
				Signature struct {
					Type      string `json:"type"`
					Version   string `json:"version"`
					Algorithm string `json:"algorithm"`
					Selector  string `json:"selector"`
					Signature string `json:"signature"`
					BodyHash  string `json:"body_hash"`
					Domain    string `json:"domain"`
					Headers   string `json:"headers"`
				} `json:"signature,omitempty"`
				AuthenticationResults struct {
					Type     string `json:"type"`
					Compauth struct {
						Verdict string `json:"verdict"`
						Reason  string `json:"reason"`
					} `json:"compauth"`
					Dkim        string `json:"dkim"`
					DkimDetails []struct {
						Type   string `json:"type"`
						Domain string `json:"domain"`
					} `json:"dkim_details"`
					Dmarc        string `json:"dmarc"`
					DmarcDetails struct {
						Version interface{} `json:"version"`
						Verdict string      `json:"verdict"`
						Action  string      `json:"action"`
						From    struct {
							Domain     string `json:"domain"`
							RootDomain string `json:"root_domain"`
							Sld        string `json:"sld"`
							Tld        string `json:"tld"`
							Valid      bool   `json:"valid"`
						} `json:"from"`
					} `json:"dmarc_details"`
					Spf        string `json:"spf"`
					SpfDetails struct {
						Verdict  string `json:"verdict"`
						ClientIP struct {
							IP string `json:"ip"`
						} `json:"client_ip"`
						Designator  string `json:"designator"`
						Description string `json:"description"`
					} `json:"spf_details"`
				} `json:"authentication_results,omitempty"`
				ReceivedSpf struct {
					Verdict string `json:"verdict"`
					Server  struct {
						Domain     string `json:"domain"`
						RootDomain string `json:"root_domain"`
						Sld        string `json:"sld"`
						Subdomain  string `json:"subdomain"`
						Tld        string `json:"tld"`
						Valid      bool   `json:"valid"`
					} `json:"server"`
					ClientIP struct {
						IP string `json:"ip"`
					} `json:"client_ip"`
					Designator  string `json:"designator"`
					Description string `json:"description"`
				} `json:"received_spf,omitempty"`
			} `json:"hops"`
		} `json:"headers"`
		Type struct {
			Inbound bool `json:"inbound"`
		} `json:"type"`
		Mailbox struct {
			DisplayName string `json:"display_name"`
			Email       struct {
				Email     string `json:"email"`
				LocalPart string `json:"local_part"`
				Domain    struct {
					Domain     string `json:"domain"`
					RootDomain string `json:"root_domain"`
					Sld        string `json:"sld"`
					Subdomain  string `json:"subdomain"`
					Tld        string `json:"tld"`
					Valid      bool   `json:"valid"`
				} `json:"domain"`
			} `json:"email"`
		} `json:"mailbox"`
		Recipients struct {
			To []struct {
				Email struct {
					Email     string `json:"email"`
					LocalPart string `json:"local_part"`
					Domain    struct {
						Domain     string `json:"domain"`
						RootDomain string `json:"root_domain"`
						Sld        string `json:"sld"`
						Subdomain  string `json:"subdomain"`
						Tld        string `json:"tld"`
						Valid      bool   `json:"valid"`
					} `json:"domain"`
				} `json:"email"`
			} `json:"to"`
		} `json:"recipients"`
		Sender struct {
			Email struct {
				Email     string `json:"email"`
				LocalPart string `json:"local_part"`
				Domain    struct {
					Domain     string `json:"domain"`
					RootDomain string `json:"root_domain"`
					Sld        string `json:"sld"`
					Tld        string `json:"tld"`
					Valid      bool   `json:"valid"`
				} `json:"domain"`
			} `json:"email"`
		} `json:"sender"`
		Subject struct {
			Subject string `json:"subject"`
		} `json:"subject"`
		Meta struct {
			ID            string    `json:"id"`
			CanonicalID   string    `json:"canonical_id"`
			CreatedAt     time.Time `json:"created_at"`
			SchemaVersion string    `json:"schema_version"`
		} `json:"_meta"`
	} `json:"data_model"`
}

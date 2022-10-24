package models

import "time"

type Application struct {
	Data struct {
		Uuid                             string      `json:"uuid"`
		Status                           string      `json:"status"`
		UserCreated                      string      `json:"user_created"`
		DateCreated                      string      `json:"date_created"`
		UserUpdated                      string      `json:"user_updated"`
		DateUpdated                      string      `json:"date_updated"`
		BusinessName                     string      `json:"business_name"`
		Remarks                          interface{} `json:"remarks"`
		WebsiteUrl                       string      `json:"website_url"`
		BusinessPhoneNumber              string      `json:"business_phone_number"`
		NaturalPersonPhone               string      `json:"natural_person_phone"`
		NaturalPersonEmail               string      `json:"natural_person_email"`
		CountryOfResidence               string      `json:"country_of_residence"`
		CountryOfCitizenship             string      `json:"country_of_citizenship"`
		NaturalPerson                    int         `json:"natural_person"`
		ApplicationCompany               int         `json:"application_company"`
		ApplicationBusiness              string      `json:"application_business"`
		ApplicationCompanyDowJonesResult interface{} `json:"application_company_dow_jones_result"`
		Todos                            interface{} `json:"todos"`
		ApplicationJson                  struct {
			S01BusinessApplication struct {
				UuidCheck                             string `json:"uuid_check"`
				StatusCheck                           string `json:"status_check"`
				UserCreatedCheck                      string `json:"user_created_check"`
				DateCreatedCheck                      string `json:"date_created_check"`
				UserUpdatedCheck                      string `json:"user_updated_check"`
				DateUpdatedCheck                      string `json:"date_updated_check"`
				BusinessNameCheck                     string `json:"business_name_check"`
				RemarksCheck                          string `json:"remarks_check"`
				WebsiteUrlCheck                       string `json:"website_url_check"`
				BusinessPhoneNumberCheck              string `json:"business_phone_number_check"`
				NaturalPersonPhoneCheck               string `json:"natural_person_phone_check"`
				NaturalPersonEmailCheck               string `json:"natural_person_email_check"`
				CountryOfResidenceCheck               string `json:"country_of_residence_check"`
				CountryOfCitizenshipCheck             string `json:"country_of_citizenship_check"`
				NaturalPersonCheck                    string `json:"natural_person_check"`
				ApplicationCompanyCheck               string `json:"application_company_check"`
				ApplicationBusinessCheck              string `json:"application_business_check"`
				ApplicationCompanyDowJonesResultCheck string `json:"application_company_dow_jones_result_check"`
				TodosCheck                            string `json:"todos_check"`
				PartnerChannelCheck                   string `json:"partner_channel_check"`
				PEnvironmentCheck                     string `json:"p_environment_check"`
				BusinessAddressCheck                  string `json:"business_address_check"`
				ProvisionAtCheck                      string `json:"provision_at_check"`
				SubmittedByCheck                      string `json:"submitted_by_check"`
				ReviewedByCheck                       string `json:"reviewed_by_check"`
				BusinessVolumeAvgMonthlyVolCheck      string `json:"business_volume_avg_monthly_vol_check"`
				BusinessVolumeAvgTransactionCheck     string `json:"business_volume_avg_transaction_check"`
				BusinessVolumeMaxTransactionCheck     string `json:"business_volume_max_transaction_check"`
				ToProvisionANewBusinessCheck          string `json:"to_provision_a_new_business_check"`
				IsPhoneVerifiedCheck                  string `json:"is_phone_verified_check"`
				IsEmailVerifiedCheck                  string `json:"is_email_verified_check"`
				BusinessDbaCheck                      string `json:"business_dba_check"`
				ProductsServicesDescriptionCheck      string `json:"products_services_description_check"`
				BusinessEmailCheck                    string `json:"business_email_check"`
				SelfDeclaredRoleCheck                 string `json:"self_declared_role_check"`
				ScoreDataCompletenessCheck            string `json:"score_data_completeness_check"`
				ReferralCodeCheck                     string `json:"referral_code_check"`
				InternalLogsCheck                     string `json:"internal_logs_check"`
				SalesMarketingCampaignCheck           string `json:"sales_marketing_campaign_check"`
				MerchantCategoryCheck                 string `json:"merchant_category_check"`
				MccCheck                              string `json:"mcc_check"`
				ProcessTypeCheck                      string `json:"process_type_check"`
				TempScoreCheck                        string `json:"temp_score_check"`
				ComplianceRemarksCheck                string `json:"compliance_remarks_check"`
				MembersCheck                          string `json:"members_check"`
				BusinessApplicationImagesCheck        string `json:"business_application_images_check"`
				CompanyMembersCheck                   string `json:"company_members_check"`
				LegacyApplicationFormsCheck           string `json:"legacy_application_forms_check"`
				CoapplicantsCheck                     string `json:"coapplicants_check"`
				MerchantServicesAgreementsCheck       string `json:"merchant_services_agreements_check"`
				BusinessApplicationLogsCheck          string `json:"business_application_logs_check"`
			} `json:"s01_business_application"`
			S03DBusinessDocumentBrs []struct {
				Type              string      `json:"type"`
				IsActive          bool        `json:"is_active"`
				FileId            string      `json:"file_id"`
				FiledAt           time.Time   `json:"filed_at"`
				IsDownloaded      bool        `json:"is_downloaded"`
				DownloadedAt      interface{} `json:"downloaded_at"`
				IsOcr             bool        `json:"is_ocr"`
				OcrAt             interface{} `json:"ocr_at"`
				IsIntegrated      bool        `json:"is_integrated"`
				LinkedToCompanyAt time.Time   `json:"linked_to_company_at"`
			} `json:"s03d_business_document_brs"`
			S08MidsStatus struct {
				Alipay struct {
					Online struct {
						Type           string      `json:"type"`
						IsProvision    bool        `json:"is_provision"`
						IsSuccess      bool        `json:"is_success"`
						IsActive       bool        `json:"is_active"`
						ActivationDate interface{} `json:"activation_date"`
						Data           struct {
							Id              string      `json:"id"`
							Status          string      `json:"status"`
							UserCreated     string      `json:"user_created"`
							DateCreated     time.Time   `json:"date_created"`
							UserUpdated     interface{} `json:"user_updated"`
							DateUpdated     interface{} `json:"date_updated"`
							Mid             string      `json:"mid"`
							Tid             interface{} `json:"tid"`
							PaymentAcquirer string      `json:"payment_acquirer"`
							Business        string      `json:"business"`
							ActivationDate  interface{} `json:"activation_date"`
							Remarks         interface{} `json:"remarks"`
						} `json:"data"`
						Remark string `json:"remark"`
					} `json:"online"`
					Offline struct {
						Type           string      `json:"type"`
						IsProvision    bool        `json:"is_provision"`
						IsSuccess      bool        `json:"is_success"`
						IsActive       bool        `json:"is_active"`
						ActivationDate interface{} `json:"activation_date"`
						Data           struct {
							Id              string      `json:"id"`
							Status          string      `json:"status"`
							UserCreated     string      `json:"user_created"`
							DateCreated     time.Time   `json:"date_created"`
							UserUpdated     interface{} `json:"user_updated"`
							DateUpdated     interface{} `json:"date_updated"`
							Mid             string      `json:"mid"`
							Tid             interface{} `json:"tid"`
							PaymentAcquirer string      `json:"payment_acquirer"`
							Business        string      `json:"business"`
							ActivationDate  interface{} `json:"activation_date"`
							Remarks         interface{} `json:"remarks"`
						} `json:"data"`
						Remark string `json:"remark"`
					} `json:"offline"`
				} `json:"Alipay"`
				WeChatPay struct {
					Online struct {
						Type           string      `json:"type"`
						IsProvision    bool        `json:"is_provision"`
						MID            interface{} `json:"MID"`
						IsSuccess      bool        `json:"is_success"`
						TID            interface{} `json:"TID"`
						IsActive       bool        `json:"is_active"`
						ActivationDate interface{} `json:"activation_date"`
						Data           interface{} `json:"data"`
						Remark         string      `json:"remark"`
					} `json:"online"`
					Offline struct {
						Type           string      `json:"type"`
						IsProvision    bool        `json:"is_provision"`
						MID            interface{} `json:"MID"`
						IsSuccess      bool        `json:"is_success"`
						TID            interface{} `json:"TID"`
						IsActive       bool        `json:"is_active"`
						ActivationDate interface{} `json:"activation_date"`
						Data           interface{} `json:"data"`
						Remark         string      `json:"remark"`
					} `json:"offline"`
				} `json:"WeChat Pay"`
			} `json:"s08_mids_status"`
			Pipedream struct {
				Patch struct {
					ToProvisionANewBusiness bool      `json:"to_provision_a_new_business"`
					ProvisionAt             time.Time `json:"provision_at"`
				} `json:"patch"`
			} `json:"pipedream"`
			StoreApplicationProvision struct {
				Store struct {
					Address1          string      `json:"address1"`
					Address2          string      `json:"address2"`
					BusinessName      string      `json:"business_name"`
					CallingCode       string      `json:"calling_code"`
					Chain             bool        `json:"chain"`
					City              string      `json:"city"`
					CompactTitle      interface{} `json:"compact_title"`
					CountryCode       string      `json:"country_code"`
					Currency          string      `json:"currency"`
					Deleted           bool        `json:"deleted"`
					Delivery          bool        `json:"delivery"`
					DeliveryDesc      interface{} `json:"delivery_desc"`
					DeliveryFee       float64     `json:"delivery_fee"`
					DeliveryMiles     int         `json:"delivery_miles"`
					Description       interface{} `json:"description"`
					DeliveryMinAmount int         `json:"delivery_min_amount"`
					Email             string      `json:"email"`
					Group             interface{} `json:"group"`
					GroupParentId     interface{} `json:"group_parent_id"`
					Highlight         interface{} `json:"highlight"`
					Homepage          interface{} `json:"homepage"`
					Id                int         `json:"id"`
					IframeActive      bool        `json:"iframe_active"`
					OpeningHours      interface{} `json:"opening_hours"`
					ParentId          interface{} `json:"parent_id"`
					Phone             string      `json:"phone"`
					PickupDesc        interface{} `json:"pickup_desc"`
					Policy            interface{} `json:"policy"`
					PosActive         bool        `json:"pos_active"`
					ReplyToStore      bool        `json:"reply_to_store"`
					ReportTimeZone    interface{} `json:"report_time_zone"`
					ReportingToken    string      `json:"reporting_token"`
					Slug              string      `json:"slug"`
					SoftDescriptor    interface{} `json:"soft_descriptor"`
					Sort              interface{} `json:"sort"`
					State             string      `json:"state"`
					SyncInventory     bool        `json:"sync_inventory"`
					Suspended         bool        `json:"suspended"`
					TimeSegments      struct {
					} `json:"time_segments"`
					Title                                   string        `json:"title"`
					Zipcode                                 string        `json:"zipcode"`
					TaxRate                                 float64       `json:"tax_rate"`
					IpadScreensaverUrl                      string        `json:"ipad_screensaver_url"`
					IsScreensaverOn                         bool          `json:"is_screensaver_on"`
					Timezone                                string        `json:"timezone"`
					StoreCreditEnabled                      bool          `json:"store_credit_enabled"`
					LogoUrl                                 interface{}   `json:"logo_url"`
					IconUrl                                 interface{}   `json:"icon_url"`
					DeliveryAreas                           []interface{} `json:"delivery_areas"`
					CurrentExchangeRates                    []interface{} `json:"current_exchange_rates"`
					Lat                                     float64       `json:"lat"`
					Lng                                     float64       `json:"lng"`
					Tags                                    []interface{} `json:"tags"`
					Open                                    bool          `json:"open"`
					TranslationInfo                         interface{}   `json:"translation_info"`
					ConsumerRate                            interface{}   `json:"consumer_rate"`
					Bookings                                interface{}   `json:"bookings"`
					Queueings                               interface{}   `json:"queueings"`
					Pickup                                  bool          `json:"pickup"`
					AssociateType                           interface{}   `json:"associate_type"`
					IsDefault                               interface{}   `json:"is_default"`
					OriginalUrl                             interface{}   `json:"original_url"`
					Tag                                     interface{}   `json:"tag"`
					CustomerAttributes                      interface{}   `json:"customer_attributes"`
					AvgConsumption                          interface{}   `json:"avg_consumption"`
					Distance                                int           `json:"distance"`
					Price                                   interface{}   `json:"price"`
					PlatformId                              interface{}   `json:"platform_id"`
					StoreName                               interface{}   `json:"store_name"`
					StoreAudit                              interface{}   `json:"store_audit"`
					PlatformStoreClassification             interface{}   `json:"platform_store_classification"`
					CustomerAttribute                       interface{}   `json:"customer_attribute"`
					StoreCategory                           interface{}   `json:"store_category"`
					StoreBrands                             interface{}   `json:"store_brands"`
					StoreId                                 interface{}   `json:"store_id"`
					IsBusiness                              bool          `json:"is_business"`
					DineIn                                  bool          `json:"dine_in"`
					DineInUnassigned                        bool          `json:"dine_in_unassigned"`
					StoreScore                              interface{}   `json:"store_score"`
					LoginSelectionForCustomerOrderingDevice interface{}   `json:"login_selection_for_customer_ordering_device"`
					EnterScreen                             interface{}   `json:"enter_screen"`
					ExternalId                              interface{}   `json:"external_id"`
					MasterTerminalIp                        interface{}   `json:"master_terminal_ip"`
					StorePictures                           interface{}   `json:"store_pictures"`
					StoreConfig                             interface{}   `json:"store_config"`
					StoreModule                             interface{}   `json:"store_module"`
					Favorites                               interface{}   `json:"favorites"`
					ModifierSetList                         interface{}   `json:"modifier_set_list"`
					PickUpLocations                         interface{}   `json:"pick_up_locations"`
					TaxOptions                              interface{}   `json:"tax_options"`
					Tables                                  interface{}   `json:"tables"`
				} `json:"store"`
				Success bool `json:"success"`
			} `json:"store_application_provision"`
			AlipayWechatpayProvision struct {
				Alipay struct {
					Result struct {
						ExistMids []interface{} `json:"exist_mids"`
						ApplyMid  struct {
							Online  string `json:"online"`
							Offline struct {
								Status string `json:"status"`
								MID    int    `json:"MID"`
							} `json:"offline"`
						} `json:"apply_mid"`
					} `json:"result"`
				} `json:"Alipay"`
				WeChatPay struct {
					Result struct {
						ExistMids []interface{} `json:"exist_mids"`
						ApplyMid  struct {
							Online  string `json:"online"`
							Offline struct {
								MissingParams struct {
									Companies  []interface{} `json:"companies"`
									Businesses []interface{} `json:"businesses"`
									HkBrs      []string      `json:"hk_brs"`
									Error      string        `json:"error"`
								} `json:"missing_params"`
								Status string `json:"status"`
								Err    string `json:"err"`
							} `json:"offline"`
						} `json:"apply_mid"`
					} `json:"result"`
				} `json:"WeChat Pay"`
			} `json:"alipay_wechatpay_provision"`
		} `json:"application_json"`
		PartnerChannel               string      `json:"partner_channel"`
		PEnvironment                 interface{} `json:"p_environment"`
		BusinessAddress              string      `json:"business_address"`
		ProvisionAt                  string      `json:"provision_at"`
		SubmittedBy                  int         `json:"submitted_by"`
		ReviewedBy                   interface{} `json:"reviewed_by"`
		BusinessVolumeAvgMonthlyVol  string      `json:"business_volume_avg_monthly_vol"`
		BusinessVolumeAvgTransaction string      `json:"business_volume_avg_transaction"`
		BusinessVolumeMaxTransaction string      `json:"business_volume_max_transaction"`
		ToProvisionANewBusiness      bool        `json:"to_provision_a_new_business"`
		IsPhoneVerified              bool        `json:"is_phone_verified"`
		IsEmailVerified              bool        `json:"is_email_verified"`
		BusinessDba                  string      `json:"business_dba"`
		ProductsServicesDescription  string      `json:"products_services_description"`
		BusinessEmail                interface{} `json:"business_email"`
		SelfDeclaredRole             string      `json:"self_declared_role"`
		ScoreDataCompleteness        struct {
			CompanyType    string `json:"company_type"`
			IsBrUploaded   string `json:"is_br_uploaded"`
			IsCiUploaded   string `json:"is_ci_uploaded"`
			IsEeUploaded   string `json:"is_ee_uploaded"`
			IsNar1Uploaded string `json:"is_nar1_uploaded"`
			IsNnc1Uploaded string `json:"is_nnc1_uploaded"`
		} `json:"score_data_completeness"`
		ReferralCode               string        `json:"referral_code"`
		InternalLogs               string        `json:"internal_logs"`
		SalesMarketingCampaign     string        `json:"sales_marketing_campaign"`
		MerchantCategory           string        `json:"merchant_category"`
		Mcc                        interface{}   `json:"mcc"`
		ProcessType                interface{}   `json:"process_type"`
		TempScore                  interface{}   `json:"temp_score"`
		ComplianceRemarks          interface{}   `json:"compliance_remarks"`
		BusinessApplicationImages  []interface{} `json:"business_application_images"`
		LegacyApplicationForms     []interface{} `json:"legacy_application_forms"`
		Coapplicants               []interface{} `json:"coapplicants"`
		MerchantServicesAgreements []interface{} `json:"merchant_services_agreements"`
		BusinessApplicationLogs    []string      `json:"business_application_logs"`
	} `json:"data"`
}

type Compnay struct {
	Data struct {
		Id                               int           `json:"id"`
		NameOfBusinessCorporationEnglish string        `json:"name_of_business_corporation_english"`
		RegisteredAddress                interface{}   `json:"registered_address"`
		DateOfCommencement               interface{}   `json:"date_of_commencement"`
		DateOfCessation                  interface{}   `json:"date_of_cessation"`
		DateOfExpiry                     interface{}   `json:"date_of_expiry"`
		CertificateNumber                interface{}   `json:"certificate_number"`
		FeeAndLevy                       interface{}   `json:"fee_and_levy"`
		BusinessBranchNameEnglish        interface{}   `json:"business_branch_name_english"`
		BusinessBranchNameChinese        interface{}   `json:"business_branch_name_chinese"`
		DateOfIncorporation              string        `json:"date_of_incorporation"`
		CompanyNumber                    string        `json:"company_number"`
		NameOfBusinessCorporationChinese string        `json:"name_of_business_corporation_chinese"`
		CreatedAt                        string        `json:"created_at"`
		UpdatedAt                        string        `json:"updated_at"`
		CreatedById                      interface{}   `json:"created_by_id"`
		UpdatedById                      interface{}   `json:"updated_by_id"`
		WindingUpMode                    string        `json:"winding_up_mode"`
		RegisterOfCharges                string        `json:"register_of_charges"`
		ImportantNote                    string        `json:"important_note"`
		LastUpdateBy                     string        `json:"last_update_by"`
		NatureOfBusiness                 interface{}   `json:"nature_of_business"`
		BusinessStatus                   interface{}   `json:"business_status"`
		CompanyCategory                  interface{}   `json:"company_category"`
		CompanyType                      string        `json:"company_type"`
		AdministrativeDivision           interface{}   `json:"administrative_division"`
		Remarks                          string        `json:"remarks"`
		ActiveStatus                     string        `json:"active_status"`
		TotalIssuedShares                interface{}   `json:"total_issued_shares"`
		BrNumber                         interface{}   `json:"br_number"`
		StatusRemarks                    interface{}   `json:"status_remarks"`
		Todos                            interface{}   `json:"todos"`
		Country                          string        `json:"country"`
		OrgChart                         interface{}   `json:"org_chart"`
		OtherAttributes                  interface{}   `json:"other_attributes"`
		IsBrWaiver                       interface{}   `json:"is_br_waiver"`
		NameHistories                    []int         `json:"name_histories"`
		CompanyDocs                      []interface{} `json:"company_docs"`
		DowJonesBatchScreeningsResults   []interface{} `json:"dow_jones_batch_screenings_results"`
		Businesses                       []string      `json:"businesses"`
		Members                          []int         `json:"members"`
		HkNar1S                          []string      `json:"hk_nar1s"`
		HkBrs                            []string      `json:"hk_brs"`
		HkNnc1S                          []interface{} `json:"hk_nnc1s"`
		DowJonesCompanyScreeningResults  []interface{} `json:"dow_jones_company_screening_results"`
		BusinessApplications             []string      `json:"business_applications"`
		HkElectronicExtracts             []interface{} `json:"hk_electronic_extracts"`
	} `json:"data"`
}

type Business struct {
	Data struct {
		Id                        string        `json:"id"`
		BusinessPhone             interface{}   `json:"business_phone"`
		BusinessEmail             interface{}   `json:"business_email"`
		CreatedAt                 string        `json:"created_at"`
		UpdatedAt                 string        `json:"updated_at"`
		PublishedAt               interface{}   `json:"published_at"`
		CreatedById               interface{}   `json:"created_by_id"`
		UpdatedById               interface{}   `json:"updated_by_id"`
		MerchantContactEmail      interface{}   `json:"merchant_contact_email"`
		BusinessId                int           `json:"business_id"`
		BusinessName              string        `json:"business_name"`
		DateApproval              interface{}   `json:"date_approval"`
		Status                    string        `json:"status"`
		Remarks                   interface{}   `json:"remarks"`
		SettlementEntityType      interface{}   `json:"settlement_entity_type"`
		Omnichannel               interface{}   `json:"omnichannel"`
		Company                   int           `json:"company"`
		DelegatedSettlementEntity interface{}   `json:"delegated_settlement_entity"`
		RollingReserves           interface{}   `json:"rolling_reserves"`
		Mdrs                      interface{}   `json:"mdrs"`
		RiskScore                 interface{}   `json:"risk_score"`
		RiskLevel                 interface{}   `json:"risk_level"`
		Country                   string        `json:"country"`
		JoinCode                  string        `json:"join_code"`
		PDeletedAt                interface{}   `json:"p_deleted_at"`
		Address1                  interface{}   `json:"address_1"`
		Address2                  interface{}   `json:"address_2"`
		City                      interface{}   `json:"city"`
		State                     interface{}   `json:"state"`
		ZipCode                   interface{}   `json:"zip_code"`
		BusinessHighlight         interface{}   `json:"business_highlight"`
		PartnerChannel            interface{}   `json:"partner_channel"`
		MerchantCategory          interface{}   `json:"merchant_category"`
		FeeGroup                  interface{}   `json:"fee_group"`
		BankAccounts              []interface{} `json:"bank_accounts"`
		Contacts                  []interface{} `json:"contacts"`
		MerchantIds               []string      `json:"merchant_ids"`
		Devices                   []interface{} `json:"devices"`
		Supports                  []interface{} `json:"supports"`
		TerminalServices          []interface{} `json:"terminal_services"`
		BusinessPictures          []int         `json:"business_pictures"`
		Subscriptions             []interface{} `json:"subscriptions"`
		Translations              []interface{} `json:"translations"`
		BusinessApplication       []string      `json:"business_application"`
	} `json:"data"`
}

type Br struct {
	Data struct {
		Id                        string      `json:"id"`
		Status                    string      `json:"status"`
		UserCreated               string      `json:"user_created"`
		DateCreated               time.Time   `json:"date_created"`
		UserUpdated               string      `json:"user_updated"`
		DateUpdated               time.Time   `json:"date_updated"`
		BrNumberWithoutCheckDigit string      `json:"br_number_without_check_digit"`
		BrNumberWithCheckDigit    string      `json:"br_number_with_check_digit"`
		CompanyEnglishName        string      `json:"company_english_name"`
		CompanyChineseName        string      `json:"company_chinese_name"`
		CompanyAddress            string      `json:"company_address"`
		ErrorMessage              string      `json:"error_message"`
		ExpirationDate            string      `json:"expiration_date"`
		MatchedCompany            int         `json:"matched_company"`
		MatchedAt                 time.Time   `json:"matched_at"`
		NaturalOfBusiness         interface{} `json:"natural_of_business"`
		CertificateNumber         string      `json:"certificate_number"`
		CommencementDate          string      `json:"commencement_date"`
		BranchNameEnglish         interface{} `json:"branch_name_english"`
		BranchNameChinese         interface{} `json:"branch_name_chinese"`
		FeeAndLevy                interface{} `json:"fee_and_levy"`
		Remarks                   string      `json:"remarks"`
		ConvertedJson             interface{} `json:"converted_json"`
		SubmissionDate            interface{} `json:"submission_date"`
		DocumentSavedAt           interface{} `json:"document_saved_at"`
		References                []int       `json:"references"`
	}
}

type Applymodel struct {
	SpAppid     string `json:"sp_appid"`
	SpMchid     string `json:"sp_mchid"`
	Name        string `json:"name"`
	Shortname   string `json:"shortname"`
	OfficePhone string `json:"office_phone"`
	Contact     struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"contact"`
	BusinessCategory              int    `json:"business_category"`
	MerchantCountryCode           string `json:"merchant_country_code"`
	MerchantType                  string `json:"merchant_type"`
	RegistrationCertificateNumber string `json:"registration_certificate_number"`
	RegistrationCertificateDate   string `json:"registration_certificate_date"`
	SettlementBankNumber          string `json:"settlement_bank_number"`
	Business                      struct {
		BusinessType string `json:"business_type"`
		Mcc          string `json:"mcc"`
		MiniProgram  string `json:"mini_program"`
		StoreAddress string `json:"store_address"`
	} `json:"business"`
	Director struct {
		Name   string `json:"name"`
		Number string `json:"number"`
	} `json:"director"`
}

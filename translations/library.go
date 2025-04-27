package translations

type (
	TranslationKey struct {
		MSEA string `json:"msea"`
		KMS  string `json:"kms"`
	}
	TranslationLibrary struct {
		Keys map[string]TranslationKey `json:"keys"`
	}
)

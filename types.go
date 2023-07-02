package main

type ChargeRequest struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	CallbackUrl string `json:"callback_url"`
	SuccessUrl  string `json:"success_url"`
}

type ChargeResponseData struct {
	// Id                string           `json:"id"`
	// Description       string           `json:"description"`
	// CreatedAt         int64            `json:"created_at"`
	// Status            string           `json:"status"`
	Amount int `json:"amount"`
	// CallbackUrl       string           `json:"callback_url"`
	// SuccessUrl        string           `json:"success_url"`
	// HostedCheckoutUrl string           `json:"hosted_checkout_url"`
	// OrderId           string           `json:"order_id"`
	// Currency          string           `json:"currency"`
	// SourceFiatValue   int              `json:"source_fiat_value"`
	LightningInvoice LightningInvoice `json:"lightning_invoice"`
}

type ChargeResponse struct {
	Data ChargeResponseData `json:"data"`
}

type LightningInvoice struct {
	ExperiesAt int64  `json:"expires_at"`
	PayReq     string `json:"payreq"`
}

type ListCharges struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Search   string `json:"search"`
}

type Rates struct {
	Data map[string]Rate `json:"data"`
}

type Rate struct {
	Currency string  `json:"currency"`
	BTC      float64 `json:"BTC"`
}

type ConcurencyRate struct {
	Symbol string  `json:"symbol"`
	Rate   float64 `json:"rate"`
}

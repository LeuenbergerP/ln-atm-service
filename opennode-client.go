package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const content_type string = "application/string"

var client *http.Client

type OpenNodeClient struct {
	apiKey string
	baseUrl    string
}

type Charge struct {
  Amount int `json:"amount"`
  CallbackUrl string `json:"callback_url"`
  SuccessUrl string `json:"success_url"`
  Description string `json:"description"`
  Currency string `json:"currency"`
  CustomerEmail string `json:"customer_email,omitempty"`
  NotifyEmail string `json:"notify_email"`
  CustomerName string `json:"customer_name"`
  OrderId string `json:"order_id"`
  AutoSettle bool `json:"auto_settle"`
  Ttl int `json:"ttl"`
}

type ChargeResponse struct {
  Id string `json:"id"`
  Description string `json:"description"`
  CreatedAt int64 `json:"created_at"`
  Status string `json:"status"`
  Amount int `json:"amount"`
  CallbackUrl string `json:"callback_url"`
  SuccessUrl string `json:"success_url"`
  HostedCheckoutUrl string `json:"hosted_checkout_url"`
  OrderId string `json:"order_id"`
  Currency string `json:"currency"`
  SourceFiatValue int `json:"source_fiat_value"`
  LightningInvoice LightningInvoice `json:"lightning_invoice"`
}

type LightningInvoice struct {
  ExperiesAt int64 `json:"expires_at"`
  PayReq string `json:"payreq"`
}

type ListCharges struct {
  Page int `json:"page"`
  PageSize int `json:"page_size"`
  Search string `json:"search"`
}

func NewOpenNodeClient(baseUrl string, apiKey string) *OpenNodeClient {
	return &OpenNodeClient{
		apiKey: apiKey,
		baseUrl:   baseUrl,
	}
}

func (c *OpenNodeClient) CreateCharge(charge Charge) (*ChargeResponse, error) {
  response := ChargeResponse {}
  return &response, nil
}

func (c *OpenNodeClient) ListCharges(list ListCharges) (string, error) {
	return "", nil
}

func (c *OpenNodeClient) RefundInfo(id string) (string, error) {
  return "", nil
}

func (c *OpenNodeClient) buildUrl(path string) string {
  return fmt.Sprintf(c.baseUrl, path)
}

func getJson(url string, target interface{}) error {
  resp, err := client.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close() 
  return json.NewDecoder(resp.Body).Decode(target)
}

func postJson(url string, body interface{}) error {
  json, err := json.Marshal(body)

  if err != nil {
    
  }
  requestBody := string(json)
  _= requestBody


  return nil
}

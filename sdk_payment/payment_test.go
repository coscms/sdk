package sdk_payment

import (
	"testing"

	"github.com/coscms/sdk/sdk_options"
)

type testAppInfo struct {
	AppID       string
	AppSecret   string
	ApiEndpoint string
}

func (t *testAppInfo) GetAppID() string      { return t.AppID }
func (t *testAppInfo) GetAppSecret() string   { return t.AppSecret }
func (t *testAppInfo) GetApiEndpoint() string { return t.ApiEndpoint }

func TestNewPaymentOptions(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	if opts == nil {
		t.Fatal("expected non-nil Options")
	}
}

func TestCheckoutOptionsURLValues(t *testing.T) {
	c := &CheckoutOptions{
		AppID:     "app-1",
		Price:     99.99,
		OutOrderNo: "order-001",
		Subject:   "Test Product",
		Type:      "alipay",
	}

	vals := c.URLValues()
	if vals.Get("appID") != "app-1" {
		t.Errorf("expected app-1, got %s", vals.Get("appID"))
	}
	if vals.Get("price") != "99.99" {
		t.Errorf("expected 99.99, got %s", vals.Get("price"))
	}
	if vals.Get("outOrderNo") != "order-001" {
		t.Errorf("expected order-001, got %s", vals.Get("outOrderNo"))
	}
	if vals.Get("subject") != "Test Product" {
		t.Errorf("expected Test Product, got %s", vals.Get("subject"))
	}
	if vals.Get("type") != "alipay" {
		t.Errorf("expected alipay, got %s", vals.Get("type"))
	}
	if vals.Get("isVirtualProduct") != "0" {
		t.Errorf("expected 0, got %s", vals.Get("isVirtualProduct"))
	}
}

func TestCheckoutOptionsIsVirtualProduct(t *testing.T) {
	c := &CheckoutOptions{IsVirtualProduct: true}
	vals := c.URLValues()
	if vals.Get("isVirtualProduct") != "1" {
		t.Errorf("expected 1, got %s", vals.Get("isVirtualProduct"))
	}
}

func TestCheckoutOptionsString(t *testing.T) {
	c := &CheckoutOptions{AppID: "app-1", Subject: "test"}
	s := c.String()
	if s == "" {
		t.Error("expected non-empty JSON string")
	}
}

func TestRefundOptionsURLValues(t *testing.T) {
	r := &RefundOptions{
		AppID:        "app-1",
		OutOrderNo:   "order-001",
		RefundAmount: 50.00,
		OutRefundNo:  "refund-001",
	}

	vals := r.URLValues()
	if vals.Get("appID") != "app-1" {
		t.Errorf("expected app-1, got %s", vals.Get("appID"))
	}
	if vals.Get("outOrderNo") != "order-001" {
		t.Errorf("expected order-001, got %s", vals.Get("outOrderNo"))
	}
	if vals.Get("refundAmount") != "50" {
		t.Errorf("expected 50, got %s", vals.Get("refundAmount"))
	}
	if vals.Get("outRefundNo") != "refund-001" {
		t.Errorf("expected refund-001, got %s", vals.Get("outRefundNo"))
	}
}

func TestRefundOptionsAlwaysSave(t *testing.T) {
	trueVal := true
	r := &RefundOptions{AlwaysSave: &trueVal}
	vals := r.URLValues()
	if vals.Get("alwaysSave") != "true" {
		t.Errorf("expected true, got %s", vals.Get("alwaysSave"))
	}

	r2 := &RefundOptions{}
	vals2 := r2.URLValues()
	if vals2.Get("alwaysSave") != "" {
		t.Errorf("expected empty, got %s", vals2.Get("alwaysSave"))
	}
}

func TestRefundOptionsString(t *testing.T) {
	r := &RefundOptions{AppID: "app-1"}
	s := r.String()
	if s == "" {
		t.Error("expected non-empty JSON string")
	}
}

func TestNotifyOptionsURLValues(t *testing.T) {
	n := &NotifyOptions{
		AppID:   "app-1",
		OrderNo: "order-001",
		Status:  TradeStatusSuccess,
	}

	vals := n.URLValues()
	if vals.Get("appID") != "app-1" {
		t.Errorf("expected app-1, got %s", vals.Get("appID"))
	}
	if vals.Get("status") != "success" {
		t.Errorf("expected success, got %s", vals.Get("status"))
	}
}

func TestNotifyOptionsStatusMethods(t *testing.T) {
	n := &NotifyOptions{Status: TradeStatusSuccess}
	if !n.IsSuccess() {
		t.Error("expected IsSuccess()=true")
	}
	if n.IsFailure() {
		t.Error("expected IsFailure()=false")
	}
	if n.IsCancelled() {
		t.Error("expected IsCancelled()=false")
	}

	n.Status = TradeStatusFailure
	if !n.IsFailure() {
		t.Error("expected IsFailure()=true")
	}

	n.Status = TradeStatusCancelled
	if !n.IsCancelled() {
		t.Error("expected IsCancelled()=true")
	}
}

func TestNotifyOptionsRefundFields(t *testing.T) {
	n := &NotifyOptions{
		AppID:       "app-1",
		OutRefundNo: "refund-001",
		RefundNo:    "platform-refund-001",
		RefundAmount: 25.00,
	}

	vals := n.URLValues()
	if vals.Get("outRefundNo") != "refund-001" {
		t.Errorf("expected refund-001, got %s", vals.Get("outRefundNo"))
	}
	if vals.Get("refundNo") != "platform-refund-001" {
		t.Errorf("expected platform-refund-001, got %s", vals.Get("refundNo"))
	}
	if vals.Get("refundAmount") != "25" {
		t.Errorf("expected 25, got %s", vals.Get("refundAmount"))
	}
}

func TestNotifyOptionsString(t *testing.T) {
	n := &NotifyOptions{AppID: "app-1"}
	s := n.String()
	if s == "" {
		t.Error("expected non-empty JSON string")
	}
}

func TestTradeStatusConstants(t *testing.T) {
	tests := []struct {
		status TradeStatus
		method func(TradeStatus) bool
		want   bool
	}{
		{TradeStatusSuccess, TradeStatus.IsSuccess, true},
		{TradeStatusFailure, TradeStatus.IsFailure, true},
		{TradeStatusCancelled, TradeStatus.IsCancelled, true},
		{TradeStatusExpired, TradeStatus.IsExpired, true},
		{TradeStatusDeleted, TradeStatus.IsDeleted, true},
		{TradeStatusSuccess, TradeStatus.IsFailure, false},
		{TradeStatusFailure, TradeStatus.IsSuccess, false},
	}

	for _, tt := range tests {
		got := tt.method(tt.status)
		if got != tt.want {
			t.Errorf("unexpected result for status %s", tt.status)
		}
	}
}

func TestSetDefaults(t *testing.T) {
	config := map[string]string{
		"appId":     "app-from-config",
		"notifyUrl": "https://example.com/notify",
		"productId": "prod-001",
	}

	getter := func(key string) string { return config[key] }

	// CheckoutOptions
	co := &CheckoutOptions{}
	co.SetDefaults(getter)
	if co.AppID != "app-from-config" {
		t.Errorf("expected app-from-config, got %s", co.AppID)
	}
	if co.NotifyURL != "https://example.com/notify" {
		t.Errorf("expected notify URL, got %s", co.NotifyURL)
	}

	// NotifyOptions
	no := &NotifyOptions{}
	no.SetDefaults(getter)
	if no.AppID != "app-from-config" {
		t.Errorf("expected app-from-config, got %s", no.AppID)
	}

	// RefundOptions
	ro := &RefundOptions{}
	ro.SetDefaults(getter)
	if ro.AppID != "app-from-config" {
		t.Errorf("expected app-from-config, got %s", ro.AppID)
	}
}

func TestSetDefaultsDontOverwrite(t *testing.T) {
	getter := func(key string) string { return "should-not-override" }

	co := &CheckoutOptions{AppID: "existing"}
	co.SetDefaults(getter)
	if co.AppID != "existing" {
		t.Errorf("expected existing, got %s", co.AppID)
	}
}

func TestPaymentURL(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	c := &CheckoutOptions{
		AppID:     "test-id",
		Price:     99.99,
		OutOrderNo: "order-001",
		Subject:   "Test",
		Type:      "alipay",
	}

	urlStr, err := opts.PaymentURL(c)
	if err != nil {
		t.Fatalf("PaymentURL failed: %v", err)
	}
	if urlStr == "" {
		t.Error("expected non-empty URL")
	}
}

func TestPaymentProviderListURL(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	urlStr, err := opts.PaymentProviderListURL("test-id")
	if err != nil {
		t.Fatalf("PaymentProviderListURL failed: %v", err)
	}
	if urlStr == "" {
		t.Error("expected non-empty URL")
	}
}

func TestPaymentProvierListURLDeprecated(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	urlStr, err := opts.PaymentProvierListURL("test-id")
	if err != nil {
		t.Fatalf("PaymentProvierListURL failed: %v", err)
	}
	if urlStr == "" {
		t.Error("expected non-empty URL")
	}
}

func TestPaymentURLAppIDConflict(t *testing.T) {
	app := &testAppInfo{AppID: "config-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	c := &CheckoutOptions{
		AppID:     "different-id",
		Price:     99.99,
		OutOrderNo: "order-001",
		Subject:   "Test",
		Type:      "alipay",
	}

	_, err := opts.PaymentURL(c)
	if err == nil {
		t.Error("expected error for AppID conflict")
	}
}

func TestRefundURL(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	r := &RefundOptions{
		AppID:        "test-id",
		OutOrderNo:   "order-001",
		RefundAmount: 50.00,
	}

	urlStr, err := opts.RefundURL(r)
	if err != nil {
		t.Fatalf("RefundURL failed: %v", err)
	}
	if urlStr == "" {
		t.Error("expected non-empty URL")
	}
}

func TestPaymentQueryURL(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	urlStr, vals, err := opts.PaymentQueryURLWithValues("test-id", "order-no", "out-order-no")
	if err != nil {
		t.Fatalf("PaymentQueryURLWithValues failed: %v", err)
	}
	if urlStr == "" {
		t.Error("expected non-empty URL")
	}
	if vals == nil {
		t.Error("expected non-nil values")
	}
}

func TestClientPaymentQueryURL(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	urlStr, vals, err := opts.ClientPaymentQueryURLWithValues("test-id", "out-order-no")
	if err != nil {
		t.Fatalf("ClientPaymentQueryURLWithValues failed: %v", err)
	}
	if urlStr == "" {
		t.Error("expected non-empty URL")
	}
	if vals == nil {
		t.Error("expected non-nil values")
	}
}

func TestPaystartURL(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	c := &CheckoutOptions{
		AppID:     "test-id",
		Price:     99.99,
		OutOrderNo: "order-001",
		Subject:   "Test",
		Type:      "alipay",
	}

	urlStr, err := opts.PaystartURL(c)
	if err != nil {
		t.Fatalf("PaystartURL failed: %v", err)
	}
	if urlStr == "" {
		t.Error("expected non-empty URL")
	}
}

func TestPaymentURLWithValues(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(sdk_options.TypePayment, app)

	c := &CheckoutOptions{
		AppID:     "test-id",
		Price:     99.99,
		OutOrderNo: "order-001",
		Subject:   "Test",
		Type:      "alipay",
	}

	urlStr, vals, err := opts.PaymentURLWithValues(c)
	if err != nil {
		t.Fatalf("PaymentURLWithValues failed: %v", err)
	}
	if urlStr == "" {
		t.Error("expected non-empty URL")
	}
	if vals == nil {
		t.Error("expected non-nil values")
	}
}

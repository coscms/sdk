# COSCMS SDK

Go SDK for the COSCMS open platform API. Provides signed request building, payment integration, OAuth, and utility functions.

## 安装

```bash
go get github.com/coscms/sdk
```

## 概览

SDK 包含三个子包：

| 包 | 说明 |
|---|---|
| `sdk_options` | 核心选项、签名生成、API 请求提交 |
| `sdk_payment` | 支付相关：下单、退款、通知处理、交易状态 |
| `sdk_utils`  | 工具函数：MD5 哈希、HTML 标签清除 |

---

## sdk_options — 核心

### 应用凭证

实现 `AppInfoGetter` 接口以提供应用凭证：

```go
type AppInfoGetter interface {
    GetAppID() string
    GetAppSecret() string
    GetApiEndpoint() string
}
```

内置的 `AppInfo` 结构体可直接使用：

```go
import "github.com/coscms/sdk/sdk_options"

app := &sdk_options.AppInfo{
    AppID:       "your-app-id",
    Secret:      "your-app-secret",
    ApiEndpoint: "https://api.example.com",
}
```

### 创建 Options

使用 `New()` 和可选的功能选项：

```go
opts := sdk_options.New(sdk_options.TypePayment, app,
    sdk_options.OptionUserAgent("MyApp/1.0"),
    sdk_options.OptionClientIP("127.0.0.1"),
)
```

对于 OAuth 类型：

```go
opts := sdk_options.New(sdk_options.TypeOauth, app)
```

### API 接口地址生成

```go
// 签名并构建请求地址
uri, formData, err := opts.ToURLWithGenerator(generator, "/api/v1/endpoint")
if err != nil {
    // 处理错误
}

// 完整请求 URL
requestURL := uri + "?" + formData.Encode()
```

如果表单数据中的 `appID` 与配置的 appID 不一致，会返回 `ErrAppIDConflict`。

### 安全强化密钥

通过混合 UserAgent 和 ClientIP 到密钥中防止签名重放：

```go
// 第三个参数传入 true 启用安全强化
uri, formData, err := opts.ToURLWithGenerator(generator, "/api/v1/endpoint", true)
```

### OAuth 供应商列表

```go
urlStr, err := opts.OauthProviderListURL()
```

### 自定义签名函数

```go
opts.SetSigner(func(vals url.Values, secret string) string {
    // 自定义签名逻辑
    return customSignature
})
```

### 提交 API 请求

```go
import "context"

// 返回结构化 Response
resp, err := sdk_options.Submit(ctx, apiURL, formData)

// 或解析到自定义结构体
var result myStruct
_, err := sdk_options.SubmitWithRecv(ctx, &result, apiURL, formData)

// 或获取原始 map 响应
apiResp, body, err := sdk_options.Submitx(ctx, apiURL, formData)
```

`Response` 结构包含 `Code`、`State`、`Info`、`Data` 等字段，使用 `IsSuccess()` 判断是否成功（`Code == 1`）。

### 签名函数

```go
sig := sdk_options.SignString("raw data", "secret")     // 生成签名
err := sdk_options.CheckSign("raw data", sig, "secret")  // 验证签名（返回 ErrInvalidSign 或 nil）
sig = sdk_options.GenSign(formData, "secret")            // 从 url.Values 生成签名
```

### 辅助函数

```go
// 从 map 中取值，支持后备键
val := sdk_options.GetValueByKey(data, "primaryKey", "fallbackKey")
```

---

## sdk_payment — 支付

### 创建支付客户端

```go
import (
    "github.com/coscms/sdk/sdk_options"
    "github.com/coscms/sdk/sdk_payment"
)

app := &sdk_options.AppInfo{
    AppID:       "your-app-id",
    Secret:      "your-app-secret",
    ApiEndpoint: "https://api.example.com",
}

pay := sdk_payment.New(sdk_options.TypePayment, app)
```

### 下单（Checkout）

```go
order := &sdk_payment.CheckoutOptions{
    AppID:       "your-app-id",
    Price:       99.99,
    OutOrderNo:  "order-2026052301",
    Subject:     "商品名称",
    Type:        "alipay", // 或 "wechat"
    NotifyURL:   "https://your-site.com/notify",
    ReturnURL:   "https://your-site.com/success",
    ProductID:   "prod-001",
}

// 直接支付地址（GET 方式）
urlStr, err := pay.PaymentURL(order)

// 或获取 URL 和表单数据
urlStr, formData, err := pay.PaymentURLWithValues(order)

// 支付方式选择页地址
urlStr, err := pay.PaystartURL(order)
```

`CheckoutOptions` 支持从配置填充默认值：

```go
order := &sdk_payment.CheckoutOptions{}
order.SetDefaults(func(key string) string {
    return config[key] // 从配置文件读取
})
```

### 退款（Refund）

```go
refund := &sdk_payment.RefundOptions{
    AppID:        "your-app-id",
    OutOrderNo:   "order-2026052301",
    RefundAmount: 50.00,
    OutRefundNo:  "refund-2026052301",
}

urlStr, err := pay.RefundURL(refund)

// 或获取 URL 和表单数据
urlStr, formData, err := pay.RefundURLWithValues(refund)
```

### 支付查询

```go
// 服务端查询
urlStr, formData, err := pay.PaymentQueryURLWithValues("app-id", "orderNo", "outOrderNo")

// 客户端查询
urlStr, formData, err := pay.ClientPaymentQueryURLWithValues("app-id", "outOrderNo")
```

### 退款查询

```go
urlStr, formData, err := pay.RefundQueryURLWithValues("app-id", "refundNo", "outRefundNo")
urlStr, formData, err := pay.ClientRefundQueryURLWithValues("app-id", "outRefundNo")
```

### 支付通知处理

```go
notify := &sdk_payment.NotifyOptions{
    Status: sdk_payment.TradeStatusSuccess,
}

if notify.IsSuccess() {
    // 支付成功处理
} else if notify.IsFailure() {
    // 支付失败处理
} else if notify.IsCancelled() {
    // 支付取消处理
}
```

### 交易状态

```go
// 状态常量
sdk_payment.TradeStatusSuccess   // "success"   — 成功
sdk_payment.TradeStatusFailure   // "failure"   — 失败
sdk_payment.TradeStatusCancelled // "cancelled" — 被取消
sdk_payment.TradeStatusExpired   // "expired"   — 已过期
sdk_payment.TradeStatusDeleted   // "deleted"   — 已删除

// 状态判断方法
status := sdk_payment.TradeStatus("success")
status.IsSuccess()    // true
status.IsFailure()    // false
status.IsCancelled()  // false
status.IsExpired()    // false
status.IsDeleted()    // false
```

### 支付网关供应商列表

```go
urlStr, err := pay.PaymentProviderListURL("app-id")
urlStr, err := pay.PaymentProviderListURL("app-id", "CNY", "USD") // 筛选货币
```

---

## sdk_utils — 工具函数

```go
import "github.com/coscms/sdk/sdk_utils"

// MD5 哈希
hash := sdk_utils.Md5("hello") // "5d41402abc4b2a76b9719d911017c592"

// 清除 HTML/XML 标签（含 script、style 内容、注释）
text := sdk_utils.StripTags("<b>bold</b>")         // "bold"
text := sdk_utils.StripTags("<script>alert(1)</script>visible") // "visible"
```

---

## 完整的支付示例

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/coscms/sdk/sdk_options"
    "github.com/coscms/sdk/sdk_payment"
)

func main() {
    app := &sdk_options.AppInfo{
        AppID:       "your-app-id",
        Secret:      "your-app-secret",
        ApiEndpoint: "https://api.example.com",
    }

    pay := sdk_payment.New(sdk_options.TypePayment, app)

    // 1. 创建订单
    order := &sdk_payment.CheckoutOptions{
        AppID:      "your-app-id",
        Price:      199.00,
        OutOrderNo: "order-2026052301",
        Subject:    "VIP 会员",
        Type:       "alipay",
        NotifyURL:  "https://your-site.com/payment/notify",
    }

    payURL, err := pay.PaymentURL(order)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("支付地址:", payURL)

    // 2. 处理通知
    notify := &sdk_payment.NotifyOptions{
        AppID:     "your-app-id",
        OrderNo:   "platform-order-001",
        OutOrderNo: "order-2026052301",
        Status:    sdk_payment.TradeStatusSuccess,
    }

    if notify.IsSuccess() {
        // 更新订单状态...
    }
}
```

## 开发

```bash
# 运行所有测试
go test ./...

# 运行特定包测试
go test ./sdk_payment/...
go test ./sdk_options/...
go test ./sdk_utils/...
```

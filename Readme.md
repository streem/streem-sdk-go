# Streem SDK

A Go sdk utility for creating _JWT_ token with _JWS_ signature as per _IETF_ [RFC 7515](https://tools.ietf.org/html/rfc7515) specifications. It mimics token generation implementation by _streem sdk_ for nodejs project and generates token with _Compact Serialization_ technique.

## Usage

### Init

Use `Init()` method to initialize the Streem SDK. `Init` sets the package's config. This is required before creating a `TokenBuilder`

```go
func Init(apiKeyId, apiKeySecret, apiEnvironment string) (error)
```

#### Params required

1. _apiKeyId_
2. _apiKeySecret_
3. _apiEnvironment_

_Key Id_, _Secret_ and _Environment_ corresponds to Streem account for which token builder need to be configured.

### NewTokenBuilder

To create a Streem Token, first create a `TokenBuilder`

```go
builder := streem.NewTokenBuilder()
```

Then specify the details for the currently logged-in user:

```go
// required
builder.SetUserId("someId")

// recommended
builder.SetName("T Rex")
builder.SetEmail("some@email.com")
builder.SetAvatarUrl("some.avatar.url")

// optional
addDurationToNow := builder.AddTime()
builder.SetTokenExpirationMs(addDurationToNow(time.Minute * 30))
builder.SetSessionExpirationMs(addDurationToNow(time.Hour * 12))

// If using the Group Reservation feature, set the reservation sid from the API response
builder.SetReservationSid("rsv_abc123")
```

Finally, call `build()` to generate the token string:

```go
token, err := builder.Build()
if err != nil {
    fmt.Println("err: ", err)
    return
}

fmt.Printf("Got token %s\n", token)
```

### Example

```go
err := streem.Init(apiKeyId, apiKeySecret, apiEnvironment)
if err != nil {
    fmt.Println("err: ", err)
    os.Exit(1)
}

builder := streem.NewTokenBuilder()

// required
builder.SetUserId("someId")

// recommended
builder.SetName("T Rex")
builder.SetEmail("some@email.com")
builder.SetAvatarUrl("some.avatar.url")

// optional
addDurationToNow := builder.AddTime()
builder.SetTokenExpirationMs(addDurationToNow(time.Minute * 30))
builder.SetSessionExpirationMs(addDurationToNow(time.Hour * 12))

token, err := builder.Build()
if err != nil {
    fmt.Println("err: ", err)
    return
}

fmt.Printf("Got token %s\n", token)
```

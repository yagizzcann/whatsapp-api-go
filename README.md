# whatsapp
A Whatsapp's Official API helper for Golang

## Disclaimers

1. Whatsapp's Official API is now generally available.
   To get started, you can follow [this steps](https://developers.facebook.com/docs/whatsapp/getting-started/signing-up).



## Set up

First, you need a Facebook app with the Whatsapp API activated.
You can create your first app following [this steps](https://developers.facebook.com/docs/whatsapp/getting-started/signing-up).
Get the API token, either a temporal or a permanent one.

In your server you can install the module using :

```
go get github.com/yagizzcann/whatsapp-api-go
```

Now you can write code like this:

```go
import (
"github.com/yagizzcann/whatsapp-api-go"
)

Token := "YOUR_TOKEN";

api := whatsapp.API{
URI:     whatsapp.DefaultURI,
Version: whatsapp.DefaultVersion,
Token:   Token,
WebHookVerification:"custom webhook verify token ",
}

t := api.NewText("MESSAGE_CONTENT", false)
res, err := t.Send("YOUR_PHONE_ID", "TO_PHONE_NUMBER")

m := api.NewMediaLink("https://img_url", "image")
res, err := m.Send("YOUR_PHONE_ID", "TO_PHONE_NUMBER")


```

To recieve the post requests on message, you must setup the webhook at your Facebook app.
While setting up, you will be asked a Verify Token. This can be any string you want.

The app also has a GET wizard for the webhook authentication:

```go
cq := make(chan whatsapp.WebhookMessage, 100)
r := mux.NewRouter()
r.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) { whatsapp.WebhookVerificationHandler(w, r, cq) }).Methods("GET")
r.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) { whatsapp.WebhookEventHandler(w, r, cq) }).Methods("POST")
```


And that's it! Now you have a functioning Whatsapp Bot connected to your server.



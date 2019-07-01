# go-sendpulse
A very small and basic SendPulse SMTP API client library written in Golang

This was written open sourced by the [NetVPX](https://netvpx.com) DevOps team for internal usage. This library is now maintained by Aditya Diwakar, not affiliated with NetVPX.

## Install

To use this library, simply install it using:
```
go get github.com/netvpx/go-sendpulse
```

If you want to set it up as a submodule, use:
```
git submodule add https://github.com/netvpx/go-sendpulse github.com/netvpx/go-sendpulse
```

## Usage

The library is very basic and rudementary.

#### Initializing 

In order to use the library, we include a ``Initialize()`` function.

In order to use this, you must provide the ``client_id``, ``client_secret``, ``client_name``, and ``client_email``.

``client_name`` and ``client_email`` are two fields that are presented in the email.

Set it up like so:
```go
sendpulse.Initialize(
    CLIENT_ID,
    CLIENT_SECRET,
    CLIENT_NAME,
    CLIENT_EMAIL,
)
```


#### Getting your key

Your client identifier and client secret are only useful for generating your OAuth access token which can be used to make requests.

This is managed by the library and you will never use this.

```go
func getKey() (string, error)
```

This returns a string which is the access token.

#### Sending an email

You can send an email using this library. We do not support attachments, quite yet.

```go
func SendEmail(html []byte, text []byte, subject string, to []recipient) error
```

The ``html`` byte slice can be anything and needs to be valid HTML. This library does not validate HTML but... y'know ;) You should use proper HTML.

The ``text`` byte slice is the same as above, needs to be only text and NOT html.

The ``subject`` is pretty obvious.

The ``to`` field is an array of recipients which is defined as... 
```go
sendpulse.Recipient{
    Name: "Some Name",
    Email: "example@netvpx.com",
}
```

###### Example

The above endpoint can be hit like so... which will send an email with a subject of ``Hey There`` with the body of ``Peekaboo!`` in bold if HTML is supported.

```go
package main

import (
    "github.com/netvpx/go-sendpulse"
    "log"
)

func main() {
    html := []byte("<strong>Peekaboo!</strong>")
    text := []byte("Peekaboo!")
    recipients := []Recipient{
        Recipient{
            Name: "Some Name",
            Email: "example@netvpx.com",
        },
    }
    subject := "Hey There"

    sendpulse.Initialize(
        CLIENT_ID,
        CLIENT_SECRET,
        CLIENT_NAME,
        CLIENT_EMAIL,
    )   
    err := sendpulse.SendEmail(
        html,
        text,
        subject,
        recipients,
    )

    log.Fatalln(err)
}
```


## Support

If you need help, feel free to open a GitHub issue or email ``aditya@netvpx.com`` who is the Head of R&D at NetVPX.

## Contributing

Feel free to contribute, please follow ``gofmt`` for formatting and make a pull request. No guarantees that contributions will be accepted, thanks for your interest.

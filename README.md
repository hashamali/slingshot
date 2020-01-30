# Slingshot: Forward Form Responses to Email

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Slingshot is a web application that forwards the body of a POST request to email. You provide SMTP email configuration details and Slingshot handles the rest. Deploy to Heroku's free tier to get going immediately, simply click the above button.


## API

Slingshot exposes two endpoints: `/ping` and `/form`.

* `/ping`: [GET] Returns a 204 if all is well.
* `/form`: [POST] Sends the body to the configured recipients.

#### Examples

To check if the service is up:

`curl -I --request GET 'localhost:8080/ping'`

To send some form data:

```
curl -I --location --request POST 'localhost:8080/form' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'abc1=1234'
```

And lastly, see an example of an HTML form using Slingshot at [example/index.html](example/index.html).


## Configure

Slingshot can be configured via environment variables or by editing `config.yaml`. Note that `PORT` must always be provided via environment variable.

* `PORT`: This is the port that Slingshot runs on. If no PORT environment variable is provided, port 8080 is used.
* `CORS`: This is a comma separated list of the valid domains you would like to allow access to Slingshot. **This is not required but highly recommended.**

#### Via Environment Variables

* `SMTP_HOST`: The host of the SMTP email server to send with.
* `SMTP_PORT`: The port of the SMTP email server to send with.
* `SMTP_PASSWORD`: The password for your SMTP email account.
* `EMAIL_FROM`: The email account to send from.
* `EMAIL_TO`: The emails to send to. If more than one recipient, use a comma separated list.
* `RECAPTCHA_SECRET`: Optional if reCAPTCHA token validation is needed. 

###### Example

`PORT=8089 SMTP_HOST=smtp.example.com SMTP_PORT=587 SMTP_PASSWORD=PASSWORD EMAIL_FROM=you@example.com EMAIL_TO=recipient1@example.com,recipient2@example.com make run`

#### Via `config.yaml`

* `smtpHost`: The host of the SMTP email server to send with.
* `smtpPort`: The port of the SMTP email server to send with.
* `smtpPassword`: The password for your SMTP email account.
* `emailFrom`: The email account to send from.
* `emailTo`: The emails to send to.
* `recaptchaSecret`: Optional if reCAPTCHA token validation is needed. 

###### Example

```
smtpHost: smtp.example.com
smtpPort: 587
smtpPassword: PASSWORD
emailFrom: you@example.com
emailTo:
  - recipient1@example.com
  - recipient2@example.com
```


## Develop

#### Install

`go get github.com/hashamali/slingshot`

#### Run

If you have Go set up locally, you can run Slingshot using:

`PORT=8080 make run`

Alternatively, you can use Docker. First, build the Slingshot image:

`docker build --build-arg PORT=8080 -t slingshot .`

Then, run it:

`docker run -d -p 8080:8080 slingshot`


## Gmail Integration

Slingshot works quite well with Gmail. Setup is straight-forward:

1. Create an [App Password in Gmail](https://support.google.com/mail/answer/185833).
2. Configure `config.yaml`:
```
smtpHost: smtp.gmail.com
smtpPort: 587
smtpPassword: YOUR_APP_PASSWORD
emailFrom: YOUR_GMAIL_EMAIL@gmail.com
emailTo:
  - recipient1@example.com
  - recipient2@example.com
```


## Testing

Haven't gotten around to writing tests. Will add some soon if there is any interest.

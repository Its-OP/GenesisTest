# Genesis KMA SE School 3.0
## Prostakov Oleh
&nbsp;
# Important
**The application relies on SendGrid for sending emails. In order to enable sending emails by the application, you need to provide details of your SendGrid sender to the application, using environmental variables.**

**The next data needs to be passed to the application:**
1. Sender's name, via `SENDGRID_API_SENDER_NAME`.
2. Sender's email via `SENDGRID_API_SENDER_EMAIL`.
3. API key via `SENDGRID_API_KEY`.

**Information about your senders can be taken from this link:** https://app.sendgrid.com/settings/sender_auth/senders.
**The API key can be taken from here:** https://app.sendgrid.com/settings/api_keys.

**Placeholders for all these variables have been added to `docker-compose`. If you use it to run the application, please set the real values instead of the placeholders.**




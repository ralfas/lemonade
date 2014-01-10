Lemonade
========

# Intro

Lemonade is a notification service that calls you.

# Dependencies

Lemonade uses S3 and Twilio.

# Privacy

Lemonade uses Signed URLs to give Twilio access to the message stored in S3, so only Twilio has access to it.

Also, the messages stored in S3 are created with a short expiry.

# Costs

Each notification will:
- create a short lived S3 object; [S3 pricing](http://aws.amazon.com/s3/pricing/).
- make a call using Twilio; [Twilio pricing](http://www.twilio.com/voice/pricing).

# Usage

- Store your AWS Credentials at the location specified by the `AWS_CREDENTIAL_FILE` environment variable.
- Store your Twilio Credentials at the location specified by the `TWILIO_CREDENTIAL_FILE` environment variable.
- Start up the Lemonade service. Lemonade returns the API address: http://localhost:2024/
- Start your work process and have it notify the Lemonade Call API when it completes.
- Lemonade will call you when your work process completes.


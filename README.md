## Webhooks for detecting commit events

### What it does:
- The tool listens for webhook POST requests for push actions.
- When a push occurs in a monitored non-Github repo, a Webhook POST request is sent to a specified endpoint.
- Using a free ngrok tunnel URL as the endpoint.
- The request is validated vs a configured secret set in the repo webhook settings.
- Once validated, README is updated with the current timestamp and updated README is pushed to a private Github repo.

### Why:
- More normal usage would be in CI/CD pipelines.
- Here the test use case is to log the volume in restricted non-Github repos, e.g., school repos.
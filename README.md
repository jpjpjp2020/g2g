## Commit visibility tool: Mirroring commits in ringfenced/restricted non-Github repos

### What it does:
- The tool listens for webhook POST requests for push actions sent to a server endpoint.
- When a push to the non-Github repo happens, Webhook POST request is sent to a specified endpoint (free ngrok tunnel URL).
- The request is validated vs a configured secret set in the repo webhook.
- Once validated, README is updated with the current timestamp and updated README is pushed to a private Github repo.

### Why:
- To log commit volume in restricted non-Github repos, e.g., school repos with proprietary modules and study projects.
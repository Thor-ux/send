# Production Mailer on DigitalOcean using AWS

## Build & Run

1. Fork or Git clone resp.  
2. Run Docker build & up.  
3. Emails send automatically with **warm-up, reputation, bounce handling, daily reset**.  
4. Webhook listens on port 80 for SMTP/postback events.
5. The reputation score adjusts the delay automatically: if bounces/complaints happen, the system slows down se>
6. If you send too many emails too quickly, receiving servers (Gmail, Outlook, Yahoo, etc.) may flag your IP/do>
7. Gradual increase is key — called “IP/domain warm-up.”


## 👤 Author

Built by Daniel with ☕ and patience.

```bash
docker-compose build
docker-compose up -d
docker-compose down
docker logs -f mailer
# Feedposter

An Automation Engine to watch RSS Feeds and post them to Twitter! Runs on GitHub Actions without any database 🎯



## Configuration
1. Clone the Repository 
2. Edit `feed.yaml` to include [RSS Feeds](https://www.techopedia.com/definition/24756/rss-feed) you want to watch. `feed_url` consists of your RSS Feed
```yaml
   monitor:
  - feed:
    - item:
      name: KCD Chennai
      feed_url: https://dev.to/feed/kcdchennai
```
3. Get Tokens from Twitter with Elavated Access and it to Secrets on Repository Settings. The required tokens are:
    1. ACCESS_TOKEN_SECRET 
    2. ACCESS_TOKEN 
    3. CONSUMER_SECRET 
    4. CONSUMER_KEY 
4. The workflow will be running every 6 hours 






> **Warning**
> You need Twitter access Tokens

> **Note**
> This Runs on GitHub Actions

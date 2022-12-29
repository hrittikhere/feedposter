# Feedposter [![Go Report Card](https://goreportcard.com/badge/github.com/hrittikhere/feedposter)](https://goreportcard.com/report/github.com/hrittikhere/feedposter)


An Automation Engine to watch RSS Feeds and post them to Twitter! Runs on GitHub Actions without any database 🎯



## Configuration
1. Clone the Repository 

2. Edit `feed.yaml` to include [RSS Feeds](https://www.techopedia.com/definition/24756/rss-feed) you want to watch. `feed_url` consists of the URI of your RSS Feed

```yaml
   monitor:
  - feed:
    - item:
      name: KCD Chennai
      feed_url: https://dev.to/feed/kcdchennai
```

3. Get Tokens from Twitter with Elevated Access and it to Secrets on Repository Settings. The required tokens are:
    1. ACCESS_TOKEN_SECRET 
    2. ACCESS_TOKEN 
    3. CONSUMER_SECRET 
    4. CONSUMER_KEY 

4. The workflow will be running every 6 hours without any configuration update


## Twitter API Keys
To get Twitter API keys with full functionalities, you will need to follow the steps below:

1. Go to https://developer.twitter.com/en/docs and create a developer account.
1. After creating the developer account, you will need to create a new project. Click on the "Create Project" button and give your project a name.
1. After creating the project, you will need to create a developer application. Click on the "Create App" button and give your application a name.
1. Fill out the required information for your application, including a detailed description of what you plan to use the API for.
1. Once you have completed the application form, click on the "Create" button to submit your application.
1. Your application will be reviewed by Twitter's team, and you will receive an email when it has been approved.
Once your application has been approved, you will be able to access your API keys in the "Keys and Tokens" tab of your developer dashboard.
1. Click on the "Generate" button next to the "Consumer Keys" section to generate your API key and API secret.
1. Click on the "Generate" button next to the "Access Tokens" section to generate your access token and access token secret.
1. With these keys, you will be able to use the Twitter API to access various Twitter data, such as tweets, users, and trends. Keep in mind that the Twitter API has rate limits, so you will need to be mindful of how often you are making requests to the API. For Batch Tweets, like here it **won't** be an issue and you won't hit rate limits.

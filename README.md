# Discord Engagement Tweeter 🐦

Discord bot that creates tweets from discord messages in a server channel after a message gets certain amount of reactions.

Created with Go, using [Discordgo](https://github.com/bwmarrin/discordgo) package for handling communication with Discord and Twitter API
to communicate with Twitter.

## Installation

To install the bot, clone this repository. Compile the program with Go compiler `go build`.

To run the program run `./discord-engagement-tweeter` in the directory

You should see a message printed in the terminal that says `Bot is running...`

## Usage

To create the Discord bot **[Discord Application](https://discord.com/developers/applications?new_application=true)**
To post Tweets on behalf of your account you need to create **[Twitter Developer Account](https://developer.twitter.com/en)**

Before you can start you need to set these environment variables in `.env` file. There is an example `.env.example` file that you can use as a template.

Copy the credentials from the application settings in Discord and Twitter to the `.env`.

```
DISCORD_BOT_TOKEN=
X_CONSUMER_KEY=
X_CONSUMER_SECRET=
X_ACCESS_TOKEN=
X_TOKEN_SECRET=
CHANNEL_ID= id of the Discord server channel in which the bot watches for reactions
SEND_TWEET_AFTER= number of reactions to post the tweet after
```

The last step is to invite the bot to your server.

## License

MIT

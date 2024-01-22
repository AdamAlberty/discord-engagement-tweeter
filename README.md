# Discord Engagement Tweeter 🐦

Discord bot that posts tweets based on messages in a server channel once a specific number of reactions are received for a message.

Created with **Go**, using **[Discordgo](https://github.com/bwmarrin/discordgo)** package for handling communication with Discord and Twitter API
to communicate with Twitter (using [go-oauth1](github.com/klaidas/go-oauth1) for Twitter auth).

## Installation

To install the bot, clone and `cd` to this repository. Compile the program with Go compiler `go build`.

To run the program run the outputted binary `./discord-engagement-tweeter`.

Optionally, you can run the program with custom config path `./discord-engagement-tweeter --config /path/to/config.json`

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
```

```json
{
  "db_path": "messages.txt", <- name of file to store already tweeted messages to avoid duplicate tweets
  "reaction_threshold": 2, <- amount of reactions needed for message to be posted on Twitter
  "channel_ids": ["1175797847107575808"], <- IDs of channels in which the bot listens to reactions
  "ignore_after_hours": 72 <- Number of hours after which the bot ignores the new reactions (To avoid tweeting old messages)
}
```

The last step is to invite the bot to your server.

## Troubleshooting

The most probable cause of errors is missing credentials from `.env` file or wrong config format

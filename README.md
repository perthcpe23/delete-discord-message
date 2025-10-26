# Delete Discord Messages

A command-line tool to bulk delete messages from Discord direct message channels.

## Description

This tool allows you to delete all of your messages in one or more Discord channels. It fetches all messages in a channel and deletes the ones authored by you.

## Getting Started

### Prerequisites

- Go 1.22.2 or later

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/perthcpe23/delete-discord-message.git
   ```
2. Navigate to the project directory:
   ```sh
   cd delete-discord-message
   ```

## Usage

1. **Get your Discord token.** You can get it from the Discord web application.
2. **Get the channel ID(s).** You can get it from the Discord web application.

Run the tool with your token and channel ID(s):

```sh
go run main.go <token> <channel_id1> [channel_id2 channel_id3 ...]
```

**Example:**

```sh
go run main.go YOUR_DISCORD_TOKEN 123456789012345678 098765432109876543
```

## Building

You can build the executable using the Makefile:

```sh
make build
```

This will create an executable named `delete-discord-message` in the project root.

Then you can run it directly:

```sh
./delete-discord-message <token> <channel_id1> [channel_id2 channel_id3 ...]
```

## Disclaimer

This tool is not affiliated with Discord. Use it at your own risk. The author is not responsible for any damage caused by this tool. Be careful, as deleted messages cannot be recovered.

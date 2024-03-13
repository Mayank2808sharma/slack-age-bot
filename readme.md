# slack-age-bot

`slack-age-bot` is a slack bot based on `GO` that calculates user age.

### Getting Started

Follow the steps below to get started:

1. Clone the repository using Git:

   ```bash
   git clone https://github.com/Mayank2808sharma/slack-age-bot
   ```
2. Change to the project directory:

   ```bash
   cd slack-bot-clone
   ```
3. Create a .env file:
    ```bash
        SLACK_BOT_TOKEN="your workspace slack bot token"
        SLACK_APP_TOKEN="your workspace slack app token"
    ```
4. Build the binary:

   ```bash
   go build
   ```
### Usage
```bash
@age-bot my dob is <DOB>
```
```bash
@age-bot <DOB>
```
The above cmds you can use in your slack channel to interact with the bot

#### Example:

```bash
@age-bot my dob is 2003-08-28

// bot response:- Your age is 20.
```
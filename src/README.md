# feedbackBot

## This bot simplifies user feedback and support management.

### Features:
- Enables secure communication with users without sharing accounts.
- Facilitates team communication for multiple support agents.

# Building and running

## Prerequisites

### Software Requirements:
- [Golang](https://golang.org/doc/install) installed
- [PostgreSQL](https://www.postgresql.org/download/) server installed

### Configuration Steps:
1. Clone the repository:
    ```bash
    git clone https://github.com/ti-bone/feedbackBot.git
    cd feedbackBot
    ```
2. Copy the example configuration file:
    ```bash
    cp src/config.json.example config.json
    ```
3. Edit the `config.json` file with your configuration details.
4. Build the application:
    ```bash
    rm -rf build
    mkdir build
    cd src
    go build -o ../build/feedbackBot .
    cd ../build
    ```
5. Run the PostgreSQL server as described in the [PostgreSQL documentation](https://www.postgresql.org/docs/).
6. Create new user and database and specify them under `db_dsn` field in your `config.json` file.
7. Obtain a Telegram bot token:
   - Follow [Telegram's instructions on creating new bot](https://core.telegram.org/bots/features#creating-a-new-bot) to create a new bot and get the token.
8. Set the Telegram bot token in the `bot_token` field inside your `config.json` file.
9. Add the bot to your Telegram group, you can set `logs_id` to `0`, until you obtained the ID.
10. Enable topics in the group:
    - Follow [Telegram's introduction to topics](https://telegram.org/blog/topics-in-groups-collectible-usernames#topics-in-groups) to do so.
11. Run the application:
    ```bash
    ./feedbackBot
    ```
    - It automatically creates the required database tables and indexes.
12. Obtain ID of the Telegram group where the bot is added:
    - Send `/id` command to the group chat.
    - Copy the ID and set it in your `config.json` file.
13. Restart the application.
14. Start the bot by sending `/start` command to the bot in DM.
15. Set your `is_admin` to `true` in your database(it's required for anyone, who needs to reply to users):
    - Connect to your database using `psql` or any other client.
    - Obtain your user ID by sending `/id` command to the bot in DM.
    - Run the following query:
        ```sql
        UPDATE users SET is_admin = true WHERE id = <your_user_id>;
        ```
      Otherwise, you won't be able to reply to users via your group.

### Feel free to open an issue if you have any questions or suggestions.
### Pull requests are welcome!

# Example Screenshots

## Support Agent's (left) Interface and User's Interface (right):

| ![Support Agent's UI](https://static.bytefuck.dev/feedback-admin-side.png) | ![User's UI](https://static.bytefuck.dev/feedback-user-side.png)  |
|:--------------------------------------------------------------------------:|:-----------------------------------------------------------------:|


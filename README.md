# Slack bot on Golang

Hello, my dudes! If you are not sure if it's Wednesday or not, this bot will help you figure it out. Just because it can. 

![Kermit says hello](https://kingoflimericks.com/wp-content/uploads/2019/04/Kermit-and-the-Mythology-of-Muppets.jpg)

Pls, feel free to use this code as a template for your own bots. Let's populate memes over stuffy work channels.

### Install & run
The common flow to create Slack bot for memes:
- Create Slack App in your workspace: [Slack App page](https://api.slack.com/apps?new_app=1)
- Install [slack-go](https://github.com/slack-go/slack) lib
- Git clone this repo or write API handler yourself (who am I to force you to do smth?). A link to a useful guide is posted below. 
- Daemonize, dockerize, do whaterer you want with the app
- `go run main.go` (or whatever the command is)

### Useful links
- The initial config was based on [Develop a Slack-bot using Golang Guide](https://towardsdatascience.com/develop-a-slack-bot-using-golang-1025b3e606bc). Many thanks to @percybolmer !
- [Slack-go socketmode example](https://github.com/slack-go/slack/blob/master/examples/socketmode/socketmode.go)
- [Slack API](https://api.slack.com)
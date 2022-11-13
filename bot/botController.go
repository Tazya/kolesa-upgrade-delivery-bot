package bot

func LaunchBot(modifiedBot *ModifiedBot) {

	modifiedBot.Bot.Handle("/hello", modifiedBot.HelloHandler)
	modifiedBot.Bot.Start()
}

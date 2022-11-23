package bot

func LaunchBot(modifiedBot *ModifiedBot) {

	modifiedBot.Bot.Handle("/hello", modifiedBot.HelloHandler)
	modifiedBot.Bot.Handle("/start", modifiedBot.StartHandler)
	modifiedBot.Bot.Start()
}

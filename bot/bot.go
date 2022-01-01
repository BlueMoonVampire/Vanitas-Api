package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"syl-api/types"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func BotInit() {

	b, err := gotgbot.NewBot(os.Getenv("TOKEN"), &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})

	if err != nil {
		fmt.Println(err)
	}

	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				fmt.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			Panic:       nil,
			ErrorLog:    nil,
			MaxRoutines: 0,
		},
	})

	dispatcher := updater.Dispatcher

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\n", b.User.Username)

	dispatcher.AddHandler(handlers.NewCommand("check", check))
	dispatcher.AddHandler(handlers.NewCommand("start", start))

	updater.Idle()

}

func check(b *gotgbot.Bot, ctx *ext.Context) error {

	user := ctx.Args()[1]
	// ok, _ := strconv.Atoi(user)

	res, err := http.Get("https://sylviorus-api.up.railway.app/user/" + user)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var baka types.User
	json.Unmarshal(body, &baka)

	if baka.Blacklisted {
		b.SendMessage(ctx.Message.Chat.Id, fmt.Sprintf("%v is blacklisted\nReason: %v\nEnforcer: %v\nMessage: %v", baka.User, baka.Reason, baka.Enforcer, baka.Message), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})
	} else {
		b.SendMessage(ctx.Message.Chat.Id, fmt.Sprintf("%v is not blacklisted", baka.User), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})

	}
	return nil

}

func start(b *gotgbot.Bot, ctx *ext.Context) error {

	b.SendMessage(ctx.Message.Chat.Id, "Hello, I'm Syl Bot", &gotgbot.SendMessageOpts{
		ParseMode: "Markdown",
	})

	return nil

}

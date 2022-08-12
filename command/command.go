package command

import (
	"strings"

	"github.com/eggmoid/mattermost-chatbot/config"
	"github.com/mattermost/mattermost-server/v5/model"
)

type PluginAPI interface {
	GetBundlePath() (string, error)
}

type HandlerFunc func(context *model.CommandArgs, args ...string) *model.CommandResponse

type Handler struct {
	handlers       map[string]HandlerFunc
	defaultHandler HandlerFunc
}

const (
	commonHelpText = "###### Mattermost Plugin - Slash Command Help\n\n"
	botHelpText    = commonHelpText +
		"* `/bot list` - .\n" +
		"* `/bot app` - 애플리케이션 목록을 가져옵니다.\n" +
		"* `/bot env` - 배포 환경 목록을 가져옵니다.\n" +
		"* `/bot help` - 도움말.\n"
	invalidCommand = "유효하지 않은 명령어입니다."
)

var CommandHandler = Handler{
	handlers: map[string]HandlerFunc{
		"/bot/list": listCommand,
		"/bot/app":  appCommand,
		"/bot/env":  envCommand,
		"/bot/help": helpCommand,
	},
	defaultHandler: executeDefault,
}

func GetCommands(pAPI PluginAPI) []*model.Command {
	return []*model.Command{&model.Command{
		Trigger:          "bot",
		DisplayName:      "Bot",
		Description:      "명령어 목록: list, app, env, help",
		AutoComplete:     true,
		AutoCompleteDesc: "명령어 목록: list, app, env, help",
		AutoCompleteHint: "[command]",
	}}
}

func executeDefault(context *model.CommandArgs, args ...string) *model.CommandResponse {
	out := invalidCommand + "\n\n"
	out += botHelpText

	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         out,
	}
}

func postCommandResponse(context *model.CommandArgs, text string) {
	post := &model.Post{
		UserId:    config.BotUserID,
		ChannelId: context.ChannelId,
		Message:   text,
	}
	_ = config.Mattermost.SendEphemeralPost(context.UserId, post)
}

func (ch Handler) Handle(context *model.CommandArgs, args ...string) *model.CommandResponse {
	for n := len(args); n > 0; n-- {
		h := ch.handlers[strings.Join(args[:n], "/")]
		if h != nil {
			return h(context, args[n:]...)
		}
	}
	return ch.defaultHandler(context, args...)
}

func listCommand(context *model.CommandArgs, args ...string) *model.CommandResponse {
	// TODO: 명령어 목록을 가져오는 메서드를 추가해야 함.
	return &model.CommandResponse{}
}

func appCommand(context *model.CommandArgs, args ...string) *model.CommandResponse {
	// TODO: 어플리케이션 목록을 가져오는 메서드를 추가해야 함.
	return &model.CommandResponse{}
}

func envCommand(context *model.CommandArgs, args ...string) *model.CommandResponse {
	// TODO: 배포 환경 목록을 가져오는 메서드를 추가해야 함.
	return &model.CommandResponse{}
}

func helpCommand(context *model.CommandArgs, args ...string) *model.CommandResponse {
	postCommandResponse(context, botHelpText)
	return &model.CommandResponse{}
}

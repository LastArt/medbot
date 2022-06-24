package pkg

import (
	"log"
	"medcalcbot/calculators"
	"medcalcbot/set"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TgBotCreate() {
	bot, _ := tgbotapi.NewBotAPI(BotApiKey())

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	bot.Debug = false
	log.Printf("Подключение к боту - %s", bot.Self.UserName)
	updates := bot.GetUpdatesChan(u)
	warnings := new(Warning)
	descriptions := new(Description)
	admins := new(Admin)
	faqs := new(Faq)
	norms := new(Norms)
	recomends := new(Recomend)
	for update := range updates {
		if update.Message != nil {
			if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
				switch update.Message.Text {
				case "/start":
					userName := update.Message.From.FirstName
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "👋 Привет "+userName+"\n"+warnings.ShowinBot(set.IntroductionInfo))
					bot.Send(msg)
				case "/help":
					repl := warnings.ShowinBot(set.InfoText)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, repl)
					bot.Send(msg)
				case "/calculators":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Калькуляторы:")
					bot.Send(msg)
					calc := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					calc.ReplyMarkup = set.Calculators
					bot.Send(calc)
				case "/recomendations":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.FilesInfo))
					bot.Send(msg)
					for update := range updates {
						if update.Message != nil {
							if update.Message.Text != "в" {
								r := CheckRequest(update.Message.Text)
								if r == warnings.ShowinBot(set.OverTwoSymbols) {
									msg := tgbotapi.NewMessage(update.Message.Chat.ID, r)
									bot.Send(msg)
								} else if r == warnings.ShowinBot(set.NotRuSymbols) {
									msg := tgbotapi.NewMessage(update.Message.Chat.ID, r)
									bot.Send(msg)
								} else {
									r := CheckRequest(update.Message.Text)
									result := recomends.Search(r)
									prefix := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ReccomendPrefix)+" по запросу - \n\n"+r+"\n\n"+result+"\n\n\n"+warnings.ShowinBot(set.CalcBackWarning))
									bot.Send(prefix)
								}
							} else {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ExitNoti))
								bot.Send(msg)
								break
							}

						}
					}
				case "/medconf": //TODO Функция вызывающая кнопки по разделам медицины, а далее запускающая скраппинг по сайту медконференций в нужном разделе
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.UnderConstruction))
					bot.Send(msg)
				case "/settings": // Доступ которым в последствие возможно будет управлять ботом (организовать через инлайн батоны)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.BOT_ADMIN_ENTER_INFO))
					bot.Send(msg)
					for upd := range updates {
						msgIn := upd.Message.Text
						res := StringValuator(msgIn)
						if len(res) == 2 {
							bln := admins.Select(res[0], res[1])
							if bln != true {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.BOT_ADMIN_ACCESS_BAD))
								bot.Send(msg)
								break
							} else {
								keys := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.BOT_WARNING_ADMIN_ACCESS))
								keys.ReplyMarkup = set.SettingsKeyboard
								bot.Send(keys)
								break
							}
						} else {
							msg.Text = warnings.ShowinBot(set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH) + "2 значения"
							bot.Send(msg)
						}
					}
				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					x := CheckRequest(update.Message.Text)
					result := norms.Search(x)
					//prefix := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.NormsPrefix)+" по запросу - "+update.Message.Text+"\n---------\n")
					//bot.Send(prefix)
					if x != warnings.ShowinBot(set.NotRuSymbols) || x != warnings.ShowinBot(set.OverTwoSymbols) {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, x)
						bot.Send(msg)
					}
					prefix := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.NormsPrefix)+" по запросу - "+update.Message.Text+"\n---------\n")
					bot.Send(prefix)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, result)
					bot.Send(msg)
				}
			} else {
				//Отправлем сообщение если пользователь отправил картинку или стикер
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "😥К сожалению я обрабатываю только слова!")
				bot.Send(msg)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			// ТУТ НАЧИНАЕТСЯ РАБОТА С ОСНОВНЫМИ РАЗДЕЛАМИ
			switch update.CallbackQuery.Data {
			case "Базовый обмен веществ": // ГОТОВО
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "📟 Выбран: "+update.CallbackQuery.Data)
				bot.Send(msg)
				info := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, descriptions.ShowinBot(set.BaseChangeText))
				bot.Send(info)
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, faqs.ShowinBot(set.CaloryCalcInfo))
				bot.Send(msg)
				for update := range updates {
					if update.Message != nil {
						msgIn := update.Message.Text
						if msgIn != "в" {
							res := NumberValuator(msgIn)
							if len(res) < 5 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else if len(res) > 5 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💬Результат:\n\n"+calculators.Bov(res[0], res[1], res[2], res[3], res[4])+"\n"+warnings.ShowinBot(set.CalcBackWarning))
								msg.ReplyToMessageID = update.Message.MessageID
								bot.Send(msg)
							}
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ExitNoti))
							bot.Send(msg)
							break
						}
					}
				}
			case "Дефицит калия": // ГОТОВО!
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "📟 Выбран калькулятор: "+update.CallbackQuery.Data)
				bot.Send(msg)
				info := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, descriptions.ShowinBot(set.KaliumDefInfo))
				bot.Send(info)
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, faqs.ShowinBot(set.KaliumCalcInfo))
				bot.Send(msg)
				for update := range updates {
					if update.Message != nil {
						msgIn := update.Message.Text
						if msgIn != "в" {
							res := NumberValuator(msgIn)
							if len(res) < 2 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else if len(res) > 2 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💬Результат: \n"+calculators.KaliumDef(res[0], res[1])+"\n"+warnings.ShowinBot(set.CalcBackWarning))
								msg.ReplyToMessageID = update.Message.MessageID
								bot.Send(msg)
							}
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ExitNoti))
							bot.Send(msg)
							break
						}
					}
				}
			case "Скорость клубочковой фильтрации": // ГОТОВО!
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "📟 Выбран калькулятор: "+update.CallbackQuery.Data+warnings.ShowinBot(set.Tabulation)+descriptions.ShowinBot(set.SkfInfo)+warnings.ShowinBot(set.Tabulation)+faqs.ShowinBot(set.SkfCalcInfo))
				bot.Send(msg)
				for update := range updates {
					if update.Message != nil {
						msgIn := update.Message.Text
						if msgIn != "в" {
							res := NumberValuator(msgIn)
							if len(res) < 5 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else if len(res) > 5 {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💬Результат: \n"+calculators.Skf(res[0], res[1], res[2], res[3], res[4])+"\n"+warnings.ShowinBot(set.CalcBackWarning))
								msg.ReplyToMessageID = update.Message.MessageID
								bot.Send(msg)
							}
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ExitNoti))
							bot.Send(msg)
							break
						}
					}
				}
			case "Среднее АД": // ГОТОВО!
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "📟 Выбран калькулятор: "+update.CallbackQuery.Data)
				bot.Send(msg)
				info := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, descriptions.ShowinBot(set.MidADInfo))
				bot.Send(info)
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, faqs.ShowinBot(set.MidADCalcInfo))
				bot.Send(msg)
				for update := range updates {
					if update.Message != nil {
						msgIn := update.Message.Text
						if msgIn != "в" {
							res := NumberValuator(msgIn)
							if len(res) < 2 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else if len(res) > 2 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💬Результат: \n"+calculators.MidAD(res[0], res[1])+"\n"+warnings.ShowinBot(set.CalcBackWarning))
								msg.ReplyToMessageID = update.Message.MessageID
								bot.Send(msg)
							}
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ExitNoti))
							bot.Send(msg)
							break
						}
					}
				}
			case "Индекс РОМА": // ГОТОВО!
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "📟 Выбран калькулятор: "+update.CallbackQuery.Data)
				bot.Send(msg)
				info := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, descriptions.ShowinBot(set.RomaInfo))
				bot.Send(info)
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, faqs.ShowinBot(set.RomaCalcInfo))
				bot.Send(msg)
				for update := range updates {
					if update.Message != nil {
						msgIn := update.Message.Text
						if msgIn != "в" {
							res := NumberValuator(msgIn)
							if len(res) < 3 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else if len(res) > 3 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💬Результат: \n"+calculators.RomaIndex(res[0], res[1], res[2])+"\n"+warnings.ShowinBot(set.CalcBackWarning))
								//msg.ReplyToMessageID = update.Message.MessageID
								bot.Send(msg)
							}
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ExitNoti))
							bot.Send(msg)
							break
						}
					}
				}
			case "Индекс массы тела": // ГОТОВО!
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выбран калькулятор: "+update.CallbackQuery.Data)
				bot.Send(msg)
				info := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, descriptions.ShowinBot(set.ImtInfo))
				bot.Send(info)
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, faqs.ShowinBot(set.ImtCalcInfo))
				bot.Send(msg)
				for update := range updates {
					if update.Message != nil { // Если к нам поступило сообщение от пользователя
						msgIn := update.Message.Text
						if msgIn != "в" {
							res := NumberValuator(msgIn)
							if len(res) < 2 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else if len(res) > 2 {
								msg = tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ErrorInput))
								bot.Send(msg)
							} else {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💬Результат: \n"+calculators.Bmi(res[0], res[1])+"\n"+warnings.ShowinBot(set.CalcBackWarning))
								msg.ReplyToMessageID = update.Message.MessageID
								bot.Send(msg)
							}
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, warnings.ShowinBot(set.ExitNoti))
							bot.Send(msg)
							break
						}
					}
				}
			}
		}
	}
}

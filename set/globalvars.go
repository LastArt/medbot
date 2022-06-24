package set

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var Reg string = "[0-9]+(\\.[0-9]+)?(\\\\[0-9]+(\\.[0-9]+)?)*"
var RuSymbol string = "[А-Яа-я]+?"

// Calc description info
var (
	BaseChangeText string = "BaseChangeText"
	AlcoInfo       string = "AlcoInfo"
	ImtInfo        string = "ImtInfo"
	KaliumDefInfo  string = "KaliumDefInfo"
	MidADInfo      string = "MidADInfo"
	SkfInfo        string = "SkfInfo"
	RomaInfo       string = "RomaInfo"
)

//Calc faq
var (
	CaloryCalcInfo string = "CaloryCalcInfo"
	AlcoCalcInfo   string = "AlcoCalcInfo"
	KaliumCalcInfo string = "KaliumCalcInfo"
	ImtCalcInfo    string = "ImtCalcInfo"
	MidADCalcInfo  string = "MidADCalcInfo"
	SkfCalcInfo    string = "SkfCalcInfo"
	RomaCalcInfo   string = "RomaCalcInfo"
)

//MAIN warnings
var (
	InfoText          string = "InfoText"
	CalcBackWarning   string = "CalcBackWarning"
	Tabulation        string = "Tabulation"
	ErrorInput        string = "ErrorInput"
	ExitNoti          string = "ExitNoti"
	UnderConstruction string = "UnderConstruction"
	FilesInfo         string = "FilesInfo"
	ReccomendPrefix   string = "ReccomendPrefix"
	NormsPrefix       string = "NormsPrefix"
	IntroductionInfo  string = "IntroductionInfo"
	ChangeConnString  string = "ChangeConnString"
	ChangeBotToken    string = "ChangeBotToken"
	ChangeLogPath     string = "ChangeLogPath"
	OverTwoSymbols    string = "OverTwoSymbols"
	NotRuSymbols      string = "NotRuSymbols"
	HaveNotParametrs  string = "HaveNotParametrs"
)

// OTHER Warnings
var (
	BOT_ERROR_DELETE_USER            string = "BOT_ERROR_DELETE_USER"
	BOT_ERROR_ADDTODB                string = "BOT_ERROR_ADDTODB"
	BOT_ADMIN_ACCESS_BAD             string = "BOT_ADMIN_ACCESS_BAD"
	BOT_ERROR_EDITUSER               string = "BOT_ERROR_EDITUSER"
	BOT_WARNING_DELETE_FINE          string = "BOT_WARNING_DELETE_FINE"
	BOT_WARNING_ADD_FINE             string = "BOT_WARNING_ADD_FINE"
	BOT_WARNING_ADMIN_ACCESS         string = "BOT_WARNING_ADMIN_ACCESS"
	BOT_WARNING_EDITUSER_FINE        string = "BOT_WARNING_EDITUSER_FINE"
	BOT_WARNING_ARGUMENTS_NOT_ENOUGH string = "BOT_WARNING_ARGUMENTS_NOT_ENOUGH"
	BOT_ADMIN_ENTER_INFO             string = "BOT_ADMIN_ENTER_INFO"
)

var Calculators = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Индекс массы тела", "Индекс массы тела"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Дефицит калия", "Дефицит калия"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Базовый обмен веществ", "Базовый обмен веществ"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Среднее АД", "Среднее АД"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Скорость клубочковой фильтрации", "Скорость клубочковой фильтрации"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Индекс РОМА", "Индекс РОМА"),
	),
)

var ExitButton = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("↩️ выход", "↩️ выход"),
	),
)

// =================================================================================
// Кнопки для работы с настройками бота - Клавиатура
// =================================================================================
var SettingsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🗄 База данных"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📇 Изменить путь к Лог журналу"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📚 Лабараторные показатели"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🗂 Клинические рекомендации"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👨🏻‍💻 Администраторы"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("↩️-Назад"),
	),
)

type Config struct {
	TelegramBotApi   string
	ConnectionString string
	LogJournalPath   string
}

type Client struct {
	Bot *tgbotapi.BotAPI
}

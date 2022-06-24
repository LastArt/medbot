package pkg

import (
	"database/sql"
	"encoding/json"
	"log"
	"medcalcbot/set"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Admin struct {
	Id       int
	Login    string
	Password string
}

type Warning struct {
	Id        int
	WrnName   string
	WrnValues string
}

type Faq struct {
	Id       int
	FqName   string
	FqValues string
}

type Description struct {
	Id        int
	DscName   string
	DscValues string
}

type Recomend struct {
	Id                int
	RecomendationName string
	Link              string
	TypeRecomend      string
}

type Norms struct {
	Id          int
	Params      string
	ParamsValue string
}

type DbMethods interface {
	Insert()
	Delete()
	Update()
	Select()
	Search()
	ShowinBot()
}

func NewBot(token string) *set.Client {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	return &set.Client{
		Bot: bot,
	}
}

func DbConnect() {
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

}

func BotApiKey() string {
	file, err := os.Open("set/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := set.Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return conf.TelegramBotApi
}

func LogJournalPath() string {
	file, err := os.Open("set/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := set.Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return conf.LogJournalPath
}

func ConnString() string {
	file, err := os.Open("./set/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := set.Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return conf.ConnectionString
}

func (a Admin) Select(login, pass string) bool {
	result := false
	var lg, pwd string
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM admins WHERE login=? AND password=?", login, pass)
	if err != nil {
		Logger(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&a.Id, &a.Login, &a.Password)
		if err != nil {
			Logger(err)
			continue
		}
		lg = a.Login
		pwd = a.Password
	}
	if lg == "" && pwd == "" {
		result = false
	} else {
		result = true
	}
	return result
}

func (w Warning) ShowinBot(s string) string {
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT warningValue FROM warnings WHERE warningName = ?", s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var str string
	for rows.Next() {
		err := rows.Scan(&str)
		if err != nil {
			log.Println("ОШИБКА - ", err)
			continue
		}
		str += w.WrnValues
	}
	return str
}

func (d Description) ShowinBot(s string) string {
	db, err := sql.Open("mysql", ConnString())
	defer db.Close()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT descrValue FROM calc_description WHERE descrName = ?", s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&d.DscValues)
		if err != nil {
			log.Println("ОШИБКА - ", err)
			continue
		}
	}
	return d.DscValues
}

func (f Faq) ShowinBot(s string) string {
	db, err := sql.Open("mysql", ConnString())
	defer db.Close()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT faqValue FROM calc_faq WHERE faqName = ?", s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var str string
	for rows.Next() {
		err := rows.Scan(&str)
		if err != nil {
			log.Println("ОШИБКА - ", err)
			continue
		}
		str += f.FqValues
	}
	return str
}

func (n Norms) Search(inputParams string) string {
	db, err := sql.Open("mysql", ConnString())
	defer db.Close()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id, parametrs, value FROM norms WHERE parametrs LIKE  ?", "%"+inputParams+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var str string
	nrm := []Norms{}
	for rows.Next() {
		err := rows.Scan(&n.Id, &n.Params, &n.ParamsValue)
		if err != nil {
			log.Println("ОШИБКА - ", err)
			continue
		}
		nrm = append(nrm, n)
		str += "\n№: " + strconv.Itoa(n.Id) + "\nПоказатель: " + n.Params + "\n Значение: " + n.ParamsValue + "\n============================\n"
	}
	return str
}

func (r Recomend) Search(inputRecomends string) string {
	db, err := sql.Open("mysql", ConnString())
	defer db.Close()
	if err != nil {
		panic(err)
	}
	warnings := new(Warning)
	var str string
	res := CheckRequest(inputRecomends)

	if res == warnings.ShowinBot(set.NotRuSymbols) {
		str = "Некорректный ввод, пишите русскими буквами!"
	} else {
		rows, err := db.Query("SELECT id,recom_name, link, type FROM recom WHERE recom_name LIKE ?", "%"+inputRecomends+"%")
		if err != nil {
		}
		defer rows.Close()

		rcm := []Recomend{}
		for rows.Next() {
			err := rows.Scan(&r.Id, &r.RecomendationName, &r.Link, &r.TypeRecomend)
			if err != nil {
				log.Println("ОШИБКА - ", err)
				continue
			}
			rcm = append(rcm, r)
			str += "\n№ " + strconv.Itoa(r.Id) + "\nНозология: " + r.RecomendationName + "\nСсылка: " + r.Link + "\nИсточник: " + r.TypeRecomend + "\n============================\n"
		}
	}

	return str
}

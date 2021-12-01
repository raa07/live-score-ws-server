package main

import (
	"encoding/json"
	//"encoding/json"
	"fmt"
	"github.com/raa07/live-score-ws-server/data_provider"
)

func main() {
	//config, err := loadConfig()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	//var ctx = context.Background()
	//
	//var redisClient = redis.NewClient(&redis.Options{
	//	Addr: "localhost:6379",
	//})
	//err = NewWSServer(config.Server)

	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	serieData := data_provider.SerieData{}
	jsonString := "{\"selected_match\":1,\"dire_wins\":0,\"radiant_wins\":0,\"serie\":{\"id_external\":603946,\"is_deactivated\":0,\"type\":3,\"id_team_winner\":0,\"teamRadiant\":{\"id\":1,\"id_external\":8449479,\"name\":\"Team GL\",\"logo\":\"unknown_team.png\"},\"teamDire\":{\"id\":1,\"id_external\":8449479,\"name\":\"Team GL\",\"logo\":\"unknown_team.png\"}},\"matches\":{\"1\":{\"id_match\":1,\"buildings\":{\"radiant\":\"111111111\",\"dire\":\"111111111\",\"dire_barracks\":\"11111111\",\"radiant_barracks\":\"11111111\"},\"result\":{\"id_match\":1,\"gold\":-65,\"kills_dire\":0,\"kills_radiant\":0,\"buildings\":\"4784201\",\"barracks\":\"1111111111111111\",\"towers\":\"111111111111111111\",\"game_time\":\"11:05\"},\"playersDire\":[{\"id_acc\":250358373,\"main_name\":\"positive player\",\"hero_image\":\"lina_full.png\",\"hero_real_name\":\"LINA\",\"items\":[]},{\"id_acc\":1106500482,\"main_name\":\"GL.Jaden\",\"hero_image\":\"legion_commander_full.png\",\"hero_real_name\":\"LEGION COMMANDER\",\"items\":[]},{\"id_acc\":173851224,\"main_name\":\"666\",\"hero_image\":\"bane_full.png\",\"hero_real_name\":\"BANE\",\"items\":[]},{\"id_acc\":1117866283,\"main_name\":\"GL.Rengoku\",\"hero_image\":\"hoodwink_full.png\",\"hero_real_name\":\"HOODWINK\",\"items\":[]},{\"id_acc\":836108854,\"main_name\":\"\\u8c0b\\u6740\",\"hero_image\":\"medusa_full.png\",\"hero_real_name\":\"MEDUSA\",\"items\":[]}],\"playersRadiant\":[{\"id_acc\":1251664901,\"main_name\":\"tttokyo\",\"hero_image\":\"night_stalker_full.png\",\"hero_real_name\":\"NIGHT STALKER\",\"items\":[]},{\"id_acc\":1251319262,\"main_name\":\"CS.impulse\",\"hero_image\":\"invoker_full.png\",\"hero_real_name\":\"INVOKER\",\"items\":[]},{\"id_acc\":1239198638,\"main_name\":\"standin.\\u00d3\\u00f0inn\",\"hero_image\":\"lion_full.png\",\"hero_real_name\":\"LION\",\"items\":[]},{\"id_acc\":1158036449,\"main_name\":\"Progon\",\"hero_image\":\"tidehunter_full.png\",\"hero_real_name\":\"TIDEHUNTER\",\"items\":[]},{\"id_acc\":1251984026,\"main_name\":\"CS.daylight\",\"hero_image\":\"bloodseeker_full.png\",\"hero_real_name\":\"BLOODSEEKER\",\"items\":[]}],\"radiantBans\":[{\"hero_logo\":\"necrolyte_full.png\"},{\"hero_logo\":\"phoenix_full.png\"},{\"hero_logo\":\"storm_spirit_full.png\"},{\"hero_logo\":\"ember_spirit_full.png\"},{\"hero_logo\":\"void_spirit_full.png\"},{\"hero_logo\":\"terrorblade_full.png\"},{\"hero_logo\":\"windrunner_full.png\"}],\"direBans\":[{\"hero_logo\":\"necrolyte_full.png\"},{\"hero_logo\":\"phoenix_full.png\"},{\"hero_logo\":\"storm_spirit_full.png\"},{\"hero_logo\":\"ember_spirit_full.png\"},{\"hero_logo\":\"void_spirit_full.png\"},{\"hero_logo\":\"terrorblade_full.png\"},{\"hero_logo\":\"windrunner_full.png\"}]}}}\n"
	err := json.Unmarshal([]byte(jsonString), &serieData)
	if err != nil {
		fmt.Println(err)
	}
	res2B, _ := json.Marshal(serieData)
	fmt.Println(string(res2B))

}

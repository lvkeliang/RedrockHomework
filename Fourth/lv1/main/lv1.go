package main

import (
	"fmt"
	"strings"
)

var callskills map[string]string
var attackskills map[string]string

var call func(a string)
var attack func(a string)

var skill, skillSystem, skillname string
var releaseSkillFunc func(string)

func main() {
	//定义召唤系技能字典
	callskills = make(map[string]string)
	callskills["Call_Piccachu"] = "Come out! Piccachu!!!\n/＼7　　　 ∠＿/\n/　│　　 ／　／\n│　Z ＿,＜　／　　 /`ヽ\n│　　　　　ヽ　　 /　　〉\nY　　　　　`　 /　　/\nｲ●　､　●　　⊂⊃〈　　/\n()　 へ　　　　|　＼〈\n>ｰ ､_　 ィ　 │ ／／\n/ へ　　 /　ﾉ＜| ＼＼\nヽ_ﾉ　　(_／　 │／／\n7　　　　　　　|／\n＞―r￣￣`ｰ―＿\n\n"
	callskills["Call_Alpaca"] = "Come out! Alpaca!!!\n Δ~~~~Δ \nξ •ェ• ξ \n ξ　~　ξ \n ξ　　 ξ \n ξ　　　“~～~～〇 \n ξ　　　　　　 ξ \n ξ　ξ　ξ~～~ξ　ξ \n 　ξ_ξξ_ξ　ξ_ξξ_ξ "

	//定义召唤系技能
	call = func(a string) {
		fmt.Printf(callskills[a])
	}

	//定义攻击系技能字典
	attackskills = make(map[string]string)
	attackskills["Firevolt"] = "烧烧烧烧烧烧烧烧烧烧\n烧烧烧烧烧烧烧烧烧烧\n烧烧烧\n烧烧烧\n烧烧烧\n烧烧烧烧烧烧烧烧烧烧\n烧烧烧烧烧烧烧烧烧烧\n烧烧烧\n烧烧烧\n烧烧烧\n烧烧烧\n烧烧烧\n烧烧烧"
	attackskills["Lie"] = "  ┌───────────────┐\n　│通知：　 　    │\n　│今天晚上不上   |\n　│晚自习。 　    │\n　  (ﾖ─-∧＿∧─-E)\n　  ＼（* ´∀｀）／\n　 　 Y 　　　 Y"

	//定义攻击系技能
	attack = func(a string) {
		fmt.Printf(attackskills[a])
	}

	selectskill()
	//fmt.Printf("skill:%v\nskillSystem:%v\nskillname:%v\n", skill, skillSystem, skillname)
	ReleaseSkill(skillname, releaseSkillFunc)
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}

// 定义根据用户输入选择功能
func selectskill() {
	fmt.Println("请进行操作！")
	fmt.Scanf("%v", &skill)
	ind := strings.Index(skill, ":")
	skillSystem = skill[:ind]
	skillname = skill[ind+1:]
	if skillSystem == "call" {
		releaseSkillFunc = call
	} else if skillSystem == "attack" {
		releaseSkillFunc = attack
	}
}

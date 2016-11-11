package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
	return Re.MatchString(email)
}

func validateURL(url string) bool {
	//Re := regexp.MustCompile(`^[www.|http://|https://]*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	Re := regexp.MustCompile(`^(www.|https://|http://)*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	//[A-Za-z0-9./?=_%+\-]* al final para youtube.com/WATCH_03
	return Re.MatchString(url)
}
func validateID(id string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z]+[0-9]+$`)
	return Re.MatchString(id)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateDigits(digits string) bool {
	Re := regexp.MustCompile(`^[0-9]+$`)
	return Re.MatchString(digits)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateSteelMov(steel string) bool {
	re := regexp.MustCompile(`^(autotomize|steel\swing|iron\stail|metal\sclaw|meteor\smash|metal\ssound|iron\sdefence|doom\sdesire|gyro\sball|metal\sburst|mirror\sshot)+$`)
	return re.MatchString(steel)

	/*re:=regexp.MustCompile(`^[steel\swing|[iron\stail|[metal\sclaw||
	meteor\smash|metal\ssound|iron\sdefence|
	doom\sdesire|gyro\sball|metal\sburst|mirror\sshot]+$`)
	return re.MatchString(steel)*/
}
func validateFireMov(fire string) bool {
	re := regexp.MustCompile(`^(eruption|overheat|ember|fire\spunch|flame\sthrower|fire\sspin|fire\sblast|flame\swheel|sacred\sfire|sunny\sday|heat\swave|will\so\swhisp|blast\sburn|flare\sblitz|fire\sfang|lava\splume|magma\sstorm|flame\sburst)+$`)
	return re.MatchString(fire)
}
func validateBUGmov(bug string) bool {
	re := regexp.MustCompile(`^(megahorn|powder|infestation|twindeedle|bug\sbuzz|pin\smissile|string\sshot|leech\slife|spider\sweb|fury\scutter|tail\sglow|silver\swind|signal\sbeam|u\sturn|bug\sbite|attack\sorder|defend\sorder|heal\sorder|rage\spowder|struggle\sbug|sticky\sweb|fell\sstinger|quiver\sdance)+$`)
	return re.MatchString(bug)
}

func validateWATERmov(water string) bool {
	re := regexp.MustCompile(`^(soak|brine|dive|bubble|octazooka|whirpool|withdraw|waterfall|clamp|crabhammer|surf|hydro\spump|water\sgun|bubble\sbeam|rain\sdance|hydro\scannon|water\sspout|muddy\swater|water\spulse|aqua\stail|aqua\sjet|water\spledge|razor\sshell|water\sshuriken|origin\spulse)+$`)
	return re.MatchString(water)
}
func validateEARTHmov(earth string) bool {
	re := regexp.MustCompile(`^(fissure|earthquake|sand\sattack|bone\sclub|bonemearng|mud\sslap|bone\srush|magnitude|mud\ssport|sand\stomb|mud\sshot|earth\spower|mud\sbomb|drill\srun)+$`)
	return re.MatchString(earth)
}
func main() {
	f, _ := os.Open("C://Download//Pokemon//pokemon.txt")
	scanner := bufio.NewScanner(f)
	//parts := strings.Split(line," ")
	for scanner.Scan() {
		line := scanner.Text()

		if validateEmail(line) {
			fmt.Println("Es un Email: ", line)
		} else if validateURL(line) {
			fmt.Println("Es un URL: ", line)
		} else if validateID(line) {
			fmt.Println("Es un ID: ", line)
		} else if validateDigits(line) {
			fmt.Println("Es un Numero: ", line)
		} else if validateSteelMov(line) {
			fmt.Println("Steel attack: ", line)
		} else if validateFireMov(line) {
			fmt.Println("Fire Attack: ", line)
		} else if validateBUGmov(line) {
			fmt.Println("Bug attack: ", line)
		} else if validateWATERmov(line) {
			fmt.Println("Water attack: ", line)
		} else if validateEARTHmov(line) {
			fmt.Println("Earth attack", line)
		} else {
			fmt.Println("Error")
		}
	}

}


default: run

run:
	go run bot.go board.go fields.go monte.go
test:
	go run bot.go board.go fields.go monte.go < moves.txt
test2:
	go run bot.go board.go fields.go monte.go < moves2.txt
zip:
	zip bot.zip board.go bot.go fields.go monte.go
temp:
	go run bot.go fields.go < moves.txt

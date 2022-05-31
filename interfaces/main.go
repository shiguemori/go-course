package main

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb Bot) {
	println(eb.getGreeting())
}

func (eb EnglishBot) getGreeting() string {
	return "hi there!"
}

func (sb SpanishBot) getGreeting() string {
	return "hola!"
}

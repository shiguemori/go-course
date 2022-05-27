package main

import (
	"reflect"
	"testing"
)

var deckFortest = deck{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Diamonds", "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs"}

func Test_deal(t *testing.T) {
	type args struct {
		d        deck
		handSize int
	}
	tests := []struct {
		name  string
		args  args
		want  deck
		want1 deck
	}{
		{
			name: "Test_deal",
			args: args{
				d:        newDeck(),
				handSize: 5,
			},
			want:  deck{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades"},
			want1: deck{"Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Diamonds", "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := deal(tt.args.d, tt.args.handSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("deal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_deck_saveToFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		d       deck
		args    args
		wantErr bool
	}{
		{
			name: "Test_deck_saveToFile",
			d:    newDeck(),
			args: args{
				filename: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.saveToFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("saveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newDeck(t *testing.T) {
	tests := []struct {
		name string
		want deck
	}{
		{
			name: "Test_newDeck",
			want: deckFortest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDeckFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want deck
	}{
		{
			name: "Test_newDeckFromFile",
			args: args{
				filename: "test",
			},
			want: deckFortest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeckFromFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeckFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyFolder(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test_verifyFolder",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyFolder(); (err != nil) != tt.wantErr {
				t.Errorf("verifyFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

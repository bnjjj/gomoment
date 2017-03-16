[![GoDoc](https://godoc.org/github.com/bnjjj/gomoment?status.svg)](http://godoc.org/github.com/bnjjj/gomoment)
[![Build Status](https://travis-ci.org/bnjjj/gomoment.svg?branch=master)](https://travis-ci.org/bnjjj/gomoment)
# gomoment

_gomoment_ is a golang package to parse text that contain date or moment in french into a golang time struct

## Getting started

To use gomoment there is only 1 public function called GetDate, here is some examples:

```golang
func GetDate(moment string, duration bool, location *time.Location) (time.Time, time.Time, error)
```

+ Example for a simple moment date

```golang
begin, _, err := GetDate("Donne moi la date d'aujourd'hui", false, nil)
// begin is the date of today 0h0min
```


+ Other examples for duration

```golang
begin, end, err := GetDate("Combien de km j'ai réalisé depuis le mois dernier ?", true, nil)
// begin is today - 1 month and end is today 00:00
```

+ Example with another location

```golang
location, _ := time.LoadLocation("America/New_York")
begin, end, err := GetDate("Combien de km j'ai réalisé hier ?", true, location)
```

## Examples of text date that *gomoment* can parse


+ `avant-hier ?`
+ `la veille`
+ `hier`
+ `5 jours`
+ `7j`
+ `2 semaines`
+ `2sem`
+ `5mois`
+ `ce mois-ci`
+ `la semaine passée `
+ `cette semaine`
+ `cette annéee`
+ `5 décembre 2015`
+ `le 5 janvier`
+ `le 5/01/2017`

## References

It's the package used for [Talk to my car](http://www.talk-to-my-car.com)

## Contributions

Feel free to contribute and extend this package and if you have bugs or if you want more specs make an issue. Have fun !

-------------

Made by [Coenen Benjamin](https://twitter.com/BnJ25) with love

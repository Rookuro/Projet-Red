package main

import (
	"math/rand"
	"time"
)

var casque int
var armure int
var bottes int
var tour int = 0
var p *personnage = &personnage{}
var p1 personnage = personnage{}
var m *Monstre = &Monstre{}
var m1 Monstre = Monstre{}
var m2 Monstre = Monstre{}
var m3 Monstre = Monstre{}
var m4 Monstre = Monstre{}
var m5 Monstre = Monstre{}
var m6 Monstre = Monstre{}
var m7 Monstre = Monstre{}
var m8 Monstre = Monstre{}
var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)
var sort = 0
var sort2 = 0
var LimitInv int = 10
var epeefer = 0
var epeemidgard = 0
var hache = 0
var c = 0
var a = 0
var b = 0

# GO! Snake
Cheesy snake game implemented using GLFW and Go

This was hacked together over a day or so, so there's lots of unused code, but it basically works!

[![screenshot](https://github.com/awdavies/goSnake/raw/master/img/screenie.png)](#)

## Build Instructions.

```
$ go get -u github.com/go-gl/glfw/v3.2/glfw
$ go get -u github.com/go-gl/gl/v2.1/gl
$ go build main.go
$ ./main
```

## Controls

* Arrow keys steer your trusty snake!

* Esc exits the game.

So far there is no "you lose" screen yet, so the game just stops when you hit
an edge or run into yourself.  At which point you'll have to hit Esc and start
the game over again.

## Releases

The binary releases are built for Mac OSX and Linux (specifically amd64.  Will
get to 32 bit release when I'm feelin like it):

* [Linux 64-bit](https://github.com/awdavies/goSnake/raw/master/release/linux/main)

* [Mac OSX](https://github.com/awdavies/goSnake/raw/master/release/mac/main)

* Windows (Not yet...  Eventually, though)

### Running/Installing.

#### Mac OSX/Linux
```
$ cd ~/Downloads # Or wherever you have things downloaded.
$ chmod 755 main
$ ./main
```

#### Windows

Pending...

## Credits.

* The wonderful folks at [go-gl](https://github.com/go-gl)

* The creators of Go for giving me something to do.

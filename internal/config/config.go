package config

import (
	"flag"
)

var (
	DefSize              = 50    // стандартное количество игроков
	DefAttemptsPerPlayer = 25    // стандартное количество попыток
	DefSessionsCount     = 10000 // стандартный размер сессии
)

// Создаёт флаги для изменения стандартных параметров.
func ServerFlags() {
	flag.IntVar(&DefSize, "si", DefSize, "changes the number of size")
	flag.IntVar(&DefAttemptsPerPlayer, "a", DefAttemptsPerPlayer, "changes the number of attempts")
	flag.IntVar(&DefSessionsCount, "ss", DefSessionsCount, "changes the number of sessions")
	flag.Parse()
}

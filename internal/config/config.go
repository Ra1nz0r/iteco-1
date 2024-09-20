package config

import (
	"flag"
)

var (
	DefSize              = 50    // стандартное количество игроков
	DefAttemptsPerPlayer = 25    // стандартное количество попыток
	DefSessionsCount     = 10000 // стандартный размер сессии
)

// Создаёт флаги для запуска сервера, если в терминале переданы переменные окружения,
// то приоритет будет отдаваться им.
func ServerFlags() {
	flag.IntVar(&DefSize, "si", DefSize, "number of sessions")
	flag.IntVar(&DefAttemptsPerPlayer, "a", DefAttemptsPerPlayer, "number of sessions")
	flag.IntVar(&DefSessionsCount, "ss", DefSessionsCount, "number of sessions")
	flag.Parse()
}

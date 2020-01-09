package utils

const (
	luaAppVersion = "1.0b"
	//Android константа настройки
	Android = "Android"
	//PC константа настройки
	PC = "PC"
	//Assets константа настройки
	Assets = "Assets"
)

//GetEngineVersion функция геттер для получения версии движка
func GetEngineVersion() string {
	return luaAppVersion
}

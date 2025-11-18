# 📞 Go Call Center Backend для Asterisk

Микросервисная архитектура для мониторинга и управления звонками через Asterisk ARI (REST Interface).

---

## 🎯 Возможности

- ✅ **Мониторинг звонков в реальном времени** через WebSocket
- ✅ **Логирование всех событий** в SQLite базу данных
- ✅ **REST API** для управления звонками (в разработке)
- ✅ **Модульная архитектура** - легко расширять и поддерживать
- ✅ **Без внешних зависимостей** для Asterisk - прямая работа с WebSocket/HTTP

---

## 🏗️ Архитектура

```
┌─────────────────────────────────────────────────────┐
│                   Asterisk PBX                      │
│              (ARI WebSocket + HTTP)                 │
└──────────────────┬──────────────────────────────────┘
                   │
                   │ WebSocket (события)
                   │ HTTP REST (управление)
                   │
┌──────────────────▼──────────────────────────────────┐
│              Go Backend Application                 │
│                                                      │
│  ┌──────────────┐  ┌──────────────┐  ┌───────────┐ │
│  │  WebSocket   │  │    Events    │  │ Database  │ │
│  │    Client    │─→│  Processor   │─→│Repository │ │
│  └──────────────┘  └──────────────┘  └───────────┘ │
│                                                      │
│  ┌──────────────┐  ┌──────────────┐                │
│  │  HTTP Client │  │    Config    │                │
│  │   (ARI API)  │  │   Manager    │                │
│  └──────────────┘  └──────────────┘                │
└─────────────────────────────────────────────────────┘
                   │
                   ▼
           ┌──────────────┐
           │   SQLite DB  │
           └──────────────┘
```

---

## 📁 Структура проекта

```
go-server-asterisk/
├── cmd/
│   └── main.go                    # Точка входа приложения
├── internal/                      # Внутренние модули
│   ├── config/                    # Конфигурация
│   │   └── config.go              
│   ├── websocket/                 # WebSocket клиент для Asterisk
│   │   ├── client.go              # Подключение и чтение событий
│   │   └── handler.go             # Обработчики событий
│   ├── http/                      # HTTP клиент для ARI REST API
│   │   └── client.go              
│   ├── events/                    # Обработка событий звонков
│   │   ├── processor.go           # Процессор событий
│   │   └── handlers.go            # Обработчики типов событий
│   └── database/                  # База данных
│       ├── models.go              # Модели данных
│       ├── repository.go          # CRUD операции
│       └── connection.go          # Подключение к БД
├── go.mod                         # Go модули
├── go.sum                         # Checksums зависимостей
└── README.md                      # Этот файл
```

---

## 🚀 Быстрый старт

### 1. Требования

- **Go** 1.21+
- **Asterisk** 18+ с включенным ARI
- **Linux/Windows** система

### 2. Установка

```bash
# Клонируй репозиторий
git clone <repo-url>
cd go-server-asterisk

# Установи зависимости
go mod download
```

### 3. Настройка Asterisk

#### 3.1. Включи ARI в `/etc/asterisk/ari.conf`:

```ini
[general]
enabled = yes
pretty = yes
allowed_origins = *

[asterisk_dev]
type = user
read_only = no
password = your_secure_password
```

#### 3.2. Настрой HTTP в `/etc/asterisk/http.conf`:

```ini
[general]
enabled = yes
bindaddr = 0.0.0.0
bindport = 8088
prefix = asterisk
```

#### 3.3. Настрой dialplan в `/etc/asterisk/extensions.conf`:

```ini
[internal]
exten => _X.,1,NoOp(Call to ${EXTEN})
 same => n,Stasis(callcenter-app,${EXTEN})
 same => n,Dial(PJSIP/${EXTEN},30,tr)
 same => n,Hangup()
```

#### 3.4. Перезагрузи Asterisk:

```bash
sudo systemctl restart asterisk
# или
sudo asterisk -rx "core reload"
sudo asterisk -rx "dialplan reload"
```

### 4. Конфигурация приложения

Создай файл `.env` или установи переменные окружения:

```bash
# Asterisk ARI
ARI_URL=http://192.168.88.205:8088/asterisk/ari
ARI_USERNAME=asterisk_dev
ARI_PASSWORD=your_secure_password
ARI_APP_NAME=callcenter-app

# База данных
DB_PATH=./callcenter.db

# HTTP сервер (будущее)
HTTP_PORT=3000
```

### 5. Запуск

```bash
go run cmd/main.go
```

**Ожидаемый вывод:**

```
🚀 Запуск Call Center Server...
⚙️  Конфигурация загружена
📦 База данных инициализирована
🔌 Подключено к Asterisk ARI
✅ WebSocket подключен
✅ Слушатель событий запущен
📞 Готов к приему звонков!
```

---

## 📊 Использование

### Мониторинг звонков

После запуска приложение автоматически логирует все звонки:

```bash
📞 НОВЫЙ ЗВОНОК: 102 -> 101 (Channel: PJSIP/102-00000001)
✅ Звонок сохранен в БД (ID: 1)
🔄 ИЗМЕНЕНИЕ СОСТОЯНИЯ: Ringing
✅ Звонок отвечен: 102 -> 101
📴 ЗВОНОК ЗАВЕРШЕН (Длительность: 45 сек)
```

### Просмотр данных в БД

```bash
# Открой SQLite БД
sqlite3 callcenter.db

# Последние 10 звонков
SELECT * FROM calls ORDER BY created_at DESC LIMIT 10;

# Статистика по статусам
SELECT status, COUNT(*) FROM calls GROUP BY status;
```

---

## 🔧 Технологии

- **Go 1.21+** - основной язык
- **gorilla/websocket** - WebSocket клиент
- **SQLite** (modernc.org/sqlite) - база данных (без CGO)
- **Asterisk ARI** - REST API + WebSocket для управления звонками

---

## 📝 Модули

### 1. **config** - управление конфигурацией
- Загрузка из переменных окружения
- Валидация параметров
- Значения по умолчанию

### 2. **websocket** - WebSocket клиент
- Подключение к Asterisk ARI WebSocket
- Автоматический reconnect при обрыве
- Чтение событий в реальном времени

### 3. **events** - обработка событий
- `StasisStart` - начало звонка
- `ChannelStateChange` - изменение состояния
- `ChannelDestroyed` - завершение звонка
- `ChannelDtmfReceived` - DTMF события

### 4. **database** - работа с БД
- Модели: Call, Endpoint
- CRUD операции
- История звонков

### 5. **http** (в разработке)
- HTTP клиент для ARI REST API
- Управление каналами и мостами
- Click-to-call функционал

---

## 🎯 Roadmap

- [x] WebSocket подключение к Asterisk
- [x] Логирование событий звонков
- [x] Сохранение в SQLite
- [ ] HTTP REST API для управления
- [ ] Создание мостов (bridge) между каналами
- [ ] Click-to-call функционал
- [ ] Web интерфейс (React)
- [ ] Запись разговоров
- [ ] Real-time dashboard
- [ ] Metrics и мониторинг

---

## 🐛 Troubleshooting

### WebSocket не подключается

1. Проверь что Asterisk запущен:
   ```bash
   sudo systemctl status asterisk
   ```

2. Проверь что ARI включен:
   ```bash
   sudo asterisk -rx "http show status"
   ```

3. Проверь доступность порта:
   ```bash
   curl http://192.168.88.205:8088/asterisk/ari/asterisk/info
   ```

### События не приходят

1. Проверь dialplan:
   ```bash
   sudo asterisk -rx "dialplan show internal"
   ```
   Должен быть `Stasis(callcenter-app)`

2. Проверь что приложение создано:
   ```bash
   sudo asterisk -rvvv
   ```
   Должно быть: `Creating Stasis app 'callcenter-app'`

### Ошибка CGO/SQLite

Используется `modernc.org/sqlite` - чистый Go драйвер, CGO не нужен.

Если ошибка:
```bash
go get modernc.org/sqlite
go mod tidy
```

---

## 📄 Лицензия

MIT License

---

## 👥 Автор

Разработано для колл-центра с использованием Asterisk PBX и Go.

---

## 🔗 Полезные ссылки

- [Asterisk ARI Documentation](https://docs.asterisk.org/Asterisk_18_Documentation/API_Documentation/Asterisk_REST_Interface/)
- [Asterisk Wiki](https://wiki.asterisk.org/)
- [Go gorilla/websocket](https://github.com/gorilla/websocket)
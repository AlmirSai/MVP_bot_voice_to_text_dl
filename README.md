
# <span style="color: #4CAF50;">Telegram-AI-BotTelegram Bot</span>

![Preview](preview.png)

This is a Telegram bot that converts voice messages to text using **Whisper** and **FFmpeg**. It supports switching between different Whisper models and is optimized for handling voice transcription in Russian (`ru`).

---

## <span style="color: #2196F3;">Features</span>

- Convert voice messages to text using OpenAI Whisper.
- Switch between Whisper models via commands.
- Error handling for invalid inputs or processing failures.
- Configurable for different languages and models.

---

## <span style="color: #2196F3;">Prerequisites</span>

Before running the bot, ensure you have the following installed:

1. **FFmpeg**  
   Install FFmpeg using the appropriate method for your operating system.

   - On **MacOS**:
     ```bash
     brew install ffmpeg
     ```
   - On **Linux** (Ubuntu/Debian-based systems):
     ```bash
     sudo apt update
     sudo apt install ffmpeg
     ```

2. **Whisper**  
   Follow the installation guide for [Whisper](https://github.com/openai/whisper). You can install it using `pip`:
   ```bash
   pip install -U openai-whisper
   ```

3. **Go Programming Language**  
   Install Go from the official site: [https://go.dev](https://go.dev).

   - On **MacOS**:
     ```bash
     brew install go
     ```
   - On **Linux**:
     ```bash
     sudo apt update
     sudo apt install golang-go
     ```

4. **Telegram Bot Token**  
   Get a bot token from [BotFather](https://core.telegram.org/bots#botfather) on Telegram.

---

## <span style="color: #2196F3;">Installation</span>

1. Clone the repository:
   ```bash
   git clone git@github.com:AlmirSai/MVP_bot_voice_to_text_dl.git
   cd MVP_bot_voice_to_text_dl
   ```

2. Set up the environment:
   - Create a `.env` file and add the following variables:
   ```bash
   touch .env
   ```
   - Enter your Telegram bot token in the `.env` file:
   ```bash
   TELEGRAM_TOKEN="YOUR_TELEGRAM_BOT_TOKEN"
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the bot:
   ```bash
   go run bot/main.go
   ```

---

## <span style="color: #2196F3;">Usage</span>

1. Start the bot:
   ```bash
   go run bot/main.go
   ```

2. Add the bot to a chat on Telegram and send a voice message. The bot will transcribe it and return the text.

3. Use the command `model: type` to switch between models.

---

## <span style="color: #2196F3;">File Structure</span>

```plaintext
├── Dockerfile                  // Docker configuration
├── LICENSE                     // License information
├── Makefile                    // Makefile for building and running the bot
├── README.md                   // README file
├── ROADMAP.md                  // Roadmap for the project
├── SECURITY.md                 // Security policy
├── TODO.md                     // TODO list
├── bot                         // Bot directory
│   ├── cmd                     // Bot command directory
│   │   └── main.go             // Main bot file
│   ├── config                  // Configuration directory
│   │   ├── config.go           // Configuration file
│   │   ├── config_test.go      // Configuration test file
│   ├── handlers                // Handler directory
│   │   ├── command_handler.go  // Command handler
│   │   ├── text_handler.go     // Text handler
│   │   └── voice_handler.go    // Voice handler
│   └── utils                   // Utility directory
│       ├── executils           // Execution utilities directory
│       │   ├── executils.go    // Execution utilities
│       │   └── executils_test.go // Execution utilities test
│       ├── file_utils.go       // File utilities
│       ├── logger              // Logger directory
│       │   ├── logger.go       // Logger
│       │   ├── logger_test.go  // Logger test
│       └── speech_to_text.go   // Speech-to-text utilities
├── docker-compose.yml          // Docker Compose configuration
├── go.mod                      // Go module file
├── go.sum                      // Go module checksum file
├── mypy.ini                    // Mypy configuration file
├── pyrightconfig.json          // Pyright configuration file
├── server                      // Server directory
│   ├── app                     // Application directory
│   │   ├── __init__.py         // Application initialization
│   │   ├── config.py           // Application configuration
│   │   ├── main.py             // Application entry point
│   │   ├── routes              // Routes directory
│   │   │   ├── __init__.py     // Routes initialization
│   │   │   └── upload.py       // Upload route
│   │   ├── services            // Services directory
│   │   │   ├── __init__.py     // Services initialization
│   │   │   ├── html_parser.py  // HTML parser
│   │   │   └── json_parser.py  // JSON parser
│   │   └── templates           // Templates directory
│   │       └── upload_form.html // Upload form template
│   └── requirements.txt        // Python dependencies
├── storage                     // Storage directory
└── .github/                    // GitHub configuration(CICD)
```

---

## <span style="color: #2196F3;">Configuration</span>

### Supported Commands
- `model: tiny` - Switch to the **tiny** model.
- `model: base` - Switch to the **base** model.
- `model: small` - Switch to the **small** model.
- `model: medium` - Switch to the **medium** model.
- `model: large` - Switch to the **large** model.

---

## <span style="color: #2196F3;">Troubleshooting</span>

### Common Issues

1. **Missing Dependencies**  
   If you see an error like `Missing dependencies`, ensure FFmpeg and Whisper are installed and accessible in your PATH:
   ```bash
   ffmpeg -version
   whisper --help
   ```

2. **Permission Denied**  
   If the bot cannot execute FFmpeg or Whisper, check their permissions:
   ```bash
   chmod +x $(which ffmpeg)
   chmod +x $(which whisper)
   ```

---

## <span style="color: #2196F3;">Contributing</span>

Contributions are welcome! Feel free to submit issues, fork the repository, and create pull requests.

---

## <span style="color: #2196F3;">License</span>

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details.

---

## <span style="color: #2196F3;">Acknowledgments</span>

- [OpenAI Whisper](https://github.com/openai/whisper) for the transcription tool.
- [FFmpeg](https://ffmpeg.org/) for audio processing.
- [Telegram Bot API](https://core.telegram.org/bots/api) for the bot framework.
```

### Key Updates:
- **MacOS and Linux instructions**: Updated installation instructions for FFmpeg, Whisper, and Go using `brew` for MacOS and standard commands for Linux.
- **Hyperlinks**: Added links to installation guides for external tools.
- **Clarity**: Streamlined commands and instructions for easier understanding.

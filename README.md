# <span style="color: #4CAF50;">Voice-to-Text Telegram Bot</span>

This is a Telegram bot that converts voice messages to text using **Whisper** and **FFmpeg**. It supports switching between different Whisper models and is optimized for handling voice transcription in Russian (`ru`).

---

## <span style="color: #2196F3;">Features</span>

- Convert voice messages to text using OpenAI Whisper.
- Switch between Whisper models (`tiny`, `large`) via commands.
- Error handling for invalid inputs or processing failures.
- Configurable for different languages and models.

---

## <span style="color: #2196F3;">Prerequisites</span>

Before running the bot, ensure you have the following installed:

1. **FFmpeg**  
   Install FFmpeg:
   ```bash
   sudo apt install ffmpeg
   ```

2. **Whisper**  
   Follow the installation guide [here](https://github.com/openai/whisper).

3. **Go Programming Language**  
   Install Go from the official site: [https://go.dev](https://go.dev).

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
   - Create .env file and add following variables:
   ```bash
   touch .env
   ```
   - Enter your telegram bot token:
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

3. Use the command `model: tiny` or `model: large` to switch between models.

---

## <span style="color: #2196F3;">File Structure</span>

```plaintext
voice-to-text-bot/
├── bot/
│   ├── config/
│   │   └── config.go        # Dependency checks and configuration
│   ├── handlers/
│   │   ├── text_handler.go  # Handle text messages
│   │   └── voice_handler.go # Handle voice messages
│   ├── utils/
│   │   ├── file_utils.go    # File operations (download, cleanup)
│   │   └── speech_to_text.go# Whisper integration
│   ├── main.go              # Entry point
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

## <span style="color: #2196F3;">Configuration</span>

### Supported Commands
- `model: tiny` - Switch to the **tiny** model.
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

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## <span style="color: #2196F3;">Acknowledgments</span>

- [OpenAI Whisper](https://github.com/openai/whisper) for the transcription tool.
- [FFmpeg](https://ffmpeg.org/) for audio processing.
- [Telegram Bot API](https://core.telegram.org/bots/api) for the bot framework.

```

### What this update includes:
- Added color to the section titles for better readability and design.
- Emphasis on key actions and parts of the documentation (like the installation and usage commands).
- Cleaned up and maintained a structure that ensures ease of reading.

You can further improve this by embedding images, adding badges, or additional HTML components, but for now, this is a good start with a clean, professional look.
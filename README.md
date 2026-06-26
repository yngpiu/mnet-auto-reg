# Mnet Plus Account Auto Creator

Tool that automatically creates Mnet Plus accounts. You only enter the number of accounts — the tool does everything else.

---

## Quick Start

### Option 1: Download pre-built binary

Download the `.exe` file and double-click to run.

### Option 2: Build from source

1. Install Go from [go.dev](https://go.dev/)
2. Run these commands:

```
go build -o mnet-auto.exe .
mnet-auto
```

Done! A menu will appear on your screen.

---

## Detailed Guide

### How to open Command Prompt / Terminal?

**Windows:**
- Press `Windows` + `R` on your keyboard
- Type `cmd` then press **Enter**
- Or type `powershell` then press **Enter**

**macOS:**
- Press `Command + Space` to open Spotlight
- Type "Terminal" and press **Enter**

**Linux:**
- Press `Ctrl + Alt + T`

### How to check if Node.js is installed?

Type this command and press Enter:

```
node --version
```

- If you see something like `v18.x.x` or higher → **Already installed, skip Step 1**
- If you see an error like "command not found" → **Not installed, follow Step 1 above**

---

## After the tool starts

You will see a menu like this on your screen:

```

Mnet Plus Account Auto Creator

  Current password: nmixx0222-
  Current language: English
  Current mode: Auto

  1. Create Account
  2. Switch Mode
  3. Change Password
  4. Change Language
  5. Exit

  Select:
```

### What each option does

| Press | What it does |
| ----- | ------------ |
| **1** | Start creating accounts — tool does everything automatically |
| **2** | Change how the tool works (see "2 modes" below) |
| **3** | Change the password used for new accounts |
| **4** | Switch the menu language (English, Tiếng Việt, 简体中文, 繁體中文, 한국어) |
| **5** | Exit the tool |

### 2 modes the tool can work in

Press **2** to see the mode selection menu:

```

  Switch Mode

  1. Auto (current)
  2. Manual
  3. Exit

  Select:
```

| Mode | What it means | When to use |
| ---- | ------------- | ----------- |
| **Auto** | Tool creates a temp email + registers via API automatically | **Recommended** — you only enter the number of accounts |
| **Manual** | You type your own email, tool registers it via API | You want to use your own email address |

### How to create accounts

1. Press **1** then press Enter
2. Type the number of accounts you want (for example: `3`) then press Enter
3. The tool will automatically:
   - Create a temporary email address
   - Fill out the Mnet registration form
   - Wait for the confirmation email
   - Click the confirmation link
   - Save the account information
4. When done, you will see a summary like `Done! 3/3 accounts created successfully`

### Where are accounts saved?

All accounts are saved in a file called `accounts.txt` in the same folder where you run the tool.

Open this file with Notepad (Windows) or TextEdit (macOS) or any text editor.

The file looks like this:

```
--- 20.05.26 14:33:12 ---
abc123@tempmail.io
xyz789@tempmail.io
--- 20.05.26 15:01:45 ---
def456@tempmail.io
ghi012@tempmail.io
```

Each time you create accounts, the tool adds a new section with the date and time, then lists the emails. The file grows over time — old accounts are never deleted.

> **Important:** The emails are temporary. If you don't save them somewhere else, you won't be able to recover your accounts later.

### Changing the password

> **Important:** The default password `nmixx0222-` does **not meet** Mnet's password requirements. You should change it before creating accounts.

Password requirements:
- 8 to 20 characters long
- Has at least one UPPERCASE letter
- Has at least one lowercase letter
- Has at least one number
- Has at least one special character (like `!@#$%`)

Example of a valid password: `Nmixx0222-`

How to change:
1. Press **3** then press Enter
2. Type your new password then press Enter
3. If the password is valid, you'll see "Password changed successfully"
4. If it's not valid, the tool will tell you what's missing

### Changing the language

Press **4** to see the language selection menu:

```

  Change Language

  1. English
  2. Tiếng Việt
  3. 简体中文
  4. 繁體中文
  5. 한국어
  6. Exit

  Select:
```

Pick a number and press Enter. The menu will switch to that language immediately.

---

## Common Issues

| Problem | What to do |
| ------- | --- |
| `mnet-auto: command not found` | Download the pre-built binary or build from source |
| No confirmation email received | Wait 1-2 more minutes. If still nothing, try creating the account again |
| Account creation fails | Check your internet connection. Make sure your password meets all requirements |

---

## Updating the tool

Download the new binary or pull the source and rebuild:

```
git pull
go build -o mnet-auto.exe .
```

---

## Notes

- The tool waits 10-30 seconds randomly between accounts to avoid being blocked
- Emails are temporary. Save them somewhere safe if you need them later
- Your settings (password, language, mode) are saved automatically in a `config.json` file
- This tool is for educational purposes only

---

## License

This tool is for educational purposes only. Use at your own risk.

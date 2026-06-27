# Mnet Auto Reg

Automatically register Mnet Plus accounts. Choose the count, the tool does the rest.

## Install

### Pre-built binary (recommended)

Download from [releases](https://github.com/yngpiu/mnet-auto-reg/releases):

| File | Platform | Arch |
|---|---|---|
| `mnet-auto-x64.exe` | Windows | amd64 |
| `mnet-auto-arm64.exe` | Windows | arm64 |
| `mnet-auto-x64.bin` | Linux | amd64 |
| `mnet-auto-arm64.bin` | Linux | arm64 |
| `mnet-auto-macos-x64` | macOS | amd64 |
| `mnet-auto-macos-arm64` | macOS | arm64 |

### Build from source

```
go build -o mnet-auto.exe .
.\mnet-auto.exe
```

## Usage

```
Mnet Plus Account Auto Registration

  Current password: nmixx0222-
  Current mode: Auto

  1. Create Account
  2. Switch Mode
  3. Change Password
  4. Exit
```

- **Auto** — auto temp email + API registration, retries 3 times with domain rotation
- **Manual** — enter your own email + API registration

Change the default password before creating accounts (requirements: 8-20 chars, uppercase + lowercase + number + special char).

## Notes

- Random delay between accounts to avoid rate limiting
- Accounts saved to `accounts.txt`
- Emails are temporary — save them if you need them later

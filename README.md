# FiveM Server Bot Detection 🤖

## Overview ℹ️

This is a simple Go application designed to detect potential bots on FiveM game servers using the FiveM server API and the Steam API. It fetches player data from a FiveM server using its cfxcode and then checks if any of the players might be bots by querying their Steam profiles.
Official repository: [Link to Official Repository](https://github.com/usershhh/FiveM-Bot-Detection/tree/main)

## Installation 🚀

1. Clone the repository:

    ```bash
    git clone https://github.com/BetterCzz/FiveM-Bot-Detection.git
    ```

2. Goto to the directory:

    ```bash
    cd FiveM-Bot-Detection
    ```

3. Build:

    ```bash
    go build main.go
    ```

## Usage 💻

### Flags 🚩

- `-code`: FiveM CFX Code of the Server.

### Example 🌟

```bash
./main -code CFX_CODE
```

Replace `CFX_CODE` with the CFX Code of the FiveM server you want to check.

## Dependencies 📦

- [Go](https://golang.org/): The programming language used for the application.
- [GitHub](https://github.com/): Hosting for the repository.
- [FiveM Server API](https://servers-frontend.fivem.net/api/servers/single): Provides server data, including player lists.
- [Steam API](http://api.steampowered.com): Used to fetch Steam profiles for player identification.

## Contributing 🤝

Contributions are welcome! If you find a bug or have an enhancement in mind, please open an issue or submit a pull request.

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to adjust any sections or wording to better fit your project's specifics.

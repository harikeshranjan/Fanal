# 📂 Fanal - File Analysis CLI Tool

**Fanal** (short for **File Analysis**) is a simple and powerful Command Line Interface (CLI) tool written in Go that helps you analyze files and directories with ease. Whether you're organizing a project folder or managing large datasets, Fanal provides clear, tabular insights into your directory structure.

## 🚀 Features

| Command  | Description                                                                 |
|----------|-----------------------------------------------------------------------------|
| `desc`   | Describes the files in a directory (name, size, type, etc.) in table format |
| `search` | Search files by keyword, name, or extension                                 |
| `count`  | Count and list the number of files that match a specific keyword or pattern |
| `fsum`   | Prints the summary of a particular file in the specified directory          |
| `top`    | Prints the top `N` files in the directory                                   |

Powered by:

- [Cobra CLI](https://github.com/spf13/cobra) for command structure
- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) for elegant tabular output

## 📦 Installation

1. **Clone the repo**

   ```bash
   git clone https://github.com/your-username/fanal.git
   cd fanal
   ```

   2. **Install Dependencies**

    ```bash
    go mod tidy
    ```

3. **Build the CLI**
  
    ```bash
    go build -o fanal main.go
    ```

4. **Run the CLI**
  
    ```bash
    ./fanal --help
    ```

## 🔍 Commands in Detail

### 1. `desc` - Describe Files

Describes the contents of the given directory. Lists file name, size, extension, and type in a table.

```bash
fanal desc <directory_path>

fanal desc <directory_path> --export=<json/csv>
```

#### Example

```bash
fanal desc .

fanal desc . --export=json
```

This command will output a table with the following columns:

- **Name**: The name of the file
- **Size**: The size of the file in bytes
- **Extension**: The file extension (e.g., .txt, .jpg)
- **Last Modified**: The last modified date of the file

## Example Output

```plaintext
+---------------------+----------+-----------+---------------------+
| Name                | Size     | Extension | Last Modified       |
+---------------------+----------+-----------+---------------------+
| file1.txt           | 1024     | .txt      | 2023-10-01 12:00:00 |
| image1.jpg          | 2048     | .jpg      | 2023-10-02 14:30:00 | 
| video1.mp4          | 5120     | .mp4      | 2023-10-03 16:45:00 |
| document1.pdf       | 3072     | .pdf      | 2023-10-04 18:15:00 |
+---------------------+----------+-----------+---------------------+
```

### 2. `search` - Search Files

Search for files in the given directory by keyword, name, or extension.

```bash
fanal search <directory_path> <accurate_file_name> 
```

#### Example

```bash
fanal search . file1.txt
```

This command will search for files in the current directory that match the name `file1.txt`.
It will return a list of files that match the search criteria.

## 3. `count` - Count Files

Count the number of files in the given directory that match a specific keyword or pattern.

```bash
fanal count <directory_path> <keyword>
```

#### Example

```bash
fanal count . .txt
```

This command will count the number of files in the current directory that have the `.txt` extension.
It will return the count of matching files.

## 4. `fsum` - File Summary

It prints the full summary of the file by giving us the details about words count, letters count, file size, number of lines, extension and last modified.

```bash
fanal fsum <directory_path> <filename>
```

#### Example

```bash
fanal fsum . README.md
```

## 5. `top` - Top `N` files

It prints the top `N` files in the current or the directory specified. If the directory is empty, you get the the `The directory is empty` message

```bash
fanal top --n=<N>
```

#### Example

```bash
fanal top --n=10
```

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 📧 Contact

For any questions or feedback, please reach out to me at [ranjansinhaharikesh@gmail.com](mailto:ranjansinhaharikesh@gmail.com).

Feel free to open issues or submit pull requests on the GitHub repository.

### How to contribute

If you want to contribute to this project, please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/YourFeature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some feature'`)
5. Push to the branch (`git push origin feature/YourFeature`)
6. Open a pull request
7. Discuss and review the changes with the maintainers
8. Once approved, your changes will be merged into the main branch
9. Celebrate your contribution! 🎉

## 🛠️ Tools Used

- Go (Golang) for the CLI implementation
- Cobra for command-line argument parsing
- Tablewriter for tabular output formatting
- Git for version control
- GitHub for hosting the repository
- Markdown for documentation formatting
- Neovim for code editing

---

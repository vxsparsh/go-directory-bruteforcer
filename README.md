# Go Directory Brute Forcer Tool

This tool is a simple directory brute forcer written in Go. It scans a web server for directories specified in a wordlist file, and it can output the results to both the console and an optional output file. The tool also allows for customization of the request delay, user-agent header, and HTTP request timeout.

#### Features
- Specify target URL
- Use a wordlist file for directory paths
- Set delay between requests
- Customize User-Agent header
- Set HTTP request timeout
- Optionally write results to an output file

#### Usage

```bash
./godirb -url <targetURL> -wordlist <wordlistFile> -delay <delayInMilliseconds> -useragent <User-Agent> -timeout <timeoutInSeconds> -output <outputFile>
```

#### Command-Line Arguments

- `-url <targetURL>`: The base URL of the target server.
- `-wordlist <wordlistFile>`: Path to the wordlist file containing directory names.
- `-delay <delayInMilliseconds>`: Delay between requests in milliseconds (default: 0).
- `-useragent <User-Agent>`: User-Agent header for the HTTP requests (default: "Mozilla/5.0").
- `-timeout <timeoutInSeconds>`: Timeout for the HTTP requests in seconds (default: 10).
- `-output <outputFile>`: File to write results to (optional).

#### Example

```bash
./godirb -url http://example.com -wordlist wordlist.txt -delay 1000 -useragent "Mozilla/5.0 (compatible; MyBot/1.0)" -timeout 10 -output results.txt
```

This example will:
- Target the URL `http://example.com`
- Use `wordlist.txt` to get the directory names
- Delay each request by 1000 milliseconds (1 second)
- Use the custom User-Agent `"Mozilla/5.0 (compatible; MyBot/1.0)"`
- Set the HTTP request timeout to 10 seconds
- Write results to `results.txt`


#### Installation

Ensure you have Go installed. Clone the repository or copy the script to your local machine. or download the compiled binary from releases.


#### Example Wordlist File (`wordlist.txt`)

```
admin
backup
secret
hidden
test
```

Each line in the wordlist file represents a directory to be checked on the target server.

#### Important Notes

- **Legal and Ethical Use**: Ensure you have permission to perform directory brute forcing on the target server. Unauthorized use may violate laws and ethical guidelines.
- **Error Handling**: The tool handles errors in accessing URLs and reading the wordlist file, printing appropriate error messages to the console.
- **Resource Management**: The tool ensures proper closing of file and response body resources to prevent leaks.

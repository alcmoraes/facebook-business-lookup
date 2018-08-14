# Facebook Business Lookup

Using DuckDuckGo search, try to get the facebook URL for a business and also its phone number

## Usage (from build)

```./build/fbbl "{BUSINESS NAME}, {BUSINESS LOCATION}"```

### Example Usage

```./build/fbbl "Ghazal Arabian Food, Florianópolis, Santa Catarina"```

### Example output

```
    Ghazal Arabian Food - Home - Florianópolis, Santa Catarina ...
    https://m.facebook.com/ghazal.arabian/
    Ligar (48) 99901-6916
```

## Usage (from source)

### Requirements

```go```
```dep```

### Install

```dep ensure```

### Running
```go run main.go "{BUSINESS NAME}, {BUSINESS LOCATION}"```


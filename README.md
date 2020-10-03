Uveira
------

Command line client for MongoDB Wikipedia imports via [dumpster-dive](https://github.com/spencermountain/dumpster-dive).

## Build

```sh
$ go build
```

## Usage

In order for Uveira to be able to connect to the database the config has to be set as environment variables:

```sh
export UVEIRA_MONGO_URI="mongodb://user:password@mongodb-host:27017" 
export UVEIRA_DATABASE="enwiki"
export UVEIRA_COLLECTION="pages"
```

Only then the tool can be used:

```sh
$ uveira -help
Usage of ./uveira:
  -te string
        Query by title (exact)
  -to
        List titles only in query results
  -tr string
        Query by title (RegEx)
```

## Examples

Get a page by its exact title (case sensitive!):

```sh
$ uveira -te "Tesseract"
```

Get pages by using a RegEx query on their titles:

```sh
$ uveira -tr "^Tesseract.*"
```

Get only page titles by using a RegEx query on their titles:

```sh
$ uveira -tr "^Tesseract.*" -to
```

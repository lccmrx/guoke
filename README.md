
# Guoke (Go + Quake 3) - Log Parser

Logs are structured as:

##:##   xxxxx:   xxxxxxxxxxx

  ┗━━━┓   ┗━━━━━━━┓   ┗━━┓
      ┃           ┃      ┃
      
game stopwatch  event   info

Each game is started with the following lines:
```
 ##:## ------------------------------------------------------------
 ##:## InitGame: ....
```
and ended with the following lines:
``` 
 ##:## ShutdownGame:
 ##:## ------------------------------------------------------------
```

Every event that happens between these two flags, are about the same game match.
> Disclaimer: there were inconsistent [logs found](./q.log#97) (e.g. `InitGame` that wasn't followed by a `ShutdownGame` event). This could be due many things (user interruption on server, seg-fault, missing logs, etc). So we will only consider event logs that are between `InitGame` and `ShutdownGame` boundaries. Events that happened in a faulty state such as [this](./q.log#97-157) won't be accounted for.

The code will replay the events and get the "server" to it's latest state.

## Usage

Simply go with a simple (to run in development mode):
```bash
go run -C cmd . <logpath>
```
or go with a (to run in a productive mode): 
```bash
go build -C cmd . -o guoke
./guoke <logpath>
```

## Roadmap

- [ ] Chunk-based reading
> This will allow to rapidly paralelize the processing for the log events
- [ ] io.Reader events
> This will allow the server to connect to many log event sources, which could be the stdout of the gamer server

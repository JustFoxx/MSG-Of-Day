# MSG-Of-Day
> Simple website that shows message of day in multiple languages

## Done languages

- Golang ✅

## How does it work?
1. Run app to listen on 80 port
2. Show JSON file `msg.json` to main page
3. Read POST on `/msg` with body json `{"message":"<msg>"}`
4. Anything that will be send in `/msg` save in JSON file `msg.json`
5. Repeat from 2.

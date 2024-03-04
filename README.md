# GO Temperature By Zip Code

### Run on dev mode:

To run the project, first you need to go to copy the content of `.env.sample` to a `.env` file
and change the **WEATHER_API_KEY** value by yours:

```SHELL
$ cp .env.sample .env
```

To get a new **WEATHER_API_KEY** you need to create a new account on https://www.weatherapi.com/.

Then, with docker compose installed, just simply run:

```SHELL
docker compose up
```

Then, on any client, call `localhost:8080/{CEP}`.

### Test on production:

To test the live API you just need to call `https://go-temperature-by-zip-code-nvteiykumq-uc.a.run.app/{CEP}` replacing
`{CEP}` by any CEP on Brazil.
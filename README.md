## sample-rates

* [x] Sample golang http-serrver that expose end-points for exhange rates.


### Pre-Requisite
	
	- Please run this in your command line to ensure packages are in-place.
	  (normally these will be handled when compiling the api binary)
	
```sh

		go get -u -v github.com/go-sql-driver/mysql

```

### Compile

```sh

     git clone https://github.com/bayugyug/sample-rates.git && cd sample-rates

     git pull && make clean && make

```

### Required Data Preparation

    - Create sample mysql db
	
	- Needs to create the api database and grant the necessary permissions
	
	- Refer the testdata/dump.sql
	
```sh

	create database rates;
	create user rates;
	grant all privileges on rates.* to rates@localhost identified by 'xxxx';
	grant all privileges on rates.* to rates@127.0.0.1 identified by 'xxxx';
	flush privileges;

```

### List of End-Points-Url


```sh



        #Show the latest rates
        curl -X GET  'http://127.0.0.1:8989/rates/latest'

            @output:
            {
              "base": "EUR",
              "rates": {
                "AUD": 1.5884,
                "BGN": 1.9558,
                "BRL": 4.237,
                "CAD": 1.5154,
                "CHF": 1.1266,
                "CNY": 7.723,
                "CZK": 25.572,
                "DKK": 7.4638,
                "GBP": 0.89025,
                "HKD": 8.9594,
                "HRK": 7.4265,
                "HUF": 322.8,
                "IDR": 16135.26,
                "ILS": 4.19,
                "INR": 81.231,
                "ISK": 138.5,
                "JPY": 124.02,
                "KRW": 1283.26,
                "MXN": 21.7315,
                "MYR": 4.6765,
                "NOK": 9.7615,
                "NZD": 1.6757,
                "PHP": 59.519,
                "PLN": 4.293,
                "RON": 4.6831,
                "RUB": 76.6722,
                "SEK": 10.2403,
                "SGD": 1.5474,
                "THB": 36.443,
                "TRY": 6.228,
                "USD": 1.1424,
                "ZAR": 15.7877
              }
            }



        #Show the latest rates per date parameter
        curl -X GET  'http://127.0.0.1:8989/rates/2019-01-14'

            @output:
            {
              "base": "EUR",
              "rates": {
                "AUD": 1.5945,
                "BGN": 1.9558,
                "BRL": 4.2739,
                "CAD": 1.5226,
                "CHF": 1.1258,
                "CNY": 7.7595,
                "CZK": 25.561,
                "DKK": 7.4639,
                "GBP": 0.89263,
                "HKD": 8.9934,
                "HRK": 7.4325,
                "HUF": 321.33,
                "IDR": 16195.99,
                "ILS": 4.1945,
                "INR": 81.2195,
                "ISK": 138.5,
                "JPY": 123.93,
                "KRW": 1287.82,
                "MXN": 21.983,
                "MYR": 4.7049,
                "NOK": 9.7868,
                "NZD": 1.682,
                "PHP": 59.912,
                "PLN": 4.2925,
                "RON": 4.6815,
                "RUB": 77.0602,
                "SEK": 10.2493,
                "SGD": 1.5526,
                "THB": 36.626,
                "TRY": 6.3241,
                "USD": 1.1467,
                "ZAR": 15.9218
              }
            }



        #Show the latest rates per date parameter (invalid)
        curl -X GET  'http://127.0.0.1:8989/rates/2019-01-14x'

            @output:
            {
                "Code": 404,
                "Status": "Invalid Endpoint"
            }



        #Show the rates summary (min/max/avg)
        curl -X GET  'http://127.0.0.1:8989/rates/analyze'

            @output:
            {
              "base": "EUR",
              "rates_analyze": {
                "AUD": {
                  "min": 1.5354,
                  "max": 1.6287,
                  "avg": 1.586409836066
                },
                "BGN": {
                  "min": 1.9558,
                  "max": 1.9558,
                  "avg": 1.9558
                },
                "BRL": {
                  "min": 4.1231,
                  "max": 4.4786,
                  "avg": 4.319004918033
                },
                "CAD": {
                  "min": 1.4856,
                  "max": 1.5605,
                  "avg": 1.512327868852
                },
                "CHF": {
                  "min": 1.1219,
                  "max": 1.146,
                  "avg": 1.133537704918
                },
                "CNY": {
                  "min": 7.723,
                  "max": 7.9835,
                  "avg": 7.866606557377
                },
                "CZK": {
                  "min": 25.561,
                  "max": 26.032,
                  "avg": 25.84062295082
                },
                "DKK": {
                  "min": 7.4593,
                  "max": 7.4679,
                  "avg": 7.463175409836
                },
                "GBP": {
                  "min": 0.86945,
                  "max": 0.90423,
                  "avg": 0.889757540984
                },
                "HKD": {
                  "min": 8.8161,
                  "max": 9.0427,
                  "avg": 8.920921311475
                },
                "HRK": {
                  "min": 7.387,
                  "max": 7.4393,
                  "avg": 7.421027868852
                },
                "HUF": {
                  "min": 320.98,
                  "max": 325.1,
                  "avg": 322.60737704918
                },
                "IDR": {
                  "min": 16122.63,
                  "max": 17480.12,
                  "avg": 16655.379344262295
                },
                "ILS": {
                  "min": 4.1423,
                  "max": 4.315,
                  "avg": 4.232459016393
                },
                "INR": {
                  "min": 79.0815,
                  "max": 84.685,
                  "avg": 81.471939344262
                },
                "ISK": {
                  "min": 133,
                  "max": 142.3,
                  "avg": 138.059016393443
                },
                "JPY": {
                  "min": 122.21,
                  "max": 130.02,
                  "avg": 127.677704918033
                },
                "KRW": {
                  "min": 1260.44,
                  "max": 1306.55,
                  "avg": 1283.331475409836
                },
                "MXN": {
                  "min": 21.7315,
                  "max": 23.3643,
                  "avg": 22.70302295082
                },
                "MYR": {
                  "min": 4.6765,
                  "max": 4.7879,
                  "avg": 4.74548852459
                },
                "NOK": {
                  "min": 9.4518,
                  "max": 10.0025,
                  "avg": 9.69141147541
                },
                "NZD": {
                  "min": 1.6376,
                  "max": 1.7535,
                  "avg": 1.688227868852
                },
                "PHP": {
                  "min": 59.352,
                  "max": 62.137,
                  "avg": 60.227032786885
                },
                "PLN": {
                  "min": 4.2809,
                  "max": 4.3392,
                  "avg": 4.298652459016
                },
                "RON": {
                  "min": 4.6389,
                  "max": 4.6831,
                  "avg": 4.661142622951
                },
                "RUB": {
                  "min": 74.369,
                  "max": 79.7153,
                  "avg": 76.087991803279
                },
                "SEK": {
                  "min": 10.1753,
                  "max": 10.42,
                  "avg": 10.290703278689
                },
                "SGD": {
                  "min": 1.5474,
                  "max": 1.5863,
                  "avg": 1.562614754098
                },
                "THB": {
                  "min": 36.443,
                  "max": 37.88,
                  "avg": 37.287278688525
                },
                "TRY": {
                  "min": 5.8747,
                  "max": 6.6227,
                  "avg": 6.159029508197
                },
                "USD": {
                  "min": 1.1261,
                  "max": 1.1535,
                  "avg": 1.139321311475
                },
                "ZAR": {
                  "min": 15.5084,
                  "max": 16.7942,
                  "avg": 16.158006557377
                }
              }
            }
```


### Mini-How-To on running the api binary

	[x] Prior to running the server, db must be configured first 
	
    [x] The api can accept a json format configuration
	
	[x] Fields:
	
		- http_port = port to run the http server (default: 8989)
		
		- driver    = database details for mysql  (user/pass/dbname/host/port)
		
		- showlog   = flag for dev't log on std-out
		
	[x] Sanity check
	    
		go test ./...
	
	[x] Run from the console

```sh
	./sample-rates --config '{
                "http_port":"8989",
                    "driver":{
                    "user":"rates",
                    "pass":"rat3s",
                    "port":"3306",
                    "name":"rates",
                    "host":"127.0.0.1"},
                "showlog":true}'


```

### Notes

[EUR-EXCHANGE-RATE-URL](https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml)


### Reference
	

### License

[MIT](https://bayugyug.mit-license.org/)


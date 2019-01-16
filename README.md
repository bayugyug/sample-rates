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

        #Show the latest rates per date parameter
        curl -X GET  'http://127.0.0.1:8989/rates/2019-01-14'

        #Show the rates summary (min/max/avg)
        curl -X GET  'http://127.0.0.1:8989/rates/analyze'

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



### Reference
	

### License

[MIT](https://bayugyug.mit-license.org/)

